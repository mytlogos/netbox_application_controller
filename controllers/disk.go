package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type Disk struct {
	Storage *services.Storage
}

func (c *Disk) Post(update models.DiskUpdate) (*models.OkResponse, error) {
	_, err := c.Storage.UpdateDisk(update.Uuid, &update.Disk)

	if err != nil {
		return nil, err
	}
	return &models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}, nil
}
