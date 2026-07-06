package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	validateConfigShort = "Validate config"
	validateConfigLong  = `Validate configuration without running a sync.

For ` + "`registry: cloudquery`" + ` plugins, the spec JSON schema is fetched from
the CloudQuery Hub API (https://api.cloudquery.io). This avoids downloading
the plugin binary and works for public plugins without authentication; if a
CloudQuery API token is available (via login or CLOUDQUERY_API_KEY) it is
propagated so private plugins resolve too.

For other registries (` + "`local`, `grpc`, `docker`" + `) the plugin is still spawned
locally to obtain its schema, identical to the previous behaviour. The tables
list is not validated against the source — this validation is stricter than
the validation done during ` + "`sync`" + `, so a config passing here will also pass
sync's validation.`
	validateConfigExample = `# Validate configs
cloudquery validate-config ./directory
# Validate configs from directories and files
cloudquery validate-config ./directory ./aws.yml ./pg.yml
`
)

func newCmdValidateConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "validate-config [files or directories]",
		Short:   validateConfigShort,
		Long:    validateConfigLong,
		Example: validateConfigExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    validateConfig,
		Hidden:  false,
	}
	cmd.Flags().String("license", "", "Set offline license file. When provided, the Hub API is bypassed and plugins are spawned locally (mirrors `cloudquery sync --license`).")

	return cmd
}

func validateConfig(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}
	licenseFile, err := cmd.Flags().GetString("license")
	if err != nil {
		return err
	}

	ctx := cmd.Context()

	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	// Read without enforcing a destination: a source-only platform spec (as `init`
	// scaffolds) has none until the platform destination is auto-injected below.
	// Full validation runs via SetDestinationsAndValidate once destinations are
	// finalized, so non-platform specs are validated identically to before.
	specReader, err := specs.NewSpecReaderWithoutValidation(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	sources := specReader.Sources
	destinations := specReader.Destinations
	transformers := specReader.Transformers

	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations, transformers)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	// Resolve the plugin download credential + team at most once (mirrors sync's
	// DownloadAuth), shared by the platform opt-in block and the plugin-spawn path
	// below. Lazy: only invoked when actually needed, so pure-Hub-API validation of
	// public plugins never requires auth.
	var dlToken, teamName string
	downloadAuthResolved := false
	resolveDownloadAuth := func() error {
		if downloadAuthResolved {
			return nil
		}
		var err error
		if dlToken, teamName, err = platform.DownloadAuth(ctx, log.Logger, sources, destinations, transformers); err != nil {
			return err
		}
		downloadAuthResolved = true
		return nil
	}

	// A user-declared `platform` destination (debugging/override) is validated
	// normally below; only the CLI-injected one is skipped. Capture it before
	// injection, which may add one.
	userPlatformDest := slices.ContainsFunc(destinations, func(d *specs.Destination) bool {
		return platform.IsInjectedDestination(d.Name)
	})

	// Platform opt-in: mirror `sync`. Resolve the platform credential, reject
	// source versions the tenant can't ingest (the same window the sync-time
	// CreateExternalSync gate enforces), then auto-inject the `platform`
	// destination — so a source-only platform spec validates exactly as sync would
	// run it. Non-platform specs skip this entirely and are unaffected.
	if platform.AnySourceTargetsPlatform(sources) {
		if err := resolveDownloadAuth(); err != nil {
			return err
		}
		if err := platform.GateSources(ctx, log.Logger, dlToken, teamName, sources); err != nil {
			return err
		}
		if destinations, err = platform.MaybeInjectDestination(ctx, log.Logger, dlToken, teamName, sources, destinations); err != nil {
			return err
		}
	}

	if err := specReader.SetDestinationsAndValidate(destinations); err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	destinations = specReader.Destinations

	apiClient, err := api.NewClient(authToken.Value)
	if err != nil {
		return fmt.Errorf("failed to create Hub API client: %w", err)
	}

	// Partition entries: CloudQuery-registry plugins fetch their schema from the
	// Hub API; all other registries fall through to the original plugin-spawn path.
	sourcePluginConfigs := make([]managedplugin.Config, 0, len(sources))
	sourceRegInferred := make([]bool, 0, len(sources))
	sourceSpawnIdx := make([]int, 0, len(sources))
	destinationPluginConfigs := make([]managedplugin.Config, 0, len(destinations))
	destinationRegInferred := make([]bool, 0, len(destinations))
	destinationSpawnIdx := make([]int, 0, len(destinations))

	var initErrors []error

	useHubAPI := licenseFile == ""
	for i, source := range sources {
		if useHubAPI && source.Registry == specs.RegistryCloudQuery {
			if err := validateViaHubAPI(ctx, apiClient, source.Path, cloudquery_api.PluginKindSource, source.Version, source.Spec); err != nil {
				initErrors = append(initErrors, fmt.Errorf("failed to validate source config %v: %w", source.VersionString(), err))
			} else {
				log.Info().Str("source", source.VersionString()).Msg("validated successfully")
			}
			continue
		}
		sourcePluginConfigs = append(sourcePluginConfigs, managedplugin.Config{
			Name:       source.Name,
			Version:    source.Version,
			Path:       source.Path,
			Registry:   SpecRegistryToPlugin(source.Registry),
			DockerAuth: source.DockerRegistryAuthToken,
		})
		sourceRegInferred = append(sourceRegInferred, source.RegistryInferred())
		sourceSpawnIdx = append(sourceSpawnIdx, i)
	}

	for i, destination := range destinations {
		// Skip the auto-injected `platform` destination — CLI-generated (a
		// token-based spec with api_url derived from the token), not user config;
		// sync doesn't schema-validate it either, and the source's version was
		// already gated above. A user-declared `platform` override is NOT skipped:
		// it's validated through the normal path below.
		if platform.IsInjectedDestination(destination.Name) && !userPlatformDest {
			continue
		}
		if useHubAPI && destination.Registry == specs.RegistryCloudQuery {
			if err := validateViaHubAPI(ctx, apiClient, destination.Path, cloudquery_api.PluginKindDestination, destination.Version, destination.Spec); err != nil {
				initErrors = append(initErrors, fmt.Errorf("failed to validate destination config %v: %w", destination.VersionString(), err))
			} else {
				log.Info().Str("destination", destination.VersionString()).Msg("validated successfully")
			}
			continue
		}
		destinationPluginConfigs = append(destinationPluginConfigs, managedplugin.Config{
			Name:       destination.Name,
			Version:    destination.Version,
			Path:       destination.Path,
			Registry:   SpecRegistryToPlugin(destination.Registry),
			DockerAuth: destination.DockerRegistryAuthToken,
		})
		destinationRegInferred = append(destinationRegInferred, destination.RegistryInferred())
		destinationSpawnIdx = append(destinationSpawnIdx, i)
	}

	if len(sourcePluginConfigs) == 0 && len(destinationPluginConfigs) == 0 {
		return errors.Join(initErrors...)
	}

	// Plugin spawn is still required for non-Hub registries (local/grpc/docker).
	// Reuse the download credential + team resolved above (or resolve now if this
	// is a non-platform spec) — same as sync — and propagate it so spawned plugins
	// authenticate premium-table validation / usage against cloud (see sync.go).
	if err := resolveDownloadAuth(); err != nil {
		return fmt.Errorf("failed to resolve plugin download auth: %w", err)
	}
	platform.PropagatePluginCredential(dlToken)
	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithAuthToken(dlToken),
		managedplugin.WithTeamName(teamName),
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
	if licenseFile != "" {
		opts = append(opts, managedplugin.WithLicenseFile(licenseFile))
	}

	sourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...)
	if err != nil {
		return enrichClientError(sourceClients, sourceRegInferred, err)
	}
	defer func() {
		if err := sourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationClients, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...)
	if err != nil {
		return enrichClientError(destinationClients, destinationRegInferred, err)
	}
	defer func() {
		if err := destinationClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()

	for ci, client := range sourceClients {
		i := sourceSpawnIdx[ci]
		pluginClient := plugin.NewPluginClient(client.Conn)
		log.Info().Str("source", sources[i].VersionString()).Msg("Initializing source")
		err := validatePluginSpec(ctx, pluginClient, sources[i].Spec)
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to validate source config %v: %w", sources[i].VersionString(), err))
		} else {
			log.Info().Str("source", sources[i].VersionString()).Msg("validated successfully")
		}
	}
	for ci, client := range destinationClients {
		i := destinationSpawnIdx[ci]
		pluginClient := plugin.NewPluginClient(client.Conn)
		log.Info().Str("destination", destinations[i].VersionString()).Msg("Initializing destination")
		err = validatePluginSpec(ctx, pluginClient, destinations[i].Spec)
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to validate destination config %v: %w", destinations[i].VersionString(), err))
		} else {
			log.Info().Str("destination", destinations[i].VersionString()).Msg("validated successfully")
		}
	}

	return errors.Join(initErrors...)
}

// validateViaHubAPI fetches the JSON schema for a CloudQuery-registry plugin from
// the Hub API and validates spec against it. The plugin path is the canonical
// "team/name" hub identifier from the source/destination's `path:` field.
func validateViaHubAPI(ctx context.Context, c *cloudquery_api.ClientWithResponses, pluginPath string, kind cloudquery_api.PluginKind, version string, spec map[string]any) error {
	team, name, err := splitHubPath(pluginPath)
	if err != nil {
		return err
	}
	log.Info().Str("team", team).Str("name", name).Str("version", version).Str("kind", string(kind)).Msg("Fetching spec schema from Hub API")
	resp, err := c.GetPluginVersionWithResponse(ctx, team, kind, name, version)
	if err != nil {
		return fmt.Errorf("failed to fetch spec schema from Hub API: %w", err)
	}
	if resp.StatusCode() != http.StatusOK || resp.JSON200 == nil {
		return hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}
	if resp.JSON200.SpecJsonSchema == nil || *resp.JSON200.SpecJsonSchema == "" {
		log.Info().Str("name", name).Msg("Hub did not return a spec schema, skipping validation")
		return nil
	}
	return validateSpecAgainstSchema(*resp.JSON200.SpecJsonSchema, spec)
}

// splitHubPath parses a CloudQuery-registry path field ("team/name") into its parts.
func splitHubPath(p string) (team, name string, err error) {
	parts := strings.SplitN(p, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid cloudquery-registry path %q (expected team/name)", p)
	}
	return parts[0], parts[1], nil
}
