package provider

import (
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/spf13/cobra"
)

var (
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

func NewCmdProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "provider [subcommand]",
		Short:   providerShort,
		Long:    providerShort,
		Example: providerExample,
		Version: core.Version,
	}
	cmd.AddCommand(newCmdProviderSync(), newCmdProviderDrop(), newCmdProviderPurge(), newCmdProviderSync(), newCmdProviderDownload())
	return cmd
}
