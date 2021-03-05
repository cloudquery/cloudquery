package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

func Parse(configPath string) (*Config, error) {
	log.Debug().Str("path", configPath).Msg("reading configuration file")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Load Config data from string input
func LoadFromString(data string) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
