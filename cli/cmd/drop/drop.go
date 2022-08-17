package drop

import "github.com/spf13/cobra"

var (
	providerShort   = "Top-level command to interact with providers."
	providerExample = `
  # Sync (Upgrade or Downgrade) all providers specified in cloudquery.yml This will also create the schema.
  cloudquery plugin sync 
  # Sync plugin
  cloudquery plugin sync aws
  # Drop provider schema, running fetch again will recreate all tables unless --skip-build-tables is specified
  cloudquery provider drop aws
`
)

func NewCmdDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "drop <filename>",
		Short:   providerShort,
		Long:    providerShort,
		Example: providerExample,
		RunE:    drop,
	}
	return cmd
}

func drop(cmd *cobra.Command, args []string) error {
	return nil
}
