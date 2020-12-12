package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
)

type Options struct {
	dsn    string
	driver string
	verbose bool
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
		client, err := cloudqueryclient.New(options.driver, options.dsn, options.verbose)
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
	rootCmd.PersistentFlags().StringVar(&options.dsn, "dsn", "./cloudquery.db", "database connection string or filepath if driver is sqlite")
	rootCmd.PersistentFlags().StringVar(&options.driver, "driver", "sqlite", "database driver sqlite/postgresql/mysql/sqlserver")
	rootCmd.PersistentFlags().BoolVarP(&options.verbose, "verbose", "v", false, "verbose output")
}
