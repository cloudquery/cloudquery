package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	tablesShort   = "Generate documentation for all supported tables of source plugins specified in the spec(s)"
	tablesExample = `# Generate documentation for all supported tables of source plugins specified in the spec(s) 
cloudquery tables ./directory
# The default format is JSON, you can override it with --format
cloudquery tables ./directory --format markdown
# You can also specify an output directory. The default is ./cq-docs
cloudquery tables ./directory --output-dir ./docs
# You can also filter which tables are included in the output. The default is all, use --filter=spec to include only tables referenced in the spec
cloudquery tables ./directory --filter spec
`
)

func NewCmdTables() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tables [files or directories]",
		Short:   tablesShort,
		Long:    tablesShort,
		Example: tablesExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    tables,
	}
	cmd.Flags().String("output-dir", "cq-docs", "Base output directory for generated files")
	cmd.Flags().String("format", "json", "Output format. One of: json, markdown")
	cmd.Flags().String("filter", "all", "Filter tables. One of: all, spec")
	return cmd
}

func tables(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}
	outputDir, err := cmd.Flags().GetString("output-dir")
	if err != nil {
		return err
	}
	filter, err := cmd.Flags().GetString("filter")
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
	authToken, err := auth.GetAuthTokenIfNeeded(log.Logger, sources, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := auth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return fmt.Errorf("failed to get team name: %w", err)
	}

	pluginVersionWarner, _ := managedplugin.NewPluginVersionWarner(log.Logger, authToken.Value)
	specs.WarnOnOutdatedVersions(ctx, pluginVersionWarner, sources, nil, nil)

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
	pluginConfigs := make([]managedplugin.Config, len(sources))
	sourceRegInferred := make([]bool, len(sources))
	for i, sourceSpec := range sources {
		pluginConfigs[i] = managedplugin.Config{
			Name:       sourceSpec.Name,
			Path:       sourceSpec.Path,
			Version:    sourceSpec.Version,
			Registry:   SpecRegistryToPlugin(sourceSpec.Registry),
			DockerAuth: sourceSpec.DockerRegistryAuthToken,
		}
		sourceRegInferred[i] = sourceSpec.Registry == specs.RegistryUnset
	}

	sourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, pluginConfigs, opts...)
	if err != nil {
		return enrichClientError(sourceClients, sourceRegInferred, err)
	}
	defer func() {
		if err := sourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	for _, source := range specReader.Sources {
		cl := sourceClients.ClientByName(source.Name)
		outputPath := path.Join(outputDir, source.Name)
		fmt.Printf("Generating docs for %q to directory %q\n", source.Name, outputPath)
		versions, err := cl.Versions(ctx)
		if err != nil {
			return fmt.Errorf("failed to get versions for %s. Error: %w", source.Name, err)
		}
		maxVersion := findMaxCommonVersion(versions, []int{2, 3})
		switch maxVersion {
		case 3:
			if err := tablesV3(ctx, cl, source, outputPath, format, filter); err != nil {
				return err
			}
			fmt.Printf("Done generating docs for %q to directory %q\n", source.Name, outputPath)
		case 2:
			if err := tablesV2(ctx, cl, outputPath, format); err != nil {
				return err
			}
			fmt.Printf("Done generating docs for %q to directory %q\n", source.Name, outputPath)
		default:
			return fmt.Errorf("unsupported version %d for %s", maxVersion, source.Name)
		}
	}

	return nil
}
