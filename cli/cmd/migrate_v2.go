package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/manageddestination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/managedsource"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v1"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v2"
	"github.com/rs/zerolog/log"
)

func migrateConnectionV2(ctx context.Context, sourceClient *managedsource.Client, managedDestinationsClients manageddestination.Clients) error {
	sourceSpec := sourceClient.Spec
	destinationStrings := managedDestinationsClients.Names()
	migrateStart := time.Now().UTC()
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("End migration")

	sourcePbClient := pbSource.NewSourceClient(sourceClient.Conn)
	destinationsPbClients := make([]pbdestination.DestinationClient, len(managedDestinationsClients))
	for i := range managedDestinationsClients {
		destinationsPbClients[i] = pbdestination.NewDestinationClient(managedDestinationsClients[i].Conn)
	}
	specBytes, err := json.Marshal(sourceClient.Spec)
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

	fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
	for i := range managedDestinationsClients {
		destSpecBytes, err := json.Marshal(managedDestinationsClients[i].Spec)
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Configure(ctx, &pbdestination.Configure_Request{
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