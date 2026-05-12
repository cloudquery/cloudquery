package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	validateConfigShort   = "Validate config"
	validateConfigLong    = "Validate configuration without requiring any credentials or connections. This will not validate the tables specified in the tables list. This validation is stricter than the validation done during `sync`, but if it passes this validation it will pass the sync validation."
	validateConfigExample = `# Validate configs
cloudquery validate-config ./directory
# Validate configs from directories and files
cloudquery validate-config ./directory ./aws.yml ./pg.yml
# Validate fully offline using locally-stored plugin JSON schemas
cloudquery validate-config --schemas-dir ./schemas ./aws.yml
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
	cmd.Flags().String("schemas-dir", "", "Directory of pre-fetched <plugin-name>.json schema files. Plugins with a matching file are validated offline (no plugin spawn, no auth). Use 'cloudquery plugin spec-schema' to generate these files.")

	return cmd
}

func validateConfig(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}
	schemasDir, err := cmd.Flags().GetString("schemas-dir")
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

	// Resolve local schema files when --schemas-dir is set. Empty string means "no file, spawn plugin".
	sourceSchemaFiles := make([]string, len(sources))
	destinationSchemaFiles := make([]string, len(destinations))
	if schemasDir != "" {
		for i, source := range sources {
			sourceSchemaFiles[i] = lookupSchemaFile(schemasDir, source.Name, source.Version)
		}
		for i, destination := range destinations {
			destinationSchemaFiles[i] = lookupSchemaFile(schemasDir, destination.Name, destination.Version)
		}
	}

	// Partition plugin spawn list to those without a local schema file.
	sourcePluginConfigs := make([]managedplugin.Config, 0, len(sources))
	sourcePluginIdx := make([]int, 0, len(sources))
	sourceRegInferred := make([]bool, 0, len(sources))
	sourcesNeedingPlugin := make([]*specs.Source, 0, len(sources))
	for i, source := range sources {
		if sourceSchemaFiles[i] != "" {
			continue
		}
		sourcePluginConfigs = append(sourcePluginConfigs, managedplugin.Config{
			Name:       source.Name,
			Version:    source.Version,
			Path:       source.Path,
			Registry:   SpecRegistryToPlugin(source.Registry),
			DockerAuth: source.DockerRegistryAuthToken,
		})
		sourcePluginIdx = append(sourcePluginIdx, i)
		sourceRegInferred = append(sourceRegInferred, source.RegistryInferred())
		sourcesNeedingPlugin = append(sourcesNeedingPlugin, source)
	}
	destinationPluginConfigs := make([]managedplugin.Config, 0, len(destinations))
	destinationPluginIdx := make([]int, 0, len(destinations))
	destinationRegInferred := make([]bool, 0, len(destinations))
	destinationsNeedingPlugin := make([]*specs.Destination, 0, len(destinations))
	for i, destination := range destinations {
		if destinationSchemaFiles[i] != "" {
			continue
		}
		destinationPluginConfigs = append(destinationPluginConfigs, managedplugin.Config{
			Name:       destination.Name,
			Version:    destination.Version,
			Path:       destination.Path,
			Registry:   SpecRegistryToPlugin(destination.Registry),
			DockerAuth: destination.DockerRegistryAuthToken,
		})
		destinationPluginIdx = append(destinationPluginIdx, i)
		destinationRegInferred = append(destinationRegInferred, destination.RegistryInferred())
		destinationsNeedingPlugin = append(destinationsNeedingPlugin, destination)
	}

	var sourceClients, destinationClients managedplugin.Clients
	if len(sourcePluginConfigs) > 0 || len(destinationPluginConfigs) > 0 {
		authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sourcesNeedingPlugin, destinationsNeedingPlugin, nil)
		if err != nil {
			return fmt.Errorf("failed to get auth token: %w", err)
		}
		teamName, err := auth.GetTeamForToken(ctx, authToken)
		if err != nil {
			return fmt.Errorf("failed to get team name: %w", err)
		}
		opts := []managedplugin.Option{
			managedplugin.WithLogger(log.Logger),
			managedplugin.WithAuthToken(authToken.Value),
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

		sourceClients, err = managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...)
		if err != nil {
			return enrichClientError(sourceClients, sourceRegInferred, err)
		}
		defer func() {
			if err := sourceClients.Terminate(); err != nil {
				fmt.Println(err)
			}
		}()
		destinationClients, err = managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...)
		if err != nil {
			return enrichClientError(destinationClients, destinationRegInferred, err)
		}
		defer func() {
			if err := destinationClients.Terminate(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	var initErrors []error
	// File-based validation (offline; no plugin spawn).
	for i, source := range sources {
		if sourceSchemaFiles[i] == "" {
			continue
		}
		log.Info().Str("source", source.VersionString()).Str("schema", sourceSchemaFiles[i]).Msg("Validating source against local schema")
		schemaBytes, err := os.ReadFile(sourceSchemaFiles[i])
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to read schema file for source %v: %w", source.VersionString(), err))
			continue
		}
		if err := validateSpecAgainstSchema(string(schemaBytes), source.Spec); err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to validate source config %v: %w", source.VersionString(), err))
		} else {
			log.Info().Str("source", source.VersionString()).Msg("validated successfully")
		}
	}
	for i, destination := range destinations {
		if destinationSchemaFiles[i] == "" {
			continue
		}
		log.Info().Str("destination", destination.VersionString()).Str("schema", destinationSchemaFiles[i]).Msg("Validating destination against local schema")
		schemaBytes, err := os.ReadFile(destinationSchemaFiles[i])
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to read schema file for destination %v: %w", destination.VersionString(), err))
			continue
		}
		if err := validateSpecAgainstSchema(string(schemaBytes), destination.Spec); err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to validate destination config %v: %w", destination.VersionString(), err))
		} else {
			log.Info().Str("destination", destination.VersionString()).Msg("validated successfully")
		}
	}

	// Plugin-based validation for entries without a local schema file.
	for ci, client := range sourceClients {
		i := sourcePluginIdx[ci]
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
		i := destinationPluginIdx[ci]
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

// lookupSchemaFile resolves a plugin's pre-fetched schema file under dir.
// Prefers <name>@<version>.json so validation can pin to the configured plugin version,
// falling back to <name>.json when version is empty (e.g. for registry: local) or when
// only the unversioned file exists.
func lookupSchemaFile(dir, name, version string) string {
	if dir == "" {
		return ""
	}
	if version != "" {
		p := filepath.Join(dir, name+"@"+version+".json")
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	p := filepath.Join(dir, name+".json")
	if _, err := os.Stat(p); err == nil {
		return p
	}
	return ""
}
