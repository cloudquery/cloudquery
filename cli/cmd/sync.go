package cmd

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/manageddestination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/managedsource"
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
	sourceOpts := []managedsource.Option{
		managedsource.WithLogger(log.Logger),
	}
	destinationOpts := []manageddestination.Option{
		manageddestination.WithLogger(log.Logger),
	}
	if cqDir != "" {
		sourceOpts = append(sourceOpts, managedsource.WithDirectory(cqDir))
		destinationOpts = append(destinationOpts, manageddestination.WithDirectory(cqDir))
	}
	if disableSentry {
		sourceOpts = append(sourceOpts, managedsource.WithNoSentry())
		destinationOpts = append(destinationOpts, manageddestination.WithNoSentry())
	}

	sourcesClients, err := managedsource.NewClients(ctx, sources, sourceOpts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := sourcesClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()
	destinationsClients, err := manageddestination.NewClients(ctx, destinations, destinationOpts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := destinationsClients.Terminate(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, cl := range sourcesClients {
		maxVersion, err := cl.MaxVersion(ctx)
		if err != nil {
			return err
		}
		switch maxVersion {
		case 2:
			for _, destination := range destinationsClients {
				versions, err := destination.Versions(ctx)
				if err != nil {
					return fmt.Errorf("failed to get destination versions%w", err)
				}
				if !slices.Contains(versions, 2) {
					return fmt.Errorf("destination %s does not support version 2. please upgrade destination", destination.Spec.Name)
				}
			}
			if err := syncConnectionV2(ctx, cl, destinationsClients, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v2 source %s: %w", cl.Spec.Name, err)
			}
		case 1:
			if err := syncConnectionV1(ctx, cl, destinationsClients, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Spec.Name, err)
			}
		case 0:
			if err := syncConnectionV0_2(ctx, cl, destinationsClients, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", cl.Spec.Name, err)
			}
		case -1:
			return fmt.Errorf("please upgrade your source or use an older CLI version < v3.0.1")
		default:
			return fmt.Errorf("unknown source version %d", maxVersion)
		}
	}

	return nil
}
