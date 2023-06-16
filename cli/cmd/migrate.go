package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/manageddestination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/managedsource"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
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
	var sourceOpts []managedsource.Option
	var destinationOpts []manageddestination.Option
	if cqDir != "" {
		sourceOpts = append(sourceOpts, managedsource.WithDirectory(cqDir))
		destinationOpts = append(destinationOpts, manageddestination.WithDirectory(cqDir))
	}

	managedSourceClients, err := managedsource.NewClients(ctx, sources, sourceOpts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := managedSourceClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	managedDestinationsClients, err := manageddestination.NewClients(ctx, destinations, destinationOpts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := managedDestinationsClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, cl := range managedSourceClients {
		maxVersion, err := cl.MaxVersion(ctx)
		if err != nil {
			return err
		}
		destinationsForSource := specReader.GetDestinationNamesForSource(cl.Spec.Name)
		destinationsClientsForSource := managedDestinationsClients.ClientsByNames(destinationsForSource)
		switch maxVersion {
		case 2:
			for _, destination := range destinationsClientsForSource {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions: %w", err)
				}
				if !slices.Contains(versions, 1) {
					return fmt.Errorf("destination %[1]s does not support CloudQuery SDK version 1. Please upgrade to newer version of %[1]s", destination.Spec.Name)
				}
			}
			return migrateConnectionV2(ctx, cl, destinationsClientsForSource)
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
