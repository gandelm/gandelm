package container

import (
	"github.com/gandelm/gandelm/internal/container/config"
	"github.com/gandelm/gandelm/internal/container/github"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Containerer interface {
	Kubernetes() client.Client
	Github() github.GithubClient
	Config() config.Config
}

type Container struct {
	kubernetes client.Client
	github     github.GithubClient
	config     config.Config
}

func (c *Container) Kubernetes() client.Client {
	return c.kubernetes
}

func (c *Container) Github() github.GithubClient {
	return c.github
}

func (c *Container) Config() config.Config {
	return c.config
}

func NewContainer(kubernetes client.Client) Containerer {
	return &Container{
		kubernetes: kubernetes,
		config:     config.NewConfig(),
		// github:     github,
	}
}
