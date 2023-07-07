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

	managedSourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := managedSourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationPluginClients, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...)
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
		maxVersion := findMaxCommonVersion(versions, []int{3, 2, 1, 0})

		var destinationClientsForSource []*managedplugin.Client
		var destinationForSourceSpec []specs.Destination
		for _, destination := range destinations {
			if slices.Contains(source.Destinations, destination.Name) {
				destinationClientsForSource = append(destinationClientsForSource, destinationPluginClients.ClientByName(destination.Name))
				destinationForSourceSpec = append(destinationForSourceSpec, *destination)
			}
		}
		switch maxVersion {
		case 3:
			for _, destination := range destinationClientsForSource {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions: %w", err)
				}
				if !slices.Contains(versions, 3) {
					return fmt.Errorf("destination %[1]s does not support CloudQuery protocol version 3, required by %[2]s. Please upgrade to newer version of %[1]s", destination.Name(), source.Name)
				}
			}
			if err := migrateConnectionV3(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec); err != nil {
				return fmt.Errorf("failed to migrate v3 source %s: %w", cl.Name(), err)
			}
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
			if err := migrateConnectionV2(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec); err != nil {
				return fmt.Errorf("failed to migrate source %v@%v: %w", source.Name, source.Version, err)
			}
		case 1:
			if err := migrateConnectionV1(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec); err != nil {
				return fmt.Errorf("failed to migrate source %v@%v: %w", source.Name, source.Version, err)
			}
		case 0:
			return fmt.Errorf("please upgrade your source or use a CLI version between v3.0.1 and v3.5.3")
		case -1:
			return fmt.Errorf("please upgrade CLI to sync source %v@%v", source.Name, source.Version)
		case -2:
			return fmt.Errorf("please downgrade CLI or upgrade source to sync %v", source.Name)
		}
	}
	return nil
}
