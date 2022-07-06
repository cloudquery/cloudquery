package policy

import (
	"github.com/spf13/cobra"
)

const (
	policyShort = "Download and run CloudQuery policy"
)

func NewCmdPolicy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy SUBCOMMAND",
		Short: policyShort,
		Long:  policyShort,
	}
	cmd.AddCommand(newCmdPolicyDescribe(), newCmdPolicyDownload(), newCmdPolicyRun(), newCmdPolicySnapshot(), newCmdPolicyTest(), newCmdPolicyValidate(), newCmdPolicyPrune())
	return cmd
}
