package config

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/spf13/viper"
)

func (p *Parser) loadConfigFromSource(name string, data []byte, source SourceType) (*Config, hcl.Diagnostics) {
	body, diags := p.loadFromSource(name, data, source)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfig(body, diags)
}

func (p *Parser) LoadConfigFromSource(name string, data []byte) (*Config, hcl.Diagnostics) {
	return p.loadConfigFromSource(name, data, SourceHCL)
}

func (p *Parser) LoadConfigFromJson(name string, data []byte) (*Config, hcl.Diagnostics) {
	return p.loadConfigFromSource(name, data, SourceJSON)
}

func (p *Parser) LoadConfigFile(path string) (*Config, hcl.Diagnostics) {

	body, diags := p.LoadHCLFile(path)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfig(body, diags)
}

func (p *Parser) decodeConfig(body hcl.Body, diags hcl.Diagnostics) (*Config, hcl.Diagnostics) {

	existingProviders := make(map[string]bool)
	config := &Config{}

	content, contentDiags := body.Content(configFileSchema)
	diags = append(diags, contentDiags...)

	for _, block := range content.Blocks {
		switch block.Type {
		case "cloudquery":
			contentDiags = gohcl.DecodeBody(block.Body, &p.HCLContext, &config.CloudQuery)
			diags = append(diags, contentDiags...)
			// TODO: decode in a more generic way

			if config.CloudQuery.Connection == nil {
				config.CloudQuery.Connection = &Connection{
					DSN: "",
				}
			}

			if dsn := viper.GetString("dsn"); dsn != "" {
				config.CloudQuery.Connection.DSN = dsn
			}
			if dir := viper.GetString("plugin-dir"); dir != "." {
				if dir == "." {
					if dir, err := os.Getwd(); err == nil {
						config.CloudQuery.PluginDirectory = dir
					}
				} else {
					config.CloudQuery.PluginDirectory = dir
				}
			}
			if dir := viper.GetString("policy-dir"); dir != "" {
				if dir == "." {
					if dir, err := os.Getwd(); err != nil {
						config.CloudQuery.PolicyDirectory = dir
					}
				} else {
					config.CloudQuery.PolicyDirectory = dir
				}
			}
		case "provider":
			cfg, cfgDiags := decodeProviderBlock(block, &p.HCLContext, existingProviders)
			diags = append(diags, cfgDiags...)
			if cfg != nil {
				config.Providers = append(config.Providers, cfg)
			}
		case "modules":
			// Module manager will process this for us
			config.Modules = block.Body
		default:
			// Should never happen because the above cases should be exhaustive
			// for all block type names in our schema.
			continue
		}
	}
	return config, diags
}
