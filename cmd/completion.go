/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script (run --help for full instructions)",
	Long: `To load completions:

Bash:

$ source <(cloudquery completion bash)

# To load completions for each session, execute once:
Linux:
  $ cloudquery completion bash > /etc/bash_completion.d/cloudquery
MacOS:
  $ cloudquery completion bash > /usr/local/etc/bash_completion.d/cloudquery

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ cloudquery completion zsh > "${fpath[1]}/_cloudquery"

# You will need to start a new shell for this setup to take effect.

Fish:

$ cloudquery completion fish | source

# To load completions for each session, execute once:
$ cloudquery completion fish > ~/.config/fish/completions/cloudquery.fish

Powershell:

PS> cloudquery completion powershell | Out-String | Invoke-Expression

# To load completions for every new session, run:
PS> cloudquery completion powershell > cloudquery.ps1
# and source this file from your powershell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: handleError(func(_ context.Context, cmd *cobra.Command, args []string) error {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		return err
	}),
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
