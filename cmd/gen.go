package cmd

import (
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:       "gen",
	Short:     "Generate initial config.yml for fetch command or policy.yml for query command",
	ValidArgs: validArgs,
	Args:      cobra.RangeArgs(1, len(validArgs)),
}

func init() {
	rootCmd.AddCommand(genCmd)
}
