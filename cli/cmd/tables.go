package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
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
	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}
	opts := []managedplugin.Option{
		managedplugin.WithLogger(log.Logger),
		managedplugin.WithDirectory(cqDir),
	}
	pluginConfigs := make([]managedplugin.Config, 0, len(specReader.Sources))
	for _, sourceSpec := range specReader.Sources {
		pluginConfigs = append(pluginConfigs, managedplugin.Config{
			Name:     sourceSpec.Name,
			Path:     sourceSpec.Path,
			Version:  sourceSpec.Version,
			Registry: SpecRegistryToPlugin(sourceSpec.Registry),
		})
	}

	sourceClients, err := managedplugin.NewClients(ctx, managedplugin.PluginSource, pluginConfigs, opts...)
	if err != nil {
		return err
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
			if err := tablesV3(ctx, cl, outputPath, format); err != nil {
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
