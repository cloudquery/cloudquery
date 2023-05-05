package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/destination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	migrateShort   = "Run migration for source and destination plugins specified in configuration"
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
	return cmd
}

func migrate(cmd *cobra.Command, args []string) error {
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
	var sourceOpts []source.PluginOption
	var destinationOpts []destination.PluginOption
	if cqDir != "" {
		sourceOpts = append(sourceOpts, source.WithDirectory(cqDir))
		destinationOpts = append(destinationOpts, destination.WithDirectory(cqDir))
	}

	sourcesClients, err := source.NewClients(ctx, sources, sourceOpts...)
	if err != nil {
		return err
	}
	defer sourcesClients.Terminate()
	destinationsClients, err := destination.NewClients(ctx, destinations, destinationOpts...)
	if err != nil {
		return err
	}
	defer destinationsClients.Terminate()


	for _, cl := range sourcesClients {
		maxVersion, err := cl.MaxVersion(ctx)
		if err != nil {
			return err
		}
		destinationsForSource := specReader.GetDestinationNamesForSource(cl.Spec.Name)
		destinationsClientsForSource := destinationsClients.ClientsByNames(destinationsForSource)
		switch maxVersion {
			case 1:
				return migrateConnectionV1(ctx, cl, destinationsClientsForSource)
			case 0:
				return migrateConnectionV0(ctx, cl, destinationsClientsForSource)
			default:
				return fmt.Errorf("unknown version %d", maxVersion)
		}
	}
	return nil
}
