package generate

import (
	"fmt"

	"strings"

	"github.com/cloudquery/cloudquery/cmd/enum"
	"github.com/cloudquery/cloudquery/internal/plugin"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	initShort   = "Generate initial *.cq.yml file for sources,destionations,connections"
	initExample = `
# Downloads aws provider and generates aws.cq.yml for aws provider
cloudquery generate aws

# Downloads aws provider and generates aws.cq.yml for aws provider
cloudquery generate gcp
`
)

func NewCmdInit() *cobra.Command {
	registry := enum.NewEnum([]string{"github", "local", "grpc"}, "github")
	cmd := &cobra.Command{
		Use:     "generate <source/destination/connection> <path>",
		Aliases: []string{"gen"},
		Short:   initShort,
		Long:    initShort,
		Example: initExample,
		Args:    cobra.ExactArgs(2),
		RunE:    runGen,
	}
	cmd.Flags().Var(registry, "registry", "where to download the plugin")
	return cmd
}

func runGen(cmd *cobra.Command, args []string) error {

	pluginManager := plugin.NewPluginManager(plugin.WithLogger(log.Logger))
	switch args[0] {
	case "source":
		return genSource(cmd, args[1], pluginManager)
	case "destination":
		return genDestination(cmd, args[1], pluginManager)
	default:
		return errors.Errorf("unknown type: %s", args[0])
	}
}

func genSource(cmd *cobra.Command, path string, pm *plugin.PluginManager) error {
	registry := cmd.Flag("registry").Value.String()
	if registry == "github" && !strings.Contains(path, "/") {
		path = "cloudquery/" + path
	}
	version := "latest"
	if strings.Contains(path, "@") {
		version = strings.Split(path, "@")[1]
	}

	sourceSpec := specs.SourceSpec{
		Name:     path,
		Path:     path,
		Registry: registry,
		Version:  version,
	}
	sourceClient, err := pm.GetSourcePluginClient(cmd.Context(), sourceSpec)
	if err != nil {
		return errors.Wrap(err, "failed to get plugin client")
	}
	res, err := sourceClient.GetExampleConfig(cmd.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get example config")
	}
	fmt.Println(res)
	return nil
}

func genDestination(cmd *cobra.Command, path string, pm *plugin.PluginManager) error {
	destSpec := specs.DestinationSpec{
		Name:     path,
		Path:     path,
		Registry: cmd.Flag("registry").Value.String(),
	}
	destClient, err := pm.GetDestinationClient(cmd.Context(), destSpec, plugins.DestinationPluginOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to get plugin client")
	}
	res, err := destClient.GetExampleConfig(cmd.Context())
	if err != nil {
		return errors.Wrap(err, "failed to get example config")
	}
	fmt.Println(res)
	return nil
}
