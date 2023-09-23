package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type Ram struct {
	Storage *services.Storage
}

func (c *Ram) Post(update models.RamUpdate) (*models.OkResponse, error) {
	_, err := c.Storage.UpdateRam(update.Uuid, &update.Ram)

	if err != nil {
		return nil, err
	}

	return &models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}, nil
}
