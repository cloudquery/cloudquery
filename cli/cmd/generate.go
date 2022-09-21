package cmd

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/cli/internal/enum"
	"github.com/cloudquery/cloudquery/cli/internal/plugins"
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
	pluginManager := plugins.NewPluginManager(plugins.WithLogger(log.Logger))
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

func getSourceSpec(path string, registry specs.Registry) specs.Source {
	if registry == specs.RegistryGithub {
		if !strings.Contains(path, "/") {
			path = "cloudquery/" + path
		}

		nameParts := strings.Split(path, "/")
		versionParts := strings.Split(nameParts[1], "@")

		org := nameParts[0]
		name := versionParts[0]
		version := "latest"
		if len(versionParts) > 1 {
			version = versionParts[1]
		}
		return specs.Source{
			Name:     name,
			Path:     fmt.Sprintf("%s/%s", org, name),
			Registry: registry,
			Version:  version,
		}
	}

	return specs.Source{
		Name:     path,
		Path:     path,
		Registry: registry,
		Version:  "latest",
	}
}

func genSource(cmd *cobra.Command, path string, pm *plugins.PluginManager, registry specs.Registry, outputFile string) error {
	sourceSpec := getSourceSpec(path, registry)
	sourceSpec.SetDefaults()

	plugin, err := pm.NewSourcePlugin(cmd.Context(), &sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to create source plugin: %w", err)
	}
	defer plugin.Close()
	client := plugin.GetClient()

	name, err := client.Name(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get plugin name: %w", err)
	}
	sourceSpec.Name = name
	cfg, err := client.ExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}
	sourceSpec.Spec = cfg

	configPath := outputFile
	if configPath == "" {
		configPath = name + ".yml"
	}
	err = writeSource(configPath, sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to write source to file: %w", err)
	}
	fmt.Println("Source plugin config successfully written to " + configPath)
	return nil
}

func genDestination(cmd *cobra.Command, path string, pm *plugins.PluginManager, registry specs.Registry, outputFile string) error {
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

	cfg, err := client.GetExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %w", err)
	}
	destSpec.Spec = cfg

	configPath := outputFile
	if configPath == "" {
		configPath = name + ".yml"
	}
	err = writeDestination(configPath, destSpec)
	if err != nil {
		return fmt.Errorf("failed to write destination to file: %w", err)
	}
	fmt.Println("Destination plugin config successfully written to " + configPath)
	return nil
}

func writeSource(path string, sourceSpec specs.Source) error {
	return writeConfig(path, "source.go.tpl", sourceSpec)
}

func writeDestination(path string, destinationSpec specs.Destination) error {
	return writeConfig(path, "destination.go.tpl", destinationSpec)
}

func writeConfig(path, cfgTemplate string, spec interface{}) error {
	err := os.MkdirAll(filepath.Dir(path), 0744)
	if err != nil {
		return fmt.Errorf("failed to directory for file: %w", err)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	tpl, err := template.New(cfgTemplate).Funcs(template.FuncMap{
		"indent": indentSpaces,
	}).ParseFS(templatesFS, "templates/"+cfgTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	w := bufio.NewWriter(f)
	err = tpl.Execute(w, spec)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	return w.Flush()
}

func indentSpaces(text string, spaces int) string {
	s := strings.Repeat(" ", spaces)
	return s + strings.ReplaceAll(text, "\n", "\n"+s)
}
