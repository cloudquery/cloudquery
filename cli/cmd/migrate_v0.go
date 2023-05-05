package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/manageddestination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/managedsource"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v0"
	"github.com/rs/zerolog/log"
)

func migrateConnectionV0(ctx context.Context, sourceClient *managedsource.Client, destinationsClients manageddestination.Clients) error {
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

	tablesBytes, err := getTablesForSpec(ctx, sourcePbClient, specBytes)
	if err != nil {
		return err
	}

	fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
	for i := range destinationsClients {
		if _, err := destinationsPbClients[i].Migrate(ctx, &pbdestination.Migrate_Request{
			Tables: tablesBytes,
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
