package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

const (
	migrateShort   = "Run migration for source and destination plugins specified in configuration"
	migrateExample = `# Run migration for plugins specified in directory
cloudquery migrate ./directory
# Run migration for plugins specified in directory and config files
cloudquery migrate ./directory ./aws.yml ./pg.yml
`
)

func NewCmdMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate [files or directories]",
		Short:   migrateShort,
		Long:    migrateShort,
		Example: migrateExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    migrate,
	}
	return cmd
}

func migrate(cmd *cobra.Command, args []string) error {
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
	var opts []managedplugin.Option
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

	managedSourceClients, err := managedplugin.NewClients(ctx, sourcePluginConfigs, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := managedSourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationPluginClients, err := managedplugin.NewClients(ctx, destinationPluginConfigs, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := destinationPluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, source := range sources {
		cl := managedSourceClients.ClientByName(source.Name)
		versions, err := cl.Versions(ctx)
		if err != nil {
			return fmt.Errorf("failed to get source versions: %w", err)
		}
		maxVersion := findMaxVersion(versions)
		if maxVersion >= 3 {
			return fmt.Errorf("please upgrade CLI to latest version to sync source %s", cl.Name())
		}

		var destinationClientsForSource []*managedplugin.Client
		var destinationForSourceSpec []specs.Destination
		for _, destination := range destinations {
			if slices.Contains(source.Destinations, destination.Name) {
				destinationClientsForSource = append(destinationClientsForSource, destinationPluginClients.ClientByName(destination.Name))
				destinationForSourceSpec = append(destinationForSourceSpec, *destination)
			}
		}
		switch maxVersion {
		case 2:
			for _, destination := range destinationClientsForSource {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions: %w", err)
				}
				if !slices.Contains(versions, 1) {
					return fmt.Errorf("destination %[1]s does not support CloudQuery SDK version 1. Please upgrade to newer version of %[1]s", destination.Name())
				}
			}
			return migrateConnectionV2(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec)
		case 1:
			return migrateConnectionV1(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec)
		case 0:
			return fmt.Errorf("please upgrade your source or use an older v3.0.1 < CLI version < v3.5.3")
		default:
			return fmt.Errorf("unknown version %d", maxVersion)
		}
	}
	return nil
}
