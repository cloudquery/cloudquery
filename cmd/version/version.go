package version

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/spf13/cobra"
)

const (
	shortVersion = "Print full version info of cloudquery"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "version",
		Short:                 shortVersion,
		Long:                  shortVersion,
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", core.Version)
			fmt.Printf("Commit: %s\n", util.Commit)
			fmt.Printf("Date: %s\n", util.Date)
		},
	}
	// cmd.SetHelpTemplate(usageTemplateWithFlags)
	return cmd
}
