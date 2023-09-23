package controllers

import (
	"time"

	"github.com/google/uuid"
	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/services"
)

type User struct {
	Storage *services.Storage
}

type DeployAgent struct {
	Host string
}

type AgentConfig struct {
	Server string
	Uuid   string
}

func (c *User) PostAgentDeploy(value DeployAgent) (*models.Host, error) {
	uuid := uuid.New()
	// TODO: generate agent secret
	// TODO: deploy agent via ansible
	err := c.Storage.AddAgent(uuid.String(), models.Agent{
		DeploymentDate: time.Now(),
		LastUpdate:     nil,
		State:          "Deploying",
		Version:        "1.0.0",
	})
	if err != nil {
		return nil, err
	}

	return c.Storage.Get(uuid.String())
}

func (c *User) GetAgentAll() ([]*models.Host, error) {
	return c.Storage.GetAll()
}

func (c *User) GetAgentConfigBy(uuid string) (*AgentConfig, error) {
	_, err := c.Storage.Get(uuid)
	if err != nil {
		return nil, err
	}
	return &AgentConfig{
		Server: "http://192.168.1.13:8080",
		Uuid:   uuid,
	}, nil
}
