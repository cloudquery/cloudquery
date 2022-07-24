package init

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/spf13/cobra"
)

const (
	initShort   = "Generate initial cloudquery.yml for fetch command"
	initExample = `
  # Downloads aws provider and generates cloudquery.yml for aws provider
  cloudquery init aws

  # Downloads aws,gcp providers and generates one cloudquery.yml with both providers
  cloudquery init aws gcp`
)

func NewCmdInit() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init [choose one or more providers (aws gcp azure okta ...)]",
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    runInit,
	}
	return initCmd
}

func runInit(cmd *cobra.Command, args []string) error {
	return nil
}

func ParseProviderCLIArg(providerCLIArg string) (org string, name string, version string, err error) {
	argParts := strings.Split(providerCLIArg, "@")

	l := len(argParts)

	// e.g. aws@latest@0.1.0
	if l > 2 {
		return "", "", "", fmt.Errorf("invalid provider name@version %q", providerCLIArg)
	}

	// e.g. aws@latest
	if l == 2 && argParts[1] == "latest" {
		org, name, err = registry.ParseProviderName(argParts[0])
		return org, name, "latest", err
	}

	// e.g. aws
	if l == 1 {
		org, name, err = registry.ParseProviderName(argParts[0])
		return org, name, "latest", err
	}

	// e.g. aws@0.12.0
	org, name, err = registry.ParseProviderName(argParts[0])
	if err != nil {
		return "", "", "", err
	}

	ver, err := config.ParseVersion(argParts[1])
	if err != nil {
		return "", "", "", fmt.Errorf("invalid version %q: %w", argParts[1], err)
	}

	return org, name, config.FormatVersion(ver), nil
}
