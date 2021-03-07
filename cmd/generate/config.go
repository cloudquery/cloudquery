package generate

import (
	"github.com/cloudquery/cloudquery/client"
	"github.com/spf13/cobra"
)

var (
	configPath      = "./config.yml"
	appendProviders = false
	force           = false
	configCmd       = &cobra.Command{
		Use:   "config [choose one or more providers (aws,gcp,azure,okta,...)]",
		Short: "Generate initial config.yml for fetch command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return client.GenerateConfig(configPath, args, appendProviders, force)
		},
	}
)

func init() {
	Cmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&configPath, "path", configPath, "path to output generated config file")
	configCmd.Flags().BoolVar(&appendProviders, "append", appendProviders, "append new providers to existing config file")
	configCmd.Flags().BoolVar(&force, "force", force, "override output")
}
