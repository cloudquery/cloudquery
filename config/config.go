package config

import (
	"encoding/json"
	"fmt"
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

// Default JSON unmarshaler will not populate the Rest field
// and there are no type annotations that make this happen.
// See https://github.com/golang/go/issues/6213
//
// Function assumes config keys are lowercase
func (c *Config) UnmarshalJSON(data []byte) error {
	var v map[string][]map[string]interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	for _, d := range v["providers"] {
		prov := Provider{}
		var ok bool
		numKnownKeys := 0

		if _, ok := d["name"]; ok {
			numKnownKeys++
		}
		if _, ok := d["version"]; ok {
			numKnownKeys++
		}

		rest := make(map[string]interface{}, len(d) - numKnownKeys)

		if prov.Name, ok = d["name"].(string); !ok {
			return fmt.Errorf("Could not parse provider config")
		}
		if prov.Version, ok = d["version"].(string); !ok {
			prov.Version = "latest"
		}

		for key, value := range d {
			if key == "name" || key == "version" { continue }
			rest[key] = value
		}
		prov.Rest = rest
		c.Providers = append(c.Providers, prov)

	}

	return nil
}