package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"strconv"
	gosync "sync"
	"sync/atomic"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/ghodss/yaml"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/thoas/go-funk"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/tablenamechanger"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformer"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformerpipeline"
	"github.com/cloudquery/cloudquery/cli/v6/internal/utils"

	_ "embed"
)

type destinationsQueries struct {
	Query string   `yaml:"query"`
	Args  []string `yaml:"args"`
}

type tableQueries struct {
	Dependencies []string                         `yaml:"dependencies"`
	Description  string                           `yaml:"description"`
	Destinations []map[string]destinationsQueries `yaml:"destinations"`
}

var (
	//go:embed data/tables_queries.yml
	tablesQueriesString []byte
	tablesQueries       []map[string]tableQueries
)

func init() {
	if err := yaml.Unmarshal(tablesQueriesString, &tablesQueries); err != nil {
		panic(err)
	}
	for _, tableQueries := range tablesQueries {
		if len(tableQueries) != 1 {
			panic(fmt.Sprintf("expected 1 table query, got %d", len(tableQueries)))
		}
		tableQuery := lo.Values(tableQueries)[0]
		for _, destinationQuery := range tableQuery.Destinations {
			destinationPaths := lo.Keys(destinationQuery)
			if len(destinationPaths) != 1 {
				panic(fmt.Sprintf("expected 1 destination query, got %d", len(destinationPaths)))
			}
		}
	}
}

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
	tableDurations            func() map[string]time.Duration
}

func syncConnectionV3(ctx context.Context, syncOptions syncV3Options) (syncErr error) {
	var (
		source                    = syncOptions.source
		destinations              = syncOptions.destinations
		transformersByDestination = syncOptions.transformersByDestination
		uid                       = syncOptions.uid
		noMigrate                 = syncOptions.noMigrate
		summaryLocation           = syncOptions.summaryLocation
		shard                     = syncOptions.shard
	)

	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
	skippedFromDeleteStale := make(map[string]bool)
	tablesForDeleteStale := make(map[string]bool)

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
		statsPerTable  = utils.NewConcurrentMap[string, SyncRunTableProgressValue]()
	)
	defer func() {
		// Platform-only syncs are external syncs: the platform emits the canonical
		// sync_run_completed once ingestion settles, so sending it here too would
		// double-count. See platform.OnlyPlatformDestinations.
		if platform.OnlyPlatformDestinations(destinationSpecs) {
			return
		}
		analytics.TrackSyncCompleted(ctx, invocationUUID.UUID, analytics.SyncFinishedEvent{
			SyncStartedEvent:  syncStartedEvent,
			Errors:            totals.Errors,
			Warnings:          totals.Warnings,
			Duration:          syncTimeTook,
			ResourceCount:     totalResources,
			AbortedDueToError: syncErr,
		})
	}()

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

	sess, err := newV3Sync(ctx, v3SyncConfig{
		syncOptions:            syncOptions,
		sourceSpec:             sourceSpec,
		sourceName:             sourceName,
		syncTime:               syncTime,
		destinationSpecs:       destinationSpecs,
		statsPerTable:          statsPerTable,
		totalResources:         &totalResources,
		skippedFromDeleteStale: skippedFromDeleteStale,
		tablesForDeleteStale:   tablesForDeleteStale,
	})
	if err != nil {
		return err
	}
	sourceSpec = sess.sourceSpec
	destinationsPbClients := sess.destinationsPbClients
	destinationTransformers := sess.destinationTransformers
	writeClients := sess.writeClients
	syncGroupId := sess.syncGroupID

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
	syncClient, err := sess.sourcePbClient.Sync(ctx, syncReq)
	if err != nil {
		return err
	}

	// Read from the sync stream and write to all destinations.
	isComplete := int64(0)

	if err := sess.startPipelines(ctx); err != nil {
		return err
	}
	tableNameChanger := sess.tableNameChanger
	sourceTables := sess.sourceTables

	sess.eg.Go(func() error {
		// Close all transformation pipelines when the source is done
		defer sess.closePipelines()
		_, err := sess.relay(ctx, syncClient, syncTime)
		return err
	})
	if err := sess.eg.Wait(); err != nil { // wait for source & transformers to finish. If any fails, sync fails.
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

	sourceTableNames := lo.Keys(sourceTables)
	slices.Sort(sourceTableNames)

	resourcesPerTable := lo.Reduce(sourceTableNames, func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
		acc[tableName] = uint64(tableProgress[tableName].Rows)
		return acc
	}, map[string]uint64{})

	errorsPerTable := lo.Reduce(sourceTableNames, func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
		acc[tableName] = uint64(tableProgress[tableName].Errors)
		return acc
	}, map[string]uint64{})

	durationsPerTable := map[string]uint64{}
	if syncOptions.tableDurations != nil {
		data := syncOptions.tableDurations()
		durationsPerTable = lo.Reduce(sourceTableNames, func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
			acc[tableName] = uint64(data[tableName].Milliseconds())
			return acc
		}, durationsPerTable)
	}
	syncTimeTook = time.Since(syncTime)
	for i := range destinationsClients {
		if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
			// Table names might have changed due to transformers
			updatedTablesForDeleteStale := tableNameChanger.UpdateTableNames(destinationSpecs[i].Name, tablesForDeleteStale)
			if err := deleteStale(writeClients[i], updatedTablesForDeleteStale, sourceName, syncTime); err != nil {
				return err
			}
		}
	}

	for i := range destinationsClients {
		m := destinationsClients[i].Metrics()
		summary := syncSummary{
			Resources:           uint64(totalResources),
			SourceErrors:        sourceErrors,
			SourceWarnings:      sourceWarnings,
			SyncID:              uid,
			SyncTime:            syncTime,
			SourceName:          sourceSpec.Name,
			SyncDurationMs:      uint64(syncTimeTook.Milliseconds()),
			SourceVersion:       sourceSpec.Version,
			SourcePath:          sourceSpec.Path,
			SourceTables:        tableNameChanger.UpdateTableNamesSlice(destinationSpecs[i].Name, sourceTableNames),
			CLIVersion:          Version,
			DestinationErrors:   m.Errors,
			DestinationWarnings: m.Warnings,
			DestinationName:     destinationSpecs[i].Name,
			DestinationVersion:  destinationSpecs[i].Version,
			DestinationPath:     destinationSpecs[i].Path,
			ResourcesPerTable:   resourcesPerTable,
			ErrorsPerTable:      errorsPerTable,
			DurationsPerTableMs: durationsPerTable,
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

	for i := range destinationsClients {
		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Close(ctx, &plugin.Close_Request{}); err != nil {
			return err
		}
	}

	atomic.StoreInt64(&isComplete, 1)

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
		hintSelectMessage(destinationSpecs, statsPerTable, sourceTables)
	}

	return nil
}

type SyncRunTableProgressValue struct {
	Errors int64 `json:"errors"`
	Rows   int64 `json:"rows"`
}

func hintSelectMessage(destinationSpecs []specs.Destination, statsPerTable *utils.ConcurrentMap[string, SyncRunTableProgressValue], sourceTables map[string]bool) {
	val, _ := config.GetValue("first_sync_completed")
	firstSyncCompleted, _ := strconv.ParseBool(val)
	if firstSyncCompleted {
		return
	}

	if auth.NewTokenClient().GetTokenType() != auth.BearerToken {
		return
	}

	ensureDependencies := func(tableQuery tableQueries) bool {
		for _, dependency := range tableQuery.Dependencies {
			_, ok := sourceTables[dependency]
			if !ok {
				return false
			}
		}
		return true
	}

	destPaths := lo.SliceToMap(destinationSpecs, func(spec specs.Destination) (string, specs.Destination) {
		return spec.Path, spec
	})

	for _, tablePair := range tablesQueries {
		tables := lo.Keys(tablePair)
		if len(tables) == 0 {
			continue
		}
		tableName := tables[0]
		rows, ok := statsPerTable.Get(tableName)
		if !ok || rows.Rows == 0 {
			continue
		}

		tableQuery := tablePair[tableName]

		if !ensureDependencies(tableQuery) {
			continue
		}

		for _, destinationQuery := range tableQuery.Destinations {
			destinationPaths := lo.Keys(destinationQuery)
			if len(destinationPaths) == 0 {
				continue
			}
			destinationPath := destinationPaths[0]
			destSpec, ok := destPaths[destinationPath]
			if !ok {
				continue
			}

			fmt.Println()
			fmt.Println("🎉 Success!")
			fmt.Println()
			fmt.Println(tableQuery.Description)
			fmt.Println()

			args := lo.Map(destinationQuery[destinationPath].Args, func(arg string, _ int) any {
				return funk.Get(destSpec.Spec, arg, funk.WithAllowZero())
			})

			fmt.Printf(destinationQuery[destinationPath].Query+"\n", args...)

			if err := config.SetValue("first_sync_completed", "true"); err != nil {
				log.Debug().Err(err).Msg("Failed to set first_sync_completed")
			}
			return
		}
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

// v3SyncConfig is what newV3Sync needs from its caller: the run's identity and
// the shared bookkeeping the caller reports on once the sync ends.
type v3SyncConfig struct {
	syncOptions      syncV3Options
	sourceSpec       specs.Source
	sourceName       string
	syncTime         time.Time
	destinationSpecs []specs.Destination

	statsPerTable  *utils.ConcurrentMap[string, SyncRunTableProgressValue]
	totalResources *int64

	skippedFromDeleteStale map[string]bool
	tablesForDeleteStale   map[string]bool
}

// v3Sync is the initialized wiring for one v3 sync: the source, destinations,
// backend and transformers, their open write and transform streams, the
// transformer pipelines records are fanned out through, and the bookkeeping the
// relay loop fills in. `sync_mode: full` and `sync_mode: events` share it; they
// differ only in how they drive the source.
type v3Sync struct {
	// sourceSpec is the spec after `@@plugins.*.connection` variable replacement.
	sourceSpec          specs.Source
	sourceName          string
	sourceClient        *managedplugin.Client
	destinationSpecs    []specs.Destination
	destinationsClients []*managedplugin.Client
	noMigrate           bool

	sourcePbClient                    plugin.PluginClient
	destinationsPbClients             []plugin.PluginClient
	transformerPbClientsByDestination map[string][]plugin.PluginClient
	transformClientsByDestination     map[string][]plugin.Plugin_TransformClient

	destinationTransformers []*transformer.RecordTransformer
	writeClients            []safeWriteClient
	syncGroupID             []string

	isStateBackendEnabled bool
	tableNameChanger      *tablenamechanger.TableNameChanger
	statsPerTable         *utils.ConcurrentMap[string, SyncRunTableProgressValue]

	sourceTables           map[string]bool
	skippedFromDeleteStale map[string]bool
	tablesForDeleteStale   map[string]bool

	totalResources *int64
	newResources   int64

	// sentToPipeline and writtenToDestination count the records handed to each
	// destination's transformer pipeline and the records that came out of it and
	// were written. The events loop waits for the two to match before it
	// acknowledges a batch, because a pipeline hands records on asynchronously.
	sentToPipeline       []int64
	writtenToDestination []int64

	bar                       progressBar
	eg                        *errgroup.Group
	gctx                      context.Context
	pipelineByDestinationName map[string]*transformerpipeline.TransformerPipeline
}

// newV3Sync creates the plugin clients, initializes every plugin and opens the
// destination write streams and the transformer transform streams.
func newV3Sync(ctx context.Context, cfg v3SyncConfig) (*v3Sync, error) {
	var (
		source                    = cfg.syncOptions.source
		destinations              = cfg.syncOptions.destinations
		transformersByDestination = cfg.syncOptions.transformersByDestination
		backend                   = cfg.syncOptions.backend
		uid                       = cfg.syncOptions.uid
		cqColumnsNotNull          = cfg.syncOptions.cqColumnsNotNull
	)
	sourceSpec := cfg.sourceSpec
	destinationSpecs := cfg.destinationSpecs

	s := &v3Sync{
		sourceName:                        cfg.sourceName,
		sourceClient:                      source.client,
		destinationSpecs:                  destinationSpecs,
		destinationsClients:               make([]*managedplugin.Client, len(destinations)),
		noMigrate:                         cfg.syncOptions.noMigrate,
		destinationsPbClients:             make([]plugin.PluginClient, len(destinations)),
		transformerPbClientsByDestination: map[string][]plugin.PluginClient{},
		transformClientsByDestination:     map[string][]plugin.Plugin_TransformClient{},
		destinationTransformers:           make([]*transformer.RecordTransformer, len(destinations)),
		writeClients:                      make([]safeWriteClient, len(destinations)),
		syncGroupID:                       make([]string, len(destinations)),
		statsPerTable:                     cfg.statsPerTable,
		sourceTables:                      map[string]bool{},
		skippedFromDeleteStale:            cfg.skippedFromDeleteStale,
		tablesForDeleteStale:              cfg.tablesForDeleteStale,
		totalResources:                    cfg.totalResources,
		sentToPipeline:                    make([]int64, len(destinations)),
		writtenToDestination:              make([]int64, len(destinations)),
	}
	for i := range destinations {
		s.destinationsClients[i] = destinations[i].client
	}

	variables := specs.Variables{
		Plugins: make(map[string]specs.PluginVariables),
	}
	s.sourcePbClient = plugin.NewPluginClient(source.client.Conn)
	for name, transformers := range transformersByDestination {
		for _, tf := range transformers {
			s.transformerPbClientsByDestination[name] = append(s.transformerPbClientsByDestination[name], plugin.NewPluginClient(tf.client.Conn))
		}
	}

	backendPbClient := plugin.PluginClient(nil)
	for i := range destinations {
		s.destinationsPbClients[i] = plugin.NewPluginClient(destinations[i].client.Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(cfg.sourceName),
			transformer.WithSyncTimeColumn(cfg.syncTime),
		}
		if cqColumnsNotNull {
			opts = append(opts, transformer.WithCQColumnsNotNull())
		}
		if destinationSpecs[i].SyncGroupId != "" {
			s.syncGroupID[i] = destinationSpecs[i].RenderedSyncGroupId(cfg.syncTime, uid)
			opts = append(opts, transformer.WithSyncGroupIdColumn(s.syncGroupID[i]))
		}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithRemoveUniqueConstraints())
		} else if destinationSpecs[i].PKMode == specs.PKModeCQID {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithCQIDPrimaryKey())
		}
		s.destinationTransformers[i] = transformer.NewRecordTransformer(opts...)
		connection := destinations[i].client.ConnectionString()
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
		if err := initPlugin(ctx, s.destinationsPbClients[i], destinationSpec.Spec, false, uid); err != nil {
			if platform.IsInjectedDestination(destinationSpec.Name) {
				// Surface the platform's own message (e.g. an unsupported-version
				// 422) instead of the gRPC/plugin-sdk wrapper chain.
				return nil, errors.New(platform.CleanInitError(err))
			}
			return nil, fmt.Errorf("failed to init destination %v: %w", destinationSpec.Name, err)
		}
	}
	if backend != nil {
		if err := initPlugin(ctx, backendPbClient, backend.spec.Spec, false, uid); err != nil {
			return nil, fmt.Errorf("failed to init backend %v: %w", backend.spec.Name, err)
		}
	}
	for name, transformers := range transformersByDestination {
		for i, tf := range transformers {
			if err := initPlugin(ctx, s.transformerPbClientsByDestination[name][i], tf.spec.Spec, false, uid); err != nil {
				return nil, fmt.Errorf("failed to init transformer %v: %w", tf.spec.Name, err)
			}
		}
	}

	// replace @@plugins.name.connection with the actual GRPC connection string from the client
	// NOTE: if this becomes a stable feature, it can move out of sync_v3 and into sync.go
	specBytes, err := json.Marshal(sourceSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal source spec JSON before variable replacement: %w", err)
	}
	specBytesExpanded, err := specs.ReplaceVariables(string(specBytes), variables, source.shouldReplaceLocalhost)
	if err != nil {
		return nil, fmt.Errorf("failed to replace variables: %w", err)
	}
	if err := json.Unmarshal([]byte(specBytesExpanded), &sourceSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source spec JSON after variable replacement: %w", err)
	}

	if err := initPlugin(ctx, s.sourcePbClient, sourceSpec.Spec, false, uid); err != nil {
		return nil, fmt.Errorf("failed to init source %v: %w", sourceSpec.Name, err)
	}
	s.sourceSpec = sourceSpec
	s.isStateBackendEnabled = sourceSpec.BackendOptions != nil && sourceSpec.BackendOptions.TableName != ""

	for i := range s.destinationsPbClients {
		writeClient, err := s.destinationsPbClients[i].Write(ctx)
		if err != nil {
			return nil, err
		}
		s.writeClients[i] = newSafeWriteClient(writeClient)
	}
	for name, transformerPbClients := range s.transformerPbClientsByDestination {
		for _, transformerPbClient := range transformerPbClients {
			transformClient, err := transformerPbClient.Transform(ctx)
			if err != nil {
				return nil, err
			}
			s.transformClientsByDestination[name] = append(s.transformClientsByDestination[name], transformClient)
		}
	}

	return s, nil
}

// startPipelines starts the progress bar and one transformer pipeline per
// destination. Records are relayed only after this returns.
func (s *v3Sync) startPipelines(ctx context.Context) error {
	s.bar = progressBar(noopProgressBar{})
	if !logConsole {
		s.bar = progressbar.NewOptions(-1,
			progressbar.OptionSetDescription("Syncing resources..."),
			progressbar.OptionSetItsString("resources"),
			progressbar.OptionShowIts(),
			progressbar.OptionSetElapsedTime(true),
			progressbar.OptionShowCount(),
			progressbar.OptionClearOnFinish(),
		)
	}

	// Transformers can change table names. We need to keep track of table name changes
	// in case we do things that depend on table names.
	//
	// Note that transformers run per-destination, so we need to keep track of table name changes per-destination.
	s.tableNameChanger = tablenamechanger.New(s.destinationSpecs)

	// Note: we want to stop this errorgroup if ctx is cancelled, but we don't want to cancel ctx if gctx is cancelled.
	// gctx is always cancelled when the errorgroup returns, and this isn't necessarily an error.
	eg, gctx := errgroup.WithContext(ctx)
	s.eg = eg
	s.pipelineByDestinationName = map[string]*transformerpipeline.TransformerPipeline{}

	// Each destination has its own transformer pipeline
	for i := range s.destinationsPbClients {
		destinationName := s.destinationSpecs[i].Name

		// Start a pipeline of transformers that will receive & transform the source records
		var (
			pipeline *transformerpipeline.TransformerPipeline
			err      error
		)
		pipeline, gctx, err = transformerpipeline.New(gctx, s.transformClientsByDestination[destinationName])
		if err != nil {
			return fmt.Errorf("failed to create transformer pipeline: %w", err)
		}
		writeClient := s.writeClients[i]
		written := &s.writtenToDestination[i]
		err = pipeline.OnOutput(func(recordBytes []byte) error {
			wr := &plugin.Write_Request{
				Message: &plugin.Write_Request_Insert{
					Insert: &plugin.Write_MessageInsert{
						Record: recordBytes,
					},
				},
			}
			if err := writeClient.Send(wr); err != nil {
				return handleSendError(err, writeClient, "insert")
			}
			atomic.AddInt64(written, 1)
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to create register pipeline output: %w", err)
		}
		eg.Go(pipeline.RunBlocking) // each transformer runs in its own goroutine
		s.pipelineByDestinationName[destinationName] = pipeline
	}
	s.gctx = gctx

	// Add a ticker to update the progress bar every 100ms
	t := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-gctx.Done():
				t.Stop()
				change := atomic.SwapInt64(&s.newResources, 0)
				_ = s.bar.Add(int(change))
				if err := s.bar.Finish(); err != nil {
					log.Warn().Err(err).Msg("Failed to finish progress bar")
				}
				return
			case <-t.C:
				change := atomic.SwapInt64(&s.newResources, 0)
				_ = s.bar.Add(int(change))
			}
		}
	}()

	return nil
}

// closePipelines closes every transformer pipeline. It must be called once the
// last record has been relayed, and never between two relay passes.
func (s *v3Sync) closePipelines() {
	for _, pipeline := range s.pipelineByDestinationName {
		pipeline.Close()
	}
}

// relayTotals reports what one pass over a source Sync stream moved to the
// destinations.
type relayTotals struct {
	// Rows is the number of rows in the relayed Insert messages.
	Rows int64
	// Deletes is the number of DeleteRecord messages relayed per destination.
	Deletes int64
	// Errors is the number of Error messages received from the source.
	Errors int64
	// Failures are the Error messages themselves, in the order they arrived.
	Failures []relayFailure
}

type relayFailure struct {
	Table string
	Error string
}

// relay reads a source Sync stream to its end and writes every message it
// carries to each destination, stamping the records with syncTime. The full sync
// passes the run's start time; an events refresh batch passes its own batch time.
func (s *v3Sync) relay(ctx context.Context, syncClient plugin.Plugin_SyncClient, syncTime time.Time) (relayTotals, error) {
	var totals relayTotals
	for {
		r, err := syncClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return totals, fmt.Errorf("unexpected error from sync client receive: %w", err)
		}
		syncResponseMsg := r.GetMessage()
		switch m := syncResponseMsg.(type) {
		case *plugin.Sync_Response_Insert:
			recordBytes := m.Insert.Record
			record, err := plugin.NewRecordFromBytes(recordBytes)
			if err != nil {
				return totals, fmt.Errorf("failed to get record from bytes: %w", err)
			}

			atomic.AddInt64(&s.newResources, record.NumRows())
			atomic.AddInt64(s.totalResources, record.NumRows())
			totals.Rows += record.NumRows()
			tableName, _ := record.Schema().Metadata().GetValue(schema.MetadataTableName)
			s.sourceTables[tableName] = true
			stats, _ := s.statsPerTable.Get(tableName)
			stats.Rows += record.NumRows()
			s.statsPerTable.Add(tableName, stats)
			for i := range s.destinationsPbClients {
				destinationName := s.destinationSpecs[i].Name
				transformedRecord := s.destinationTransformers[i].TransformWithSyncTime(record, syncTime)
				transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
				if err != nil {
					return totals, fmt.Errorf("failed to transform record bytes: %w", err)
				}
				atomic.AddInt64(&s.sentToPipeline[i], 1)
				if err := s.pipelineByDestinationName[destinationName].Send(transformedRecordBytes); err != nil {
					atomic.AddInt64(&s.sentToPipeline[i], -1)
					if errors.Is(err, transformerpipeline.ErrPipelineClosed) {
						// If the pipeline is closed, we should stop processing records. The error that cause the pipeline to close
						// Will be returned by the pipeline itself.
						return totals, nil
					}
					return totals, err
				}
			}
		case *plugin.Sync_Response_DeleteRecord:
			totals.Deletes++
			for i := range s.destinationsPbClients {
				destinationName := s.destinationSpecs[i].Name

				// Create a minimal schema with the table name, similar to MigrateTable
				originalSchema := createTableNameSchema(m.DeleteRecord.TableName)

				// Apply the same transformation pipeline as MigrateTable
				transformedSchema := s.destinationTransformers[i].TransformSchema(originalSchema)
				transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
				if err != nil {
					return totals, err
				}

				for _, transformerPbClient := range s.transformerPbClientsByDestination[destinationName] {
					resp, err := transformerPbClient.TransformSchema(ctx, &plugin.TransformSchema_Request{Schema: transformedSchemaBytes})
					if err != nil {
						return totals, err
					}
					transformedSchemaBytes = resp.Schema
				}

				tableName, err := getTransformedTableNameFromSchema(transformedSchemaBytes)
				if err != nil {
					// Fall back to tableNameChanger if extraction fails
					log.Warn().Err(err).Str("table", m.DeleteRecord.TableName).Str("destination", destinationName).Msg("Failed to extract table name from transformed schema, falling back to tableNameChanger")
					tableName = s.tableNameChanger.UpdateTableName(destinationName, m.DeleteRecord.TableName)
				}

				wr := &plugin.Write_Request{}
				wr.Message = &plugin.Write_Request_DeleteRecord{
					DeleteRecord: &plugin.Write_MessageDeleteRecord{
						TableName:      tableName,
						TableRelations: m.DeleteRecord.TableRelations,
						WhereClause:    m.DeleteRecord.WhereClause,
					},
				}
				if err := s.writeClients[i].Send(wr); err != nil {
					return totals, handleSendError(err, s.writeClients[i], "delete")
				}
			}
		case *plugin.Sync_Response_MigrateTable:
			sc, err := plugin.NewSchemaFromBytes(m.MigrateTable.Table)
			if err != nil {
				return totals, err
			}
			table, err := schema.NewTableFromArrowSchema(sc)
			if err != nil {
				return totals, err
			}

			s.sourceTables[table.Name] = true
			// This works since we sync and send migrate messages for parents before children
			if s.isStateBackendEnabled && (table.IsIncremental || (table.Parent != nil && s.skippedFromDeleteStale[table.Parent.Name])) {
				s.skippedFromDeleteStale[table.Name] = true
			} else {
				s.tablesForDeleteStale[table.Name] = true
			}
			if s.noMigrate {
				continue
			}
			for i := range s.destinationsPbClients {
				destinationName := s.destinationSpecs[i].Name
				transformedSchema := s.destinationTransformers[i].TransformSchema(sc)
				transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
				if err != nil {
					return totals, err
				}
				// Sequentially apply schema transformations from transformers
				for _, transformerPbClient := range s.transformerPbClientsByDestination[destinationName] {
					resp, err := transformerPbClient.TransformSchema(ctx, &plugin.TransformSchema_Request{Schema: transformedSchemaBytes})
					if err != nil {
						return totals, err
					}
					transformedSchemaBytes = resp.Schema
				}

				// Table name might have changed due to a transformation.
				if err := s.tableNameChanger.LearnTableNameChange(destinationName, table.Name, transformedSchemaBytes); err != nil {
					return totals, err
				}

				wr := &plugin.Write_Request{}
				wr.Message = &plugin.Write_Request_MigrateTable{
					MigrateTable: &plugin.Write_MessageMigrateTable{
						MigrateForce: s.destinationSpecs[i].MigrateMode == specs.MigrateModeForced,
						Table:        transformedSchemaBytes,
					},
				}
				if err := s.writeClients[i].Send(wr); err != nil {
					return totals, handleSendError(err, s.writeClients[i], "migrate")
				}
			}
		case *plugin.Sync_Response_Error:
			totals.Errors++
			totals.Failures = append(totals.Failures, relayFailure{Table: m.Error.TableName, Error: m.Error.Error})
			stats, _ := s.statsPerTable.Get(m.Error.TableName)
			stats.Errors++
			s.statsPerTable.Add(m.Error.TableName, stats)
		default:
			return totals, fmt.Errorf("unknown message type: %T", m)
		}
	}
	return totals, nil
}

// writesSettled reports whether every record handed to a transformer pipeline
// has come out of it and been written to its destination.
func (s *v3Sync) writesSettled() bool {
	for i := range s.sentToPipeline {
		if atomic.LoadInt64(&s.writtenToDestination[i]) < atomic.LoadInt64(&s.sentToPipeline[i]) {
			return false
		}
	}
	return true
}

// destinationWriteTimeout bounds waitForWrites. It is deliberately not a spec
// field: exceeding it means the barrier's assumption is broken, not that the
// destination is slow in a way the user should tune around.
const destinationWriteTimeout = 60 * time.Second

// errWritesNotSettled is returned when the destination write barrier does not
// clear. Its wording matches destinationWriteTimeout.
var errWritesNotSettled = errors.New("destination writes did not settle within 60s; a transformer that drops or merges records is not supported in sync_mode: events")

// waitForWrites blocks until every record relayed so far has been written to
// every destination. A transformer pipeline hands records on asynchronously, so
// the events loop needs this barrier before it deletes stale rows or
// acknowledges a batch of changes.
//
// The barrier counts records into each pipeline and out of it, so it assumes a
// transformer emits exactly one record for every record it is given. A
// transformer that drops, merges or splits records never lets the two counts
// meet, which is what destinationWriteTimeout bounds.
//
// runCtx is the context the user's interrupt cancels and streamCtx is the one
// the plugin streams hang off; either ending aborts the wait. A run that is
// already stopping is not treated as an interruption: its final batch still has
// to settle before it can be acknowledged, and the timeout bounds that.
func (s *v3Sync) waitForWrites(runCtx, streamCtx context.Context) error {
	return s.waitForWritesWithin(runCtx, streamCtx, destinationWriteTimeout)
}

func (s *v3Sync) waitForWritesWithin(runCtx, streamCtx context.Context, timeout time.Duration) error {
	interrupted := runCtx.Done()
	if runCtx.Err() != nil {
		interrupted = nil
	}
	deadline := time.NewTimer(timeout)
	defer deadline.Stop()
	t := time.NewTicker(2 * time.Millisecond)
	defer t.Stop()
	for !s.writesSettled() {
		select {
		case <-s.gctx.Done():
			// A transformer pipeline or a destination write failed. Close the
			// pipelines so that the errorgroup reports the error that caused it.
			s.closePipelines()
			if err := s.eg.Wait(); err != nil {
				return err
			}
			return errors.New("destination pipeline stopped before every record was written")
		case <-interrupted:
			return fmt.Errorf("%w: %w", errWritesNotSettled, runCtx.Err())
		case <-streamCtx.Done():
			return fmt.Errorf("%w: %w", errWritesNotSettled, streamCtx.Err())
		case <-deadline.C:
			return errWritesNotSettled
		case <-t.C:
		}
	}
	return nil
}
