package manifests

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PluginProperties struct {
	Source      bool `json:"source"`
	Destination bool `json:"destination"`
}

type PluginTarget struct {
	Name string `json:"name"`
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

type Manifest struct {
	Version    int              `yaml:"version"`
	Kind       string           `yaml:"kind"`
	Targets    []PluginTarget   `yaml:"targets"`
	Properties PluginProperties `yaml:"properties"`
}

func Read(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := &Manifest{}
	err = yaml.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}
	return m, err
}
