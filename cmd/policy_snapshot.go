package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const policySnapshotHelpMsg = `Take database snapshot of all tables included in a CloudQuery policy`

var (
	snapshotPolicyCmd = &cobra.Command{
		Use:   "snapshot",
		Short: policySnapshotHelpMsg,
		Long:  policySnapshotHelpMsg,
		Args:  cobra.ExactArgs(2),
		Run: handleCommand(func(ctx context.Context, c *console.Client, cmd *cobra.Command, args []string) error {
			return c.SnapshotPolicy(ctx, args[0], args[1])
		}),
	}
)

func init() {
	snapshotPolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
	policyCmd.AddCommand(snapshotPolicyCmd)
}
