package config

import (
	"encoding/json"
	"fmt"

	"github.com/creasty/defaults"
)

type Config struct {
	Providers []Provider
}

type Provider struct {
	Name    string
	Version string                 `default:"latest"`
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

// Default JSON unmarshaler will not populate the Rest field
// and there are no type annotations that make this happen.
// See https://github.com/golang/go/issues/6213
//
// Function assumes config keys are lowercase
func (c *Config) UnmarshalJSON(data []byte) error {
	var conf map[string][]map[string]interface{}

	if err := json.Unmarshal(data, &conf); err != nil {
		return err
	}

	for _, provMap := range conf["providers"] {
		prov := Provider{}
		var ok bool
		numKnownKeys := 0

		if _, ok := provMap["name"]; ok {
			numKnownKeys++
		}
		if _, ok := provMap["version"]; ok {
			numKnownKeys++
		}

		rest := make(map[string]interface{}, len(provMap)-numKnownKeys)

		if prov.Name, ok = provMap["name"].(string); !ok {
			return fmt.Errorf("Could not parse provider config")
		}
		if prov.Version, ok = provMap["version"].(string); !ok {
			prov.Version = "latest"
		}

		for key, value := range provMap {
			if key == "name" || key == "version" {
				continue
			}
			rest[key] = value
		}
		prov.Rest = rest
		c.Providers = append(c.Providers, prov)

	}

	return nil
}
