package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/plugin/destination"
	"github.com/cloudquery/cloudquery/cli/internal/plugin/source"
	"github.com/cloudquery/plugin-pb-go/metrics"
	pbdestination "github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v1"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)



func syncConnectionV1(ctx context.Context, sourceClient *source.Client, destinationsClients destination.Clients, uid string, noMigrate bool) error {
	syncTime := time.Now().UTC()
	sourceSpec := sourceClient.Spec
	destinationStrings := destinationsClients.Names()
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

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

	if !noMigrate {
		migrateStart := time.Now().UTC()
		fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
		for i := range destinationsClients {
			if _, err := destinationsPbClients[i].Migrate(ctx, &pbdestination.Migrate_Request{
				Tables: tablesRes.Tables,
			}); err != nil {
				return err
			}
		}
		migrateTimeTook := time.Since(migrateStart)
		fmt.Printf("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.VersionString()).
			Strs("destinations", destinationStrings).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncClient, err := sourcePbClient.Sync(ctx, &pbSource.Sync_Request{})
	if err != nil {
		return err
	}
	writeClients := make([]pbdestination.Destination_Write2Client, len(destinationsPbClients))
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write2(ctx)
		if err != nil {
			return err
		}
		writeClients[i].Send(&pbdestination.Write2_Request{
			Source: sourceClient.Spec.Name,
			Tables: tablesRes.Tables,
			Timestamp: timestamppb.New(syncTime),
		})
	}
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)
	for {
		r, err := syncClient.Recv()
		if err == io.EOF {
			return nil
		}
		_ = bar.Add(1)
		for i := range destinationsPbClients {
			if err := writeClients[i].Send(&pbdestination.Write2_Request{
				Resource: r.Resource,
			}); err != nil {
				return err
			}
		}
	}
	getMetricsRes, err := sourcePbClient.GetMetrics(ctx, &pbSource.GetMetrics_Request{})
	if err != nil {
		return err
	}
	var m metrics.Metrics
	if err := json.Unmarshal(getMetricsRes.Metrics, &m); err != nil {
		return err
	}

	syncTimeTook := time.Since(syncTime)
	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", m.TotalResources(), m.TotalErrors(), m.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	return nil
}

