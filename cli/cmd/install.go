package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	installShort   = "Install required plugin images from your configuration"
	installExample = `# Install required plugins specified in directory
cloudquery install ./directory
# Install required plugins specified in directory and config files
cloudquery install ./directory ./aws.yml ./pg.yml
`
)

func newCmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install [files or directories]",
		Short:   installShort,
		Long:    installShort,
		Example: installExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    install,
	}
	return cmd
}

func install(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sources := specReader.Sources
	destinations := specReader.Destinations
	opts := []managedplugin.Option{managedplugin.WithNoExec()}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}

	sourcePluginConfigs := make([]managedplugin.Config, 0, len(sources))
	for _, source := range sources {
		sourcePluginConfigs = append(sourcePluginConfigs, managedplugin.Config{
			Name:     source.Name,
			Version:  source.Version,
			Path:     source.Path,
			Registry: SpecRegistryToPlugin(source.Registry),
		})
	}
	destinationPluginConfigs := make([]managedplugin.Config, 0, len(destinations))
	for _, destination := range destinations {
		destinationPluginConfigs = append(destinationPluginConfigs, managedplugin.Config{
			Name:     destination.Name,
			Version:  destination.Version,
			Path:     destination.Path,
			Registry: SpecRegistryToPlugin(destination.Registry),
		})
	}

	if _, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...); err != nil {
		return err
	}
	if _, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...); err != nil {
		return err
	}

	return nil
}
