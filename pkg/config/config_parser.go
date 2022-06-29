package config

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed schema.json
var configSchemaYAML []byte

// configSchemaHCL is the HCL schema for the top-level of a config file. We use
// the low-level HCL API for this level so we can easily deal with each
// block type separately with its own decoding logic.
var configSchemaHCL = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "cloudquery",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type: "modules", // deprecated
		},
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

func (p *Parser) LoadConfigFromSource(name string, data []byte) (*Config, diag.Diagnostics) {
	if IsNameYAML(name) {
		newData := os.Expand(string(data), p.getVariableValue)
		return decodeConfigYAML(strings.NewReader(newData))
	}

	body, diags := p.LoadFromSource(name, data)
	if body == nil {
		return nil, diags
	}
	return p.decodeConfigHCL(body, diags)
}

func (p *Parser) LoadConfigFile(path string) (*Config, diag.Diagnostics) {
	contents, diags := p.LoadFile(path)
	if diags.HasErrors() {
		return nil, diags
	}
	return p.LoadConfigFromSource(path, contents)
}

func IsNameYAML(name string) bool {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".json", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func ValidateCQBlock(cq *CloudQuery) diag.Diagnostics {
	var diags diag.Diagnostics

	if cq.Connection == nil {
		cq.Connection = &Connection{}
	}

	if err := handleConnectionConfig(cq.Connection); err != nil {
		diags = diags.Add(diag.FromError(err, diag.USER, diag.WithSummary("Invalid DSN configuration")))
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
			diags = diags.Add(diag.FromError(fmt.Errorf("Provider %s version %s is invalid", cp.Name, cp.Version), diag.USER, diag.WithDetails("Please set to 'latest' version or valid semantic versioning starting with vX.Y.Z")))
		}
	}
	return diags
}

func ProcessValidateProviderBlock(plist []*Provider) (Providers, diag.Diagnostics) {
	var diags diag.Diagnostics
	existingProviders := make(map[string]struct{}, len(plist))

	var ret Providers

	for _, v := range plist {
		if v.Alias != "" {
			if _, ok := existingProviders[v.Alias]; ok {
				diags = diags.Add(diag.FromError(fmt.Errorf("provider with alias %s for provider %s already exists, give it a different alias", v.Alias, v.Name), diag.USER, diag.WithSummary("Duplicate Alias")))
				continue
			}
			existingProviders[v.Alias] = struct{}{}
		} else {
			if _, ok := existingProviders[v.Name]; ok {
				diags = diags.Add(diag.FromError(fmt.Errorf("provider with name %s already exists, use alias in provider configuration block", v.Name), diag.USER, diag.WithSummary("Provider Alias Required")))
				continue
			}
			existingProviders[v.Name] = struct{}{}
			v.Alias = v.Name
		}
		var err error
		v.Configuration, err = yaml.Marshal(v.ConfigKeys["configuration"])
		if err != nil {
			diags = diags.Add(diag.FromError(err, diag.INTERNAL, diag.WithSummary("ConfigKeys marshal failed")))
			continue
		}
		ret = append(ret, v)
	}

	return ret, diags
}

func decodeConfigYAML(r io.Reader) (*Config, diag.Diagnostics) {
	var yc struct {
		CloudQuery CloudQuery  `yaml:"cloudquery" json:"cloudquery"`
		Providers  []*Provider `yaml:"providers" json:"providers"`
	}

	lgc := logging.GlobalConfig
	yc.CloudQuery.Logger = &lgc

	if err := yaml.NewDecoder(r).Decode(&yc); err != nil {
		return nil, diag.FromError(err, diag.USER, diag.WithSummary("Failed to parse yaml"))
	}

	schemaLoader := gojsonschema.NewBytesLoader(configSchemaYAML)
	documentLoader := gojsonschema.NewGoLoader(yc)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, diag.FromError(err, diag.USER, diag.WithSummary("Failed to validate config"))
	}
	if !result.Valid() {
		errs := result.Errors()
		if len(errs) == 0 {
			return nil, diag.FromError(errors.New("Failed to validate config with schema"), diag.USER, diag.WithSummary("Invalid configuration"))
		}
		diags := diag.Diagnostics{}
		for _, e := range errs {
			diags = diags.Add(
				diag.FromError(errors.New(e.String()), diag.USER, diag.WithDetails("%s", e.Description()), diag.WithSummary("Config field %q has error of type %s", e.Field(), e.Type())),
			)
		}
		return nil, diags
	}

	c := &Config{
		CloudQuery: yc.CloudQuery,
		format:     cqproto.ConfigYAML,
	}

	var diags diag.Diagnostics
	c.Providers, diags = ProcessValidateProviderBlock(yc.Providers)
	diags = diags.Add(ValidateCQBlock(&c.CloudQuery))
	if diags.HasErrors() {
		return nil, diags
	}

	return c, diags
}

func (p *Parser) decodeConfigHCL(body hcl.Body, diags diag.Diagnostics) (*Config, diag.Diagnostics) {
	existingProviders := make(map[string]bool)
	config := &Config{
		format: cqproto.ConfigHCL,
	}

	content, contentDiags := body.Content(configSchemaHCL)
	diags = diags.Add(hclToSdkDiags(contentDiags))

	hasPolicyBlock := false

	for _, block := range content.Blocks {
		switch block.Type {
		case "cloudquery":
			cliLoggingConfig := logging.GlobalConfig
			cqBlock, cqDiags := decodeCloudQueryBlock(block, &p.HCLContext)
			diags = diags.Add(hclToSdkDiags(cqDiags))
			diags = diags.Add(ValidateCQBlock(&cqBlock))

			logging.Reconfigure(*cqBlock.Logger, cliLoggingConfig)
			config.CloudQuery = cqBlock
		case "provider":
			cfg, cfgDiags := decodeProviderBlock(block, &p.HCLContext, existingProviders)
			diags = diags.Add(hclToSdkDiags(cfgDiags))
			if cfg != nil {
				config.Providers = append(config.Providers, cfg)
			}
		case "policy":
			hasPolicyBlock = true
		case "modules":
			// deprecated - ignore
			continue
		default:
			// Should never happen because the above cases should be exhaustive
			// for all block type names in our schema.
			continue
		}
	}

	if hasPolicyBlock {
		diags = diags.Add(diag.FromError(errors.New("Deprecated 'policy' block in config file"), diag.USER, diag.WithSeverity(diag.WARNING), diag.WithDetails("Specifying 'policy' blocks in 'config.hcl' has been deprecated. See https://docs.cloudquery.io/docs/tutorials/policies/policies-overview for instructions on running policies (either from cloudquery-hub or a local file).")))
	}

	return config, diags
}

func decodeCloudQueryBlock(block *hcl.Block, ctx *hcl.EvalContext) (CloudQuery, hcl.Diagnostics) {
	var cq CloudQuery
	// Pre-populate with existing values
	cq.Logger = &logging.GlobalConfig
	return cq, gohcl.DecodeBody(block.Body, ctx, &cq)
}
