package config

import (
	"fmt"
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
		case "policy":
			cfg, cfgDiags := decodePolicyConfigBlock(block, &p.HCLContext)
			diags = append(diags, cfgDiags...)
			if cfg != nil {
				config.Policies = append(config.Policies, cfg)
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

// ReadModuleConfigProfiles separates the module config from the modules block, where block identifier is the module name.
func ReadModuleConfigProfiles(module string, block hcl.Body) (map[string]hcl.Body, error) {
	if block == nil {
		return nil, nil
	}

	content, _, diags := block.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       module,
				LabelNames: []string{"name"},
			},
		},
	})
	if diags.HasErrors() {
		return nil, diags
	}

	ret := make(map[string]hcl.Body, len(content.Blocks))
	for i := range content.Blocks {
		if _, ok := ret[content.Blocks[i].Labels[0]]; ok {
			return nil, hcl.Diagnostics{
				{
					Severity: hcl.DiagError,
					Summary:  "Duplicate profile name",
					Detail:   fmt.Sprintf("Profile name %q already defined", content.Blocks[i].Labels[0]),
					Subject:  content.Blocks[i].DefRange.Ptr(),
				},
			}
		}

		ret[content.Blocks[i].Labels[0]] = content.Blocks[i].Body
	}
	return ret, nil
}
