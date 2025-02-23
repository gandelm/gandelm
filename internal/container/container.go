package container

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Containerer interface {
	Kubernetes() client.Client
	Github() Github
	Config() Config
}

type Github interface {
	ListBranch(ctx context.Context) ([]string, error)
	HookAction(ctx context.Context, target string)
}

type Config interface {
	Namespace() string
}

type Container struct {
	kubernetes client.Client
	github     Github
	config     Config
}

func (c *Container) Kubernetes() client.Client {
	return c.kubernetes
}

func (c *Container) Github() Github {
	return c.github
}

func (c *Container) Config() Config {
	return c.config
}

func NewContainer(kubernetes client.Client, config Config, github Github) Containerer {
	return &Container{
		kubernetes: kubernetes,
		config:     config,
		github:     github,
	}
}
