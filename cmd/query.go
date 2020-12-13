package cmd

import (
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
)

var queryDSN string
var queryDriver string
var queryConfigPath string

var queryCmd = &cobra.Command{
	Use:     "query",
	Short:   "Run queries specified in a policy file",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cloudqueryclient.New(driver, dsn, verbose)
		if err != nil {
			return err
		}
		return client.RunQuery(queryConfigPath)

	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVar(&queryDSN, "dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite")
	queryCmd.Flags().StringVar(&queryDriver, "driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver")
	queryCmd.Flags().StringVar(&queryConfigPath, "path", "./policy.yml", "path to a policy file. can be generated with 'gen policy' command")
}
