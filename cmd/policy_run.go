package cmd

import (
	"context"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/signalcontext"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	policyRunCmd = &cobra.Command{
		Use:   "run [policy hub path]",
		Short: "Runs a policy",
		Long: `Examples:
# Download policy from Policy Hub
./cloudquery policy download cq-aws

# Run policy
./cloudquery policy run cq-aws

See https://hub.cloudquery.io for additional policies.

`,
		Version: Version,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			subpath := cmd.PersistentFlags().Lookup("subpolicy")
			return c.RunPolicy(ctx, args, subpath)
		},
	}
)

func init() {
	policyRunCmd.PersistentFlags().String("subpolicy", "", "Forces the policy run command to only execute this sub policy/view/query")
	policyCmd.AddCommand(policyRunCmd)
}
