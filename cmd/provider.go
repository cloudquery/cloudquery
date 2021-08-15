package cmd

import (
	"github.com/spf13/cobra"
)

var providerHelpMsg = "Top-level command to interact with providers."

var providerCmd = &cobra.Command{
	Use:   "provider [subcommand]",
	Short: providerHelpMsg,
	Long:  providerHelpMsg,
	Example: `
  # Downloads all providers specified in config.hcl:
  ./cloudquery provider download
`,
	Version: Version,
}

func init() {
	rootCmd.AddCommand(providerCmd)
}
