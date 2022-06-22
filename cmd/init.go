package cmd

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/internal/file"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/google/uuid"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const initHelpMsg = "Generate initial config.hcl for fetch command"

var (
	initCmd = &cobra.Command{
		Use:   "init [choose one or more providers (aws gcp azure okta ...)]",
		Short: initHelpMsg,
		Long:  initHelpMsg,
		Example: `
  # Downloads aws provider and generates config.hcl for aws provider
  cloudquery init aws

  # Downloads aws,gcp providers and generates one config.hcl with both providers
  cloudquery init aws gcp`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Initialize(cmd.Context(), args)
		},
	}
)

func Initialize(ctx context.Context, providers []string) error {
	fs := afero.NewOsFs()

	configPath := getConfigFile() // by definition, this will get us an existing file if possible

	if info, _ := fs.Stat(configPath); info != nil {
		ui.ColorizedOutput(ui.ColorError, "Error: Config file %s already exists\n", configPath)
		return diag.FromError(fmt.Errorf("config file %q already exists", configPath), diag.USER)
	}

	requiredProviders := make([]*config.RequiredProvider, len(providers))
	for i, p := range providers {
		organization, providerName, provVersion, err := registry.ParseProviderNameWithVersion(p)
		if err != nil {
			return fmt.Errorf("could not parse requested provider: %w", err)
		}
		rp := config.RequiredProvider{
			Name:    providerName,
			Version: provVersion,
		}
		if organization != registry.DefaultOrganization {
			source := fmt.Sprintf("%s/%s", organization, providerName)
			rp.Source = &source
		}
		requiredProviders[i] = &rp
		providers[i] = providerName // overwrite "provider@version" with just "provider"
	}

	mainConfig := config.Config{
		CloudQuery: config.CloudQuery{
			Providers: requiredProviders,
			Connection: &config.Connection{
				Username: "postgres",
				Password: "pass",
				Host:     "localhost",
				Port:     5432,
				Database: "postgres",
				SSLMode:  "disable",
			},
		},
	}
	if diags := config.ValidateCQBlock(&mainConfig.CloudQuery); diags.HasErrors() {
		return diags
	}

	cCfg := mainConfig
	cCfg.CloudQuery.Connection.DSN = "" // Don't connect
	c, err := console.CreateClientFromConfig(ctx, &cCfg, uuid.Nil)
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}

	var b []byte
	if config.IsNameYAML(configPath) {
		b, err = generateYAMLConfig(ctx, c, providers, mainConfig)
	} else {
		b, err = generateHCLConfig(ctx, c, providers, mainConfig)
	}
	if err != nil {
		return err
	}
	_ = afero.WriteFile(fs, configPath, b, 0644)
	ui.ColorizedOutput(ui.ColorSuccess, "configuration generated successfully to %s\n", configPath)
	return nil
}

func generateYAMLConfig(ctx context.Context, c *console.Client, providers []string, mainConfig config.Config) ([]byte, error) {
	cqConfig := struct {
		CloudQuery config.CloudQuery `yaml:"cloudquery" json:"cloudquery"`
	}{
		CloudQuery: mainConfig.CloudQuery,
	}
	b, err := yaml.Marshal(cqConfig)
	if err != nil {
		return nil, diag.WrapError(err)
	}

	var cqConfigRaw = struct {
		CQ yaml.Node `yaml:"cloudquery"`
	}{}
	if err := yaml.Unmarshal(b, &cqConfigRaw); err != nil {
		return nil, diag.WrapError(err)
	}

	provNode := &yaml.Node{
		Kind:        yaml.SequenceNode,
		HeadComment: "provider configurations",
	}

	for _, p := range providers {
		pCfg, diags := core.GetProviderConfiguration(ctx, c.PluginManager, &core.GetProviderConfigOptions{
			Provider: c.ConvertRequiredToRegistry(p),
			Format:   cqproto.ConfigYAML,
		})
		if pCfg != nil && pCfg.Format != cqproto.ConfigYAML {
			diags = diags.Add(diag.FromError(fmt.Errorf("provider %s doesn't support YAML config. Fallback to HCL or upgrade provider", p), diag.USER, diag.WithDetails("Use `cloudquery init <providers> --config config.hcl` to use HCL config format")))
		}
		if diags.HasErrors() {
			return nil, diags
		}

		var yCfg yaml.Node
		if err := yaml.Unmarshal(pCfg.Config, &yCfg); err != nil {
			return nil, diag.WrapError(err)
		}

		provNode.Content = append(provNode.Content, &yaml.Node{
			Kind: yaml.MappingNode,
			Content: append([]*yaml.Node{
				{
					Kind:  yaml.ScalarNode,
					Value: "name",
				},
				{
					Kind:  yaml.ScalarNode,
					Value: p,
				},
			}, yCfg.Content[0].Content...),
		})
	}

	nd := struct {
		Data map[string]*yaml.Node `yaml:",inline"`
	}{
		Data: map[string]*yaml.Node{
			"cloudquery": &cqConfigRaw.CQ,
			"providers":  provNode,
		},
	}

	return yaml.Marshal(&nd)
}

func generateHCLConfig(ctx context.Context, c *console.Client, providers []string, mainConfig config.Config) ([]byte, error) {
	f := hclwrite.NewEmptyFile()
	rootBody := f.Body()
	cqBlock := gohcl.EncodeAsBlock(&mainConfig.CloudQuery, "cloudquery")

	// Remove deprecated "plugin_directory" and "policy_directory"
	cqBlock.Body().RemoveAttribute("plugin_directory")
	cqBlock.Body().RemoveAttribute("policy_directory")

	// Update connection block to remove unwanted keys
	if b := cqBlock.Body().FirstMatchingBlock("connection", nil); b != nil {
		bd := b.Body()
		bd.RemoveAttribute("dsn")
		bd.RemoveAttribute("type")
		bd.RemoveAttribute("extras")
	}

	rootBody.AppendBlock(cqBlock)
	rootBody.AppendNewline()
	rootBody.AppendUnstructuredTokens(hclwrite.Tokens{
		{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte("// All Provider Configurations"),
		},
	})
	rootBody.AppendNewline()
	var buffer bytes.Buffer
	buffer.WriteString("// Configuration AutoGenerated by CloudQuery CLI\n")
	if n, err := buffer.Write(f.Bytes()); n != len(f.Bytes()) || err != nil {
		return nil, err
	}
	for _, p := range providers {
		pCfg, diags := core.GetProviderConfiguration(ctx, c.PluginManager, &core.GetProviderConfigOptions{
			Provider: c.ConvertRequiredToRegistry(p),
			Format:   cqproto.ConfigHCL,
		})
		if pCfg != nil && pCfg.Format != cqproto.ConfigHCL {
			diags = diags.Add(diag.FromError(fmt.Errorf("provider %s doesn't support HCL config. Please upgrade provider", p), diag.USER))
		}
		if diags.HasErrors() {
			return nil, diags
		}
		buffer.Write(pCfg.Config)
		buffer.WriteString("\n")
	}

	return hclwrite.Format(buffer.Bytes()), nil
}

func init() {
	initCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(initCmd)
}

// getConfigFile returns the config filename
// if it ends with ".*", .hcl and .yml extensions are tried in order to find the existing file, if available
func getConfigFile() string {
	configPath := viper.GetString("configPath")
	if !strings.HasSuffix(configPath, ".*") {
		return configPath
	}

	fs := file.NewOsFs()
	noSuffix := strings.TrimSuffix(configPath, ".*")
	for _, tryExt := range []string{".hcl", ".yml"} {
		tryFn := noSuffix + tryExt
		if _, err := fs.Stat(tryFn); err == nil {
			return tryFn
		}
	}

	return noSuffix + ".hcl"
}
