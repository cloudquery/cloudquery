package cmd

import (
	"github.com/spf13/cobra"
)

const globalOptionsTemplate = `The following are global options and can be passed to any commands
{{if .HasAvailableLocalFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Prints list of global CLI options (applies to all commands)",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.UsageFunc()(cmd)
	},
}

func init() {
	optionsCmd.SetUsageTemplate(globalOptionsTemplate)
	rootCmd.AddCommand(optionsCmd)
}
