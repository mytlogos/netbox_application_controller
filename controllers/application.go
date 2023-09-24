package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type App struct {
	Storage *services.Storage
}

func (c *App) Post(update models.AppUpdate) models.OkResponse {
	c.Storage.UpdateApp(update.Uuid, &update.App)
	return models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}
}
