package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	pruneShort   = "Prune policy executions from the database which are older than the relative time specified"
	pruneExample = `
# Prune the policy executions which are older than the relative time specified
cloudquery policy prune 24h`
)

func newCmdPolicyPrune() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "prune",
		Short:   pruneShort,
		Long:    pruneShort,
		Example: pruneExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			retentionPeriod := args[0]
			diags := c.PrunePolicyExecutions(cmd.Context(), retentionPeriod)
			errors.CaptureDiagnostics(diags, map[string]string{"command": "policy_prune"})
			if diags.HasErrors() {
				return fmt.Errorf("policy prune has one or more errors, check logs")
			}
			return nil
		},
	}
	return cmd
}
