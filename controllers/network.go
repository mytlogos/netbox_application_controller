package controllers

import (
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type Network struct {
	Storage *services.Storage
}

func (c *Network) Post(update models.NetworkUpdate) (*models.OkResponse, error) {
	_, err := c.Storage.UpdateNetwork(update.Uuid, &update.Network)

	if err != nil {
		return nil, err
	}
	return &models.OkResponse{
		Status: "Ok",
		Ok:     true,
	}, nil
}
