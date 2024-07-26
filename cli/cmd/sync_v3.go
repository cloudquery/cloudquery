package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/internal/api"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/internal/transformer"
	"github.com/cloudquery/cloudquery/cli/internal/transformerpipeline"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/vnteamopen/godebouncer"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/google/uuid"
)

const (
	MetadataMessageTypeKey          = "cq:extension:message_type"
	MetadataMessageTypeMigrateValue = "migrate"
)

type v3source struct {
	client *managedplugin.Client
	spec   specs.Source
}

type v3destination struct {
	client *managedplugin.Client
	spec   specs.Destination
}

type v3transformer struct {
	client *managedplugin.Client
	spec   specs.Transformer
}

func getProgressAPIClient() (*cloudquery_api.ClientWithResponses, error) {
	authClient := auth.NewTokenClient()
	if authClient.GetTokenType() != auth.SyncRunAPIKey {
		return nil, nil
	}

	token, err := authClient.GetToken()
	if err != nil {
		return nil, err
	}
	return api.NewClient(token.Value)
}

// nolint:dupl
func syncConnectionV3(ctx context.Context, source v3source, destinations []v3destination, transformersByDestination map[string][]v3transformer, backend *v3destination, uid string, noMigrate bool, summaryLocation string) (syncErr error) {
	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
	skippedFromDeleteStale := make(map[string]bool, 0)
	tablesForDeleteStale := make(map[string]bool, 0)

	sourceSpec := source.spec
	sourceClient := source.client
	destinationSpecs := make([]specs.Destination, len(destinations))
	destinationsClients := make([]*managedplugin.Client, len(destinations))
	for i := range destinations {
		destinationSpecs[i] = destinations[i].spec
		destinationsClients[i] = destinations[i].client
	}

	syncStartedEvent := analytics.SyncStartedEvent{
		Source:       sourceSpec,
		Destinations: destinationSpecs,
	}
	analytics.TrackSyncStarted(ctx, invocationUUID, syncStartedEvent)
	var (
		syncTimeTook   time.Duration
		totalResources = int64(0)
		totals         = sourceClient.Metrics()
	)
	defer func() {
		analytics.TrackSyncCompleted(ctx, invocationUUID, analytics.SyncFinishedEvent{
			SyncStartedEvent:  syncStartedEvent,
			Errors:            totals.Errors,
			Warnings:          totals.Warnings,
			Duration:          syncTimeTook,
			ResourceCount:     totalResources,
			AbortedDueToError: syncErr,
		})
	}()

	progressAPIClient, err := getProgressAPIClient()
	if err != nil {
		return fmt.Errorf("failed to get API client: %w", err)
	}

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
	sourceName := sourceSpec.Name
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}

	// Get all distinct transformer version strings
	transformerStrings := []string{}
	_transformerSet := make(map[string]struct{})
	for _, transformers := range transformersByDestination {
		for _, tf := range transformers {
			name := tf.spec.Name
			if _, ok := _transformerSet[name]; !ok {
				transformerStrings = append(transformerStrings, tf.spec.VersionString())
				_transformerSet[name] = struct{}{}
			}
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Strs("transformers", transformerStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Strs("transformers", transformerStrings).Time("sync_time", syncTime).Msg("End sync")

	variables := specs.Variables{
		Plugins: make(map[string]specs.PluginVariables),
	}
	sourcePbClient := plugin.NewPluginClient(sourceClient.Conn)
	destinationsPbClients := make([]plugin.PluginClient, len(destinationsClients))
	transformerPbClientsByDestination := map[string][]plugin.PluginClient{}
	for name, transformers := range transformersByDestination {
		for _, tf := range transformers {
			transformerPbClientsByDestination[name] = append(transformerPbClientsByDestination[name], plugin.NewPluginClient(tf.client.Conn))
		}
	}

	destinationTransformers := make([]*transformer.RecordTransformer, len(destinationsClients))
	backendPbClient := plugin.PluginClient(nil)
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(sourceName),
			transformer.WithSyncTimeColumn(syncTime),
		}
		if destinationSpecs[i].SyncGroupId != "" {
			opts = append(opts, transformer.WithSyncGroupIdColumn(destinationSpecs[i].RenderedSyncGroupId(syncTime, uid)))
		}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithRemoveUniqueConstraints())
		} else if destinationSpecs[i].PKMode == specs.PKModeCQID {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithCQIDPrimaryKey())
		}
		destinationTransformers[i] = transformer.NewRecordTransformer(opts...)
		connection := destinationsClients[i].ConnectionString()
		variables.Plugins[destinationSpecs[i].Name] = specs.PluginVariables{
			Connection: connection,
		}
	}
	if backend != nil {
		backendPbClient = plugin.NewPluginClient(backend.client.Conn)
		connection := backend.client.ConnectionString()
		variables.Plugins[backend.spec.Name] = specs.PluginVariables{
			Connection: connection,
		}
	}

	// initialize destinations first, so that their connections may be used as backends by the source
	for i, destinationSpec := range destinationSpecs {
		if err := initPlugin(ctx, destinationsPbClients[i], destinationSpec.Spec, false, uid); err != nil {
			return fmt.Errorf("failed to init destination %v: %w", destinationSpec.Name, err)
		}
	}
	if backend != nil {
		if err := initPlugin(ctx, backendPbClient, backend.spec.Spec, false, uid); err != nil {
			return fmt.Errorf("failed to init backend %v: %w", backend.spec.Name, err)
		}
	}
	for name, transformers := range transformersByDestination {
		for i, tf := range transformers {
			if err := initPlugin(ctx, transformerPbClientsByDestination[name][i], tf.spec.Spec, false, uid); err != nil {
				return fmt.Errorf("failed to init transformer %v: %w", tf.spec.Name, err)
			}
		}
	}

	// replace @@plugins.name.connection with the actual GRPC connection string from the client
	// NOTE: if this becomes a stable feature, it can move out of sync_v3 and into sync.go
	specBytes, err := json.Marshal(sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to marshal source spec JSON before variable replacement: %w", err)
	}
	specBytesExpanded, err := specs.ReplaceVariables(string(specBytes), variables)
	if err != nil {
		return fmt.Errorf("failed to replace variables: %w", err)
	}
	if err := json.Unmarshal([]byte(specBytesExpanded), &sourceSpec); err != nil {
		return fmt.Errorf("failed to unmarshal source spec JSON after variable replacement: %w", err)
	}

	if err = initPlugin(ctx, sourcePbClient, sourceSpec.Spec, false, uid); err != nil {
		return fmt.Errorf("failed to init source %v: %w", sourceSpec.Name, err)
	}

	writeClients := make([]plugin.Plugin_WriteClient, len(destinationsPbClients))
	writeClientsByName := map[string]plugin.Plugin_WriteClient{}
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
		writeClientsByName[destinationSpecs[i].Name] = writeClients[i]
	}
	transformClientsByDestination := map[string][]plugin.Plugin_TransformClient{}
	for name, transformerPbClients := range transformerPbClientsByDestination {
		for _, transformerPbClient := range transformerPbClients {
			pbClient, err := transformerPbClient.Transform(ctx)
			if err != nil {
				return err
			}
			transformClientsByDestination[name] = append(transformClientsByDestination[name], pbClient)
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncReq := &plugin.Sync_Request{
		Tables:              sourceSpec.Tables,
		SkipTables:          sourceSpec.SkipTables,
		SkipDependentTables: *sourceSpec.SkipDependentTables,
		DeterministicCqId:   sourceSpec.DeterministicCQID,
	}
	if sourceSpec.BackendOptions != nil {
		syncReq.Backend = &plugin.Sync_BackendOptions{
			TableName:  sourceSpec.BackendOptions.TableName,
			Connection: sourceSpec.BackendOptions.Connection,
		}
	}
	syncClient, err := sourcePbClient.Sync(ctx, syncReq)
	if err != nil {
		return err
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

	// Add a ticker to update the progress bar every 100ms
	t := time.NewTicker(100 * time.Millisecond)
	newResources := int64(0)
	defer t.Stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				change := atomic.SwapInt64(&newResources, 0)
				_ = bar.Add(int(change))
				return
			case <-t.C:
				change := atomic.SwapInt64(&newResources, 0)
				_ = bar.Add(int(change))
			}
		}
	}()

	isStateBackendEnabled := sourceSpec.BackendOptions != nil && sourceSpec.BackendOptions.TableName != ""

	// Read from the sync stream and write to all destinations.
	isComplete := int64(0)

	var remoteProgressReporter *godebouncer.Debouncer
	if progressAPIClient != nil {
		teamName, syncName, syncRunId := os.Getenv("_CQ_TEAM_NAME"), os.Getenv("_CQ_SYNC_NAME"), os.Getenv("_CQ_SYNC_RUN_ID")
		if teamName == "" || syncName == "" || syncRunId == "" {
			return fmt.Errorf("failed to get team, sync or sync run from environment variables. got team: %s, sync name: %s, sync run ID: %s", teamName, syncName, syncRunId)
		}
		syncRunUUID, err := uuid.Parse(syncRunId)
		if err != nil {
			return fmt.Errorf("failed to parse sync run ID: %w", err)
		}
		remoteProgressReporter = godebouncer.NewWithOptions(godebouncer.WithTimeDuration(10*time.Second), godebouncer.WithTriggered(func() {
			totals := sourceClient.Metrics()
			for i := range destinationsClients {
				m := destinationsClients[i].Metrics()
				totals.Warnings += m.Warnings
				totals.Errors += m.Errors
			}
			status := cloudquery_api.SyncRunStatusStarted
			if atomic.LoadInt64(&isComplete) == 1 {
				status = cloudquery_api.SyncRunStatusCompleted
			}
			obj := cloudquery_api.CreateSyncRunProgressJSONRequestBody{
				Rows:     atomic.LoadInt64(&totalResources),
				Errors:   int64(totals.Errors),
				Warnings: int64(totals.Warnings),
				Status:   &status,
			}

			log.Debug().Interface("body", obj).Msg("Sending sync progress to API")
			res, err := progressAPIClient.CreateSyncRunProgressWithResponse(ctx, teamName, syncName, syncRunUUID, obj)
			if err != nil {
				log.Warn().Err(err).Msg("Failed to send sync progress to API")
				return
			}
			if res.StatusCode() != http.StatusNoContent {
				log.Warn().Str("status", res.Status()).Int("code", res.StatusCode()).Msg("Failed to send sync progress to API")
			}
		}), godebouncer.WithOptions(godebouncer.Options{Trailing: true, Leading: true}))
		defer remoteProgressReporter.Cancel()
	}

	// Note: we want to stop this errorgroup if ctx is cancelled, but we don't want to cancel ctx if gctx is cancelled.
	// gctx is always cancelled when the errorgroup returns, and this isn't necessarily an error.
	eg, gctx := errgroup.WithContext(ctx)
	pipelineByDestinationName := map[string]*transformerpipeline.TransformerPipeline{}

	// Each destination has its own transformer pipeline
	for i := range destinationsPbClients {
		destinationName := destinationSpecs[i].Name

		// Start a pipeline of transformers that will receive & transform the source records
		var (
			pipeline *transformerpipeline.TransformerPipeline
			err      error
		)
		pipeline, gctx, err = transformerpipeline.New(gctx, transformClientsByDestination[destinationName])
		if err != nil {
			return fmt.Errorf("failed to create transformer pipeline: %w", err)
		}
		err = pipeline.OnOutput(func(recordBytes []byte) error {
			record, err := plugin.NewRecordFromBytes(recordBytes)
			if err != nil {
				return fmt.Errorf("failed to get record from bytes: %w", err)
			}

			wr := &plugin.Write_Request{}

			// This could either be an `Insert` or a `MigrateTable` message. To answer this,
			// we look to see if the schema's metadata has a MetadataMessageTypeKey key
			msgType := "insert"
			if isMigrateTableRecord(record) {
				msgType = MetadataMessageTypeMigrateValue
				schemaBytes, err := plugin.SchemaToBytes(record.Schema())
				if err != nil {
					return fmt.Errorf("failed to get schema bytes: %w", err)
				}
				wr.Message = &plugin.Write_Request_MigrateTable{
					MigrateTable: &plugin.Write_MessageMigrateTable{
						MigrateForce: destinationSpecs[i].MigrateMode == specs.MigrateModeForced,
						Table:        schemaBytes,
					},
				}
			} else {
				wr.Message = &plugin.Write_Request_Insert{
					Insert: &plugin.Write_MessageInsert{
						Record: recordBytes,
					},
				}
			}
			if err := writeClientsByName[destinationName].Send(wr); err != nil {
				return handleSendError(err, writeClientsByName[destinationName], msgType)
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to create register pipeline output: %w", err)
		}
		eg.Go(pipeline.RunBlocking) // each transformer runs in its own goroutine
		pipelineByDestinationName[destinationName] = pipeline
	}

	eg.Go(func() error {
		// Close all transformation pipelines when the source is done
		defer func() {
			for _, pipeline := range pipelineByDestinationName {
				if err := pipeline.Close(); err != nil {
					log.Warn().Err(err).Msg("Failed to close transformer pipeline")
				}
			}
		}()
		for {
			r, err := syncClient.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return fmt.Errorf("unexpected error from sync client receive: %w", err)
			}
			syncResponseMsg := r.GetMessage()
			switch m := syncResponseMsg.(type) {
			case *plugin.Sync_Response_Insert:
				recordBytes := m.Insert.Record
				record, err := plugin.NewRecordFromBytes(recordBytes)
				if err != nil {
					return fmt.Errorf("failed to get record from bytes: %w", err)
				}

				atomic.AddInt64(&newResources, record.NumRows())
				atomic.AddInt64(&totalResources, record.NumRows())
				if remoteProgressReporter != nil {
					remoteProgressReporter.SendSignal()
				}
				for i := range destinationsPbClients {
					destinationName := destinationSpecs[i].Name
					transformedRecord := destinationTransformers[i].Transform(record)
					transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
					if err != nil {
						return fmt.Errorf("failed to transform record bytes: %w", err)
					}
					if err := pipelineByDestinationName[destinationName].Send(transformedRecordBytes); err != nil {
						return err
					}
				}
			case *plugin.Sync_Response_DeleteRecord:
				for i := range destinationsPbClients {
					wr := &plugin.Write_Request{}
					// Transformations aren't required here because DeleteRecord is only in V3
					wr.Message = &plugin.Write_Request_DeleteRecord{
						DeleteRecord: &plugin.Write_MessageDeleteRecord{
							TableName:      m.DeleteRecord.TableName,
							TableRelations: m.DeleteRecord.TableRelations,
							WhereClause:    m.DeleteRecord.WhereClause,
						},
					}
					if err := writeClients[i].Send(wr); err != nil {
						return handleSendError(err, writeClients[i], "delete")
					}
				}
			case *plugin.Sync_Response_MigrateTable:
				sc, err := plugin.NewSchemaFromBytes(m.MigrateTable.Table)
				if err != nil {
					return err
				}
				table, err := schema.NewTableFromArrowSchema(sc)
				if err != nil {
					return err
				}

				// This works since we sync and send migrate messages for parents before children
				if isStateBackendEnabled && (table.IsIncremental || (table.Parent != nil && skippedFromDeleteStale[table.Parent.Name])) {
					skippedFromDeleteStale[table.Name] = true
				} else {
					tablesForDeleteStale[table.Name] = true
				}
				if noMigrate {
					continue
				}
				for i := range destinationsPbClients {
					destinationName := destinationSpecs[i].Name
					transformedSchema := destinationTransformers[i].TransformSchema(sc)

					// Since Transformer plugins only have `.Transform()` which transforms records,
					// but in this case we must transform the schema, we create an empty record with the schema.
					//
					// The transformer won't care about this difference, but the destination still needs to
					// receive the proper message (MigrateTable), so we inject a metadata key to remember that.
					transformedRecord := makeEmptyRecordWithMigrateMetadata(transformedSchema)
					transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
					if err != nil {
						return fmt.Errorf("failed to transform record bytes: %w", err)
					}
					if err := pipelineByDestinationName[destinationName].Send(transformedRecordBytes); err != nil {
						return err
					}
				}
			default:
				return fmt.Errorf("unknown message type: %T", m)
			}
		}
		return nil
	})
	if err := eg.Wait(); err != nil { // wait for source & transformers to finish. If any fails, sync fails.
		return err
	}

	err = syncClient.CloseSend()
	if err != nil {
		return err
	}
	sourceWarnings := totals.Warnings
	sourceErrors := totals.Errors
	var metadataDataErrors error
	for i := range destinationsClients {
		m := destinationsClients[i].Metrics()
		summary := syncSummary{
			Resources:           uint64(totalResources),
			SourceErrors:        sourceErrors,
			SourceWarnings:      sourceWarnings,
			SyncID:              uid,
			SyncTime:            syncTime,
			SourceName:          sourceSpec.Name,
			SourceVersion:       sourceSpec.Version,
			SourcePath:          sourceSpec.Path,
			CLIVersion:          Version,
			DestinationErrors:   m.Errors,
			DestinationWarnings: m.Warnings,
			DestinationName:     destinationSpecs[i].Name,
			DestinationVersion:  destinationSpecs[i].Version,
			DestinationPath:     destinationSpecs[i].Path,
		}

		if err := persistSummary(summaryLocation, summary); err != nil {
			log.Warn().Err(err).Msg("Failed to persist sync summary")
		}

		log.Info().Interface("summary", summary).Msg("Sync summary")
		if !destinationSpecs[i].SyncSummary {
			continue
		}
		// Only send the summary to the destination that matches the current destination
		if err := sendSummary(writeClients[i], destinationSpecs[i], destinationsClients[i], destinationTransformers[i], &summary, noMigrate); err != nil {
			metadataDataErrors = errors.Join(metadataDataErrors, err)
		}
	}
	if metadataDataErrors != nil {
		return metadataDataErrors
	}

	for i := range destinationsClients {
		if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
			if err := deleteStale(writeClients[i], tablesForDeleteStale, sourceName, syncTime); err != nil {
				return err
			}
		}
		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Close(ctx, &plugin.Close_Request{}); err != nil {
			return err
		}
	}

	err = bar.Finish()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to finish progress bar")
	}
	atomic.StoreInt64(&isComplete, 1)
	syncTimeTook = time.Since(syncTime)
	exitReason = ExitReasonCompleted

	msg := "Sync completed successfully"
	if totals.Errors > 0 {
		msg = "Sync completed with errors, see logs for details"
	}
	fmt.Printf("%s. Resources: %d, Errors: %d, Warnings: %d, Time: %s\n", msg, totalResources, totals.Errors, totals.Warnings, syncTimeTook.Truncate(time.Second).String())
	log.Info().
		Int64("resources", totalResources).
		Uint64("errors", totals.Errors).
		Uint64("warnings", totals.Warnings).
		Str("duration", syncTimeTook.Truncate(time.Second).String()).
		Str("result", msg).
		Msg("Sync summary")

	if remoteProgressReporter != nil {
		remoteProgressReporter.SendSignal()
		<-remoteProgressReporter.Done()
	}
	return nil
}

func deleteStale(client plugin.Plugin_WriteClient, tables map[string]bool, sourceName string, syncTime time.Time) error {
	for tableName := range tables {
		if err := client.Send(&plugin.Write_Request{
			Message: &plugin.Write_Request_Delete{
				Delete: &plugin.Write_MessageDeleteStale{
					SourceName: sourceName,
					SyncTime:   timestamppb.New(syncTime),
					TableName:  tableName,
				},
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

func makeEmptyRecordWithMigrateMetadata(s *arrow.Schema) arrow.Record {
	// Add a temporary message type to the metadata to indicate that this record is for migrating a table
	mdMap := s.Metadata().ToMap()
	mdMap[MetadataMessageTypeKey] = MetadataMessageTypeMigrateValue
	newMd := arrow.MetadataFrom(mdMap)
	mem := memory.NewGoAllocator()

	cols := []arrow.Array{}
	for _, field := range s.Fields() {
		cols = append(cols, array.NewBuilder(mem, field.Type).NewArray())
	}
	return array.NewRecord(arrow.NewSchema(s.Fields(), &newMd), cols, 0)
}

func isMigrateTableRecord(record arrow.Record) bool {
	value, ok := record.Schema().Metadata().GetValue(MetadataMessageTypeKey)
	return ok && value == MetadataMessageTypeMigrateValue
}
