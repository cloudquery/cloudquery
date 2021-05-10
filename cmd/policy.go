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
	policyCmd = &cobra.Command{
		Use:   "policy",
		Short: "Execute a policy on CloudQuery fetched data",
		Long: `Examples:
# Fetch to PostgreSQL
./cloudquery policy --path=<PATH_TO_POLICY_FILE> --output=<PATH_TO_OUTPUT_POLICY_RESULT>

`,
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			configPath := viper.GetString("configPath")
			ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
			c, err := console.CreateClient(ctx, configPath)
			if err != nil {
				return err
			}
			defer c.Client().Close()
			return c.ExecutePolicy(ctx, policyPath, policyOutput)
		},
	}
	policyPath   string
	policyOutput string
)

func init() {
	flags := policyCmd.Flags()
	flags.StringVar(&policyPath, "path", "./policy.yml", "path to a policy file. can be generated with 'gen policy' command (env CQ_POLICY_PATH)")
	flags.StringVar(&policyOutput, "output", "", "output path to store results as json file (env CQ_OUTPUT)")
	rootCmd.AddCommand(policyCmd)
}
