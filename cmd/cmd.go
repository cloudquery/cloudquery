package cmd

import (
	"fmt"
    "time"
	"github.com/spf13/cobra"
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"os"
)

type Options struct {
	dsn    string
	driver string
}

var options Options
// Injected with at build time with -ldflags "-X github.com/cloudquery/cloudquery/cmd.Variable=Value"
var Version = "development"
var Commit = "development"
var Date = time.Now().String()

var rootCmd = &cobra.Command{
	Use:     "cloudquery",
	Short:   "cloudquery CLI",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cloudqueryclient.New(options.driver, options.dsn)
		if err != nil {
			return err
		}
		return client.Run("config.yml")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&options.dsn, "--dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite")
	rootCmd.PersistentFlags().StringVar(&options.driver, "--driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver")
}
