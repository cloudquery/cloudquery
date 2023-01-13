package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Version   = "development"
	rootShort = "CloudQuery Scaffold CLI"
	rootLong  = `CloudQuery Scaffold CLI

Open source data integration at scale.

Find more information at:
	https://www.cloudquery.io`
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "scaffold",
		Short:   rootShort,
		Long:    rootLong,
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.AddCommand(
		newCmdScaffoldSource(),
	)
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.DisableAutoGenTag = true

	return cmd
}
