package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	pbBase "github.com/cloudquery/plugin-pb-go/pb/base/v0"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v1"
	"github.com/rs/zerolog/log"
)

func migrateConnectionV1(ctx context.Context, sourceClient *managedplugin.Client, managedDestinationsClients managedplugin.Clients, sourceSpec specs.Source, destinationSpecs []specs.Destination) error {
	migrateStart := time.Now().UTC()
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Time("migrate_time", migrateStart).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Time("migrate_time", migrateStart).Msg("End migration")

	sourcePbClient := pbSource.NewSourceClient(sourceClient.Conn)
	destinationsPbClients := make([]pbdestination.DestinationClient, len(managedDestinationsClients))
	for i := range managedDestinationsClients {
		destinationsPbClients[i] = pbdestination.NewDestinationClient(managedDestinationsClients[i].Conn)
	}
	specBytes, err := json.Marshal(CLISourceSpecToPbSpec(sourceSpec))
	if err != nil {
		return err
	}
	if _, err := sourcePbClient.Init(ctx, &pbSource.Init_Request{
		Spec: specBytes,
	}); err != nil {
		return fmt.Errorf("failed to Init source: %w", err)
	}
	tablesRes, err := sourcePbClient.GetDynamicTables(ctx, &pbSource.GetDynamicTables_Request{})
	if err != nil {
		return fmt.Errorf("failed to GetDynamicTables: %w", err)
	}

	fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), sourceSpec.Destinations)
	for i := range managedDestinationsClients {
		destSpecBytes, err := json.Marshal(CLIDestinationSpecToPbSpec(destinationSpecs[i]))
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Configure(ctx, &pbBase.Configure_Request{
			Config: destSpecBytes,
		}); err != nil {
			return fmt.Errorf("failed to call Migrate: %w", err)
		}
		if _, err := destinationsPbClients[i].Migrate(ctx, &pbdestination.Migrate_Request{
			Tables: tablesRes.Tables,
		}); err != nil {
			return fmt.Errorf("failed to call Migrate: %w", err)
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
