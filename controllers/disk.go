package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type Disk struct {
	Storage *services.Storage
}

func (c *Disk) Post(update models.DiskUpdate) models.OkResponse {
	c.Storage.UpdateDisk(update.Uuid, &update.Disk)
	return models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}
}
