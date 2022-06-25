package provider

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/spf13/cobra"
)

func NewCmdProviderDrop(o providerOptions) *cobra.Command {
	var providerForce bool
	providerDropHelpMsg := "Drops provider schema from database"
	cmd := &cobra.Command{
		Use:   "drop [provider]",
		Short: providerDropHelpMsg,
		Long:  providerDropHelpMsg,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), util.GetConfigFile(o.Config), false, nil, util.InstanceId)
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

	return cmd
}
