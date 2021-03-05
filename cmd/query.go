package cmd

import (
	"github.com/cloudquery/cloudquery/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var queryCmd = &cobra.Command{
	Use:     "query",
	Short:   "Run queries specified in a policy file",
	Version: Version,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlag("dsn", cmd.Flags().Lookup("dsn")); err != nil {
			return err
		}
		if err := viper.BindPFlag("driver", cmd.Flags().Lookup("driver")); err != nil {
			return err
		}
		if err := viper.BindPFlag("policy_path", cmd.Flags().Lookup("path")); err != nil {
			return err
		}
		if err := viper.BindPFlag("output", cmd.Flags().Lookup("output")); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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
	flags := queryCmd.Flags()
	flags.String("dsn", "", "database connection string (env CQ_DSN) (example: 'host=localhost user=postgres password=pass DB.name=postgres port=5432')")
	flags.String("driver", "postgresql", "database driver postgresql/neo4j (env CQ_DRIVER)")
	flags.String("path", "./policy.yml", "path to a policy file. can be generated with 'gen policy' command (env CQ_POLICY_PATH)")
	flags.String("output", "", "output path to store results as json file (env CQ_OUTPUT)")

	if err := cobra.MarkFlagRequired(flags, "dsn"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(queryCmd)
}
