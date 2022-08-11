package fetch

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/cmd/utils"
	"github.com/cloudquery/cloudquery/cli/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/pkg/config"
	"github.com/cloudquery/cloudquery/cli/pkg/core"
	"github.com/cloudquery/cloudquery/cli/pkg/errors"
	"github.com/cloudquery/cloudquery/cli/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/cli/pkg/ui/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	fetchShort = "Fetch resources from configured providers"
	fetchLong  = `Fetch resources from configured providers
	
	This requires a cloudquery.yml file which can be generated by "cloudquery init"
	`
	fetchExample = `  # Fetch configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery fetch`
)

// sendProviderTelemetryEvents sends all collected telemetry events from the provider fetch response.
// It will panic if fr argument is nil.
func sendProviderTelemetryEvents(providers registry.Providers, fr *core.FetchResponse) {
	for _, e := range fr.TelemetryEvents {
		analytics.Capture(e.Category, providers, e, nil)
	}
	for _, pfs := range fr.ProviderFetchSummary {
		for _, rfs := range pfs.FetchedResources {
			for _, e := range rfs.TelemetryEvents {
				analytics.Capture(e.Category, providers, e, nil)
			}
		}
	}
}

func NewCmdFetch() *cobra.Command {
	fetchCmd := &cobra.Command{
		Use:     "fetch",
		Short:   fetchShort,
		Long:    fetchLong,
		Example: fetchExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgMutator := filterConfigProviders(args)
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, cfgMutator, utils.InstanceId)
			if err != nil {
				return err
			}
			result, diags := c.Fetch(cmd.Context())
			errors.CaptureDiagnostics(diags, map[string]string{"command": "fetch"})
			if result != nil {
				for _, p := range result.ProviderFetchSummary {
					analytics.Capture("fetch", c.Providers, p, diags, "fetch_id", result.FetchId)
				}
				sendProviderTelemetryEvents(c.Providers, result)
			}
			if diags.HasErrors() {
				return fmt.Errorf("provider has one or more errors, check logs")
			}
			return nil
		},
	}
	fetchCmd.Flags().Bool("skip-schema-upgrade", false, "skip schema upgrade of provider fetch, disabling this flag might cause issues")
	_ = viper.BindPFlag("skip-schema-upgrade", fetchCmd.Flags().Lookup("skip-schema-upgrade"))
	fetchCmd.Flags().Bool("redact-diags", false, "show redacted diagnostics only")
	_ = viper.BindPFlag("redact-diags", fetchCmd.Flags().Lookup("redact-diags"))
	_ = fetchCmd.Flags().MarkHidden("redact-diags")
	return fetchCmd
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
				copy(pMap[prov], resources)
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
