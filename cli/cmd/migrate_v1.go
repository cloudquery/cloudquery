package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/destination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v1"
	"github.com/rs/zerolog/log"
)


func migrateConnectionV1(ctx context.Context, sourceClient source.Client, destinationsClients destination.Clients, uid string, noMigrate bool) error {
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
	if _, err := sourcePbClient.Init(ctx, &pbSource.Init_Request{
		Spec: specBytes,
	}); err != nil {
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