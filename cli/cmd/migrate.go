package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
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
		if err := migrateConnection(ctx, cqDir, *sourceSpec, destinationsSpecs); err != nil {
			return fmt.Errorf("failed to migrate source %s: %w", sourceSpec.Name, err)
		}
	}

	return nil
}

func migrateConnection(ctx context.Context, cqDir string, sourceSpec specs.Source, destinationsSpecs []specs.Destination) error {
	destinationNames := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationNames[i] = destinationsSpecs[i].Name
	}
	migrateTime := time.Now().UTC()

	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("migrate_time", migrateTime).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("migrate_time", migrateTime).Msg("End migration")

	sourceClient, err := clients.NewSourceClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version,
		clients.WithSourceLogger(log.Logger),
		clients.WithSourceDirectory(cqDir),
	)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer func() {
		if err := sourceClient.Terminate(); err != nil {
			log.Error().Err(err).Msg("Failed to terminate source client")
			fmt.Println("failed to terminate source client:", err)
		}
	}()

	destClients, err := newDestinationClients(ctx, sourceSpec, destinationsSpecs, cqDir)
	if err != nil {
		return err
	}
	defer destClients.Close()

	selectedTables, tablesForSpecSupported, err := getTablesForSpec(ctx, sourceClient, sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
	}
	tableCount := len(selectedTables.FlattenTables())

	// Print a count of the tables that will be migrated.
	if tablesForSpecSupported {
		word := "tables"
		if tableCount == 1 {
			word = "table"
		}
		fmt.Printf("Source %s will migrate %d %s.\n", sourceSpec.Name, tableCount, word)
	}

	fmt.Println("Starting migration for:", sourceSpec.Name, "->", sourceSpec.Destinations)
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Starting migration")

	for i, destinationSpec := range destinationsSpecs {
		if err := destClients[i].Migrate(ctx, selectedTables); err != nil {
			return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destinationSpec.Name, err)
		}
	}
	tt := time.Since(migrateTime)
	fmt.Printf("Migration completed successfully.\n")
	log.Info().Str("source", sourceSpec.Name).
		Strs("destinations", sourceSpec.Destinations).
		Int("num_tables", tableCount).
		Float64("time_took", tt.Seconds()).
		Msg("Migration completed successfully")

	return nil
}
