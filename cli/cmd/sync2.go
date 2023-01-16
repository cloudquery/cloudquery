package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

func syncConnectionV2(ctx context.Context, cqDir string, sourceClient *clients.SourceClient, sourceSpec specs.Source, destinationsSpecs []specs.Destination, uid string, noMigrate bool) error {
	var err error
	destinationNames := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationNames[i] = destinationsSpecs[i].Name
	}
	syncTime := time.Now().UTC()

	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("End sync")

	destClients, err := newDestinationClients(ctx, sourceSpec, destinationsSpecs, cqDir)
	if err != nil {
		return err
	}
	defer destClients.Close()

	selectedTables, tablesForSpecSupported, err := getTablesForSpec(ctx, sourceClient, sourceSpec)
	if err != nil {
		return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
	}

	tableCount := len(selectedTables.FlattenTables())

	// Print a count of the tables that will be synced / migrated.
	if tablesForSpecSupported {
		word := "tables"
		if tableCount == 1 {
			word = "table"
		}
		if noMigrate {
			fmt.Printf("Source %s will sync %d %s.\n", sourceSpec.Name, tableCount, word)
		} else {
			fmt.Printf("Source %s will migrate and sync %d %s.\n", sourceSpec.Name, tableCount, word)
		}
	}

	if !noMigrate {
		fmt.Println("Starting migration for:", sourceSpec.Name, "->", sourceSpec.Destinations)
		log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Start migration")
		migrateStart := time.Now()

		for i, destinationSpec := range destinationsSpecs {
			// Currently we migrate all tables, but this is subject to change once policies
			// are adapted to handle non-existent tables in some way.
			if err := destClients[i].Migrate(ctx, selectedTables); err != nil {
				return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destinationSpec.Name, err)
			}
		}
		migrateTimeTook := time.Since(migrateStart)
		fmt.Printf("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.Name).
			Strs("destinations", sourceSpec.Destinations).
			Int("num_tables", tableCount).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	resources := make(chan []byte)
	g, gctx := errgroup.WithContext(ctx)
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Start fetching resources")
	fmt.Println("Starting sync for:", sourceSpec.Name, "->", sourceSpec.Destinations)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync2(gctx, sourceSpec, resources); err != nil {
			if isUnknownConcurrencyFieldError(err) {
				return fmt.Errorf("unsupported version of source %s@%s. Please update to the latest version from https://cloudquery.io/docs/plugins/sources", sourceSpec.Name, sourceSpec.Version)
			}
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
		return nil
	})

	destSubscriptions := make([]chan []byte, len(sourceSpec.Destinations))
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
	for i, destination := range sourceSpec.Destinations {
		i := i
		destination := destination
		g.Go(func() error {
			var destFailedWrites uint64
			var err error
			if err = destClients[i].Write2(gctx, sourceSpec, selectedTables, syncTime, destSubscriptions[i]); err != nil {
				return fmt.Errorf("failed to write for %s->%s: %w", sourceSpec.Name, destination, err)
			}
			// call Close on destination client using the outer context, so that it happens even if writes get cancelled
			if err := destClients[i].Close(ctx); err != nil {
				return fmt.Errorf("failed to close destination client for %s->%s: %w", sourceSpec.Name, destination, err)
			}
			failedWrites += destFailedWrites
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
		_ = bar.Finish()
		return err
	}
	_ = bar.Finish()
	syncTimeTook := time.Since(syncTime)

	metrics, err := sourceClient.GetMetrics(ctx)
	if err != nil {
		return fmt.Errorf("failed to get metrics for source %s: %w", sourceSpec.Name, err)
	}

	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", metrics.TotalResources(), metrics.TotalErrors(), metrics.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	// fmt.Printf("Summary: resources: %d, errors: %d, panic: %d, failed_writes: %d, time: %s\n", summary.Resources, summary.Errors, summary.Panics, failedWrites, tt.Truncate(time.Second).String())
	// log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).
	// Uint64("resources", totalResources).Uint64("errors", summary.Errors).Uint64("panic", summary.Panics).Uint64("failed_writes", failedWrites).Float64("time_took", tt.Seconds()).Msg("Sync completed successfully")

	// Send analytics, if activated. We only send if the source plugin registry is GitHub, mostly to avoid sending data from development machines.
	if analyticsClient != nil && sourceSpec.Registry == specs.RegistryGithub {
		log.Info().Msg("Sending sync summary to " + analyticsHost)
		if err := analyticsClient.SendSyncMetrics(ctx, sourceSpec, destinationsSpecs, uid, metrics); err != nil {
			log.Warn().Err(err).Msg("Failed to send sync summary")
		}
	}
	return nil
}

// getTablesForSpec first tries the newer GetTablesForSpec call, but if it is not available, falls back to
// GetTables. The returned `supported` value indicates whether GetTablesForSpec was supported by the server.
func getTablesForSpec(ctx context.Context, sourceClient *clients.SourceClient, sourceSpec specs.Source) (tables schema.Tables, supported bool, err error) {
	tables, err = sourceClient.GetTablesForSpec(ctx, &sourceSpec)
	if clients.IsUnimplemented(err) {
		// the plugin server does not support GetTablesForSpec. Fall back to GetTables.
		tables, err = sourceClient.GetTables(ctx)
		return tables, false, err
	} else if err != nil {
		// the method is supported, but failed for some other reason
		return tables, true, err
	}

	allTables, err := sourceClient.GetTables(ctx)
	if err != nil {
		return tables, true, fmt.Errorf("failed to get all tables for source %s: %w", sourceSpec.Name, err)
	}

	// make sure selected tables only includes top-level tables; we don't want a flattened list
	// (a bug in early versions of GetTablesForSpec returned a flattened list)
	tables = topLevelTables(allTables, tables)
	return tables, true, err
}

// returns only the top-level tables in the given tables list, i.e. tables
// with no parents
func topLevelTables(allTables, tables schema.Tables) schema.Tables {
	var top schema.Tables
	for _, t := range tables {
		if allTables.GetTopLevel(t.Name) == nil {
			continue
		}
		top = append(top, t)
	}
	return top
}
