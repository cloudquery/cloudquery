package config

import (
	"github.com/creasty/defaults"
)


const configHeader = "providers:"

type Config struct {
	Providers []Provider
}

type Provider struct {
	Name    string
	Version string `default:"latest"`
	Rest    map[string]interface{} `yaml:",inline"`
}

func (p *Provider) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = defaults.Set(p)

	type plain Provider
	if err := unmarshal((*plain)(p)); err != nil {
		return err
	}
	return nil
}

// Check if provider already exists in configuration file
func (c *Config) ProviderExists(providerName string) bool {
	for _, p := range c.Providers {
		if p.Name == providerName {
			return true
		}
	}
	return false
}