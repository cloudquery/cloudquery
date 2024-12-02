package cmd

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	migrateShort   = "Update schema of your destinations based on the latest changes in sources from your configuration"
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
	cmd.Flags().String("license", "", "set offline license file")
	cmd.Flags().Bool("cq-columns-not-null", false, "Force CloudQuery internal columns to be NOT NULL. This feature is in Preview. Please provide feedback to help us improve it.")
	_ = cmd.Flags().MarkHidden("cq-columns-not-null")
	return cmd
}

func migrate(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	licenseFile, err := cmd.Flags().GetString("license")
	if err != nil {
		return err
	}

	cqColumnsNotNull, err := cmd.Flags().GetBool("cq-columns-not-null")
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
	transformers := specReader.Transformers

	transformerSpecsByName := make(map[string]specs.Transformer)
	for _, transformer := range transformers {
		transformerSpecsByName[transformer.Name] = *transformer
	}

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

	managedSourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, sourcePluginConfigs, opts...)
	if err != nil {
		return enrichClientError(managedSourceClients, sourceRegInferred, err)
	}
	defer func() {
		if err := managedSourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationPluginClients, err := managedplugin.NewClients(ctx, managedplugin.PluginDestination, destinationPluginConfigs, opts...)
	if err != nil {
		return enrichClientError(destinationPluginClients, destinationRegInferred, err)
	}
	defer func() {
		if err := destinationPluginClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	transformerPluginClients, err := managedplugin.NewClients(ctx, managedplugin.PluginTransformer, transformerPluginConfigs, opts...)
	if err != nil {
		return enrichClientError(transformerPluginClients, transformerRegInferred, err)
	}
	defer func() {
		if err := transformerPluginClients.Terminate(); err != nil {
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
		transformersForDestination := make(map[string][]*managedplugin.Client)
		for _, destination := range destinations {
			if slices.Contains(source.Destinations, destination.Name) {
				destinationClientsForSource = append(destinationClientsForSource, destinationPluginClients.ClientByName(destination.Name))
				destinationForSourceSpec = append(destinationForSourceSpec, *destination)
				for _, transformerName := range destination.Transformers {
					transformersForDestination[destination.Name] = append(transformersForDestination[destination.Name], transformerPluginClients.ClientByName(transformerName))
				}
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
					return fmt.Errorf("destination plugin %[1]s does not support CloudQuery protocol version 3, required by the %[2]s source plugin. Please upgrade to a newer version of the %[1]s destination plugin", destination.Name(), source.Name)
				}
			}

			migrateOptions := migrateV3Options{
				sourceClient:               cl,
				destinationsClients:        destinationClientsForSource,
				sourceSpec:                 *source,
				destinationSpecs:           destinationForSourceSpec,
				transformersForDestination: transformersForDestination,
				transformerSpecsByName:     transformerSpecsByName,
				cqColumnsNotNull:           cqColumnsNotNull,
			}
			if err := migrateConnectionV3(ctx, migrateOptions); err != nil {
				return fmt.Errorf("failed to migrate v3 source %s: %w", cl.Name(), err)
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
			if err := migrateConnectionV2(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec, destinationsVersions); err != nil {
				return fmt.Errorf("failed to migrate source %v@%v: %w", source.Name, source.Version, err)
			}
		case 1:
			if err := migrateConnectionV1(ctx, cl, destinationClientsForSource, *source, destinationForSourceSpec); err != nil {
				return fmt.Errorf("failed to migrate source %v@%v: %w", source.Name, source.Version, err)
			}
		case 0:
			return errors.New("please upgrade your source or use a CLI version between v3.0.1 and v3.5.3")
		case -1:
			return fmt.Errorf("please upgrade CLI to sync source %v@%v", source.Name, source.Version)
		case -2:
			return fmt.Errorf("please downgrade CLI or upgrade source to sync %v", source.Name)
		}
	}
	return nil
}
