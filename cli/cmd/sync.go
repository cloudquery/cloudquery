package cmd

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	syncShort   = "Sync resources from configured source plugins to destinations"
	syncExample = `# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml
`
	unknownFieldErrorPrefix = "code = InvalidArgument desc = failed to decode spec: json: unknown field "
)

func NewCmdSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [files or directories]",
		Short:   syncShort,
		Long:    syncShort,
		Example: syncExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    sync,
	}
	cmd.Flags().Bool("no-migrate", false, "Disable auto-migration before sync. By default, sync runs a migration before syncing resources.")
	return cmd
}

func findMaxVersion(versions []int) int {
	max := -1
	if len(versions) == 0 {
		return max
	}
	for _, v := range versions {
		if v > max {
			max = v
		}
	}
	return max
}

func sync(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	noMigrate, err := cmd.Flags().GetBool("no-migrate")
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

	invocationUUID, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate invocation uuid: %w", err)
	}
	sources := specReader.Sources
	destinations := specReader.Destinations
	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
	}
	if cqDir != "" {
		opts = append(opts, managedplugin.WithDirectory(cqDir))
	}
	if disableSentry {
		opts = append(opts, managedplugin.WithNoSentry())
	}
	sourcePluginConfigs := make([]managedplugin.Config, 0, len(sources))
	for _, source := range sources {
		sourcePluginConfigs = append(sourcePluginConfigs, managedplugin.Config{
			Name:     source.Name,
			Registry: SpecRegistryToPlugin(source.Registry),
			Version:  source.Version,
			Path:     source.Path,
		})
	}

	destinationPluginConfigs := make([]managedplugin.Config, 0, len(destinations))
	for _, destination := range destinations {
		destinationPluginConfigs = append(destinationPluginConfigs, managedplugin.Config{
			Name:     destination.Name,
			Registry: SpecRegistryToPlugin(destination.Registry),
			Version:  destination.Version,
			Path:     destination.Path,
		})
	}

	sourcePluginClients, err := managedplugin.NewClients(ctx, sourcePluginConfigs, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := sourcePluginClients.Terminate(); err != nil {
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
		cl := sourcePluginClients.ClientByName(source.Name)
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
			if err := syncConnectionV2(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v2 source %s: %w", cl.Name(), err)
			}
		case 1:
			if err := syncConnectionV1(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Name(), err)
			}
		case 0:
			return fmt.Errorf("please upgrade your source or use an older v3.0.1 < CLI version < v3.5.3")
		case -1:
			return fmt.Errorf("please upgrade your source or use an older CLI version < v3.0.1")
		default:
			return fmt.Errorf("unknown source version %d", maxVersion)
		}
	}

	return nil
}
