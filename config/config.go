package config

import "github.com/gandelm/gandelm/internal/container/config"

type Env struct {
	Github struct {
		URL   string `env:"GITHUB_URL" envDefault:"https://github.com/gandelm/gandelm"`
		Token string `env:"GITHUB_TOKEN" envDefault:""`
	}

	Kubernetes struct {
		Namespace string `env:"KUBERNETES_NAMESPACE" envDefault:"default"`
	}
}

func (e *Env) BuildConfig() *config.Config {
	return config.NewConfig(e.Kubernetes.Namespace)
}
