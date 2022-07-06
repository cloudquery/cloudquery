package options

import (
	"github.com/spf13/cobra"
)

const (
	globalOptionsTemplate = `The following are global options and can be passed to any commands{{if .HasAvailableLocalFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
	optionsShort = "Prints list of global CLI options (applies to all commands)"
)

func NewCmdOptions() *cobra.Command {
	optionsCmd := &cobra.Command{
		Use:   "options",
		Short: optionsShort,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.UsageFunc()(cmd)
		},
	}
	optionsCmd.SetUsageTemplate(globalOptionsTemplate)
	return optionsCmd
}
