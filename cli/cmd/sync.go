package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/clients"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	fetchShort   = "Sync resources from configured source plugins to destinations"
	fetchExample = `# Sync resources from configuration in a directory
cloudquery sync ./directory
# Sync resources from directories and files
cloudquery sync ./directory ./aws.yml ./pg.yml
`
)

func NewCmdSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sync [files or directories]",
		Short:   fetchShort,
		Long:    fetchShort,
		Example: fetchExample,
		Args:    cobra.MinimumNArgs(1),
		RunE:    sync,
	}
	return cmd
}

func sync(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	log.Info().Strs("args", args).Msg("Loading spec(s)")
	fmt.Printf("Loading spec(s) from %s\n", strings.Join(args, ", "))
	specReader, err := specs.NewSpecReader(args)
	if err != nil {
		return fmt.Errorf("failed to load spec(s) from %s. Error: %w", strings.Join(args, ", "), err)
	}

	for _, sourceSpec := range specReader.Sources {
		if len(sourceSpec.Destinations) == 0 {
			return fmt.Errorf("no destinations found for source %s", sourceSpec.Name)
		}
		var destinationsSpecs []specs.Destination
		for _, destination := range sourceSpec.Destinations {
			spec := specReader.Destinations[destination]
			if spec == nil {
				return fmt.Errorf("failed to find destination %s in source %s", destination, sourceSpec.Name)
			}
			destinationsSpecs = append(destinationsSpecs, *spec)
		}
		if err := syncConnection(ctx, *sourceSpec, destinationsSpecs); err != nil {
			return fmt.Errorf("failed to sync source %s: %w", sourceSpec.Name, err)
		}
	}

	return nil
}

func syncConnection(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination) error {
	destinationNames := make([]string, len(destinationsSpecs))
	for i := range destinationsSpecs {
		destinationNames[i] = destinationsSpecs[i].Name
	}
	syncTime := time.Now().UTC()

	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationNames).Time("sync_time", syncTime).Msg("End sync")

	sourceClient, err := clients.NewSourceClient(ctx, sourceSpec.Registry, sourceSpec.Path, sourceSpec.Version,
		clients.WithSourceLogger(log.Logger),
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

	destClients := make([]*clients.DestinationClient, len(sourceSpec.Destinations))
	destSubscriptions := make([]chan []byte, len(sourceSpec.Destinations))
	for i := range destSubscriptions {
		destSubscriptions[i] = make(chan []byte)
	}
	defer func() {
		for _, destClient := range destClients {
			if destClient != nil {
				if err := destClient.Terminate(); err != nil {
					log.Error().Err(err).Msg("Failed to terminate destination client")
					fmt.Println("failed to terminate destination client: ", err)
				}
			}
		}
	}()
	for i, destinationSpec := range destinationsSpecs {
		destClients[i], err = clients.NewDestinationClient(ctx, destinationSpec.Registry, destinationSpec.Path, destinationSpec.Version,
			clients.WithDestinationLogger(log.Logger),
		)
		if err != nil {
			return fmt.Errorf("failed to create destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		if err := destClients[i].Initialize(ctx, destinationSpec); err != nil {
			return fmt.Errorf("failed to initialize destination plugin client for %s: %w", destinationSpec.Name, err)
		}
		tables, err := sourceClient.GetTables(ctx)
		if err != nil {
			return fmt.Errorf("failed to get tables for source %s: %w", sourceSpec.Name, err)
		}

		if err := destClients[i].Migrate(ctx, tables); err != nil {
			return fmt.Errorf("failed to migrate source %s on destination %s : %w", sourceSpec.Name, destinationSpec.Name, err)
		}
	}

	resources := make(chan []byte)
	g, gctx := errgroup.WithContext(ctx)
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).Msg("Start fetching resources")
	fmt.Println("Starting sync for: ", sourceSpec.Name, "->", sourceSpec.Destinations)
	g.Go(func() error {
		defer close(resources)
		if err := sourceClient.Sync(gctx, sourceSpec, resources); err != nil {
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
	totalResources := 0
	for i, destination := range sourceSpec.Destinations {
		i := i
		destination := destination
		g.Go(func() error {
			var destFailedWrites uint64
			var err error
			if destFailedWrites, err = destClients[i].Write(gctx, sourceSpec.Name, syncTime, destSubscriptions[i]); err != nil {
				return fmt.Errorf("failed to write for %s->%s: %w", sourceSpec.Name, destination, err)
			}
			if destClients[i].Close(ctx); err != nil {
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
	tt := time.Since(syncTime)

	fmt.Println("Sync completed successfully.")
	fmt.Printf("Summary: resources: %d, errors: %d, panic: %d, failed_writes: %d, time: %s\n", totalResources, summary.Errors, summary.Panics, failedWrites, tt.Truncate(time.Second).String())
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", sourceSpec.Destinations).
		Int("resources", totalResources).Uint64("errors", summary.Errors).Uint64("panic", summary.Panics).Uint64("failed_writes", failedWrites).Float64("time_took", tt.Seconds()).Msg("Sync completed successfully")
	return nil
}
