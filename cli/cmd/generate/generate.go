package generate

import (
	"fmt"

	"strings"

	"github.com/cloudquery/cloudquery/cli/cmd/enum"
	"github.com/cloudquery/cloudquery/cli/internal/plugin"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	initShort   = "Generate initial *.cq.yml file for sources and destionations plugins"
	initExample = `
# Downloads aws provider and generates aws.cq.yml for aws provider
cloudquery generate source aws

# Downloads aws provider and generates aws.cq.yml for aws provider
cloudquery generate source --registry grpc "localhost:7777"
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
	return cmd
}

func runGenerate(cmd *cobra.Command, args []string) error {
	pluginManager := plugin.NewPluginManager(plugin.WithLogger(log.Logger))
	registry, err := specs.RegistryFromString(cmd.Flag("registry").Value.String())
	if err != nil {
		return fmt.Errorf("runGen: invalid registry %w", err)
	}
	switch args[0] {
	case "source":
		return genSource(cmd, args[1], pluginManager, registry)
	case "destination":
		return genDestination(cmd, args[1], pluginManager, registry)
	default:
		return errors.Errorf("unknown type: %s", args[0])
	}
}

func genSource(cmd *cobra.Command, path string, pm *plugin.PluginManager, registry specs.Registry) error {
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
	res, err := plugin.GetClient().ExampleConfig(cmd.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get example config")
	}
	fmt.Println(res)
	return nil
}

func genDestination(cmd *cobra.Command, path string, pm *plugin.PluginManager, registry specs.Registry) error {
	destSpec := specs.Destination{
		Name:     path,
		Path:     path,
		Registry: registry,
	}
	destPlugin, err := pm.NewDestinationPlugin(cmd.Context(), destSpec)
	if err != nil {
		return errors.Wrap(err, "failed to get plugin client")
	}
	defer destPlugin.Close()
	res, err := destPlugin.GetClient().GetExampleConfig(cmd.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get example config")
	}
	fmt.Println(res)
	return nil
}
