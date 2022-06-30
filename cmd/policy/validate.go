package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	validateShort = "Validate policy for any issues and diagnostics"
)

func newCmdPolicyValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: validateShort,
		Long:  validateShort,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			diags := c.ValidatePolicy(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_validate"})
			if diags.HasErrors() {
				return fmt.Errorf("policy validate has one or more errors, check logs")
			}
			return nil
		},
	}
	return cmd
}
