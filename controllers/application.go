package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type App struct {
	Storage *services.Storage
}

func (c *App) Post(update models.CpuUpdate) models.OkResponse {
	c.Storage.UpdateCpu(update.Uuid, &update.Cpu)
	return models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}
}
