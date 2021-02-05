package cmd

import (
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize CloudQuery by downloading appropriate providers",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag("config_path", cmd.Flags().Lookup("path"))
		configPath := viper.GetString("config_path")
		return cloudqueryclient.Init(configPath)
	},
}

func init() {
	initCmd.Flags().String("path", "./config.yml", "path to configuration file. can be generated with 'gen config' command (env: CQ_CONFIG_PATH)")
	rootCmd.AddCommand(initCmd)
}
