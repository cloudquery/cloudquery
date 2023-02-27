package cmd

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/cloudquery/plugin-sdk/clients/source/v1"
	pluginsSource "github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

func syncConnectionV1(ctx context.Context, cqDir string, sourceSpec specs.Source, destinationsSpecs []specs.Destination, uid string, noMigrate bool) error {
	var metrics *pluginsSource.Metrics
	exitReason := "unknown"
	canceled := false
	defer func() {
		// Send analytics, if activated.
		if analyticsClient != nil {
			log.Info().Msg("Sending sync summary to " + analyticsClient.Host())
			if err := analyticsClient.SendSyncMetrics(context.Background(), sourceSpec, destinationsSpecs, uid, metrics, exitReason); err != nil {
				log.Warn().Err(err).Msg("Failed to send sync summary")
			}
		}
	}()
	opts := []source.ClientOption{
		source.WithLogger(log.Logger),
		source.WithDirectory(cqDir),
	}
	if disableSentry {
		opts = append(opts, source.WithNoSentry())
	}
	sourceCtx, sourceCancel := context.WithCancel(context.Background())
	defer sourceCancel()

	sourceClient, err := source.NewClient(sourceCtx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version, opts...)
	if err != nil {
		exitReason = "failed to get source plugin client"
		sourceCancel()
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer func() {
		if canceled {
			return
		}
		if err := sourceClient.Terminate(); err != nil {
			log.Error().Err(err).Msg("Failed to terminate source client")
			fmt.Println("failed to terminate source client: ", err)
		}
	}()

	syncTime := time.Now().UTC()
	destinationStrings := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationStrings[i] = destinationsSpecs[i].VersionString()
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	destCtx, destCancel := context.WithCancel(context.Background())
	defer destCancel()

	destClients, err := newDestinationClientsV0(destCtx, sourceSpec, destinationsSpecs, cqDir)
	if err != nil {
		exitReason = "failed to get destination plugin client"
		return err
	}
	defer func() {
		if canceled {
			return
		}
		destClients.Close()
	}()

	go func() {
		<-ctx.Done()
		if metrics == nil {
			// If we didn't get metrics because sync got interrupted, try to get them
			// now, before closing the source client.
			metrics, err = sourceClient.GetMetrics(sourceCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get metrics")
			}
		}
		canceled = true
		sourceCancel()
		destCancel()
	}()

	if err := sourceClient.Init(sourceCtx, sourceSpec); err != nil {
		exitReason = "failed to init"
		return fmt.Errorf("failed to init source %s: %w", sourceSpec.VersionString(), err)
	}

	tables, err := sourceClient.GetDynamicTables(sourceCtx)
	if err != nil {
		exitReason = "failed to get tables"
		return fmt.Errorf("failed to get dynamic tables for source %s: %w", sourceSpec.VersionString(), err)
	}

	tableCount := len(tables.FlattenTables())

	if !noMigrate {
		fmt.Printf("Starting migration with %d tables for: %s -> %s\n", tableCount, sourceSpec.VersionString(), destinationStrings)
		log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start migration")
		migrateStart := time.Now()

		for i, destinationSpec := range destinationsSpecs {
			if err := destClients[i].Migrate(sourceCtx, tables); err != nil {
				exitReason = "failed to migrate"
				return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.VersionString(), destinationSpec.VersionString(), err)
			}
		}
		migrateTimeTook := time.Since(migrateStart)
		fmt.Printf("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.VersionString()).
			Strs("destinations", destinationStrings).
			Int("num_tables", tableCount).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	resources := make(chan []byte)
	g, gctx := errgroup.WithContext(sourceCtx)
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync(gctx, resources); err != nil {
			if isUnknownConcurrencyFieldError(err) {
				return fmt.Errorf("unsupported version of source %s. Please update to the latest version from https://cloudquery.io/docs/plugins/sources", sourceSpec.VersionString())
			}
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.VersionString(), err)
		}
		return nil
	})

	destSubscriptions := make([]chan []byte, len(destinationsSpecs))
	for i := range destSubscriptions {
		destSubscriptions[i] = make(chan []byte)
	}
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Syncing resources..."),
		progressbar.OptionSetItsString("resources"),
		progressbar.OptionShowIts(),
		progressbar.OptionSetElapsedTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionClearOnFinish(),
	)
	failedWrites := uint64(0)
	totalResources := uint64(0)
	for i, destination := range destinationsSpecs {
		i := i
		destination := destination
		g.Go(func() error {
			var destFailedWrites uint64
			var err error
			if err = destClients[i].Write2(gctx, sourceSpec, tables, syncTime, destSubscriptions[i]); err != nil {
				return fmt.Errorf("failed to write for %s -> %s: %w", sourceSpec.VersionString(), destination.VersionString(), err)
			}
			// call Close on destination client using the outer context, so that it happens even if writes get cancelled
			if err := destClients[i].Close(ctx); err != nil {
				return fmt.Errorf("failed to close destination client for %s -> %s: %w", sourceSpec.VersionString(), destination.VersionString(), err)
			}
			atomic.AddUint64(&failedWrites, destFailedWrites)
			return nil
		})
	}

	g.Go(func() error {
		t := time.NewTicker(1 * time.Second)
		defer func() {
			for i := range destSubscriptions {
				close(destSubscriptions[i])
			}
			t.Stop()
		}()
		for {
			select {
			case resource, ok := <-resources:
				if !ok {
					return nil
				}
				totalResources++
				_ = bar.Add(1)
				for i := range destSubscriptions {
					select {
					case <-gctx.Done():
						return gctx.Err()
					case destSubscriptions[i] <- resource:
					}
				}
			case <-t.C:
				_ = bar.Add(0)
			case <-gctx.Done():
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		exitReason = "sync failed"
		if canceled {
			exitReason = "sync canceled"
		}
		_ = bar.Finish()
		return err
	}
	_ = bar.Finish()
	syncTimeTook := time.Since(syncTime)

	metrics, err = sourceClient.GetMetrics(sourceCtx)
	if err != nil {
		exitReason = "failed to get metrics"
		return fmt.Errorf("failed to get metrics for source %s: %w", sourceSpec.VersionString(), err)
	}

	exitReason = "success"
	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", metrics.TotalResources(), metrics.TotalErrors(), metrics.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	return nil
}
