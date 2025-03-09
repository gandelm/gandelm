package config

import (
	"github.com/gandelm/gandelm/internal/container"
	"github.com/gandelm/gandelm/internal/container/command"
	"github.com/gandelm/gandelm/internal/container/config"
	"github.com/gandelm/gandelm/internal/container/github"
)

type Env struct {
	Github struct {
		Enabled      bool   `env:"GITHUB_ENABLED" envDefault:"false"` // TODO: implement this
		Token        string `env:"GITHUB_TOKEN" envDefault:""`        // TODO: implement this
		Organization string `env:"GITHUB_ORGANIZATION" envDefault:"gandelm"`
		Repository   string `env:"GITHUB_REPOSITORY" envDefault:"gandelm"`
	}

	Kubernetes struct {
		Namespace string `env:"KUBERNETES_NAMESPACE" envDefault:"default"`
	}

	Helm struct {
		ManifestPath string `env:"HELM_MANIFEST_PATH" envDefault:"./manifests"`
		PackageName  string `env:"HELM_PACKAGE_NAME" envDefault:""`
	}
}

func (e *Env) BuildConfig() *config.Config {
	return config.NewConfig(e.Kubernetes.Namespace)
}

func (e *Env) BuildGithub() container.Github {
	return github.NewGithub(e.Github.Organization, e.Github.Repository)
}

func (e *Env) BuildCommander() container.Commander {
	return command.NewCommand(e.Helm.ManifestPath, e.Helm.PackageName)
}
