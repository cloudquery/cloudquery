package cmd

import (
	"github.com/spf13/cobra"
)

var (
	policyCmd = &cobra.Command{
		Use:   "policy [subcommand]",
		Short: "Policy command that unifies policy subcommands.",
		Long: `Examples:
# Download policy from Policy Hub
./cloudquery policy download cq-aws 

# Run cis-v1.3.0 policy
./cloudquery policy run cq-aws cis-v1.3.0

`,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(policyCmd)
}
