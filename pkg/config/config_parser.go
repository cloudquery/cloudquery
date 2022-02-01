package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/creasty/defaults"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/policy"

	"github.com/hashicorp/hcl/v2"
)

func (p *Parser) LoadConfigFromSource(name string, data []byte) (*Config, hcl.Diagnostics) {
	if strings.HasSuffix(name, ".json") {
		// we dropped support for json so error out with an explainable message
		return nil, hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  `json is not supported please use hcl format`,
			Detail:   `json is not supported please use hcl format`,
		}}
	}
	body, diags := p.LoadFromSource(name, data)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfig(body, diags)
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
			cqBlock, cqDiags := decodeCloudQueryBlock(block, &p.HCLContext)
			diags = diags.Extend(cqDiags)
			config.CloudQuery = cqBlock
		case "provider":
			cfg, cfgDiags := decodeProviderBlock(block, &p.HCLContext, existingProviders)
			diags = append(diags, cfgDiags...)
			if cfg != nil {
				config.Providers = append(config.Providers, cfg)
			}
		case "policy":
			cfg, cfgDiags := policy.DecodePolicyBlock(block, &p.HCLContext)
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

func decodeCloudQueryBlock(block *hcl.Block, ctx *hcl.EvalContext) (CloudQuery, hcl.Diagnostics) {
	var cq CloudQuery
	var diags hcl.Diagnostics
	diags = diags.Extend(gohcl.DecodeBody(block.Body, ctx, &cq))

	// TODO: decode in a more generic way
	if cq.Connection == nil {
		cq.Connection = &Connection{
			DSN: "",
		}
	}
	if cq.History != nil {
		if err := defaults.Set(cq.History); err != nil {
			diags = append(diags, &hcl.Diagnostic{Severity: hcl.DiagError, Summary: "failed to set defaults in history"})
		}
	}
	if dsn := viper.GetString("dsn"); dsn != "" {
		cq.Connection.DSN = dsn
	}

	datadir := viper.GetString("data-dir")

	if dir := viper.GetString("plugin-dir"); dir != "" {
		if dir == "." {
			if dir, err := os.Getwd(); err == nil {
				cq.PluginDirectory = dir
			}
		} else {
			cq.PluginDirectory = dir
		}
	} else if datadir != "" {
		cq.PluginDirectory = filepath.Join(datadir, "providers")
	}

	if dir := viper.GetString("policy-dir"); dir != "" {
		if dir == "." {
			if dir, err := os.Getwd(); err == nil {
				cq.PolicyDirectory = dir
			}
		} else {
			cq.PolicyDirectory = dir
		}
	} else if datadir != "" {
		cq.PolicyDirectory = filepath.Join(datadir, "policies")
	}

	// validate provider versions
	for _, cp := range cq.Providers {
		if cp.Version != "latest" && !strings.HasPrefix(cp.Version, "v") {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf("Provider %s version %s is invalid", cp.Name, cp.Version),
				Detail:   "Please set to 'latest' version or valid semantic versioning starting with vX.Y.Z",
				Subject:  &block.DefRange,
			})
		}
	}
	return cq, diags
}
