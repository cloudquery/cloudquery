package cmd

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/spf13/cobra"
)

const versionShort = "Print full version info of cloudquery"

func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "version",
		Short:                 versionShort,
		Long:                  versionShort,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", core.Version)
			fmt.Printf("Commit: %s\n", Commit)
			fmt.Printf("Date: %s\n", Date)
		},
	}
	return cmd
}
