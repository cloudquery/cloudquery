package cmd

import (
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
)


var	dsn    string
var	driver string
var verbose bool
var fetchConfigPath string

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Fetch data from configured cloud APIs to specified SQL database",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cloudqueryclient.New(driver, dsn, verbose)
		if err != nil {
			return err
		}
		return client.Run(fetchConfigPath)

	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&dsn, "dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite")
	fetchCmd.Flags().StringVar(&driver, "driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver")
	fetchCmd.Flags().StringVar(&fetchConfigPath, "path", "./config.yml", "path to configuration file. can be generated with 'gen config' command")
	fetchCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}