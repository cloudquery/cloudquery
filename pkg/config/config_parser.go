package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/spf13/viper"
)

// configFileSchema is the schema for the top-level of a config file. We use
// the low-level HCL API for this level so we can easily deal with each
// block type separately with its own decoding logic.
var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "cloudquery",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

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

	hasPolicyBlock := false

	for _, block := range content.Blocks {
		switch block.Type {
		case "cloudquery":
			cliLoggingConfig := logging.GlobalConfig
			cqBlock, cqDiags := decodeCloudQueryBlock(block, &p.HCLContext)
			logging.Reconfigure(*cqBlock.Logger, cliLoggingConfig)
			diags = diags.Extend(cqDiags)
			config.CloudQuery = cqBlock
		case "provider":
			cfg, cfgDiags := decodeProviderBlock(block, &p.HCLContext, existingProviders)
			diags = append(diags, cfgDiags...)
			if cfg != nil {
				config.Providers = append(config.Providers, cfg)
			}
		case "policy":
			hasPolicyBlock = true
		default:
			// Should never happen because the above cases should be exhaustive
			// for all block type names in our schema.
			continue
		}
	}

	if hasPolicyBlock {
		diags = append(diags,
			&hcl.Diagnostic{
				Severity: hcl.DiagWarning,
				Summary:  "Deprecated 'policy' block in config file",
				Detail:   "Specifying 'policy' blocks in 'config.hcl' has been deprecated. See https://docs.cloudquery.io/docs/tutorials/policies/policies-overview for instructions on running policies (either from cloudquery-hub or a local file).",
			},
		)
	}

	return config, diags
}

func decodeCloudQueryBlock(block *hcl.Block, ctx *hcl.EvalContext) (CloudQuery, hcl.Diagnostics) {
	var cq CloudQuery
	// Pre-populate with existing values
	cq.Logger = &logging.GlobalConfig
	var diags hcl.Diagnostics
	diags = diags.Extend(gohcl.DecodeBody(block.Body, ctx, &cq))

	// TODO: decode in a more generic way
	if cq.Connection == nil {
		cq.Connection = &Connection{}
	}

	if err := handleConnectionConfig(cq.Connection); err != nil {
		diags = append(diags, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Invalid DSN configuration",
			Detail:   err.Error(),
			Subject:  &block.DefRange,
		})
	}

	datadir := viper.GetString("data-dir")

	if datadir != "" {
		cq.PluginDirectory = filepath.Join(datadir, "providers")
	}

	if datadir != "" {
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
