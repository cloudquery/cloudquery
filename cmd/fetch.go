package cmd

import (
	"github.com/cloudquery/cloudquery/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Fetch data from configured cloud APIs to specified SQL database",
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
		err = viper.BindPFlag("config_path", cmd.Flags().Lookup("path"))
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		driver := viper.GetString("driver")
		dsn := viper.GetString("dsn")
		configPath := viper.GetString("config_path")

		client, err := client.New(driver, dsn)
		if err != nil {
			return err
		}
		return client.Run(configPath)

	},
}

func init() {
	fetchCmd.Flags().String( "dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite (env: CQ_DSN)")
	fetchCmd.Flags().String("driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver/neo4j (env: CQ_DRIVER)")
	fetchCmd.Flags().String("path", "./config.yml", "path to configuration file. can be generated with 'gen config' command (env: CQ_CONFIG_PATH)")
	fetchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	// This is for debug purposes to run plugin as is when developing a new provider or trying to debug it.
	// This eliminates the gRPC communication and provide easier way to debug providers.
	fetchCmd.Flags().Bool("runself", false, "run provider without gRPC communication (for debug purposes)")

	rootCmd.AddCommand(fetchCmd)
}
