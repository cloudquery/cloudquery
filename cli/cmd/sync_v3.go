package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/internal/transformer"
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
	sourceName := sourceSpec.Name
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	sourcePbClient := plugin.NewPluginClient(sourceClient.Conn)
	destinationsPbClients := make([]plugin.PluginClient, len(destinationsClients))
	destinationTransformers := make([]*transformer.RecordTransformer, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(sourceName),
			transformer.WithSyncTimeColumn(syncTime),
		}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs())
		}
		destinationTransformers[i] = transformer.NewRecordTransformer(opts...)
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
	totalResources := 0
	for {
		r, err := syncClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		_ = bar.Add(1)
		syncResponseMsg := r.GetMessage()
		wr := &plugin.Write_Request{}
		switch m := syncResponseMsg.(type) {
		case *plugin.Sync_Response_Insert:
			record, err := NewRecordFromBytes(m.Insert.Record)
			if err != nil {
				return err
			}
			totalResources += int(record.NumRows())
			for i := range destinationsPbClients {
				transformedRecord := destinationTransformers[i].Transform(record)
				transformedRecordBytes, err := RecordToBytes(transformedRecord)
				if err != nil {
					return err
				}
				wr.Message = &plugin.Write_Request_Insert{
					Insert: &plugin.MessageInsert{
						Record: transformedRecordBytes,
					},
				}
				if err := writeClients[i].Send(wr); err != nil {
					return err
				}
			}
		case *plugin.Sync_Response_MigrateTable:
			wr.Message = &plugin.Write_Request_MigrateTable{
				MigrateTable: m.MigrateTable,
			}
		case *plugin.Sync_Response_Delete:
			wr.Message = &plugin.Write_Request_Delete{
				Delete: m.Delete,
			}
		default:
			return fmt.Errorf("unknown message type: %T", m)
		}
	}
	for i := range destinationsClients {
		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
	}

	totals := sourceClient.Metrics()
	for i := range destinationsClients {
		m := destinationsClients[i].Metrics()
		totals.Warnings += m.Warnings
		totals.Errors += m.Errors
	}

	err = bar.Finish()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to finish progress bar")
	}
	syncTimeTook := time.Since(syncTime)
	exitReason = ExitReasonCompleted

	msg := "Sync completed successfully"
	if totals.Errors > 0 {
		msg = "Sync completed with errors, see logs for details"
	}
	fmt.Printf("%s. Resources: %d, Errors: %d, Warnings: %d, Time: %s\n", msg, totalResources, totals.Errors, totals.Warnings, syncTimeTook.Truncate(time.Second).String())
	return nil
}
