package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
)

// nolint:dupl
func syncConnectionV3(ctx context.Context, sourceClient *managedplugin.Client, destinationsClients managedplugin.Clients, sourceSpec specs.Source, destinationSpecs []specs.Destination, uid string, noMigrate bool) error {
	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
	defer func() {
		if analyticsClient != nil {
			log.Info().Msg("Sending sync summary to " + analyticsClient.Host())
			if err := analyticsClient.SendSyncMetrics(context.Background(), sourceSpec, destinationSpecs, uid, &mt, exitReason); err != nil {
				log.Warn().Err(err).Msg("Failed to send sync summary")
			}
		}
	}()
	syncTime := time.Now().UTC()
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	sourcePbClient := plugin.NewPluginClient(sourceClient.Conn)
	destinationsPbClients := make([]plugin.PluginClient, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
	}

	specBytes, err := json.Marshal(sourceSpec.Spec)
	if err != nil {
		return err
	}
	if _, err := sourcePbClient.Init(ctx, &plugin.Init_Request{
		Spec: specBytes,
	}); err != nil {
		return err
	}
	for i := range destinationsClients {
		// TODO: for backwards-compatibility check for old fields like `batch_size` and move them into the spec, log a warning
		destSpecBytes, err := json.Marshal(destinationSpecs[i].Spec)
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Init(ctx, &plugin.Init_Request{
			Spec: destSpecBytes,
		}); err != nil {
			return err
		}
	}

	writeClients := make([]plugin.Plugin_WriteClient, len(destinationsPbClients))
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
		// TODO(v4): necessary?
		if err := writeClients[i].Send(&plugin.Write_Request{
			Message: &plugin.Write_Request_Options{
				Options: &plugin.WriteOptions{
					MigrateForce: destinationSpecs[i].MigrateMode == specs.MigrateModeForced,
				},
			},
		}); err != nil {
			return err
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	// TODO(v4): figure out backends
	syncClient, err := sourcePbClient.Sync(ctx, &plugin.Sync_Request{
		Tables:     sourceSpec.Tables,
		SkipTables: sourceSpec.SkipTables,
		//StateBackend: &plugin.StateBackendSpec{
		//	Name:     sourceSpec.Backend,
		//	Path:     "",
		//	Version:  "",
		//	Registry: 0,
		//	Spec:     sourceSpec.BackendSpec,
		//},
	})
	if err != nil {
		return err
	}

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)

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
		m := r.GetMessage()
		wr := &plugin.Write_Request{}
		switch m.(type) {
		case *plugin.Sync_Response_Insert:
			wr.Message = &plugin.Write_Request_Insert{
				Insert: m.(*plugin.Sync_Response_Insert).Insert,
			}
		case *plugin.Sync_Response_MigrateTable:
			wr.Message = &plugin.Write_Request_MigrateTable{
				MigrateTable: m.(*plugin.Sync_Response_MigrateTable).MigrateTable,
			}
		case *plugin.Sync_Response_Delete:
			wr.Message = &plugin.Write_Request_Delete{
				Delete: m.(*plugin.Sync_Response_Delete).Delete,
			}
		default:
			return fmt.Errorf("unknown message type: %T", m)
		}
		for i := range destinationsPbClients {
			if err := writeClients[i].Send(wr); err != nil {
				return err
			}
		}
	}
	for i := range destinationsClients {
		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
	}

	// TODO(v4): we seem to be missing GetMetrics
	//getMetricsRes, err := sourcePbClient.GetMetrics(ctx, &source.GetMetrics_Request{})
	//if err != nil {
	//	return err
	//}
	//if err := json.Unmarshal(getMetricsRes.Metrics, &mt); err != nil {
	//	return err
	//}

	err = bar.Finish()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to finish progress bar")
	}
	syncTimeTook := time.Since(syncTime)
	exitReason = ExitReasonCompleted
	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", mt.TotalResources(), mt.TotalErrors(), mt.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	return nil
}
