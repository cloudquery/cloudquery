package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/base/v0"
	"github.com/cloudquery/plugin-pb-go/pb/destination/v0"
	"github.com/cloudquery/plugin-pb-go/pb/source/v1"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ExitReason string

// nolint:dupl
func syncConnectionV1(ctx context.Context, sourceClient *managedplugin.Client, destinationsClients managedplugin.Clients, sourceSpec specs.Source, destinationSpecs []specs.Destination, uid string, noMigrate bool) error {
	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
	defer func() {
		if oldAnalyticsClient != nil {
			log.Info().Msg("Sending sync summary to " + oldAnalyticsClient.Host())
			if err := oldAnalyticsClient.SendSyncMetrics(context.Background(), sourceSpec, destinationSpecs, uid, &mt, exitReason); err != nil {
				log.Warn().Err(err).Msg("Failed to send sync summary")
			}
		}
	}()
	// https://github.com/golang/go/issues/41087
	syncTime := time.Now().UTC().Truncate(time.Microsecond)
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	sourcePbClient := source.NewSourceClient(sourceClient.Conn)
	destinationsPbClients := make([]destination.DestinationClient, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = destination.NewDestinationClient(destinationsClients[i].Conn)
	}

	specBytes, err := json.Marshal(CLISourceSpecToPbSpec(sourceSpec))
	if err != nil {
		return err
	}
	if _, err := sourcePbClient.Init(ctx, &source.Init_Request{
		Spec: specBytes,
	}); err != nil {
		return err
	}
	tablesRes, err := sourcePbClient.GetDynamicTables(ctx, &source.GetDynamicTables_Request{})
	if err != nil {
		return err
	}
	for i := range destinationsClients {
		destSpecBytes, err := json.Marshal(CLIDestinationSpecToPbSpec(destinationSpecs[i]))
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Configure(ctx, &base.Configure_Request{
			Config: destSpecBytes,
		}); err != nil {
			return err
		}
	}

	if !noMigrate {
		migrateStart := time.Now().UTC()
		fmt.Printf("Starting migration for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
		for i := range destinationsClients {
			if _, err := destinationsPbClients[i].Migrate(ctx, &destination.Migrate_Request{
				Tables: tablesRes.Tables,
			}); err != nil {
				return err
			}
		}
		migrateTimeTook := time.Since(migrateStart)
		fmt.Print("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.VersionString()).
			Strs("destinations", destinationStrings).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncClient, err := sourcePbClient.Sync(ctx, &source.Sync_Request{})
	if err != nil {
		return err
	}
	writeClients := make([]destination.Destination_Write2Client, len(destinationsPbClients))
	defer func() {
		for i, wc := range writeClients {
			if wc == nil {
				continue
			}
			if _, closeErr := wc.CloseAndRecv(); closeErr != nil {
				log.Err(closeErr).Str("destination", destinationSpecs[i].Name).Msg("Failed to close write stream")
			}
		}
	}()

	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write2(ctx)
		if err != nil {
			return err
		}
		if err := writeClients[i].Send(&destination.Write2_Request{
			Source:    sourceSpec.Name,
			Tables:    tablesRes.Tables,
			Timestamp: timestamppb.New(syncTime),
		}); err != nil {
			return err
		}
	}

	bar := progressBar(noopProgressBar{})
	if !logConsole {
		bar = progressbar.NewOptions(-1,
			progressbar.OptionSetDescription("Syncing resources..."),
			progressbar.OptionSetItsString("resources"),
			progressbar.OptionShowIts(),
			progressbar.OptionSetElapsedTime(true),
			progressbar.OptionShowCount(),
			progressbar.OptionClearOnFinish(),
		)
	}

	// Add a ticker to update the progress bar every second.
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				_ = bar.Add(0)
			}
		}
	}()

	// Read from the sync stream and write to all destinations.
	for {
		r, err := syncClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		_ = bar.Add(1)
		for i := range destinationsPbClients {
			if err := writeClients[i].Send(&destination.Write2_Request{
				Resource: r.Resource,
			}); err != nil {
				return err
			}
		}
	}

	for i := range destinationsClients {
		if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
			_, err := destinationsPbClients[i].DeleteStale(ctx, &destination.DeleteStale_Request{
				Tables:    tablesRes.Tables,
				Source:    sourceSpec.Name,
				Timestamp: timestamppb.New(syncTime),
			})
			if err != nil {
				return err
			}
		}
	}

	getMetricsRes, err := sourcePbClient.GetMetrics(ctx, &source.GetMetrics_Request{})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(getMetricsRes.Metrics, &mt); err != nil {
		return err
	}

	err = bar.Finish()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to finish progress bar")
	}
	syncTimeTook := time.Since(syncTime)
	exitReason = ExitReasonCompleted
	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", mt.TotalResources(), mt.TotalErrors(), mt.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	return nil
}
