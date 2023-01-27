package cmd

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/cloudquery/plugin-sdk/clients/discovery/v0"
	"github.com/cloudquery/plugin-sdk/registry"
	"github.com/cloudquery/plugin-sdk/specs"
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

	for _, sourceSpec := range specReader.Sources {
		if len(sourceSpec.Destinations) == 0 {
			return fmt.Errorf("no destinations found for source %s", sourceSpec.Name)
		}
		var destinationsSpecs []specs.Destination
		for _, destination := range sourceSpec.Destinations {
			spec := specReader.Destinations[destination]
			if spec == nil {
				return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
			}
			destinationsSpecs = append(destinationsSpecs, *spec)
		}

		discoveryClient, err := discovery.NewClient(ctx, sourceSpec.Registry, registry.PluginTypeSource, sourceSpec.Path, sourceSpec.Version, discovery.WithDirectory(cqDir))
		if err != nil {
			return fmt.Errorf("failed to create discovery client for source %s: %w", sourceSpec.Name, err)
		}

		versions, err := discoveryClient.GetVersions(ctx)
		if err != nil {
			if discoveryErr := discoveryClient.Terminate(); discoveryErr != nil {
				log.Error().Err(discoveryErr).Msg("failed to terminate discovery client")
				fmt.Println("failed to terminate discovery client:", discoveryErr)
			}
			// If we get an error here, we assume that the plugin is not a v1 plugin and we try to sync it as a v0 plugin
			if err := syncConnectionV0(ctx, cqDir, *sourceSpec, destinationsSpecs, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
			}
			continue
		}
		if err := discoveryClient.Terminate(); err != nil {
			return fmt.Errorf("failed to terminate discovery client: %w", err)
		}

		if slices.Index(versions, "v1") != -1 {
			if err := syncConnectionV1(ctx, cqDir, *sourceSpec, destinationsSpecs, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v1 source %s: %w", sourceSpec.Name, err)
			}
			continue
		}

		if slices.Index(versions, "v0") != -1 {
			if err := syncConnectionV0(ctx, cqDir, *sourceSpec, destinationsSpecs, invocationUUID.String(), noMigrate); err != nil {
				return fmt.Errorf("failed to sync v0 source %s: %w", sourceSpec.Name, err)
			}
			continue
		}

		return fmt.Errorf("failed to sync source %s, unknown versions %v", sourceSpec.Name, versions)
	}

	return nil
}
