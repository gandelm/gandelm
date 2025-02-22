package config

type Config interface {
	Namespace() string
}

type ConfigImpl struct {
	namespace string
}

func (c *ConfigImpl) Namespace() string {
	return c.namespace
}

func NewConfig() Config {
	return &ConfigImpl{
		namespace: "default",
	}
}
