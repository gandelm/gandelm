package config

type Config struct {
	namespace string
}

func (c *Config) Namespace() string {
	return c.namespace
}

func NewConfig(namespace string) *Config {
	return &Config{namespace: namespace}
}
