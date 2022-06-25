package provider

import (
	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/spf13/cobra"
)

const (
	providerShort   = "Top-level command to interact with providers."
	providerExample = `
  # Sync (Upgrade or Downgrade) all providers specified in config.hcl This will also create the schema.
  cloudquery provider sync 
  # Sync one or more providers
  cloudquery provider sync aws, gcp
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  cloudquery provider drop aws
`
)

type providerOptions struct {
	Config string
}

func NewCmdProvider() *cobra.Command {
	o := providerOptions{}
	cmd := &cobra.Command{
		Use:     "provider [subcommand]",
		Short:   providerShort,
		Long:    providerShort,
		Example: providerExample,
		Version: core.Version,
	}
	cmd.PersistentFlags().StringVar(&o.Config, "config", "./config.*", util.ConfigHelp)
	cmd.AddCommand(NewCmdProviderSync(o), NewCmdProviderDrop(o), NewCmdProviderPurge(o))
	return cmd
}

// func init() {

// 	providerDropCmd.Flags().BoolVar(&providerForce, "force", false, "Really drop tables for the provider")
// 	providerCmd.AddCommand(providerDownloadCmd, providerSyncCmd, providerDropCmd, providerRemoveStaleCmd)
// 	providerCmd.SetUsageTemplate(usageTemplateWithFlags)
// 	rootCmd.AddCommand(providerCmd)
// }
