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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlag("configPath", cmd.Flags().Lookup("path")); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := viper.GetString("configPath")
		c, err := client.New( "", "")
		if err != nil {
			return err
		}
		return c.Initialize(configPath)
	},
}

func init() {
	initCmd.Flags().String("path", "./config.yml", "path to configuration file. can be generated with 'gen config' command (env: CQ_CONFIG_PATH)")
	rootCmd.AddCommand(initCmd)
}
