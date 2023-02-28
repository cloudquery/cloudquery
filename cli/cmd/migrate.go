package cmd

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/clients/discovery/v0"
	"github.com/cloudquery/plugin-sdk/registry"
	"github.com/cloudquery/plugin-sdk/specs"
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

	for _, sourceSpec := range specReader.Sources {
		if len(sourceSpec.Destinations) == 0 {
			return fmt.Errorf("no destinations found for source %s", sourceSpec.Name)
		}
		var destinationsSpecs []specs.Destination
		for _, destination := range sourceSpec.Destinations {
			spec := specReader.GetDestinationByName(destination)
			if spec == nil {
				return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
			}
			destinationsSpecs = append(destinationsSpecs, *spec)
		}
		discoveryClient, err := discovery.NewClient(ctx, sourceSpec.Registry, registry.PluginTypeSource, sourceSpec.Path, sourceSpec.Version)
		if err != nil {
			return fmt.Errorf("failed to create discovery client for source %s: %w", sourceSpec.Name, err)
		}
		versions, err := discoveryClient.GetVersions(ctx)
		if err != nil {
			if discoveryErr := discoveryClient.Terminate(); discoveryErr != nil {
				log.Error().Err(discoveryErr).Msg("failed to terminate discovery client")
				fmt.Println("failed to terminate discovery client:", discoveryErr)
			}
			if err := migrateConnectionV0(ctx, cqDir, *sourceSpec, destinationsSpecs); err != nil {
				return fmt.Errorf("failed to migrate source %s: %w", sourceSpec.Name, err)
			}
			continue
		}

		if err := discoveryClient.Terminate(); err != nil {
			return fmt.Errorf("failed to terminate discovery client: %w", err)
		}

		if slices.Index(versions, "v1") != -1 {
			if err := migrateConnectionV1(ctx, cqDir, *sourceSpec, destinationsSpecs); err != nil {
				return fmt.Errorf("failed to migrate source %s: %w", sourceSpec.Name, err)
			}
			continue
		}

		if slices.Index(versions, "v0") != -1 {
			if err := migrateConnectionV0(ctx, cqDir, *sourceSpec, destinationsSpecs); err != nil {
				return fmt.Errorf("failed to migrate source %s: %w", sourceSpec.Name, err)
			}
			continue
		}

		return fmt.Errorf("failed to migrate source %s, unknown versions %v", sourceSpec.Name, versions)
	}

	return nil
}
