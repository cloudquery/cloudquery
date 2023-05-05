package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v1"
	"github.com/cloudquery/plugin-pb-go/specs"
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
	opts := []source.PluginOption{
		source.WithLogger(log.Logger),
		source.WithDirectory(cqDir),
	}
	sourceClients, err := source.NewClients(ctx, specReader.Sources, opts...)
	if err != nil {
		return err
	}
	defer sourceClients.Terminate()
	for _, sourceClient := range sourceClients {
		outputPath := path.Join(outputDir, sourceClient.Spec.Name)
		pbSourceClient := pbSource.NewSourceClient(sourceClient.Conn)
		if _, err := pbSourceClient.GenDocs(ctx, &pbSource.GenDocs_Request{
			Format:    pbSource.GenDocs_FORMAT(pbSource.GenDocs_FORMAT_value[format]),
			Path: outputPath,
		}); err != nil {
			return err
		}
	}

	return nil
}
