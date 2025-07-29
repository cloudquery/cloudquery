package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	gosync "sync"
	"sync/atomic"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/vnteamopen/godebouncer"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/tablenamechanger"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformer"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformerpipeline"
	"github.com/cloudquery/cloudquery/cli/v6/internal/utils"
)

type v3source struct {
	client *managedplugin.Client
	spec   specs.Source

	// If true, source is running in docker and communicates
	// to a backend in the host system. It must replace localhost
	// with host.docker.internal.
	shouldReplaceLocalhost bool
}

type v3destination struct {
	client *managedplugin.Client
	spec   specs.Destination
}

type v3transformer struct {
	client *managedplugin.Client
	spec   specs.Transformer
}

type safeWriteClient struct {
	client grpc.ClientStreamingClient[plugin.Write_Request, plugin.Write_Response]
	mu     *gosync.Mutex
}

// It is not safe to call write.Send on the same stream in different goroutines
func (s safeWriteClient) Send(req *plugin.Write_Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.Send(req)
}

// It is also not safe to call CloseSend concurrently with SendMsg
func (s safeWriteClient) CloseAndRecv() (*plugin.Write_Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.client.CloseAndRecv()
}

func newSafeWriteClient(client grpc.ClientStreamingClient[plugin.Write_Request, plugin.Write_Response]) safeWriteClient {
	return safeWriteClient{client: client, mu: &gosync.Mutex{}}
}

type shard struct {
	num   int32
	total int32
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

type syncV3Options struct {
	source                    v3source
	destinations              []v3destination
	transformersByDestination map[string][]v3transformer
	backend                   *v3destination
	uid                       string
	noMigrate                 bool
	summaryLocation           string
	shard                     *shard
	cqColumnsNotNull          bool
}

func syncConnectionV3(ctx context.Context, syncOptions syncV3Options) (syncErr error) {
	var (
		source                    = syncOptions.source
		destinations              = syncOptions.destinations
		transformersByDestination = syncOptions.transformersByDestination
		backend                   = syncOptions.backend
		uid                       = syncOptions.uid
		noMigrate                 = syncOptions.noMigrate
		summaryLocation           = syncOptions.summaryLocation
		shard                     = syncOptions.shard
		cqColumnsNotNull          = syncOptions.cqColumnsNotNull
	)

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
	if shard != nil {
		syncStartedEvent.ShardNum = int(shard.num)
		syncStartedEvent.ShardTotal = int(shard.total)
	}
	analytics.TrackSyncStarted(ctx, invocationUUID.UUID, syncStartedEvent)
	var (
		syncTimeTook   time.Duration
		totalResources = int64(0)
		totals         = sourceClient.Metrics()
		statsPerTable  = utils.NewConcurrentMap[string, cloudquery_api.SyncRunTableProgressValue]()
	)
	defer func() {
		analytics.TrackSyncCompleted(ctx, invocationUUID.UUID, analytics.SyncFinishedEvent{
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
	if shard != nil {
		sourceName = fmt.Sprintf("%s_%d_%d", sourceName, shard.num, shard.total)
	}
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
	syncGroupId := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(sourceName),
			transformer.WithSyncTimeColumn(syncTime),
		}
		if cqColumnsNotNull {
			opts = append(opts, transformer.WithCQColumnsNotNull())
		}
		if destinationSpecs[i].SyncGroupId != "" {
			syncGroupId[i] = destinationSpecs[i].RenderedSyncGroupId(syncTime, uid)
			opts = append(opts, transformer.WithSyncGroupIdColumn(syncGroupId[i]))
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
	specBytesExpanded, err := specs.ReplaceVariables(string(specBytes), variables, source.shouldReplaceLocalhost)
	if err != nil {
		return fmt.Errorf("failed to replace variables: %w", err)
	}
	if err := json.Unmarshal([]byte(specBytesExpanded), &sourceSpec); err != nil {
		return fmt.Errorf("failed to unmarshal source spec JSON after variable replacement: %w", err)
	}

	if err = initPlugin(ctx, sourcePbClient, sourceSpec.Spec, false, uid); err != nil {
		return fmt.Errorf("failed to init source %v: %w", sourceSpec.Name, err)
	}

	writeClients := make([]safeWriteClient, len(destinationsPbClients))
	writeClientsByName := map[string]safeWriteClient{}
	for i := range destinationsPbClients {
		writeClient, err := destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
		writeClients[i] = newSafeWriteClient(writeClient)
		writeClientsByName[destinationSpecs[i].Name] = writeClients[i]
	}
	transformClientsByDestination := map[string][]plugin.Plugin_TransformClient{}
	for name, transformerPbClients := range transformerPbClientsByDestination {
		for _, transformerPbClient := range transformerPbClients {
			transformClient, err := transformerPbClient.Transform(ctx)
			if err != nil {
				return err
			}
			transformClientsByDestination[name] = append(transformClientsByDestination[name], transformClient)
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncReq := &plugin.Sync_Request{
		Tables:              sourceSpec.Tables,
		SkipTables:          sourceSpec.SkipTables,
		SkipDependentTables: *sourceSpec.SkipDependentTables,
		DeterministicCqId:   sourceSpec.DeterministicCQID,
		WithErrorMessages:   true,
	}
	if sourceSpec.BackendOptions != nil {
		syncReq.Backend = &plugin.Sync_BackendOptions{
			TableName:  sourceSpec.BackendOptions.TableName,
			Connection: sourceSpec.BackendOptions.Connection,
		}
	}
	if shard != nil {
		syncReq.Shard = &plugin.Sync_Request_Shard{
			Num:   shard.num,
			Total: shard.total,
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

	isStateBackendEnabled := sourceSpec.BackendOptions != nil && sourceSpec.BackendOptions.TableName != ""

	// Read from the sync stream and write to all destinations.
	isComplete := int64(0)

	sourceTables, err := getTables(ctx, sourcePbClient, &plugin.GetTables_Request{
		Tables:              sourceSpec.Tables,
		SkipTables:          sourceSpec.SkipTables,
		SkipDependentTables: *sourceSpec.SkipDependentTables})

	if err != nil {
		return err
	}
	// Pre init stats per table
	for _, table := range sourceTables {
		initialStats := cloudquery_api.SyncRunTableProgressValue{
			Rows:   0,
			Errors: 0,
		}
		statsPerTable.Add(table.Name, initialStats)
	}

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

		triggerFunc := func() {
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
			tableProgress := cloudquery_api.SyncRunTableProgress(statsPerTable.GetAll())
			obj := cloudquery_api.CreateSyncRunProgressJSONRequestBody{
				Rows:          atomic.LoadInt64(&totalResources),
				Errors:        int64(totals.Errors),
				Warnings:      int64(totals.Warnings),
				Status:        &status,
				TableProgress: &tableProgress,
			}
			if shard != nil {
				obj.ShardNum = &shard.num
				obj.ShardTotal = &shard.total
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
		}
		remoteProgressReporter = godebouncer.NewWithOptions(godebouncer.WithTimeDuration(10*time.Second), godebouncer.WithTriggered(triggerFunc), godebouncer.WithOptions(godebouncer.Options{Trailing: true, Leading: true}))

		defer func() {
			remoteProgressReporter.Cancel()
			triggerFunc()
		}()
	}

	// Transformers can change table names. We need to keep track of table name changes
	// in case we do things that depend on table names.
	//
	// Note that transformers run per-destination, so we need to keep track of table name changes per-destination.
	tableNameChanger := tablenamechanger.New(destinationSpecs)

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
			wr := &plugin.Write_Request{
				Message: &plugin.Write_Request_Insert{
					Insert: &plugin.Write_MessageInsert{
						Record: recordBytes,
					},
				},
			}
			if err := writeClientsByName[destinationName].Send(wr); err != nil {
				return handleSendError(err, writeClientsByName[destinationName], "insert")
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to create register pipeline output: %w", err)
		}
		eg.Go(pipeline.RunBlocking) // each transformer runs in its own goroutine
		pipelineByDestinationName[destinationName] = pipeline
	}

	// Add a ticker to update the progress bar every 100ms
	t := time.NewTicker(100 * time.Millisecond)
	newResources := int64(0)
	go func() {
		for {
			select {
			case <-gctx.Done():
				t.Stop()
				change := atomic.SwapInt64(&newResources, 0)
				_ = bar.Add(int(change))
				if err := bar.Finish(); err != nil {
					log.Warn().Err(err).Msg("Failed to finish progress bar")
				}
				return
			case <-t.C:
				change := atomic.SwapInt64(&newResources, 0)
				_ = bar.Add(int(change))
			}
		}
	}()

	eg.Go(func() error {
		// Close all transformation pipelines when the source is done
		defer func() {
			for _, pipeline := range pipelineByDestinationName {
				pipeline.Close()
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
				tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)
				stats, _ := statsPerTable.Get(tableName)
				stats.Rows += record.NumRows()
				statsPerTable.Add(tableName, stats)
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
						if errors.Is(err, transformerpipeline.ErrPipelineClosed) {
							// If the pipeline is closed, we should stop processing records. The error that cause the pipeline to close
							// Will be returned by the pipeline itself.
							return nil
						}
						return err
					}
				}
			case *plugin.Sync_Response_DeleteRecord:
				for i := range destinationsPbClients {
					destinationName := destinationSpecs[i].Name

					// Create a minimal schema with the table name, similar to MigrateTable
					originalSchema := createTableNameSchema(m.DeleteRecord.TableName)

					// Apply the same transformation pipeline as MigrateTable
					transformedSchema := destinationTransformers[i].TransformSchema(originalSchema)
					transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
					if err != nil {
						return err
					}

					for _, transformerPbClient := range transformerPbClientsByDestination[destinationName] {
						resp, err := transformerPbClient.TransformSchema(ctx, &plugin.TransformSchema_Request{Schema: transformedSchemaBytes})
						if err != nil {
							return err
						}
						transformedSchemaBytes = resp.Schema
					}

					tableName, err := getTransformedTableNameFromSchema(transformedSchemaBytes)
					if err != nil {
						// Fall back to tableNameChanger if extraction fails
						log.Warn().Err(err).Str("table", m.DeleteRecord.TableName).Str("destination", destinationName).Msg("Failed to extract table name from transformed schema, falling back to tableNameChanger")
						tableName = tableNameChanger.UpdateTableName(destinationName, m.DeleteRecord.TableName)
					}

					wr := &plugin.Write_Request{}
					wr.Message = &plugin.Write_Request_DeleteRecord{
						DeleteRecord: &plugin.Write_MessageDeleteRecord{
							TableName:      tableName,
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
					transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
					if err != nil {
						return err
					}
					// Sequentially apply schema transformations from transformers
					for _, transformerPbClient := range transformerPbClientsByDestination[destinationName] {
						resp, err := transformerPbClient.TransformSchema(ctx, &plugin.TransformSchema_Request{Schema: transformedSchemaBytes})
						if err != nil {
							return err
						}
						transformedSchemaBytes = resp.Schema
					}

					// Table name might have changed due to a transformation.
					if err := tableNameChanger.LearnTableNameChange(destinationName, table.Name, transformedSchemaBytes); err != nil {
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
			case *plugin.Sync_Response_Error:
				log.Error().Str("table", m.Error.TableName).Msg(m.Error.Error)
				stats, _ := statsPerTable.Get(m.Error.TableName)
				stats.Errors++
				statsPerTable.Add(m.Error.TableName, stats)
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
	totals = sourceClient.Metrics()
	sourceWarnings := totals.Warnings
	var sourceErrors uint64
	for _, val := range statsPerTable.GetAll() {
		sourceErrors += uint64(val.Errors)
	}
	if totals.Errors > sourceErrors {
		sourceErrors = totals.Errors
	}
	var metadataDataErrors error

	tableProgress := statsPerTable.GetAll()
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
			SourceTables:        tableNameChanger.UpdateTableNamesSlice(destinationSpecs[i].Name, sourceTables.TableNames()),
			CLIVersion:          Version,
			DestinationErrors:   m.Errors,
			DestinationWarnings: m.Warnings,
			DestinationName:     destinationSpecs[i].Name,
			DestinationVersion:  destinationSpecs[i].Version,
			DestinationPath:     destinationSpecs[i].Path,
			ResourcesPerTable: lo.Reduce(lo.Keys(tableProgress), func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
				acc[tableName] = uint64(tableProgress[tableName].Rows)
				return acc
			}, map[string]uint64{}),
			ErrorsPerTable: lo.Reduce(lo.Keys(tableProgress), func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
				acc[tableName] = uint64(tableProgress[tableName].Errors)
				return acc
			}, map[string]uint64{}),
		}

		if destinationSpecs[i].SyncGroupId != "" {
			summary.SyncGroupID = lo.ToPtr(syncGroupId[i])
		}

		if shard != nil {
			summary.ShardNum = lo.ToPtr(int(shard.num))
			summary.ShardTotal = lo.ToPtr(int(shard.total))
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

	destPaths := make(map[string]bool, len(destinationsClients))
	for i := range destinationsClients {
		destPaths[destinationSpecs[i].Path] = true

		if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
			// Table names might have changed due to transformers
			updatedTablesForDeleteStale := tableNameChanger.UpdateTableNames(destinationSpecs[i].Name, tablesForDeleteStale)
			if err := deleteStale(writeClients[i], updatedTablesForDeleteStale, sourceName, syncTime); err != nil {
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

	atomic.StoreInt64(&isComplete, 1)
	syncTimeTook = time.Since(syncTime)
	exitReason = ExitReasonCompleted

	msg := "Sync completed successfully"
	if sourceErrors > 0 {
		msg = "Sync completed with errors, see logs for details"
	}
	fmt.Printf("%s. Resources: %d, Errors: %d, Warnings: %d, Time: %s\n", msg, totalResources, sourceErrors, totals.Warnings, syncTimeTook.Truncate(time.Second).String())
	log.Info().
		Int64("resources", totalResources).
		Uint64("errors", sourceErrors).
		Uint64("warnings", totals.Warnings).
		Str("duration", syncTimeTook.Truncate(time.Second).String()).
		Str("result", msg).
		Msg("Sync summary")

	if totalResources > 0 {
		hintSelectMessage(sourceSpec.Path, destPaths)
	}

	return nil
}

func hintSelectMessage(sourcePath string, destinationPaths map[string]bool) {
	if sourcePath != "cloudquery/aws" {
		return
	}

	val, _ := config.GetValue("first_sync_completed")
	firstSyncCompleted, _ := strconv.ParseBool(val)
	if firstSyncCompleted {
		return
	}

	if auth.NewTokenClient().GetTokenType() != auth.BearerToken {
		return
	}

	switch {
	case destinationPaths["cloudquery/postgresql"]:
		fmt.Println("🎉 Success!")
		fmt.Println()
		fmt.Println("Run the following command to get your oldest 10 EC2 instances:")
		fmt.Println()
		fmt.Println(`SELECT account_id, instance_id, region, launch_time FROM aws_ec2_instances ORDER BY launch_time ASC LIMIT 10`)
	case destinationPaths["cloudquery/sqlite"]:
		fmt.Println("🎉 Success!")
		fmt.Println()
		fmt.Println("Run the following command to get your oldest 10 EC2 instances:")
		fmt.Println()
		fmt.Println(`sqlite3 db.sql "SELECT account_id, instance_id, region, launch_time FROM aws_ec2_instances ORDER BY launch_time ASC LIMIT 10"`)
	default:
		return
	}

	if err := config.SetValue("first_sync_completed", "true"); err != nil {
		log.Debug().Err(err).Msg("Failed to set first_sync_completed")
	}
}

func deleteStale(client safeWriteClient, tables map[string]bool, sourceName string, syncTime time.Time) error {
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

// createTableNameSchema creates a minimal Arrow schema containing only table metadata
// This schema can be processed through the same transformation pipeline as MigrateTable
func createTableNameSchema(tableName string) *arrow.Schema {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{tableName})

	// Create a minimal schema with one dummy string field
	return arrow.NewSchema(
		[]arrow.Field{
			{Name: "_dummy", Type: arrow.BinaryTypes.String},
		},
		&md,
	)
}

// getTransformedTableNameFromSchema extracts the table name from a transformed schema
func getTransformedTableNameFromSchema(transformedSchemaBytes []byte) (string, error) {
	finalSchema, err := plugin.NewSchemaFromBytes(transformedSchemaBytes)
	if err != nil {
		return "", errors.Join(errors.New("failed to create schema from bytes"), err)
	}

	tableName, ok := finalSchema.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return "", errors.New("table name not found in transformed schema metadata")
	}

	return tableName, nil
}
