package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

type validateOptions struct {
}

func NewCmdPolicyValidate(parentOptions policyOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: policyValidateHelpMsg,
		Long:  policyValidateHelpMsg,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), util.GetConfigFile(parentOptions.Config), false, nil, util.InstanceId)
			if err != nil {
				return err
			}
			diags := c.ValidatePolicy(cmd.Context(), args[0])
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_validate"})
			return fmt.Errorf("policy validate has one or more errors, check logs")
		},
	}
	return cmd
}
