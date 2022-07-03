package policy

import (
	"github.com/cloudquery/cloudquery/cmd/utils"
	"github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/spf13/cobra"
)

const (
	snapshotShort = `Take database snapshot of all tables included in a CloudQuery policy`
)

func newCmdPolicySnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: snapshotShort,
		Long:  snapshotShort,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := console.CreateClient(cmd.Context(), utils.GetConfigFile(), false, nil, utils.InstanceId)
			if err != nil {
				return err
			}
			err = c.SnapshotPolicy(cmd.Context(), args[0], args[1])
			errors.CaptureError(err, map[string]string{"command": "policy_snapshot"})
			return err
		},
	}
	return cmd
}
