package cmd

import (
	"errors"
	"fmt"
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

	return cmd
}

func validateConfig(cmd *cobra.Command, args []string) error {
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

	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, destinations, nil)
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

	sourcePluginConfigs := make([]managedplugin.Config, len(sources))
	sourceRegInferred := make([]bool, len(sources))
	for i, source := range sources {
		sourcePluginConfigs[i] = managedplugin.Config{
			Name:       source.Name,
			Version:    source.Version,
			Path:       source.Path,
			Registry:   SpecRegistryToPlugin(source.Registry),
			DockerAuth: source.DockerRegistryAuthToken,
		}
		sourceRegInferred[i] = source.RegistryInferred()
	}
	destinationPluginConfigs := make([]managedplugin.Config, len(destinations))
	destinationRegInferred := make([]bool, len(destinations))
	for i, destination := range destinations {
		destinationPluginConfigs[i] = managedplugin.Config{
			Name:       destination.Name,
			Version:    destination.Version,
			Path:       destination.Path,
			Registry:   SpecRegistryToPlugin(destination.Registry),
			DockerAuth: destination.DockerRegistryAuthToken,
		}
		destinationRegInferred[i] = destination.RegistryInferred()
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

	var initErrors []error
	for i, client := range sourceClients {
		pluginClient := plugin.NewPluginClient(client.Conn)
		log.Info().Str("source", sources[i].VersionString()).Msg("Initializing source")
		err := validatePluginSpec(ctx, pluginClient, sources[i].Spec)
		if err != nil {
			initErrors = append(initErrors, fmt.Errorf("failed to validate source config %v: %w", sources[i].VersionString(), err))
		} else {
			log.Info().Str("source", sources[i].VersionString()).Msg("validated successfully")
		}
	}
	for i, client := range destinationClients {
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
