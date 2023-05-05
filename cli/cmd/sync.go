package cmd

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/destination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	syncShort   = "Sync resources from configured source plugins to destinations"
	syncExample = `# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml
`
	unknownFieldErrorPrefix = "code = InvalidArgument desc = failed to decode spec: json: unknown field "
)

func NewCmdSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [files or directories]",
		Short:   syncShort,
		Long:    syncShort,
		Example: syncExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    sync,
	}
	cmd.Flags().Bool("no-migrate", false, "Disable auto-migration before sync. By default, sync runs a migration before syncing resources.")
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	cqDir, err := cmd.Flags().GetString("cq-dir")
	if err != nil {
		return err
	}

	noMigrate, err := cmd.Flags().GetBool("no-migrate")
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

	invocationUUID, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate invocation uuid: %w", err)
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
		switch maxVersion {
		case 1:
			if err := syncConnectionV1(ctx, cl, destinationsClients, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Spec.Name, err)
			}
		case 0:
			if err := syncConnectionV0_2(ctx, cl, destinationsClients, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Spec.Name, err)
			}
		case -1:
		default:
			return fmt.Errorf("unknown source version %d", maxVersion)
		}
	}


	
	return nil
}
