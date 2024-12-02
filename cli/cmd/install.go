package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	pluginInstallShort   = "Install required plugin images from your configuration"
	pluginInstallExample = `# Install required plugins specified in directory
cloudquery plugin install ./directory
# Install required plugins specified in directory and config files
cloudquery plugin install ./directory ./aws.yml ./pg.yml
`
)

func newCmdPluginInstall(deprecated bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install [files or directories]",
		Short:   pluginInstallShort,
		Long:    pluginInstallShort,
		Example: pluginInstallExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    installPlugin,
	}
	if deprecated {
		cmd.Deprecated = "use `cloudquery plugin install` instead"
	}
	return cmd
}

func installPlugin(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewRelaxedSpecReader(args)
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
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name: %w", err)
	}

	pluginVersionWarner, _ := managedplugin.NewPluginVersionWarner(log.Logger, authToken.Value)
	specs.WarnOnOutdatedVersions(ctx, pluginVersionWarner, sources, destinations, transformers)

	opts := []managedplugin.Option{
		managedplugin.WithNoExec(),
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

	transformerPluginConfigs := make([]managedplugin.Config, len(transformers))
	transformerRegInferred := make([]bool, len(transformers))
	for i, transformer := range transformers {
		transformerPluginConfigs[i] = managedplugin.Config{
			Name:       transformer.Name,
			Version:    transformer.Version,
			Path:       transformer.Path,
			Registry:   SpecRegistryToPlugin(transformer.Registry),
			DockerAuth: transformer.DockerRegistryAuthToken,
		}
		transformerRegInferred[i] = transformer.RegistryInferred()
	}

	if clist, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...); err != nil {
		return enrichClientError(clist, sourceRegInferred, err)
	}
	if clist, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...); err != nil {
		return enrichClientError(clist, destinationRegInferred, err)
	}
	if clist, err := managedplugin.NewClients(ctx, managedplugin.PluginTransformer, transformerPluginConfigs, opts...); err != nil {
		return enrichClientError(clist, transformerRegInferred, err)
	}

	return nil
}

// enrichClientError tries to add Hint messages to errors from managed client.
// It will also do an inferred-registry check on the failed client (which is one more than the last one on the list) and add the registry hint if needed.
func enrichClientError(clientsList managedplugin.Clients, inferredList []bool, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, managedplugin.ErrLoginRequired) {
		return fmt.Errorf("%w. Hint: You must be logged in via `cloudquery login` or you must use a valid API Key which can be generated at `cloud.cloudquery.io`", err)
	}
	if errors.Is(err, managedplugin.ErrTeamRequired) {
		return fmt.Errorf("%w. Hint: use `cloudquery switch` to set a team", err)
	}

	if !strings.Contains(strings.ToLower(err.Error()), "not found") {
		return err
	}
	l := len(clientsList)
	il := len(inferredList)
	if l > il {
		return err // shouldn't happen
	}
	if !inferredList[l] {
		return err
	}

	return fmt.Errorf("%w. Hint: make sure to use the latest plugin version from hub.cloudquery.io or to keep using an outdated version add `registry: github` to your configuration", err)
}
