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
	Long: `Examples: 
# Fetch to SQLite
./cloudquery fetch

# Fetch to MySQL
./cloudquery fetch --driver mysql --dsn "root:pass@tcp(127.0.0.1:3306)/dbname"

# Fetch to PostgreSQL
./cloudquery fetch --driver postgresql --dsn "host=localhost user=postgres password=pass DB.name=postgres port=5432"

# Fetch to SQL Server
./cloudquery fetch --driver sqlserver --dsn "sqlserver://sa:yourStrong(!)Password@localhost:1433?database=cloudquery"
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
	rootCmd.AddCommand(fetchCmd)
}
