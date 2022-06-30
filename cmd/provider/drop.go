package provider

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/cobra"
)

var (
	providerForce bool
	dropShort     = "Drops provider schema from database"
)

func newCmdProviderDrop() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "drop [provider]",
		Short: dropShort,
		Long:  dropShort,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			if !providerForce {
				ui.ColorizedOutput(ui.ColorWarning, "WARNING! This will drop all tables for the given provider. If you wish to continue, use the --force flag.\n")
				return diag.FromError(fmt.Errorf("if you wish to continue, use the --force flag"), diag.USER)
			}
			diags := c.DropProvider(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "provider_drop"})
			if diags.HasErrors() {
				return fmt.Errorf("failed to drop provider %s", args[0])
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&providerForce, "force", false, "Really drop tables for the provider")
	return cmd
}
