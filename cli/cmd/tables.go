package cmd

import (
	"fmt"
	"path"
	"strings"

	source "github.com/cloudquery/plugin-sdk/clients/source/v1"
	"github.com/cloudquery/plugin-sdk/specs"
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

	for _, sourceSpec := range specReader.Sources {
		opts := []source.ClientOption{
			source.WithLogger(log.Logger),
			source.WithDirectory(cqDir),
		}
		sourceClient, err := source.NewClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version, opts...)
		if err != nil {
			return fmt.Errorf("failed to create source client. Error: %w", err)
		}

		outputPath := path.Join(outputDir, sourceSpec.Name)
		fmt.Printf("Generating docs for: %s to %s\n", sourceSpec.VersionString(), outputPath)
		err = sourceClient.GenDocs(ctx, outputPath, format)
		if err != nil {
			return fmt.Errorf("failed to generate docs. Error: %w", err)
		}
	}

	return nil
}
