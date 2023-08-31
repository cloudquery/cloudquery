package cmd

import (
	"fmt"
	"math"
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

// findMaxCommonVersion finds the max common version between protocol versions supported by a plugin and those supported by the CLI.
// If all plugin versions are lower than min CLI supported version, it returns -1.
// If all plugin versions are higher than max CLI supported version, it returns -2.
// In this way it is possible tell whether the source or the CLI needs to be updated:
// if -1, the source needs to be updated or the CLI downgraded;
// if -2, the CLI needs to be updated or the source downgraded.
func findMaxCommonVersion(pluginSupported []int, cliSupported []int) int {
	if len(pluginSupported) == 0 {
		return -1
	}

	minCLISupported, maxCLISupported := math.MaxInt32, -1
	for _, v := range cliSupported {
		if v < minCLISupported {
			minCLISupported = v
		}
		if v > maxCLISupported {
			maxCLISupported = v
		}
	}

	minVersion := math.MaxInt32
	maxCommon := -1
	for _, v := range pluginSupported {
		if v < minVersion {
			minVersion = v
		}
		if v > maxCommon && slices.Contains(cliSupported, v) {
			maxCommon = v
		}
	}
	if maxCommon == -1 && minVersion > maxCLISupported {
		return -2
	}
	return maxCommon
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
	sourcePluginClients := make(managedplugin.Clients, 0)
	defer func() {
		if err := sourcePluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	for _, source := range sources {
		opts := []managedplugin.Option{
			managedplugin.WithLogger(log.Logger),
		}
		if cqDir != "" {
			opts = append(opts, managedplugin.WithDirectory(cqDir))
		}
		if disableSentry {
			opts = append(opts, managedplugin.WithNoSentry())
		}
		if source.OtelEndpoint != "" {
			opts = append(opts, managedplugin.WithOtelEndpoint(source.OtelEndpoint))
		}
		if source.OtelEndpointInsecure {
			opts = append(opts, managedplugin.WithOtelEndpointInsecure())
		}
		cfg := managedplugin.Config{
			Name:     source.Name,
			Registry: SpecRegistryToPlugin(source.Registry),
			Version:  source.Version,
			Path:     source.Path,
		}
		sourcePluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginSource, cfg, opts...)
		if err != nil {
			return err
		}
		sourcePluginClients = append(sourcePluginClients, sourcePluginClient)
	}

	destinationPluginClients := make(managedplugin.Clients, 0)
	defer func() {
		if err := destinationPluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	for _, destination := range destinations {
		opts := []managedplugin.Option{
			managedplugin.WithLogger(log.Logger),
		}
		if cqDir != "" {
			opts = append(opts, managedplugin.WithDirectory(cqDir))
		}
		if disableSentry {
			opts = append(opts, managedplugin.WithNoSentry())
		}
		cfg := managedplugin.Config{
			Name:     destination.Name,
			Registry: SpecRegistryToPlugin(destination.Registry),
			Version:  destination.Version,
			Path:     destination.Path,
		}
		destPluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginDestination, cfg, opts...)
		if err != nil {
			return err
		}
		destinationPluginClients = append(destinationPluginClients, destPluginClient)
	}

	for _, source := range sources {
		cl := sourcePluginClients.ClientByName(source.Name)
		versions, err := cl.Versions(ctx)
		if err != nil {
			return fmt.Errorf("failed to get source versions: %w", err)
		}
		maxVersion := findMaxCommonVersion(versions, []int{0, 1, 2, 3})

		var destinationClientsForSource []*managedplugin.Client
		var destinationForSourceSpec []specs.Destination
		var backendClientForSource *managedplugin.Client
		var destinationForSourceBackendSpec *specs.Destination
		for _, destination := range destinations {
			if slices.Contains(source.Destinations, destination.Name) {
				destinationClientsForSource = append(destinationClientsForSource, destinationPluginClients.ClientByName(destination.Name))
				destinationForSourceSpec = append(destinationForSourceSpec, *destination)
				continue
			}

			// if the destination is specified as a backend, but not used as a destination, then we initialize it separately
			if source.BackendOptions != nil && strings.Contains(source.BackendOptions.Connection, "@@plugins."+destination.Name+".") {
				backendClientForSource = destinationPluginClients.ClientByName(destination.Name)
				destinationForSourceBackendSpec = destination
			}
		}
		switch maxVersion {
		case 3:
			// for backwards-compatibility, check for old fields and move them into the spec, log a warning
			warnings := specReader.GetSourceWarningsByName(source.Name)
			for field, msg := range warnings {
				log.Warn().Str("source", source.Name).Str("field", field).Msg(msg)
			}
			if _, found := warnings["scheduler"]; found {
				source.Spec["scheduler"] = source.Scheduler.String() // nolint:staticcheck // use of deprecated field
			}
			if _, found := warnings["concurrency"]; found {
				source.Spec["concurrency"] = source.Concurrency // nolint:staticcheck // use of deprecated field
			}
			for i, destination := range destinationClientsForSource {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions: %w", err)
				}
				if !slices.Contains(versions, 3) {
					return fmt.Errorf("destination plugin %[1]s does not support CloudQuery protocol version 3, required by the %[2]s source plugin. Please upgrade to a newer version of the %[1]s destination plugin", destination.Name(), source.Name)
				}
				destWarnings := specReader.GetDestinationWarningsByName(source.Name)
				for field, msg := range destWarnings {
					log.Warn().Str("destination", destination.Name()).Str("field", field).Msg(msg)
				}
				if _, found := destWarnings["batch_size"]; found {
					destinationForSourceSpec[i].Spec["batch_size"] = destinationForSourceSpec[i].BatchSize // nolint:staticcheck // use of deprecated field
				}
				if _, found := destWarnings["batch_size_bytes"]; found {
					destinationForSourceSpec[i].Spec["batch_size_bytes"] = destinationForSourceSpec[i].BatchSizeBytes // nolint:staticcheck // use of deprecated field
				}
			}

			src := v3source{
				client: cl,
				spec:   *source,
			}
			dests := make([]v3destination, 0, len(destinationClientsForSource))
			for i, destination := range destinationClientsForSource {
				dests = append(dests, v3destination{
					client: destination,
					spec:   destinationForSourceSpec[i],
				})
			}
			var backend *v3destination
			if backendClientForSource != nil && destinationForSourceBackendSpec != nil {
				backend = &v3destination{
					client: backendClientForSource,
					spec:   *destinationForSourceBackendSpec,
				}
			}
			if err := syncConnectionV3(ctx, src, dests, backend, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v3 source %s: %w", cl.Name(), err)
			}
		case 2:
			destinationsVersions := make([][]int, 0, len(destinationClientsForSource))
			for _, destination := range destinationClientsForSource {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions: %w", err)
				}
				if !slices.Contains(versions, 1) {
					return fmt.Errorf("destination plugin %[1]s does not support CloudQuery SDK version 1. Please upgrade to a newer version of the %[1]s destination plugin", destination.Name())
				}
				destinationsVersions = append(destinationsVersions, versions)
			}
			if err := syncConnectionV2(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec, invocationUUID.String(), noMigrate, destinationsVersions); err != nil {
				return fmt.Errorf("failed to sync v2 source %s: %w", cl.Name(), err)
			}
		case 1:
			if err := syncConnectionV1(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Name(), err)
			}
		case 0:
			return fmt.Errorf("please upgrade source %v or use an older CLI version, between v3.0.1 and v3.5.3", source.Name)
		case -1:
			return fmt.Errorf("please upgrade source %v or use an older CLI version, < v3.0.1", source.Name)
		case -2:
			return fmt.Errorf("please upgrade CLI or downgrade source to sync %v", source.Name)
		default:
			return fmt.Errorf("unknown source version %d", maxVersion)
		}
	}

	return nil
}
