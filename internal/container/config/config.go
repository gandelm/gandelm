package config

import (
	"github.com/gandelm/gandelm/internal/container"
)

type Config struct {
	namespace string
}

func (c *Config) Namespace() string {
	return c.namespace
}

func NewConfig() container.Config {
	return &Config{
		namespace: "default",
	}
}
