package policy

import (
	"github.com/spf13/cobra"
)

const (
	policyDeprecated = "Please use psql directly to run policies. See https://docs.cloudquery.io/docs/policies"
)

func NewCmdPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "policy SUBCOMMAND",
		Deprecated: policyDeprecated,
	}
	return cmd
}
