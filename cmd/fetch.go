package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/errors"

	"github.com/cloudquery/cloudquery/internal/analytics"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
)

type FetchEvent struct {
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch resources from configured providers",
	Long: `Fetch resources from configured providers

  This requires a config.hcl file which can be generated by "cloudquery init"
	`,
	Example: `  # Fetch configured providers to PostgreSQL as configured in config.hcl
  cloudquery fetch`,
	Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
		result, diags := c.Fetch(ctx)
		errors.CaptureDiagnostics(diags, nil)
		for _, p := range result.ProviderFetchSummary {
			analytics.Capture("fetch", c.Providers, p, diags, "fetch_id", result.FetchId)
		}

		if viper.GetBool("fail-on-error") && diags.HasErrors() {
			return fmt.Errorf("provider has one or more errors, check logs")
		}
		return nil
	}),
}

func init() {
	fetchCmd.SetUsageTemplate(usageTemplateWithFlags)
	fetchCmd.PersistentFlags().Bool("fail-on-error", false, "CloudQuery should return a failure error code if provider has any error")
	_ = viper.BindPFlag("fail-on-error", fetchCmd.PersistentFlags().Lookup("fail-on-error"))
	fetchCmd.Flags().Bool("skip-schema-upgrade", false, "skip schema upgrade of provider fetch, disabling this flag might cause issues")
	_ = viper.BindPFlag("skip-schema-upgrade", fetchCmd.Flags().Lookup("skip-schema-upgrade"))
	fetchCmd.Flags().Bool("redact-diags", false, "show redacted diagnostics only")
	_ = viper.BindPFlag("redact-diags", fetchCmd.Flags().Lookup("redact-diags"))
	_ = fetchCmd.Flags().MarkHidden("redact-diags")
	rootCmd.AddCommand(fetchCmd)
}

// filterConfigProviders gets a list of "providerAlias:resource1,resource2" items and updates the given config, removing non-matching providers
// valid usages:
// "aws" or "aws:*" (all resources specified in the config)
// "aws:ec2.instances,s3.buckets" (only ec2.instances and s3.buckets)
func filterConfigProviders(list []string) func(*config.Config) error {
	return func(cfg *config.Config) error {
		if len(list) == 0 || cfg == nil || len(cfg.Providers) == 0 || len(cfg.CloudQuery.Providers) == 0 {
			return nil
		}

		pMap := make(map[string][]string, len(list)) // provider vs resources
		for _, item := range list {
			parts := strings.SplitN(item, ":", 2)
			prov := parts[0]
			if len(parts) == 2 && parts[1] != "*" {
				resources := strings.Split(parts[1], ",")
				pMap[prov] = make([]string, len(resources))
				for i, res := range resources {
					pMap[prov][i] = res
				}
			} else {
				pMap[prov] = nil
			}
		}

		requiredProviders := make(map[string]struct{})
		for i, p := range cfg.Providers {
			var (
				resList []string
				ok      bool
			)

			if p.Alias != "" {
				resList, ok = pMap[p.Alias]
			} else {
				resList, ok = pMap[p.Name]
			}
			if !ok {
				cfg.Providers[i] = nil
				continue
			}

			requiredProviders[p.Name] = struct{}{}

			if len(resList) > 0 {
				cfg.Providers[i].Resources = resList
			}
		}

		// Remove non-required providers and zero unused pointers afterwards
		{
			i := 0
			for _, p := range cfg.CloudQuery.Providers {
				if _, ok := requiredProviders[p.Name]; ok {
					cfg.CloudQuery.Providers[i] = p
					i++
				}
			}
			for j := i; j < len(cfg.CloudQuery.Providers); j++ {
				cfg.CloudQuery.Providers[j] = nil
			}
			cfg.CloudQuery.Providers = cfg.CloudQuery.Providers[:i]
		}
		{
			i := 0
			for _, p := range cfg.Providers {
				if p != nil {
					cfg.Providers[i] = p
					i++
				}
			}
			cfg.Providers = cfg.Providers[:i]
		}

		if len(cfg.CloudQuery.Providers) == 0 || len(cfg.Providers) == 0 {
			return fmt.Errorf("nothing to fetch")
		}

		return nil
	}
}
