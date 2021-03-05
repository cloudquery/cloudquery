package cmd

import (
	"github.com/cloudquery/cloudquery/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var queryCmd = &cobra.Command{
	Use:     "query",
	Short:   "Run queries specified in a policy file",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag("dsn", cmd.Flags().Lookup("dsn"))
		viper.BindPFlag("driver", cmd.Flags().Lookup("driver"))
		driver := viper.GetString("driver")
		dsn := viper.GetString("dsn")
		queryConfigPath := viper.GetString("policy_path")
		queryOutputJsonPath := viper.GetString("output")
		client, err := client.New(driver, dsn, nil)
		if err != nil {
			return err
		}
		return client.RunQuery(queryConfigPath, queryOutputJsonPath)

	},
}

func init() {
	queryCmd.Flags().String("dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite (env CQ_DSN)")
	queryCmd.Flags().String("driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver (env CQ_DRIVER)")
	queryCmd.Flags().String("path", "./policy.yml", "path to a policy file. can be generated with 'gen policy' command (env CQ_POLICY_PATH)")
	queryCmd.Flags().String("output", "", "output path to store results as json file (env CQ_OUTPUT)")
	viper.BindPFlag("policy_path", queryCmd.Flags().Lookup("path"))
	viper.BindPFlag("output", queryCmd.Flags().Lookup("output"))

	rootCmd.AddCommand(queryCmd)
}
