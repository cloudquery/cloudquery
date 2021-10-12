package cmd

import (
	"github.com/spf13/cobra"
)

const moduleHelpMsg = `CloudQuery module commands`

var (
	moduleCmd = &cobra.Command{
		Use:   "module SUBCOMMAND",
		Short: moduleHelpMsg,
		Long:  moduleHelpMsg,
	}
)

func init() {
	moduleCmd.SetUsageTemplate(usageTemplateWithFlags)
	rootCmd.AddCommand(moduleCmd)
}
