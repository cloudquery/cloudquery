package init

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/spf13/cobra"
)

const (
	initShort   = "Generate initial provider.cq.yml for fetch command"
	initExample = `
  # Downloads aws provider and generates aws.cq.yml for aws provider
  cloudquery init aws
	
	# Downloads aws provider and generates aws.cq.yml for aws provider
	cloudquery init gcp

	`
)

func NewCmdInit() *cobra.Command {
	initCmd := &cobra.Command{
		Use:     "init [choose one or more providers (aws gcp azure okta ...)]",
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.ExactArgs(1),
		RunE:    runInit,
	}
	return initCmd
}

func runInit(cmd *cobra.Command, args []string) error {
	hub := registry.NewRegistryHub(firebase.CloudQueryRegistryURL)
	pm, err := plugin.NewManager(hub, plugin.WithAllowReattach())
	if err != nil {
		return fmt.Errorf("failed to create plugin manager: %w", err)
	}
	providerName := args[0]
	org, name, version, err := ParseProviderCLIArg(providerName)
	if err != nil {
		return fmt.Errorf("failed to parse provider name: %w", err)
	}

	providerBinary, err := hub.Download(cmd.Context(),
		registry.Provider{
			Name:    name,
			Version: version,
			Source:  org,
		},
		false)
	if err != nil {
		return fmt.Errorf("failed to download provider %q: %w", providerName, err)
	}
	fmt.Println("Downloaded provider: ", providerBinary.Name)
	p, err := pm.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    name,
			Version: version,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create plugin %q: %w", providerName, err)
	}
	res, err := p.Provider().GetProviderConfig(cmd.Context(), &cqproto.GetProviderConfigRequest{})
	if err != nil {
		return fmt.Errorf("failed to get provider config %q: %w", providerName, err)
	}
	err = os.WriteFile(fmt.Sprintf("./%s.cq.yml", providerName), res.Config, 0644)
	if err != nil {
		return fmt.Errorf("failed to write provider config %q: %w", providerName, err)
	}

	fmt.Println("Generated provider config: ", fmt.Sprintf("./%s.cq.yml", providerName))

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
