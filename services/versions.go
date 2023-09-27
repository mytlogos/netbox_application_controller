package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/mytlogos/netbox_application_controller/services/backend"
	"golang.org/x/exp/slices"
)

type VersionSearcher struct {
	backend *backend.Backend
	argus   *ArgusConfig
}

type ArgusConfig struct {
	Server string
}

type VersionSearcherConfig struct {
	Backend *backend.Backend
	Argus   *ArgusConfig
}

func NewSearcher(config VersionSearcherConfig) *VersionSearcher {
	if config.Argus == nil {
		panic("VersionSearcher requires a version backend config, 'Argus' has no config")
	}
	copy := *config.Argus
	return &VersionSearcher{
		backend: config.Backend,
		argus:   &copy,
	}
}

func (v *VersionSearcher) makeRequest(path string, unmarshal interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/%s", v.argus.Server, path))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, unmarshal)
}

func (v *VersionSearcher) getNames() ([]string, error) {
	var names struct {
		Order []string
	}
	err := v.makeRequest("service/order", &names)
	return names.Order, err
}

func (v *VersionSearcher) getSummary(name string) (argusServiceSummary, error) {
	var summary argusServiceSummary
	err := v.makeRequest(fmt.Sprintf("service/summary/%s", name), &summary)
	return summary, err
}

func (v *VersionSearcher) Search() error {
	applications, err := v.backend.GetApplications()

	if err != nil {
		return err
	}

	names, err := v.getNames()

	if err != nil {
		return err
	}

	for _, app := range applications {
		name := app.GetVersionManagerName()

		if name == "" || !slices.Contains(names, name) {
			continue
		}
		summary, err := v.getSummary(name)

		if err != nil {
			continue
		}

		changed := false

		if app.GetVersion() != summary.Status.DeployedVersion {
			app.SetVersion(summary.Status.DeployedVersion)
			changed = true
		}

		if app.GetLastVersionUpgrade() != summary.Status.DeployedVersionTimestamp {
			app.SetLastVersionUpgrade(summary.Status.DeployedVersionTimestamp)
			changed = true
		}

		if app.GetLatestVersion() != summary.Status.LatestVersion {
			app.LatestVersion = &summary.Status.LatestVersion
			changed = true
		}

		if app.GetChangelogUrl() != summary.URL && summary.URL != "" {
			app.ChangelogUrl = &summary.URL
			changed = true
		}

		if changed {
			_, err = v.backend.UpdateApplication(app)

			if err != nil {
				slog.Error("error while updating app", "err", err)
			}
		}
	}
	return nil
}

type argusServiceSummary struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	URL                string `json:"url"`
	Icon               string `json:"icon"`
	IconLinkTo         string `json:"icon_link_to"`
	HasDeployedVersion bool   `json:"has_deployed_version"`
	Command            int    `json:"command"`
	Webhook            int    `json:"webhook"`
	Status             struct {
		ApprovedVersion          string    `json:"approved_version"`
		DeployedVersion          string    `json:"deployed_version"`
		DeployedVersionTimestamp time.Time `json:"deployed_version_timestamp"`
		LatestVersion            string    `json:"latest_version"`
		LatestVersionTimestamp   time.Time `json:"latest_version_timestamp"`
		LastQueried              time.Time `json:"last_queried"`
	} `json:"status"`
}
