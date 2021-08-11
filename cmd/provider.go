package cmd

import (
	"github.com/spf13/cobra"
)

var providerHelpMsg = "Provider command that unifies provider subcommands."

var providerCmd = &cobra.Command{
	Use:   "provider [subcommand]",
	Short: providerHelpMsg,
	Long:  providerHelpMsg,
	Example: `
  # Downloads all providers mentioned in the configuration file:
  ./cloudquery provider download
`,
	Version: Version,
}

func init() {
	rootCmd.AddCommand(providerCmd)
}
