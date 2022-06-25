package policy

import (
	"github.com/cloudquery/cloudquery/cmd/util"
	"github.com/spf13/cobra"
)

const (
	policyHelpMsg = `Download and run CloudQuery policy`

	policyValidateHelpMsg = "Validate policy for any issues and diagnostics"
)

type policyOptions struct {
	Config string
}

func NewCmdPolicy() *cobra.Command {
	o := policyOptions{}
	cmd := &cobra.Command{
		Use:   "policy SUBCOMMAND",
		Short: policyHelpMsg,
		Long:  policyHelpMsg,
	}
	cmd.AddCommand(NewCmdPolicyDownload(o), NewCmdPolicyValidate(o), NewCmdPolicyPrune(o))
	cmd.Flags().StringVar(&o.Config, "config", "./config.*", util.ConfigHelp)
	return cmd
}

// func init() {
// 	describePolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
// 	policyDownloadCmd.SetUsageTemplate(usageTemplateWithFlags)

// 	flags := policyRunCmd.Flags()
// 	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
// 	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")
// 	flags.BoolVar(&storeResults, "enable-db-persistence", false, "Enable storage of policy output in database")
// 	flags.Bool("disable-fetch-check", false, "Disable checking if a respective fetch happened before running policies")
// 	policyRunCmd.SetUsageTemplate(usageTemplateWithFlags)
// 	policyCmd.AddCommand(policyRunCmd)

// 	snapshotPolicyCmd.SetUsageTemplate(usageTemplateWithFlags)
// 	policyCmd.AddCommand(snapshotPolicyCmd)

// 	flags = policyTestCmd.Flags()
// 	flags.StringVar(&outputDir, "output-dir", "", "Generates a new file for each policy at the given dir with the output")
// 	flags.BoolVar(&noResults, "no-results", false, "Do not show policies results")

// }
