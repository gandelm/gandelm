package container

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Containerer interface {
	Kubernetes() client.Client
	Github() Github
	Config() Config
	Command() Commander
}

type Github interface {
	ListBranch(ctx context.Context) ([]string, error)
	HookAction(ctx context.Context, eventType string, input map[string]string) (any, error)
}

type Config interface {
	Namespace() string
}

type Commander interface {
	HelmInstall(namespace, releaseName string) error
	HelmUpgrade(namespace, releaseName string) error
	HelmUnInstall(namespace, releaseName string) error
	GitClone(org, repo string) error
}

type Container struct {
	kubernetes client.Client
	github     Github
	config     Config
	commander  Commander
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

func (c *Container) Command() Commander {
	return c.commander
}

func NewContainer(
	kubernetes client.Client,
	config Config,
	github Github,
	command Commander,
) Containerer {
	return &Container{
		kubernetes: kubernetes,
		config:     config,
		github:     github,
		commander:  command,
	}
}
