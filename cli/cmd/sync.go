package cmd

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/otel"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
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
# Shard the sync process into 4 shards and run the first shard
cloudquery sync spec.yml --shard 1/4
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
	cmd.Flags().String("shard", "", "Allows splitting the sync process into multiple shards. This feature is in Preview. Please provide feedback to help us improve it. For a list of supported plugins visit https://docs.cloudquery.io/docs/advanced-topics/running-cloudquery-in-parallel")
	cmd.Flags().Bool("cq-columns-not-null", false, "Force CloudQuery internal columns to be NOT NULL. This feature is in Preview. Please provide feedback to help us improve it.")
	_ = cmd.Flags().MarkHidden("cq-columns-not-null")

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

func parseShard(cmd *cobra.Command) (*shard, error) {
	shardFlag, err := cmd.Flags().GetString("shard")
	if err != nil {
		return nil, err
	}
	if shardFlag == "" {
		return nil, nil
	}

	parts := strings.Split(shardFlag, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid shard format: %s. Valid format is num/total, e.g. 1/4", shardFlag)
	}

	num, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid shard format: %s. Shard num should be a valid integer", shardFlag)
	}

	total, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid shard format: %s. Total shards should be a valid integer", shardFlag)
	}

	if num < 1 || total < 1 {
		return nil, fmt.Errorf("invalid shard format: %s. Shard num and total shards should be greater than 0", shardFlag)
	}

	if num > total {
		return nil, fmt.Errorf("invalid shard format: %s. Shard num should be less than or equal to total shards", shardFlag)
	}

	return &shard{num: int32(num), total: int32(total)}, nil
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

	shard, err := parseShard(cmd)
	if err != nil {
		return err
	}

	cqColumnsNotNull, err := cmd.Flags().GetBool("cq-columns-not-null")
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

	tableMetricsLocation, err := cmd.Flags().GetString("tables-metrics-location")
	if err != nil {
		return err
	}
	var otelReceiver *otel.OtelReceiver
	if tableMetricsLocation != "" {
		var sourcesWithOtelEndpoint []string
		for _, source := range sources {
			if source.OtelEndpoint != "" {
				sourcesWithOtelEndpoint = append(sourcesWithOtelEndpoint, source.Name)
			}
		}
		if len(sourcesWithOtelEndpoint) > 0 {
			return fmt.Errorf("the `--tables-metrics-location` flag is not supported for sources with `otel_endpoint` configured. Either remove the `--tables-metrics-location` flag or do not configure `otel_endpoint` for the following sources: %s", strings.Join(sourcesWithOtelEndpoint, ", "))
		}
		otelReceiver, err = otel.StartOtelReceiver(ctx, otel.WithMetricsFilename(tableMetricsLocation))
		if err == nil {
			defer otelReceiver.Shutdown(ctx)
		}
	}

	sourcePluginClients := make(managedplugin.Clients, 0)
	defer func() {
		if err := sourcePluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations, transformers)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name from token: %w", err)
	}

	pluginVersionWarner, _ := managedplugin.NewPluginVersionWarner(log.Logger, authToken.Value)
	specs.WarnOnOutdatedVersions(ctx, pluginVersionWarner, sources, destinations, transformers)

	// in a cloud sync environment, we pass only the relevant environment variables to the plugin
	osEnviron := os.Environ()

	// To force backend destinations to use TCP if the sources are using Docker
	backendsForDockerSource := map[string]struct{}{}    // destination plugin names
	dockerSourcesUsingBackends := map[string]struct{}{} // source plugin names

	for _, source := range sources {
		if otelReceiver != nil {
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
		// To force backend destinations to use TCP if the sources are using Docker
		if source.Registry == specs.RegistryDocker && source.BackendOptions.PluginName() != "" {
			opts = append(opts, managedplugin.WithDockerExtraHosts([]string{"host.docker.internal:host-gateway"}))
			backendsForDockerSource[source.BackendOptions.PluginName()] = struct{}{}
			dockerSourcesUsingBackends[source.Name] = struct{}{}
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
		if _, ok := backendsForDockerSource[destination.Name]; ok {
			opts = append(opts, managedplugin.WithUseTCP())
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
		transPluginClient, err := managedplugin.NewClient(ctx, managedplugin.PluginTransformer, cfg, opts...)
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
					transformerWarnings := specReader.GetTransformerWarningsByName(source.Name)
					for field, msg := range transformerWarnings {
						log.Warn().Str("transformer", transformer.Name()).Str("field", field).Msg(msg)
					}
				}
			}

			_, shouldReplaceLocalhost := dockerSourcesUsingBackends[source.Name]
			src := v3source{
				client:                 cl,
				spec:                   *source,
				shouldReplaceLocalhost: shouldReplaceLocalhost,
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

			syncOptions := syncV3Options{
				source:                    src,
				destinations:              dests,
				transformersByDestination: transfs,
				backend:                   backend,
				uid:                       invocationUUID.String(),
				noMigrate:                 noMigrate,
				summaryLocation:           summaryLocation,
				shard:                     shard,
				cqColumnsNotNull:          cqColumnsNotNull,
			}
			if err := syncConnectionV3(ctx, syncOptions); err != nil {
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
	env := make([]string, 0, len(environ))
	cleanName := strings.ReplaceAll(pluginName, "-", "_")
	prefix := strings.ToUpper("__" + kind + "_" + cleanName + "__")

	globalEnvironmentVariables := map[string]string{}
	specificEnvironmentVariables := map[string]bool{}

	for _, v := range environ {
		switch {
		case strings.HasPrefix(v, "CLOUDQUERY_API_KEY="),
			strings.HasPrefix(v, "AWS_"):
			k := getEnvKey(v)
			globalEnvironmentVariables[k] = v
		case strings.HasPrefix(v, "_CQ_TEAM_NAME="),
			strings.HasPrefix(v, "HOME="):
			env = append(env, v)
		case strings.HasPrefix(v, prefix):
			cleanEnv := strings.TrimPrefix(v, prefix)
			env = append(env, cleanEnv)
			if strings.HasPrefix(cleanEnv, "CLOUDQUERY_API_KEY=") ||
				strings.HasPrefix(cleanEnv, "AWS_") {
				k := getEnvKey(cleanEnv)
				specificEnvironmentVariables[k] = true
			}
		}
	}
	for k, v := range globalEnvironmentVariables {
		if _, ok := specificEnvironmentVariables[k]; !ok {
			env = append(env, v)
		}
	}
	return env
}

func getEnvKey(v string) string {
	parts := strings.SplitN(v, "=", 2)
	return parts[0]
}
