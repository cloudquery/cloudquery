package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/clients/source/v0"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

func syncConnectionV0_1(ctx context.Context, cqDir string, sourceClient *source.Client, sourceSpec specs.Source, destinationsSpecs []specs.Destination, uid string, noMigrate bool) error {
	var err error
	destinationNames := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationNames[i] = destinationsSpecs[i].Name
	}
	syncTime := time.Now().UTC()

	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("End sync")

	sourceClient, err = source.NewClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version,
		source.WithLogger(log.Logger),
		source.WithDirectory(cqDir),
	)
	if err != nil {
		return fmt.Errorf("failed to get source plugin client for %s: %w", sourceSpec.Name, err)
	}
	defer func() {
		if err := sourceClient.Terminate(); err != nil {
			log.Error().Err(err).Msg("Failed to terminate source client")
			fmt.Println("failed to terminate source client: ", err)
		}
	}()

	destClients, err := newDestinationClientsV0(ctx, sourceSpec, destinationsSpecs, cqDir)
	if err != nil {
		return err
	}
	defer destClients.Close()

	if !noMigrate {
		fmt.Println("Starting migration for:", sourceSpec.Name, "->", sourceSpec.Destinations)
		log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Start migration")
		tables, err := sourceClient.GetTables(ctx)
		if err != nil {
			return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
		}
		for i, destinationSpec := range destinationsSpecs {
			if err := destClients[i].Migrate(ctx, tables); err != nil {
				return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destinationSpec.Name, err)
			}
		}
		migrateTimeTook := time.Since(syncTime)
		fmt.Printf("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.Name).
			Strs("destinations", sourceSpec.Destinations).
			Int("num_tables", len(tables)).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	resources := make(chan []byte)
	g, gctx := errgroup.WithContext(ctx)
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Start fetching resources")
	fmt.Println("Starting sync for:", sourceSpec.Name, "->", sourceSpec.Destinations)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync(gctx, sourceSpec, resources); err != nil {
			if isUnknownConcurrencyFieldError(err) {
				return fmt.Errorf("unsupported version of source %s@%s. Please update to the latest version from https://cloudquery.io/docs/plugins/sources", sourceSpec.Name, sourceSpec.Version)
			}
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
		return nil
	})

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
	destSubscriptions := make([]chan []byte, len(sourceSpec.Destinations))
	for i := range destSubscriptions {
		destSubscriptions[i] = make(chan []byte)
	}
	for i, destination := range sourceSpec.Destinations {
		i := i
		destination := destination
		g.Go(func() error {
			var destFailedWrites uint64
			var err error
			if destFailedWrites, err = destClients[i].Write(gctx, sourceSpec.Name, syncTime, destSubscriptions[i]); err != nil {
				return fmt.Errorf("failed to write for %s->%s: %w", sourceSpec.Name, destination, err)
			}
			if err := destClients[i].Close(ctx); err != nil {
				return fmt.Errorf("failed to close destination client for %s->%s: %w", sourceSpec.Name, destination, err)
			}
			failedWrites += destFailedWrites
			return nil
		})
	}

	g.Go(func() error {
		for resource := range resources {
			totalResources++
			_ = bar.Add(1)
			for i := range destSubscriptions {
				select {
				case <-gctx.Done():
					return gctx.Err()
				case destSubscriptions[i] <- resource:
				}
			}
		}
		for i := range destSubscriptions {
			close(destSubscriptions[i])
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		_ = bar.Finish()
		return err
	}
	summary, err := sourceClient.GetSyncSummary(ctx)
	if err != nil {
		return fmt.Errorf("failed to get sync summary: %w", err)
	}
	_ = bar.Finish()
	syncTimeTook := time.Since(syncTime)

	fmt.Println("Sync completed successfully.")
	fmt.Printf("Summary: resources: %d, errors: %d, panic: %d, failed_writes: %d, time: %s\n", summary.Resources, summary.Errors, summary.Panics, failedWrites, syncTimeTook.Truncate(time.Second).String())
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).
		Uint64("resources", totalResources).Uint64("errors", summary.Errors).Uint64("panic", summary.Panics).Uint64("failed_writes", failedWrites).Float64("time_took", syncTimeTook.Seconds()).Msg("Sync completed successfully")

	// Send analytics, if activated. We only send if the source plugin registry is GitHub, mostly to avoid sending data from development machines.
	if analyticsClient != nil && sourceSpec.Registry == specs.RegistryGithub {
		log.Info().Msg("Sending sync summary to " + analyticsHost)
		if err := analyticsClient.SendSyncSummary(ctx, sourceSpec, destinationsSpecs, uid, *summary); err != nil {
			log.Warn().Err(err).Msg("Failed to send sync summary")
		}
	}
	return nil
}

func isUnknownConcurrencyFieldError(err error) bool {
	return strings.Contains(err.Error(), unknownFieldErrorPrefix+`"table_concurrency"`) || strings.Contains(err.Error(), unknownFieldErrorPrefix+`"resource_concurrency"`)
}
