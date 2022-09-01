package cmd

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/cloudquery/cloudquery/cli/internal/plugin"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	initShort   = "Generate initial config file for source and destination plugins"
	initExample = `
# Downloads aws provider and writes config for aws provider to stdout
cloudquery generate source aws

# Downloads aws provider and generates initial config in aws.yml
cloudquery generate source --registry grpc --output aws.yml "localhost:7777"
`
)

func NewCmdGenerate() *cobra.Command {
	registry := enum.NewEnum([]string{"github", "local", "grpc"}, "github")
	cmd := &cobra.Command{
		Use:     "generate <source/destination> <path>",
		Aliases: []string{"gen"},
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.ExactArgs(2),
		RunE:    runGenerate,
	}
	cmd.Flags().Var(registry, "registry", "where to download the plugin")
	cmd.Flags().StringP("output", "O", "", "destination file to write to (defaults to <name_of_plugin>.yml)")
	return cmd
}

func runGenerate(cmd *cobra.Command, args []string) error {
	pluginManager := plugin.NewPluginManager(plugin.WithLogger(log.Logger))
	registry, err := specs.RegistryFromString(cmd.Flag("registry").Value.String())
	if err != nil {
		return fmt.Errorf("runGen: invalid registry %w", err)
	}
	outputFile := cmd.Flag("output").Value.String()
	switch args[0] {
	case "source":
		return genSource(cmd, args[1], pluginManager, registry, outputFile)
	case "destination":
		return genDestination(cmd, args[1], pluginManager, registry, outputFile)
	default:
		return errors.Errorf("unknown type: %s", args[0])
	}
}

func genSource(cmd *cobra.Command, path string, pm *plugin.PluginManager, registry specs.Registry, outputFile string) error {
	if registry == specs.RegistryGithub && !strings.Contains(path, "/") {
		path = "cloudquery/" + path
	}
	version := "latest"
	if strings.Contains(path, "@") {
		version = strings.Split(path, "@")[1]
	}

	sourceSpec := specs.Source{
		Name:     path,
		Path:     path,
		Registry: registry,
		Version:  version,
	}
	plugin, err := pm.NewSourcePlugin(cmd.Context(), sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to create source plugin: %w", err)
	}
	defer plugin.Close()
	client := plugin.GetClient()
	cfg, err := client.ExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}
	configPath := outputFile
	if outputFile == "" {
		name, err := client.Name(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get plugin name: %w", err)
		}
		configPath = name + ".yml"
	}
	err = writeFile(configPath, cfg)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	fmt.Println("Source plugin config successfully written to " + configPath)
	return nil
}

func genDestination(cmd *cobra.Command, path string, pm *plugin.PluginManager, registry specs.Registry, outputFile string) error {
	destSpec := specs.Destination{
		Name:     path,
		Path:     path,
		Registry: registry,
	}
	destPlugin, err := pm.NewDestinationPlugin(cmd.Context(), destSpec)
	if err != nil {
		return fmt.Errorf("failed to get plugin client: %w", err)
	}
	defer destPlugin.Close()
	client := destPlugin.GetClient()
	cfg, err := client.GetExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}
	configPath := outputFile
	if outputFile == "" {
		name, err := client.Name(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to get plugin name: %w", err)
		}
		configPath = name + ".yml"
	}
	err = writeFile(configPath, cfg)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	fmt.Println("Destination plugin config successfully written to " + configPath)
	return nil
}

func writeFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	if err != nil {
		return err
	}
	return w.Flush()
}
