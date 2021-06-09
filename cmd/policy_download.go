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
	policyDownloadCmd = &cobra.Command{
		Use:   "download [policy hub path]",
		Short: "Download a policy from the CloudQuery Policy Hub",
		Long: `Examples:
# Download policy from Policy Hub
./cloudquery policy download cq-aws/cis-v1.3.0

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
			return c.DownloadPolicy(ctx, args)
		},
	}
)

func init() {
	policyCmd.AddCommand(policyDownloadCmd)
}
