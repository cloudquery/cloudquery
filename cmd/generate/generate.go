package generate

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cmd/enum"
	"github.com/cloudquery/cloudquery/internal/plugin"
	"github.com/cloudquery/cq-provider-sdk/plugins"
	"github.com/cloudquery/cq-provider-sdk/spec"
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
	registry := enum.NewEnum([]string{"hub", "local", "grpc"}, "hub")
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

	pluginManager := plugin.NewPluginManager()
	switch args[0] {
	case "source":
		return genSource(cmd, args[1], pluginManager)
	case "destination":
		return genDestination(cmd, args[1], pluginManager)
	default:
		return fmt.Errorf("unknown type: %s", args[0])
	}
}

func genSource(cmd *cobra.Command, path string, pm *plugin.PluginManager) error {
	sourceSpec := spec.SourceSpec{
		Name:     path,
		Path:     path,
		Registry: cmd.Flag("registry").Value.String(),
	}
	sourceClient, err := pm.GetSourcePluginClient(cmd.Context(), sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to get plugin client: %v", err)
	}
	res, err := sourceClient.GetExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %v", err)
	}
	fmt.Println(res)
	return nil
}

func genDestination(cmd *cobra.Command, path string, pm *plugin.PluginManager) error {
	destSpec := spec.DestinationSpec{
		Name:     path,
		Path:     path,
		Registry: cmd.Flag("registry").Value.String(),
	}
	destClient, err := pm.GetDestinationClient(cmd.Context(), destSpec, plugins.DestinationPluginOptions{})
	if err != nil {
		return fmt.Errorf("failed to get plugin client: %v", err)
	}
	res, err := destClient.GetExampleConfig(cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get example config: %v", err)
	}
	fmt.Println(res)
	return nil
}
