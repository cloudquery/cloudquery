package policy

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	downloadDeprecated = "Use 'policy run' command directly instead. If you need to download a policy to your machine, you can use 'git clone <URL FOR GIT REPO>'. See https://docs.cloudquery.io/docs/cli/policy/sources for more information on how to use the `policy run` command "
)

func newCmdPolicyDownload() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "download",
		Deprecated: downloadDeprecated,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("'policy download' command has been deprecated. Use the 'policy run' command directly instead")
		},
	}
	return cmd
}
