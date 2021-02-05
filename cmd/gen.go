package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:       "gen",
	Short:     "Generate initial config.yml for fetch command or policy.yml for query command",
}

func init() {
	rootCmd.AddCommand(genCmd)
}
