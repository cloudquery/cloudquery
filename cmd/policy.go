package cmd

import (
	"github.com/spf13/cobra"
)

const policyHelpMsg = `Download and run CloudQuery policy`

var (
	policyCmd = &cobra.Command{
		Use:   "policy SUBCOMMAND",
		Short: policyHelpMsg,
		Long:  policyHelpMsg,
	}
)

func init() {
	policyCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(policyCmd)
}
