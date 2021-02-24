package client

import (
	"bytes"
	"fmt"
	"github.com/cloudquery/cloudquery/plugin"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// Configuration provider header
const headerConfig = `providers:`

type Config struct {
	Providers []ProviderData
}

type ProviderData struct {
	Name    string
	Version string `default:"latest"`
	Rest    map[string]interface{} `yaml:",inline"`
}

func Init(configPath string) error {

	log.Debug().Str("path", configPath).Msg("reading configuration file")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cqProvidersDir := workingDir + "/.cq/providers/"
	err = os.MkdirAll(cqProvidersDir, os.ModePerm)
	if err != nil {
		return err
	}
	for _, provider := range config.Providers {
		if provider.Name == "" {
			return fmt.Errorf("provider must contain key: name")
		}
		log.Debug().Str("path", configPath).Msg("reading configuration file")
		err := downloadProviderPlugin(provider.Name, provider.Version)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateConfig(configPath string, providers []string, force bool) error {
	var s bytes.Buffer
	_, err := s.WriteString(headerConfig)
	if err != nil {
		return err
	}

	for _, provider := range providers {
		err := downloadProviderPlugin(provider, "latest")
		if err != nil {
			return err
		}
		p, err := plugin.GetManager().GetOrCreateProvider(provider, "latest")
		if err != nil {
			return err
		}
		log.Debug().Str("provider", provider).Msg("Building provider configuration yaml")
		configYaml, err := p.GenConfig()
		if err != nil {
			_ = plugin.GetManager().KillProvider(provider)
			return err
		}
		s.WriteString(configYaml)
		if err := plugin.GetManager().KillProvider(provider); err != nil {
			return err
		}
	}
	s.WriteString("\n")

	if _, err := os.Stat(configPath); err == nil && !force {
		log.Error().Str("path", configPath).Msg("configuration file already exists. Either delete it, specify other path via --path or use --force")
		return os.ErrExist
	} else if os.IsNotExist(err) || force {
		return ioutil.WriteFile(configPath, s.Bytes(), 0644)
	} else {
		return err
	}
}
