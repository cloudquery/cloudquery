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
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := viper.GetString("configPath")
		ctx, _ := signalcontext.WithInterrupt(context.Background(), logging.NewZHcLog(&log.Logger, ""))
		c, err := console.CreateClient(ctx, configPath)
		if err != nil {
			return err
		}
		defer c.Client().Close()
		return c.Fetch(ctx)
	},
}
