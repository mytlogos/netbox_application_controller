package services

import (
	"encoding/json"

	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/netbox"
	"github.com/mytlogos/netbox_application_controller/services/backend"

	"github.com/r3labs/diff/v3"
)

type Differ struct {
	backend   *backend.Backend
	agentUuid string
	device    *netbox.DeviceWithConfigContext
}

type DifferConfig struct {
	Backend   *backend.Backend
	AgentUuid string
}

func NewDiffer(config DifferConfig) *Differ {
	return &Differ{
		backend:   config.Backend,
		agentUuid: config.AgentUuid,
	}
}

func (d *Differ) loadBackendDoc() (*models.Host, error) {
	_, err := d.backend.GetDeviceByUuid(d.agentUuid)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *Differ) DiffDocument(doc *models.Host, beDoc *models.Host) (string, error) {
	differ, err := diff.NewDiffer()

	if err != nil {
		return "", err
	}

	changes, err := differ.Diff(doc, beDoc)

	if err != nil {
		return "", err
	}

	data, err := json.MarshalIndent(changes, "", " ")

	if err != nil {
		return "", err
	}

	return string(data), nil
}
