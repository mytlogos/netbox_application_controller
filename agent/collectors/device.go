package collectors

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"

	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/shirou/gopsutil/host"
)

type DeviceCollector struct {
}

func (c *DeviceCollector) GetName() string {
	return "Device"
}

func (c *DeviceCollector) GetEndpoint() string {
	return "/api/agent/device"
}

func (c *DeviceCollector) getProduct(hostname string) (string, error) {
	cmd := exec.Command("lshw", "-c", "network", "-json")
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	var lshw []map[string]interface{}
	err = json.Unmarshal(data, &lshw)

	if err != nil {
		return "", err
	}

	// try to get the product if available on an item with hostname as id
	// but do not try to force it
	// only certain devices have a product name
	for _, item := range lshw {
		id := item["id"].(string)

		if strings.EqualFold(id, hostname) {
			product, ok := item["product"]

			if !ok {
				return "", nil
			} else {
				return product.(string), nil
			}
		}
	}

	return "", nil
}

func (c *DeviceCollector) Collect(uuid string) (interface{}, error) {
	is, err := host.Info()

	if err != nil {
		return nil, err
	}

	product, err := c.getProduct(is.Hostname)

	if err != nil {
		log.Println(err)
	}

	return models.DeviceUpdate{
		BasicUpdate: models.BasicUpdate{
			Uuid: uuid,
		},
		Device: models.Device{
			Name:            is.Hostname,
			HostId:          is.HostID,
			Arch:            is.KernelArch,
			KernelVersion:   is.KernelVersion,
			Platform:        is.Platform,
			PlatformFamily:  is.PlatformFamily,
			PlatformVersion: is.PlatformVersion,
			OS:              is.OS,
			Product:         product,
			Boottime:        is.BootTime,
			Uptime:          is.Uptime,
		},
	}, nil
}
