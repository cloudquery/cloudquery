package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/destination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v0"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/rs/zerolog/log"
)

func migrateConnectionV0(ctx context.Context, sourceClient source.Client, destinationsClients destination.Clients, uid string, noMigrate bool) error {
	sourceSpec := sourceClient.Spec
	destinationStrings := destinationsClients.Names()
	migrateStart := time.Now().UTC()
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("End migration")

	sourcePbClient := pbSource.NewSourceClient(sourceClient.Conn)
	destinationsPbClients := make([]pbdestination.DestinationClient, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = pbdestination.NewDestinationClient(destinationsClients[i].Conn)
	}
	specBytes, err := json.Marshal(sourceClient.Spec)
	if err != nil {
		return err
	}

	tablesRes, err := sourcePbClient.GetDynamicTables(ctx, &pbSource.GetDynamicTables_Request{})
	if err != nil {
		return err
	}

	fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
	for i := range destinationsClients {
		if _, err := destinationsPbClients[i].Migrate(ctx, &pbdestination.Migrate_Request{
			Tables: tablesRes.Tables,
		}); err != nil {
			return err
		}
	}
	migrateTimeTook := time.Since(migrateStart)
	fmt.Println("Migration completed successfully.")
	log.Info().Str("source", sourceSpec.Name).
		Strs("destinations", sourceSpec.Destinations).
		Float64("time_took", migrateTimeTook.Seconds()).
		Msg("Migration completed successfully")
	return nil
}

func migrateConnectionV0Old(ctx context.Context, cqDir string, sourceSpec specs.Source, destinationsSpecs []specs.Destination) error {
	destinationNames := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationNames[i] = destinationsSpecs[i].Name
	}
	migrateTime := time.Now().UTC()

	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("migrate_time", migrateTime).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("migrate_time", migrateTime).Msg("End migration")

	sourceClient, err := source.NewClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version,
		source.WithLogger(log.Logger),
		source.WithDirectory(cqDir),
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

	destClients, err := newDestinationClientsV0(ctx, sourceSpec, destinationsSpecs, cqDir)
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
