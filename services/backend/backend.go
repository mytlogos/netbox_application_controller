package backend

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"regexp"
	"strings"

	"github.com/mytlogos/netbox_application_controller/netbox"
)

type contextType string

var (
	CustomFieldAgent          = "agent"
	CustomFieldUuid           = "agentuuid"
	ContextCustomFieldQueries = contextType("netbox-customfield-query")
	ErrValueNotFound          = fmt.Errorf("netbox error: value not found")
)

type BackendConfig struct {
	Server string
	Apikey string
}

type customTransport struct {
}

func (t *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	mapping, ok := r.Context().Value(ContextCustomFieldQueries).(map[string]string)

	if ok && mapping != nil {
		query := r.URL.Query()

		for fieldName, fieldValue := range mapping {
			query.Add(fieldName, fieldValue)
		}

		// Encode the parameters. (copied from client.go)
		queryParamSplit := regexp.MustCompile(`(^|&)([^&]+)`)
		queryDescape := strings.NewReplacer("%5B", "[", "%5D", "]")

		r.URL.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
			pieces := strings.Split(s, "=")
			pieces[0] = queryDescape.Replace(pieces[0])
			return strings.Join(pieces, "=")
		})
	}
	return http.DefaultTransport.RoundTrip(r)
}

func NewBackend(config BackendConfig) *Backend {
	netboxConfig := netbox.NewConfiguration()
	netboxConfig.Debug = true
	netboxConfig.Servers = netbox.ServerConfigurations{
		{
			URL:         config.Server,
			Description: "NetBox",
		},
	}
	// use default transport to support custom field queries from context
	netboxConfig.HTTPClient = &http.Client{
		Transport: &customTransport{},
	}
	return &Backend{
		client: netbox.NewAPIClient(netboxConfig),
		apiKey: map[string]netbox.APIKey{
			"tokenAuth": {
				Key:    config.Apikey,
				Prefix: "Token",
			},
		},
	}
}

type Backend struct {
	client *netbox.APIClient
	apiKey map[string]netbox.APIKey
}

func (b *Backend) getContext() context.Context {
	return context.WithValue(context.Background(), netbox.ContextAPIKeys, b.apiKey)
}

func (b *Backend) GetDevices() ([]netbox.DeviceWithConfigContext, error) {
	ctx := b.getContext()
	list, _, err := b.client.DcimApi.DcimDevicesList(ctx).Execute()

	if err != nil {
		return nil, err
	}

	return list.GetResults(), nil
}

func (b *Backend) GetPlatforms() ([]netbox.Platform, error) {
	ctx := b.getContext()
	list, _, err := b.client.DcimApi.DcimPlatformsList(ctx).Execute()

	if err != nil {
		return nil, err
	}

	return list.GetResults(), nil
}

func (b *Backend) CreatePlatform(platform netbox.Platform) (*netbox.Platform, error) {
	ctx := b.getContext()
	result, _, err := b.client.DcimApi.DcimPlatformsCreate(ctx).WritablePlatformRequest(
		netbox.WritablePlatformRequest{
			Name: platform.Name,
			Slug: platform.Slug,
		},
	).Execute()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b *Backend) GetDeviceByUuid(uuid string) (*netbox.DeviceWithConfigContext, error) {
	ctx := b.getContext()
	ctx = context.WithValue(ctx, ContextCustomFieldQueries, map[string]string{
		CustomFieldUuid: uuid,
	})
	list, _, err := b.client.DcimApi.DcimDevicesList(ctx).Execute()

	if err != nil {
		return nil, err
	}

	values := list.GetResults()

	if len(values) == 1 {
		return &values[0], nil
	}
	return nil, ErrValueNotFound
}

func (b *Backend) CreateDevice(device *netbox.DeviceWithConfigContext) (*netbox.DeviceWithConfigContext, error) {
	ctx := b.getContext()

	var platformId netbox.NullableInt32

	if device.Platform.Get() != nil {
		platformId = *netbox.NewNullableInt32(&device.Platform.Get().Id)
	}
	req := b.client.DcimApi.DcimDevicesCreate(ctx)
	req = req.WritableDeviceWithConfigContextRequest(netbox.WritableDeviceWithConfigContextRequest{
		Name: device.Name,
		// DeviceType:   device.DeviceType.Id,
		// Role:         device.Role.Id,
		Platform:     platformId,
		CustomFields: device.CustomFields,
	})
	created, _, err := req.Execute()
	return created, err
}

func (b *Backend) UpdateDevice(id int32, device *netbox.DeviceWithConfigContext) (*netbox.DeviceWithConfigContext, error) {
	ctx := b.getContext()

	var platformId netbox.NullableInt32

	if device.Platform.Get() != nil {
		platformId = *netbox.NewNullableInt32(&device.Platform.Get().Id)
	}

	req := b.client.DcimApi.DcimDevicesUpdate(ctx, id)
	req = req.WritableDeviceWithConfigContextRequest(netbox.WritableDeviceWithConfigContextRequest{
		Name:         device.Name,
		DeviceType:   device.DeviceType.Id,
		Role:         device.Role.Id,
		Site:         device.Site.Id,
		Platform:     platformId,
		CustomFields: device.CustomFields,
	})
	created, _, err := req.Execute()
	return created, err
}

func (b *Backend) GetInterfaces(deviceId int32) ([]netbox.Interface, error) {
	ctx := b.getContext()
	req := b.client.DcimApi.DcimInterfacesList(ctx)
	req = req.DeviceId([]int32{deviceId})
	list, _, err := req.Execute()
	return list.GetResults(), err
}

func (b *Backend) CreateIPAddress(value netbox.IPAddress) (*netbox.IPAddress, error) {
	ctx := b.getContext()
	req := b.client.IpamApi.IpamIpAddressesCreate(ctx)
	req = req.WritableIPAddressRequest(netbox.WritableIPAddressRequest{
		Address:            value.Address,
		DnsName:            value.DnsName,
		Status:             value.Status.Value,
		AssignedObjectType: value.AssignedObjectType,
		AssignedObjectId:   value.AssignedObjectId,
	})
	netInterface, _, err := req.Execute()
	return netInterface, err
}

func (b *Backend) UpdateIPAddress(value netbox.IPAddress) (*netbox.IPAddress, error) {
	ctx := b.getContext()
	req := b.client.IpamApi.IpamIpAddressesUpdate(ctx, value.Id)
	req = req.WritableIPAddressRequest(netbox.WritableIPAddressRequest{
		Address:            value.Address,
		DnsName:            value.DnsName,
		Status:             value.Status.Value,
		AssignedObjectType: value.AssignedObjectType,
		AssignedObjectId:   value.AssignedObjectId,
	})
	netInterface, _, err := req.Execute()
	return netInterface, err
}

func (b *Backend) DeleteIP(ipId int32) error {
	ctx := b.getContext()
	req := b.client.IpamApi.IpamIpAddressesDestroy(ctx, ipId)
	_, err := req.Execute()
	return err
}

func (b *Backend) GetIPAddressByInterface(interfaceId int32) ([]netbox.IPAddress, error) {
	ctx := b.getContext()
	req := b.client.IpamApi.IpamIpAddressesList(ctx)
	req = req.InterfaceId([]int32{interfaceId})
	netInterface, _, err := req.Execute()
	return netInterface.GetResults(), err
}

func (b *Backend) GetIPAddressByDevice(deviceId int32) ([]netbox.IPAddress, error) {
	ctx := b.getContext()
	req := b.client.IpamApi.IpamIpAddressesList(ctx)
	req = req.DeviceId([]int32{deviceId})
	netInterface, _, err := req.Execute()
	return netInterface.GetResults(), err
}

func (b *Backend) CreateInterface(deviceId int32, value netbox.Interface) (*netbox.Interface, error) {
	ctx := b.getContext()
	req := b.client.DcimApi.DcimInterfacesCreate(ctx)
	req = req.WritableInterfaceRequest(netbox.WritableInterfaceRequest{
		Device:     deviceId,
		Name:       value.Name,
		Type:       *value.Type.Value,
		MacAddress: value.MacAddress,
		Vdcs:       []int32{},
	})
	netInterface, _, err := req.Execute()
	return netInterface, err
}

func (b *Backend) UpdateInterface(value netbox.Interface) (*netbox.Interface, error) {
	ctx := b.getContext()
	req := b.client.DcimApi.DcimInterfacesPartialUpdate(ctx, value.Id)
	req = req.PatchedWritableInterfaceRequest(netbox.PatchedWritableInterfaceRequest{
		Device:     &value.Device.Id,
		Name:       &value.Name,
		Type:       value.Type.Value,
		MacAddress: value.MacAddress,
		Enabled:    value.Enabled,
		Mtu:        value.Mtu,
	})
	netInterface, _, err := req.Execute()
	return netInterface, err
}

func (b *Backend) DeleteInterface(interfaceId int32) error {
	ctx := b.getContext()
	req := b.client.DcimApi.DcimInterfacesDestroy(ctx, interfaceId)
	_, err := req.Execute()
	return err
}

func (b *Backend) GetApplicationGroups() ([]netbox.ApplicationGroup, error) {
	ctx := b.getContext()
	result, _, err := b.client.PluginsApi.PluginsApplicationsApplicationGroupsList(ctx).Execute()
	return result.GetResults(), err
}

func (b *Backend) CreateApplicationGroup(group netbox.ApplicationGroup) (*netbox.ApplicationGroup, error) {
	ctx := b.getContext()
	result, _, err := b.client.PluginsApi.
		PluginsApplicationsApplicationGroupsCreate(ctx).
		ApplicationGroupRequest(netbox.ApplicationGroupRequest{
			Name: group.Name,
		}).
		Execute()
	return result, err
}

func (b *Backend) GetApplications() ([]netbox.Application, error) {
	ctx := b.getContext()
	result, _, err := b.client.PluginsApi.PluginsApplicationsApplicationsList(ctx).Execute()
	return result.GetResults(), err
}

func (b *Backend) CreateApplication(group netbox.Application) (*netbox.Application, error) {
	ctx := b.getContext()
	result, _, err := b.client.PluginsApi.
		PluginsApplicationsApplicationsCreate(ctx).
		ApplicationRequest(netbox.ApplicationRequest{
			Name:   group.Name,
			Status: group.Status,
		}).
		Execute()
	return result, err
}

func (b *Backend) InitBackend() error {
	// create necessary backend custom fields
	ctx := b.getContext()
	list, _, err := b.client.ExtrasApi.ExtrasCustomFieldsList(ctx).Execute()

	if err != nil {
		return err
	}
	results := list.GetResults()

	agentFieldExists := false
	agentUuidFieldExists := false

	for _, customField := range results {
		if customField.Name == CustomFieldAgent {
			agentFieldExists = true
		}
		if customField.Name == CustomFieldUuid {
			agentUuidFieldExists = true
		}
	}

	if !agentFieldExists {
		slog.Info("custom field for agent missing on device, recreating")
		req := b.client.ExtrasApi.ExtrasCustomFieldsCreate(ctx)
		description := "Agent Storage"
		fieldType := "json"

		req = req.WritableCustomFieldRequest(netbox.WritableCustomFieldRequest{
			Name:         CustomFieldAgent,
			ContentTypes: []string{"dcim.device"},
			Description:  &description,
			Type:         &fieldType,
		})
		_, resp, err := req.Execute()

		if err != nil {
			data, _ := io.ReadAll(resp.Body)
			log.Println(string(data))
			return err
		}
	}

	if !agentUuidFieldExists {
		slog.Info("custom field for agent uuid missing on device, recreating")
		req := b.client.ExtrasApi.ExtrasCustomFieldsCreate(ctx)
		description := "Custom Field for querying device by agent uuid"
		fieldType := "text"

		req = req.WritableCustomFieldRequest(netbox.WritableCustomFieldRequest{
			Name:         CustomFieldUuid,
			ContentTypes: []string{"dcim.device"},
			Description:  &description,
			Type:         &fieldType,
		})
		_, resp, err := req.Execute()

		if err != nil {
			data, _ := io.ReadAll(resp.Body)
			log.Println(string(data))
			return err
		}
	}

	return nil
}