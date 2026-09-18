package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"slices"
	"sync/atomic"
	"time"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/platform"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/utils"
)

// eventCounters counts what one summary interval did. Follow mode resets it
// after every summary; catch-up mode reports it once for the whole run.
type eventCounters struct {
	received      uint64
	refreshed     uint64
	deleted       uint64
	refreshErrors uint64
	resources     uint64
}

// coalesced is the number of received changes that did not turn into a refresh,
// either because a later change superseded them or because they were deletes.
func (c eventCounters) coalesced() uint64 {
	if c.received < c.refreshed {
		return 0
	}
	return c.received - c.refreshed
}

// watchMessage is one message off the Watch stream, or the error that ended it.
type watchMessage struct {
	change  *plugin.Watch_Change
	drained bool
	err     error
}

// batchEnd says why collectBatch stopped collecting.
type batchEnd int

const (
	// batchReady means batch_max_changes or batch_max_wait was reached.
	batchReady batchEnd = iota
	// batchDrained means the source has no more changes and closed the stream.
	batchDrained
	// batchStopped means the run was asked to stop.
	batchStopped
	// batchFailed means the Watch stream failed.
	batchFailed
)

// syncEventsV3 runs a `sync_mode: events` sync: it opens the source's Watch
// stream, optionally runs a full sync first, and then refreshes the resources
// the source reports as changed, in batches, acknowledging each batch only once
// every destination has written it.
func syncEventsV3(ctx context.Context, syncOptions syncV3Options) (syncErr error) {
	var (
		source          = syncOptions.source
		destinations    = syncOptions.destinations
		uid             = syncOptions.uid
		summaryLocation = syncOptions.summaryLocation
		shard           = syncOptions.shard
	)

	sourceSpec := source.spec
	if sourceSpec.SyncMode != specs.SyncModeEvents {
		return fmt.Errorf("syncEventsV3 called for sync_mode %q", sourceSpec.SyncMode)
	}
	eventsSpec := specs.EventsOptions{}
	if sourceSpec.Events != nil {
		eventsSpec = *sourceSpec.Events
	}
	eventsSpec.SetDefaults()

	sourceClient := source.client
	destinationSpecs := make([]specs.Destination, len(destinations))
	destinationsClients := make([]*managedplugin.Client, len(destinations))
	for i := range destinations {
		destinationSpecs[i] = destinations[i].spec
		destinationsClients[i] = destinations[i].client
	}

	var mt metrics.Metrics
	var exitReason = ExitReasonStopped

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
	runStart := time.Now().UTC().Truncate(time.Microsecond)
	sourceName := sourceSpec.Name
	if shard != nil {
		sourceName = fmt.Sprintf("%s_%d_%d", sourceName, shard.num, shard.total)
	}
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", runStart).Bool("follow", eventsSpec.Follow).Msg("Start events sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", runStart).Msg("End events sync")

	// A follow run has to survive the interrupt that stops it: it still has to
	// flush and acknowledge the batch it is holding. The plugin streams therefore
	// hang off a context we cancel ourselves, and ctx only tells us to stop.
	// Until the run reaches the batch loop, an interrupt still aborts it.
	streamCtx, cancelStreams := context.WithCancel(context.WithoutCancel(ctx))
	defer cancelStreams()
	var stopIsGraceful atomic.Bool
	go func() {
		select {
		case <-ctx.Done():
			if !stopIsGraceful.Load() {
				cancelStreams()
			}
		case <-streamCtx.Done():
		}
	}()

	sess, err := newV3Sync(streamCtx, v3SyncConfig{
		syncOptions:            syncOptions,
		sourceSpec:             sourceSpec,
		sourceName:             sourceName,
		syncTime:               runStart,
		destinationSpecs:       destinationSpecs,
		statsPerTable:          statsPerTable,
		totalResources:         &totalResources,
		skippedFromDeleteStale: make(map[string]bool),
		tablesForDeleteStale:   make(map[string]bool),
	})
	if err != nil {
		return err
	}
	sourceSpec = sess.sourceSpec

	if err := sess.startPipelines(streamCtx); err != nil {
		return err
	}
	// The pipelines stay open for the whole run: every refresh batch is relayed
	// through them, so they can only be closed once the run is over.
	defer func() {
		sess.closePipelines()
		if err := sess.eg.Wait(); err != nil && syncErr == nil {
			syncErr = err
		}
	}()

	fmt.Printf("Starting events sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	watchClient, ready, err := startWatch(streamCtx, sess.sourcePbClient, sourceSpec, shard, !eventsSpec.Follow)
	if err != nil {
		return err
	}

	needFull := needsFullSync(eventsSpec, ready, time.Now().UTC())
	log.Info().
		Bool("needs_full_sync", needFull).
		Bool("has_checkpoint", ready.GetHasCheckpoint()).
		Time("checkpoint_time", ready.GetCheckpointTime().AsTime()).
		Str("full_sync", eventsSpec.FullSync).
		Msg("Source watch stream is ready")

	if needFull {
		if err := sess.runFullSync(streamCtx, sourceSpec, shard, runStart); err != nil {
			return err
		}
		// Every record of the full sync has to be on its way to the destinations
		// before stale rows are deleted, and DeleteStale runs here and nowhere
		// else in an events run.
		if err := sess.waitForWrites(streamCtx); err != nil {
			return err
		}
		for i := range destinationsClients {
			if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
				// Table names might have changed due to transformers
				updatedTablesForDeleteStale := sess.tableNameChanger.UpdateTableNames(destinationSpecs[i].Name, sess.tablesForDeleteStale)
				if err := deleteStale(sess.writeClients[i], updatedTablesForDeleteStale, sourceName, runStart); err != nil {
					return err
				}
			}
		}
	}

	if err := watchClient.Send(positionRequest(needFull, ready, eventsSpec, runStart, time.Now().UTC())); err != nil {
		return watchError(sourceSpec.Name, err)
	}

	incoming := make(chan watchMessage)
	go receiveWatch(streamCtx, watchClient, incoming)

	// From here on an interrupt is a request to stop cleanly, not to abort.
	stopIsGraceful.Store(true)

	var (
		counters      eventCounters
		intervalStart = runStart
	)
	if needFull {
		counters.resources = uint64(atomic.LoadInt64(&totalResources))
	}
	summaryTicker := time.NewTicker(eventsSpec.SummaryInterval.Duration())
	defer summaryTicker.Stop()

	flushSummary := func(end time.Time) {
		sess.writeSummaries(summaryWrite{
			summaryLocation: summaryLocation,
			uid:             uid,
			runStart:        runStart,
			intervalStart:   intervalStart,
			intervalEnd:     end,
			checkpointTime:  ready.GetCheckpointTime().AsTime(),
			hasCheckpoint:   ready.GetHasCheckpoint(),
			shard:           shard,
			counters:        counters,
			tableDurations:  syncOptions.tableDurations,
		})
		counters = eventCounters{}
		intervalStart = end
	}

	for {
		changes, end, err := collectBatch(ctx, incoming, eventsSpec.BatchMaxChanges, eventsSpec.BatchMaxWait.Duration())
		if err != nil {
			return watchError(sourceSpec.Name, err)
		}
		counters.received += uint64(len(changes))

		if len(changes) > 0 {
			batchTime := time.Now().UTC().Truncate(time.Microsecond)
			batchCounts, err := sess.refreshBatch(streamCtx, sourceSpec, shard, changes, batchTime)
			if err != nil {
				return err
			}
			counters.refreshed += batchCounts.refreshed
			counters.deleted += batchCounts.deleted
			counters.refreshErrors += batchCounts.refreshErrors
			counters.resources += batchCounts.resources

			// Acknowledge only once every destination has written the batch, so
			// that a crash before that replays the changes instead of losing them.
			if err := sess.waitForWrites(streamCtx); err != nil {
				return err
			}
			lastEventID := changes[len(changes)-1].GetEventId()
			if err := watchClient.Send(&plugin.Watch_Request{
				Message: &plugin.Watch_Request_Ack{Ack: &plugin.Watch_Ack{EventId: lastEventID}},
			}); err != nil {
				return watchError(sourceSpec.Name, err)
			}
			log.Debug().Int("changes", len(changes)).Uint64("refreshed", batchCounts.refreshed).Str("event_id", lastEventID).Msg("Acknowledged events batch")
		}

		switch end {
		case batchDrained, batchStopped:
			syncTimeTook = time.Since(runStart)
			totals = sourceClient.Metrics()
			flushSummary(time.Now().UTC())
			if err := sess.closeDestinations(streamCtx); err != nil {
				return err
			}
			exitReason = ExitReasonCompleted
			reason := "caught up"
			if end == batchStopped {
				reason = "stopped"
			}
			fmt.Printf("Events sync %s. Resources: %d, Errors: %d, Warnings: %d, Time: %s\n", reason, atomic.LoadInt64(&totalResources), totals.Errors, totals.Warnings, syncTimeTook.Truncate(time.Second).String())
			return nil
		case batchFailed:
			return watchError(sourceSpec.Name, errors.New("watch stream ended unexpectedly"))
		case batchReady:
		}

		select {
		case <-summaryTicker.C:
			flushSummary(time.Now().UTC())
		default:
		}
	}
}

// needsFullSync decides whether an events run starts with a full sync, from the
// `events` options and the checkpoint the source reported in Ready.
func needsFullSync(events specs.EventsOptions, ready *plugin.Watch_Ready, now time.Time) bool {
	if events.FullSync == specs.FullSyncAlways {
		return true
	}
	if !ready.GetHasCheckpoint() {
		return events.FullSync == specs.FullSyncOnFirstRun
	}
	if events.MaxCheckpointAge == nil {
		return false
	}
	return now.Sub(ready.GetCheckpointTime().AsTime()) > events.MaxCheckpointAge.Duration()
}

// positionRequest builds the Position that tells the source where to resume:
// after a full sync, from the moment that sync started; on a first run that will
// never do a full sync, from now, so that the backlog is not replayed; otherwise
// from the source's own checkpoint.
func positionRequest(needFull bool, ready *plugin.Watch_Ready, events specs.EventsOptions, runStart, now time.Time) *plugin.Watch_Request {
	position := &plugin.Watch_Position{}
	switch {
	case needFull:
		position.Since = timestamppb.New(runStart)
	case !ready.GetHasCheckpoint() && events.FullSync == specs.FullSyncNever:
		position.Since = timestamppb.New(now)
	}
	return &plugin.Watch_Request{Message: &plugin.Watch_Request_Position{Position: position}}
}

// startWatch opens the Watch stream, sends Start and waits for Ready.
func startWatch(ctx context.Context, client plugin.PluginClient, sourceSpec specs.Source, shard *shard, drain bool) (plugin.Plugin_WatchClient, *plugin.Watch_Ready, error) {
	watchClient, err := client.Watch(ctx)
	if err != nil {
		return nil, nil, watchError(sourceSpec.Name, err)
	}
	start := &plugin.Watch_Start{
		Tables:     sourceSpec.Tables,
		SkipTables: sourceSpec.SkipTables,
		Drain:      drain,
	}
	if sourceSpec.BackendOptions != nil {
		start.Backend = &plugin.Sync_BackendOptions{
			TableName:  sourceSpec.BackendOptions.TableName,
			Connection: sourceSpec.BackendOptions.Connection,
		}
	}
	if shard != nil {
		start.Shard = &plugin.Sync_Request_Shard{Num: shard.num, Total: shard.total}
	}
	// On a plugin without Watch the Send succeeds and the Unimplemented status
	// only surfaces on the first Recv, so send errors other than EOF are the
	// only ones worth reporting here.
	if err := watchClient.Send(&plugin.Watch_Request{Message: &plugin.Watch_Request_Start{Start: start}}); err != nil && !errors.Is(err, io.EOF) {
		return nil, nil, watchError(sourceSpec.Name, err)
	}
	resp, err := watchClient.Recv()
	if err != nil {
		return nil, nil, watchError(sourceSpec.Name, err)
	}
	message, ok := resp.GetMessage().(*plugin.Watch_Response_Ready)
	if !ok {
		return nil, nil, fmt.Errorf("expected a Ready message from source plugin %s, got %T", sourceSpec.Name, resp.GetMessage())
	}
	return watchClient, message.Ready, nil
}

// watchError turns the Unimplemented a plugin without a Watch RPC returns into
// an actionable message, and passes every other error through.
func watchError(sourceName string, err error) error {
	if status.Code(err) == codes.Unimplemented {
		return fmt.Errorf("source plugin %s does not support sync_mode: events", sourceName)
	}
	return fmt.Errorf("failed to watch source plugin %s: %w", sourceName, err)
}

// receiveWatch forwards the Watch stream onto a channel so that the batch loop
// can put a deadline on it. It stops at the first error or at Drained.
func receiveWatch(ctx context.Context, watchClient plugin.Plugin_WatchClient, out chan<- watchMessage) {
	send := func(msg watchMessage) bool {
		select {
		case out <- msg:
			return true
		case <-ctx.Done():
			return false
		}
	}
	for {
		resp, err := watchClient.Recv()
		if err != nil {
			send(watchMessage{err: err})
			return
		}
		switch m := resp.GetMessage().(type) {
		case *plugin.Watch_Response_Change:
			if !send(watchMessage{change: m.Change}) {
				return
			}
		case *plugin.Watch_Response_Drained:
			send(watchMessage{drained: true})
			return
		default:
			send(watchMessage{err: fmt.Errorf("unexpected watch response %T", resp.GetMessage())})
			return
		}
	}
}

// collectBatch reads changes until the batch is full, until maxWait has passed
// since the batch's first change, or until the stream drains or fails or the run
// is asked to stop. It always returns the changes it collected, in receive order,
// so the caller can flush and acknowledge them.
func collectBatch(ctx context.Context, in <-chan watchMessage, maxChanges int, maxWait time.Duration) ([]*plugin.Watch_Change, batchEnd, error) {
	batch := make([]*plugin.Watch_Change, 0, maxChanges)
	var (
		timer    *time.Timer
		deadline <-chan time.Time
	)
	defer func() {
		if timer != nil {
			timer.Stop()
		}
	}()
	for {
		select {
		case msg, ok := <-in:
			switch {
			case !ok:
				return batch, batchFailed, nil
			case msg.err != nil:
				return batch, batchFailed, msg.err
			case msg.drained:
				return batch, batchDrained, nil
			}
			batch = append(batch, msg.change)
			if len(batch) >= maxChanges {
				return batch, batchReady, nil
			}
			if timer == nil {
				timer = time.NewTimer(maxWait)
				deadline = timer.C
			}
		case <-deadline:
			return batch, batchReady, nil
		case <-ctx.Done():
			return batch, batchStopped, nil
		}
	}
}

// coalesceChanges collapses a batch to one resource per (table, primary keys),
// keeping the last operation seen for that resource. The order follows each
// resource's first appearance, so the refresh request is deterministic.
func coalesceChanges(changes []*plugin.Watch_Change) ([]*plugin.Sync_Resource, error) {
	positions := make(map[string]int, len(changes))
	resources := make([]*plugin.Sync_Resource, 0, len(changes))
	for _, change := range changes {
		keys, err := plugin.NewRecordFromBytes(change.GetKeys())
		if err != nil {
			return nil, fmt.Errorf("failed to decode the keys of event %s on table %s: %w", change.GetEventId(), change.GetTable(), err)
		}
		key := change.GetTable() + "\x00" + schema.KeysString(keys)
		resource := &plugin.Sync_Resource{Table: change.GetTable(), Keys: change.GetKeys(), Op: change.GetOp()}
		if i, ok := positions[key]; ok {
			resources[i] = resource
			continue
		}
		positions[key] = len(resources)
		resources = append(resources, resource)
	}
	return resources, nil
}

// batchCounts is what one refresh batch moved to the destinations.
type batchCounts struct {
	refreshed     uint64
	deleted       uint64
	refreshErrors uint64
	resources     uint64
}

// refreshBatch coalesces the changes, asks the source to refresh them and relays
// the result to the destinations stamped with the batch's own sync time.
func (s *v3Sync) refreshBatch(ctx context.Context, sourceSpec specs.Source, shard *shard, changes []*plugin.Watch_Change, batchTime time.Time) (batchCounts, error) {
	var counts batchCounts

	resources, err := coalesceChanges(changes)
	if err != nil {
		return counts, err
	}
	for _, resource := range resources {
		if resource.GetOp() == plugin.Op_OP_UPSERT {
			counts.refreshed++
		}
	}

	syncReq := &plugin.Sync_Request{
		Refresh:           resources,
		DeterministicCqId: sourceSpec.DeterministicCQID,
		WithErrorMessages: true,
	}
	if sourceSpec.BackendOptions != nil {
		syncReq.Backend = &plugin.Sync_BackendOptions{
			TableName:  sourceSpec.BackendOptions.TableName,
			Connection: sourceSpec.BackendOptions.Connection,
		}
	}
	if shard != nil {
		syncReq.Shard = &plugin.Sync_Request_Shard{Num: shard.num, Total: shard.total}
	}

	syncClient, err := s.sourcePbClient.Sync(ctx, syncReq)
	if err != nil {
		return counts, fmt.Errorf("failed to refresh %d resources: %w", len(resources), err)
	}
	totals, err := s.relay(ctx, syncClient, batchTime)
	if err != nil {
		return counts, fmt.Errorf("failed to refresh %d resources: %w", len(resources), err)
	}
	if err := syncClient.CloseSend(); err != nil {
		return counts, err
	}
	counts.deleted = uint64(totals.Deletes)
	counts.refreshErrors = uint64(totals.Errors)
	counts.resources = uint64(totals.Rows)
	return counts, nil
}

// runFullSync syncs every selected table, exactly as `sync_mode: full` does.
func (s *v3Sync) runFullSync(ctx context.Context, sourceSpec specs.Source, shard *shard, syncTime time.Time) error {
	log.Info().Str("source", sourceSpec.VersionString()).Msg("Start fetching resources")
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
		syncReq.Shard = &plugin.Sync_Request_Shard{Num: shard.num, Total: shard.total}
	}
	syncClient, err := s.sourcePbClient.Sync(ctx, syncReq)
	if err != nil {
		return err
	}
	if _, err := s.relay(ctx, syncClient, syncTime); err != nil {
		return err
	}
	return syncClient.CloseSend()
}

// closeDestinations closes every destination write stream and the destination
// plugins themselves, as the full sync does when it finishes.
func (s *v3Sync) closeDestinations(ctx context.Context) error {
	for i := range s.destinationsPbClients {
		if _, err := s.writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
		if _, err := s.destinationsPbClients[i].Close(ctx, &plugin.Close_Request{}); err != nil {
			return err
		}
	}
	return nil
}

// summaryWrite is one summary row's worth of run and interval state.
type summaryWrite struct {
	summaryLocation string
	uid             string
	runStart        time.Time
	intervalStart   time.Time
	intervalEnd     time.Time
	checkpointTime  time.Time
	hasCheckpoint   bool
	shard           *shard
	counters        eventCounters
	tableDurations  func() map[string]time.Duration
}

// writeSummaries persists and sends one summary row per destination. Summary
// failures are logged and never fail the run, because a follow run writes a
// summary on every interval.
func (s *v3Sync) writeSummaries(w summaryWrite) {
	tableProgress := s.statsPerTable.GetAll()
	sourceTableNames := lo.Keys(s.sourceTables)
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
	if w.tableDurations != nil {
		data := w.tableDurations()
		durationsPerTable = lo.Reduce(sourceTableNames, func(acc map[string]uint64, tableName string, _ int) map[string]uint64 {
			acc[tableName] = uint64(data[tableName].Milliseconds())
			return acc
		}, durationsPerTable)
	}

	var sourceErrors uint64
	for _, val := range tableProgress {
		sourceErrors += uint64(val.Errors)
	}

	for i := range s.destinationsPbClients {
		summary := syncSummary{
			Resources:           w.counters.resources,
			SourceErrors:        sourceErrors,
			SyncID:              w.uid,
			SyncTime:            w.runStart,
			SourceName:          s.sourceSpec.Name,
			SyncDurationMs:      uint64(w.intervalEnd.Sub(w.runStart).Milliseconds()),
			SourceVersion:       s.sourceSpec.Version,
			SourcePath:          s.sourceSpec.Path,
			SourceTables:        s.tableNameChanger.UpdateTableNamesSlice(s.destinationSpecs[i].Name, sourceTableNames),
			CLIVersion:          Version,
			DestinationName:     s.destinationSpecs[i].Name,
			DestinationVersion:  s.destinationSpecs[i].Version,
			DestinationPath:     s.destinationSpecs[i].Path,
			ResourcesPerTable:   resourcesPerTable,
			ErrorsPerTable:      errorsPerTable,
			DurationsPerTableMs: durationsPerTable,
			IntervalStart:       lo.ToPtr(w.intervalStart),
			IntervalEnd:         lo.ToPtr(w.intervalEnd),
			EventsReceived:      w.counters.received,
			EventsCoalesced:     w.counters.coalesced(),
			EventsRefreshed:     w.counters.refreshed,
			EventsDeleted:       w.counters.deleted,
			RefreshErrors:       w.counters.refreshErrors,
		}
		if w.hasCheckpoint {
			summary.CheckpointTime = lo.ToPtr(w.checkpointTime)
		}
		if s.syncGroupID[i] != "" {
			summary.SyncGroupID = lo.ToPtr(s.syncGroupID[i])
		}
		if w.shard != nil {
			summary.ShardNum = lo.ToPtr(int(w.shard.num))
			summary.ShardTotal = lo.ToPtr(int(w.shard.total))
		}

		if err := persistSummary(w.summaryLocation, summary); err != nil {
			log.Warn().Err(err).Msg("Failed to persist sync summary")
		}
		log.Info().Interface("summary", summary).Msg("Sync summary")
		if !s.destinationSpecs[i].SyncSummary {
			continue
		}
		if err := sendSummary(s.writeClients[i], s.destinationSpecs[i], s.destinationsClients[i], s.destinationTransformers[i], &summary, s.noMigrate); err != nil {
			log.Warn().Err(err).Str("destination", s.destinationSpecs[i].Name).Msg("Failed to send sync summary")
		}
	}
}
