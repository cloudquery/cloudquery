package config

import "github.com/hashicorp/hcl/v2/hclsimple"

func Parse(configPath string) (*Config, error) {
	var config Config
	if err := hclsimple.DecodeFile(configPath, nil, &config); err != nil {
		return nil, err
	}
	return &config, nil
}