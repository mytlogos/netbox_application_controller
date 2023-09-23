package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type Device struct {
	Storage *services.Storage
}

func (c *Device) Post(update models.DeviceUpdate) (*models.OkResponse, error) {
	update.Uptime = 0
	_, err := c.Storage.UpdateDevice(update.Uuid, &update.Device)

	if err != nil {
		return nil, err
	}
	return &models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}, nil
}
