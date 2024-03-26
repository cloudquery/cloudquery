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

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery/cli/internal/api"
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/vnteamopen/godebouncer"
	"google.golang.org/protobuf/types/known/timestamppb"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/google/uuid"
)

type v3source struct {
	client *managedplugin.Client
	spec   specs.Source
}

type v3destination struct {
	client *managedplugin.Client
	spec   specs.Destination
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
func syncConnectionV3(ctx context.Context, source v3source, destinations []v3destination, backend *v3destination, uid string, noMigrate bool) error {
	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
	tablesForDeleteStale := make(map[string]bool, 0)

	sourceSpec := source.spec
	sourceClient := source.client
	destinationSpecs := make([]specs.Destination, len(destinations))
	destinationsClients := make([]*managedplugin.Client, len(destinations))
	for i := range destinations {
		destinationSpecs[i] = destinations[i].spec
		destinationsClients[i] = destinations[i].client
	}

	progressAPIClient, err := getProgressAPIClient()
	if err != nil {
		return fmt.Errorf("failed to get API client: %w", err)
	}

	defer func() {
		if analyticsClient != nil {
			log.Info().Msg("Sending sync summary to " + analyticsClient.Host())
			if err := analyticsClient.SendSyncMetrics(context.Background(), sourceSpec, destinationSpecs, uid, &mt, exitReason); err != nil {
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
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	variables := specs.Variables{
		Plugins: make(map[string]specs.PluginVariables),
	}
	sourcePbClient := plugin.NewPluginClient(sourceClient.Conn)
	destinationsPbClients := make([]plugin.PluginClient, len(destinationsClients))
	destinationTransformers := make([]*transformer.RecordTransformer, len(destinationsClients))
	backendPbClient := plugin.PluginClient(nil)
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(sourceName),
			transformer.WithSyncTimeColumn(syncTime),
		}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs())
			if sourceSpec.DeterministicCQID {
				opts = append(opts, transformer.WithRemoveUniqueConstraints())
			}
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
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncReq := &plugin.Sync_Request{
		Tables:              sourceSpec.Tables,
		SkipTables:          sourceSpec.SkipTables,
		SkipDependentTables: sourceSpec.SkipDependentTables,
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

	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)

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
	totalResources := int64(0)
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
	}

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
			record, err := plugin.NewRecordFromBytes(m.Insert.Record)
			if err != nil {
				return fmt.Errorf("failed to get record from bytes: %w", err)
			}

			atomic.AddInt64(&newResources, record.NumRows())
			atomic.AddInt64(&totalResources, record.NumRows())
			if remoteProgressReporter != nil {
				remoteProgressReporter.SendSignal()
			}
			for i := range destinationsPbClients {
				transformedRecord := destinationTransformers[i].Transform(record)
				transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
				if err != nil {
					return fmt.Errorf("failed to transform record bytes: %w", err)
				}
				wr := &plugin.Write_Request{}
				wr.Message = &plugin.Write_Request_Insert{
					Insert: &plugin.Write_MessageInsert{
						Record: transformedRecordBytes,
					},
				}
				if err := writeClients[i].Send(wr); err != nil {
					return handleSendError(err, writeClients[i], "insert")
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
			tableName := tableNameFromSchema(sc)

			if !isStateBackendEnabled || !tableIsIncremental(sc) {
				tablesForDeleteStale[tableName] = true
			}
			if noMigrate {
				continue
			}
			for i := range destinationsPbClients {
				transformedSchema := destinationTransformers[i].TransformSchema(sc)
				transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
				if err != nil {
					return err
				}
				wr := &plugin.Write_Request{}
				wr.Message = &plugin.Write_Request_MigrateTable{
					MigrateTable: &plugin.Write_MessageMigrateTable{
						MigrateForce: destinationSpecs[i].MigrateMode == specs.MigrateModeForced,
						Table:        transformedSchemaBytes,
					},
				}
				if err := writeClients[i].Send(wr); err != nil {
					return handleSendError(err, writeClients[i], "migrate")
				}
			}
		default:
			return fmt.Errorf("unknown message type: %T", m)
		}
	}
	err = syncClient.CloseSend()
	if err != nil {
		return err
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
	atomic.StoreInt64(&isComplete, 1)
	syncTimeTook := time.Since(syncTime)
	exitReason = ExitReasonCompleted

	msg := "Sync completed successfully"
	if totals.Errors > 0 {
		msg = "Sync completed with errors, see logs for details"
	}
	fmt.Printf("%s. Resources: %d, Errors: %d, Warnings: %d, Time: %s\n", msg, totalResources, totals.Errors, totals.Warnings, syncTimeTook.Truncate(time.Second).String())

	if remoteProgressReporter != nil {
		remoteProgressReporter.SendSignal()
		<-remoteProgressReporter.Done()
	}
	return nil
}

func tableNameFromSchema(sc *arrow.Schema) string {
	tableName, _ := sc.Metadata().GetValue("cq:table_name")
	return tableName
}

func tableIsIncremental(sc *arrow.Schema) bool {
	inc, _ := sc.Metadata().GetValue("cq:extension:incremental")
	return inc == "true"
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
