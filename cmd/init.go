package cmd

import (
	"github.com/cloudquery/cloudquery/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize CloudQuery by downloading appropriate providers",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlag("configPath", cmd.Flags().Lookup("path"))
		configPath := viper.GetString("configPath")
		c, err := client.New(configPath, "", "")
		if err != nil {
			return err
		}
		return c.Initialize()
	},
}

func init() {
	initCmd.Flags().String("path", "./config.yml", "path to configuration file. can be generated with 'gen config' command (env: CQ_CONFIG_PATH)")
	rootCmd.AddCommand(initCmd)
}
