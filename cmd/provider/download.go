package provider

import (
	"github.com/spf13/cobra"
)

const (
	downloadShort      = "Downloads all providers specified in config.hcl. (Deprecated: Please use `provider sync` instead)"
	downloadDeprecated = "Please use `cloudquery provider sync` instead."
)

func newCmdProviderDownload() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "download",
		Short:      downloadShort,
		Long:       downloadShort,
		Deprecated: downloadDeprecated,
	}
	return cmd
}
