package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type CPU struct {
	Storage *services.Storage
}

func (c *CPU) Post(update models.CpuUpdate) models.OkResponse {
	c.Storage.UpdateCpu(update.Uuid, &update.Cpu)
	return models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}
}
