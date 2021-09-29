package cmd

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/client"

	"github.com/spf13/cobra"
)

const versionHelpMsg = "Print full version info of cloudquery"

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 versionHelpMsg,
	Long:                  versionHelpMsg,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", client.Version)
		fmt.Printf("Commit: %s\n", Commit)
		fmt.Printf("Date: %s\n", Date)
	},
}

func init() {
	versionCmd.SetHelpTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(versionCmd)
}
