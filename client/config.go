package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudquery/cloudquery/config"
	"github.com/cloudquery/cloudquery/plugin"
	"github.com/cloudquery/cloudquery/plugin/hub"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// GenerateConfig generates or adds provider configurations templates
// providers - list of providers to generate configuration for
// allowAppend - if the configuration file already exists, will append missing providers in an existing configuration
// force - replace the current configuration with a newly generated configuration based on given providers
func GenerateConfig(configPath string, providers []string, allowAppend bool, force bool) error {

	var cfg *config.Config
	cfg, err := config.Parse(configPath)
	if err == nil && !(allowAppend || force) {
		log.Error().Str("path", configPath).Msg("configuration file already exists. Either delete it, specify other path via --path or use --append/force to append/replace to existing providers")
		return os.ErrExist
	}
	// if configuration is null, or we are forcing a remake of the configuration
	if cfg == nil || force {
		cfg = &config.Config{Providers: make([]config.Provider, 0, len(providers))}
	}

	pluginHub := hub.NewHub(false)
	for _, provider := range providers {
		if cfg.ProviderExists(provider) {
			log.Err(err).Str("provider", provider).Msg("provider already exists in configuration, use --force to replace")
			return errors.New("provider already exists in configuration, use --force to replace")
		}

		providerCfg, err := getProviderConfig(pluginHub, provider)
		if err != nil {
			log.Err(err).Str("provider", provider).Msg("failed to get providers configuration")
			return err
		}
		cfg.Providers = append(cfg.Providers, providerCfg.Providers...)
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, data, 0644)
}

func getProviderConfig(hub *hub.Hub, providerName string) (*config.Config, error) {
	manager, err := plugin.NewManager()
	if err != nil {
		return nil, err
	}
	if err := hub.DownloadPlugin("cloudquery", providerName, "latest", true); err != nil {
		return &config.Config{}, err
	}
	p, err := manager.GetOrCreateProvider(providerName, "latest")
	if err != nil {
		return &config.Config{}, err
	}
	defer func() {
		if err := manager.KillProvider(providerName); err != nil {
			log.Warn().Err(err).Str("provider", providerName).Msg("failed to kill provider")
		}
	}()
	log.Debug().Str("provider", providerName).Msg("Building provider configuration yaml")
	configYaml, err := p.GenConfig()
	if err != nil {
		return &config.Config{}, err
	}
	return config.LoadFromString(fmt.Sprintf("providers:\n%s", configYaml))
}
