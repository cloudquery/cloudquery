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

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from configured cloud APIs to specified SQL database",
	Long: `Examples:
# Fetch to PostgreSQL
./cloudquery fetch --dsn "host=localhost user=postgres password=pass DB.name=postgres port=5432"

`,
	Version: Version,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.BindPFlag("dsn", cmd.Flags().Lookup("dsn"))
		if err != nil {
			return err
		}
		err = viper.BindPFlag("driver", cmd.Flags().Lookup("driver"))
		if err != nil {
			return err
		}
		err = viper.BindPFlag("configPath", cmd.Flags().Lookup("path"))
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := viper.GetString("configPath")
		ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
		c, err := console.CreateClient(ctx, configPath)
		if err != nil {
			return err
		}
		return c.Fetch(ctx)
	},
}

func init() {
	flags := fetchCmd.Flags()
	flags.String("dsn", "", "database connection string (env: CQ_DSN) (example: 'host=localhost user=postgres password=pass DB.name=postgres port=5432')")
	flags.String("path", "./config.hcl", "path to configuration file. can be generated with 'gen config' command (env: CQ_CONFIG_PATH)")
	//if err := cobra.MarkFlagRequired(flags, "dsn"); err != nil {
	//	log.Fatal(err)
	//}
	rootCmd.AddCommand(fetchCmd)
}
