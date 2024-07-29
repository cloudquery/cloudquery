package cmd

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/cloudquery/cloudquery/cli/internal/otel"
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
# Log tables metrics to a file
cloudquery sync ./directory ./aws.yml ./pg.yml --tables-metrics-location metrics.txt
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
	cmd.Flags().String("license", "", "set offline license file")
	cmd.Flags().String("summary-location", "", "Sync summary file location. This feature is in Preview. Please provide feedback to help us improve it.")
	cmd.Flags().String("tables-metrics-location", "", "Tables metrics file location. This feature is in Preview. Please provide feedback to help us improve it. Works with plugins released on 2024-07-10 or later.")

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

	licenseFile, err := cmd.Flags().GetString("license")
	if err != nil {
		return err
	}

	// in the cloud sync environment, we pass only the relevant environment variables to the plugin
	_, isolatePluginEnvironment := os.LookupEnv("CQ_CLOUD")

	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}

	sources := specReader.Sources
	destinations := specReader.Destinations
	transformers := specReader.Transformers

	sourcePluginClients := make(managedplugin.Clients, 0)
	defer func() {
		if err := sourcePluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name from token: %w", err)
	}

	// in a cloud sync environment, we pass only the relevant environment variables to the plugin
	osEnviron := os.Environ()

	tableMetricsLocation, err := cmd.Flags().GetString("tables-metrics-location")
	if err != nil {
		return err
	}
	var otelReceiver *otel.OtelReceiver
	if tableMetricsLocation != "" {
		otelReceiver, err = otel.StartOtelReceiver(ctx, otel.WithMetricsFilename(tableMetricsLocation))
		if err == nil {
			defer otelReceiver.Shutdown(ctx)
		}
	}

	for _, source := range sources {
		if source.OtelEndpoint == "" && otelReceiver != nil {
			source.OtelEndpoint = otelReceiver.Endpoint
			source.OtelEndpointInsecure = true
		}
		opts := []managedplugin.Option{
			managedplugin.WithLogger(log.Logger),
			managedplugin.WithOtelEndpoint(source.OtelEndpoint),
			managedplugin.WithAuthToken(authToken.Value),
			managedplugin.WithTeamName(teamName),
			managedplugin.WithLicenseFile(licenseFile),
		}
		if logConsole {
			opts = append(opts, managedplugin.WithNoProgress())
		}
		if cqDir != "" {
			opts = append(opts, managedplugin.WithDirectory(cqDir))
		}
		if disableSentry {
			opts = append(opts, managedplugin.WithNoSentry())
		}
		if source.OtelEndpointInsecure {
			opts = append(opts, managedplugin.WithOtelEndpointInsecure())
		}
		cfg := managedplugin.Config{
			Name:       source.Name,
			Registry:   SpecRegistryToPlugin(source.Registry),
			Version:    source.Version,
			Path:       source.Path,
			DockerAuth: source.DockerRegistryAuthToken,
		}
		if isolatePluginEnvironment {
			cfg.Environment = filterPluginEnv(osEnviron, source.Name, "source")
		}
		sourcePluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginSource, cfg, opts...)
		if err != nil {
			return enrichClientError(managedplugin.Clients{}, []bool{source.RegistryInferred()}, err)
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
			managedplugin.WithAuthToken(authToken.Value),
			managedplugin.WithTeamName(teamName),
			managedplugin.WithLicenseFile(licenseFile),
		}
		if logConsole {
			opts = append(opts, managedplugin.WithNoProgress())
		}
		if cqDir != "" {
			opts = append(opts, managedplugin.WithDirectory(cqDir))
		}
		if disableSentry {
			opts = append(opts, managedplugin.WithNoSentry())
		}

		cfg := managedplugin.Config{
			Name:       destination.Name,
			Registry:   SpecRegistryToPlugin(destination.Registry),
			Version:    destination.Version,
			Path:       destination.Path,
			DockerAuth: destination.DockerRegistryAuthToken,
		}
		if isolatePluginEnvironment {
			cfg.Environment = filterPluginEnv(osEnviron, destination.Name, "destination")
		}
		destPluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginDestination, cfg, opts...)
		if err != nil {
			return enrichClientError(managedplugin.Clients{}, []bool{destination.RegistryInferred()}, err)
		}
		destinationPluginClients = append(destinationPluginClients, destPluginClient)
	}

	transformerPluginClients := make(managedplugin.Clients, 0)
	defer func() {
		if err := transformerPluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	for _, transformer := range transformers {
		opts := []managedplugin.Option{
			managedplugin.WithLogger(log.Logger),
			managedplugin.WithAuthToken(authToken.Value),
			managedplugin.WithTeamName(teamName),
			managedplugin.WithLicenseFile(licenseFile),
		}
		if logConsole {
			opts = append(opts, managedplugin.WithNoProgress())
		}
		if cqDir != "" {
			opts = append(opts, managedplugin.WithDirectory(cqDir))
		}
		if disableSentry {
			opts = append(opts, managedplugin.WithNoSentry())
		}

		cfg := managedplugin.Config{
			Name:       transformer.Name,
			Registry:   SpecRegistryToPlugin(transformer.Registry),
			Version:    transformer.Version,
			Path:       transformer.Path,
			DockerAuth: transformer.DockerRegistryAuthToken,
		}
		transPluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginDestination, cfg, opts...) // TODO: s/PluginDestination/PluginTransformer/
		if err != nil {
			return enrichClientError(managedplugin.Clients{}, []bool{transformer.RegistryInferred()}, err)
		}
		transformerPluginClients = append(transformerPluginClients, transPluginClient)
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
		var transformerClientsForDestination = map[string][]*managedplugin.Client{}
		var transformerForDestinationSpec = map[string][]specs.Transformer{}
		var backendClientForSource *managedplugin.Client
		var destinationForSourceBackendSpec *specs.Destination
		for _, destination := range destinations {
			if slices.Contains(source.Destinations, destination.Name) {
				destinationClientsForSource = append(destinationClientsForSource, destinationPluginClients.ClientByName(destination.Name))
				destinationForSourceSpec = append(destinationForSourceSpec, *destination)

				// Each destination defines their own transformers
				ts := []*managedplugin.Client{}
				tsSpecs := []specs.Transformer{}
				for _, transformer := range transformers {
					if slices.Contains(destination.Transformers, transformer.Name) {
						ts = append(ts, transformerPluginClients.ClientByName(transformer.Name))
						tsSpecs = append(tsSpecs, *transformer)
					}
				}
				transformerClientsForDestination[destination.Name] = ts
				transformerForDestinationSpec[destination.Name] = tsSpecs
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
			for _, destination := range destinationClientsForSource {
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
				for _, transformer := range transformerClientsForDestination[destination.Name()] {
					versions, err := transformer.Versions(ctx)
					if err != nil {
						return fmt.Errorf("failed to get transformer versions: %w", err)
					}
					if !slices.Contains(versions, 3) {
						return fmt.Errorf("transformer plugin %[1]s does not support CloudQuery protocol version 3, required by the %[2]s source plugin. Please upgrade to a newer version of the %[1]s transformer plugin", transformer.Name(), source.Name)
					}
					destWarnings := specReader.GetTransformerWarningsByName(source.Name)
					for field, msg := range destWarnings {
						log.Warn().Str("transformer", transformer.Name()).Str("field", field).Msg(msg)
					}
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
			transfs := map[string][]v3transformer{}
			for destinationName, transformerClients := range transformerClientsForDestination {
				for i, transformer := range transformerClients {
					transfs[destinationName] = append(transfs[destinationName], v3transformer{
						client: transformer,
						spec:   transformerForDestinationSpec[destinationName][i],
					})
				}
			}

			var backend *v3destination
			if backendClientForSource != nil && destinationForSourceBackendSpec != nil {
				backend = &v3destination{
					client: backendClientForSource,
					spec:   *destinationForSourceBackendSpec,
				}
			}

			summaryLocation, err := cmd.Flags().GetString("summary-location")
			if err != nil {
				return err
			}

			if err := syncConnectionV3(ctx, src, dests, transfs, backend, invocationUUID.String(), noMigrate, summaryLocation); err != nil {
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

func filterPluginEnv(environ []string, pluginName, kind string) []string {
	env := make([]string, 0)
	cleanName := strings.ReplaceAll(pluginName, "-", "_")
	prefix := strings.ToUpper("__" + kind + "_" + cleanName + "__")
	for _, v := range environ {
		switch {
		case strings.HasPrefix(v, "CLOUDQUERY_API_KEY="),
			strings.HasPrefix(v, "_CQ_TEAM_NAME="),
			strings.HasPrefix(v, "HOME="):
			env = append(env, v)
		case strings.HasPrefix(v, prefix):
			env = append(env, strings.TrimPrefix(v, prefix))
		}
	}
	return env
}
