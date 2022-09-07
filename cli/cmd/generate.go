package cmd

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/cloudquery/plugin-sdk/clients"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/cloudquery/cloudquery/cli/internal/plugin"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

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
		return fmt.Errorf("runGen: invalid type %s", args[0])
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
	sourceSpec.SetDefaults()

	plugin, err := pm.NewSourcePlugin(cmd.Context(), sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to create source plugin: %w", err)
	}
	defer plugin.Close()
	client := plugin.GetClient()

	opts := clients.SourceExampleConfigOptions{
		Path:     path,
		Registry: registry,
	}
	name, err := client.Name(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get plugin name: %w", err)
	}
	cfg, err := client.ExampleConfig(cmd.Context(), opts)
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}

	configPath := outputFile
	if configPath == "" {
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
	destSpec.SetDefaults()

	destPlugin, err := pm.NewDestinationPlugin(cmd.Context(), destSpec)
	if err != nil {
		return fmt.Errorf("failed to get plugin client: %w", err)
	}
	defer destPlugin.Close()
	client := destPlugin.GetClient()
	name, err := client.Name(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get plugin name: %w", err)
	}
	destSpec.Name = name

	version, err := client.Version(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get plugin name: %w", err)
	}
	destSpec.Version = version

	opts := clients.DestinationExampleConfigOptions{
		Path:     path,
		Registry: registry,
	}
	cfg, err := client.GetExampleConfig(cmd.Context(), opts)
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}

	configPath := outputFile
	if configPath == "" {
		configPath = name + ".yml"
	}
	err = writeFile(configPath, cfg)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	fmt.Println("Destination plugin config successfully written to " + configPath)
	return nil
}

func writeFile(path, cfg string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString(cfg)
	if err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return w.Flush()
}
