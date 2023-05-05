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
	pbSource "github.com/cloudquery/plugin-pb-go/pb/source/v0"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func syncConnectionV0_2(ctx context.Context, sourceClient *source.Client, destinationsClients destination.Clients, uid string, noMigrate bool) error {
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

	tablesBytes, err := getTablesForSpec(ctx, sourcePbClient, specBytes)
	if err != nil {
		return err
	}

	if !noMigrate {
		migrateStart := time.Now().UTC()
		fmt.Printf("Starting migration with for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
		for i := range destinationsClients {
			if _, err := destinationsPbClients[i].Migrate(ctx, &pbdestination.Migrate_Request{
				Tables: tablesBytes,
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

	syncClient, err := sourcePbClient.Sync2(ctx, &pbSource.Sync2_Request{})
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
			Tables: tablesBytes,
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
	getMetricsRes, err := sourcePbClient.GetMetrics(ctx, &pbSource.GetSourceMetrics_Request{})
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



// getTablesForSpec first tries the newer GetTablesForSpec call, but if it is not available, falls back to
// GetTables. The returned `supported` value indicates whether GetTablesForSpec was supported by the server.
func getTablesForSpec(ctx context.Context, sourceClient pbSource.SourceClient, specSourceBytes []byte) (tables []byte, err error) {
	getTablesForSpecRes, err := sourceClient.GetTablesForSpec(ctx, &pbSource.GetTablesForSpec_Request{
		Spec: specSourceBytes,
	})
	if isUnimplemented(err) {
		// the plugin server does not support GetTablesForSpec. Fall back to GetTables.
		getTablesRes, err := sourceClient.GetTables(ctx, &pbSource.GetTables_Request{})
		if err != nil {
			return getTablesRes.Tables, fmt.Errorf("failed to call GetTables: %w", err)
		}
		return tables, nil
	} else if err != nil {
		// the method is supported, but failed for some other reason
		return nil, fmt.Errorf("failed to call GetTablesForSpec: %w", err)
	}

	return getTablesForSpecRes.Tables, nil
}