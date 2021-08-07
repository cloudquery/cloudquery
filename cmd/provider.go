package cmd

import (
	"github.com/spf13/cobra"
)

var providerCmd = &cobra.Command{
	Use:   "provider [subcommand]",
	Short: "Provider command that unifies provider subcommands.",
	Long: `Examples:
# Download all providers
./cloudquery provider download
`,
	Version: Version,
}

func init() {
	rootCmd.AddCommand(providerCmd)
}
