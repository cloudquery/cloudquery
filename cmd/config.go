package cmd

import (
	"fmt"
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var headerConfig = `providers:`

var configPath = "./config.yml"
var force = false

var configCmd = &cobra.Command{
	Use:       fmt.Sprintf("config [choose one or more providers (aws,gcp,azure,okta,...)]"),
	Short:     "Generate initial config.yml for fetch command",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cloudqueryclient.GenConfig(args)
		if err != nil {
			return err
		}

		if _, err := os.Stat(configPath); err == nil && !force {
			return fmt.Errorf("file %s already exists. Either delete it, specify other path via --path or use --force", configPath)
		} else if os.IsNotExist(err) || force {
			return ioutil.WriteFile(configPath, []byte(config), 0644)
		} else {
			return err
		}
	},
}

func init() {
	genCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&configPath, "path", configPath, "path to output generated config file")
	configCmd.Flags().BoolVar(&force, "force", force, "override output")
}
