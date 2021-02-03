package cmd

import (
	"fmt"
	"github.com/cloudquery/cloudquery/sdk"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var headerConfig = `providers:`

var validArgs = []string{"aws", "gcp", "okta", "azure", "k8s"}
var configPath = "./config.yml"
var force = false

var configCmd = &cobra.Command{
	Use:       fmt.Sprintf("config [choose one or more of: %s]", strings.Join(validArgs, ",")),
	Short:     "Generate initial config.yml for fetch command",
	ValidArgs: validArgs,
	Args:      cobra.RangeArgs(1, len(validArgs)),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return fmt.Errorf("invalid argument %s for cloudquery gen config. choose from %v", args[0], validArgs)
		}
		var s strings.Builder
		_, err = s.WriteString(headerConfig)
		if err != nil {
			return err
		}
		for _, provider := range args {
			p, err:= sdk.GetProviderPluginClient("./" + provider)
			if err != nil {
				return err
			}
			configYaml, err := p.GenConfig()
			if err != nil {
				return err
			}
			s.WriteString(configYaml)
		}
		s.WriteString("\n")
		if _, err := os.Stat(configPath); err == nil && !force {
			return fmt.Errorf("file %s already exists. Either delete it or specify other path via --path flag", configPath)
		} else if os.IsNotExist(err) || force {
			return ioutil.WriteFile(configPath, []byte(s.String()), 0644)
		} else {
			return err
		}
	},
}

func init() {
	genCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&configPath, "path", configPath, "path to output generated config file")
	configCmd.Flags().BoolVar(&force, "force", force, "override output")
}
