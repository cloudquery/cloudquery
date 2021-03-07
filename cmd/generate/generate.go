package generate

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate initial config.yml for fetch command or policy.yml for query command",
}
