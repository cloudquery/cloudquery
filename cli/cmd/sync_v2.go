package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"slices"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/destination/v1"
	pluginv3 "github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-pb-go/pb/source/v2"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getSourceV2DestV3DestinationsTransformers(destinationSpecs []specs.Destination, destinationsVersions [][]int) []*transformer.RecordTransformer {
	destinationsTransformers := make([]*transformer.RecordTransformer, 0, len(destinationsVersions))
	for i := range destinationsVersions {
		// We only need to transform to destinations that are v3
		if !slices.Contains(destinationsVersions[i], 3) {
			destinationsTransformers = append(destinationsTransformers, nil)
			continue
		}
		opts := []transformer.RecordTransformerOption{}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs(), transformer.WithRemovePKs())
			opts = append(opts, transformer.WithRemoveUniqueConstraints())
		} else if destinationSpecs[i].PKMode == specs.PKModeCQID {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithCQIDPrimaryKey())
		}
		destinationsTransformers = append(destinationsTransformers, transformer.NewRecordTransformer(opts...))
	}
	return destinationsTransformers
}

func transformSourceV2DestV3Schemas(originalSchemas [][]byte, recordTransformer *transformer.RecordTransformer) ([][]byte, error) {
	if recordTransformer == nil {
		return originalSchemas, nil
	}
	transformedSchemasBytes := make([][]byte, 0, len(originalSchemas))
	for _, s := range originalSchemas {
		schema, err := pluginv3.NewSchemaFromBytes(s)
		if err != nil {
			return nil, err
		}
		transformedSchema := recordTransformer.TransformSchema(schema)
		transformedSchemaBytes, err := pluginv3.SchemaToBytes(transformedSchema)
		if err != nil {
			return nil, err
		}
		transformedSchemasBytes = append(transformedSchemasBytes, transformedSchemaBytes)
	}
	return transformedSchemasBytes, nil
}

func transformSourceV2DestV3Resource(originalResourceBytes []byte, recordTransformer *transformer.RecordTransformer) ([]byte, error) {
	if recordTransformer == nil {
		return originalResourceBytes, nil
	}
	resource, err := pluginv3.NewRecordFromBytes(originalResourceBytes)
	if err != nil {
		return nil, err
	}
	transformedResource := recordTransformer.Transform(resource)
	transformedResourceBytes, err := pluginv3.RecordToBytes(transformedResource)
	if err != nil {
		return nil, err
	}
	return transformedResourceBytes, nil
}

// nolint:dupl
func syncConnectionV2(ctx context.Context, sourceClient *managedplugin.Client, destinationsClients managedplugin.Clients, sourceSpec specs.Source, destinationSpecs []specs.Destination, uid string, noMigrate bool, destinationsVersions [][]int) error {
	var mt metrics.Metrics
	var exitReason = ExitReasonStopped
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
	destinationStrings := make([]string, len(destinationsClients))
	for i := range destinationsClients {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}
	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("Start sync")
	defer log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Time("sync_time", syncTime).Msg("End sync")

	sourcePbClient := source.NewSourceClient(sourceClient.Conn)
	destinationsPbClients := make([]destination.DestinationClient, len(destinationsClients))
	destinationsTransformers := getSourceV2DestV3DestinationsTransformers(destinationSpecs, destinationsVersions)
	for i := range destinationsClients {
		destinationsPbClients[i] = destination.NewDestinationClient(destinationsClients[i].Conn)
	}

	specBytes, err := json.Marshal(CLISourceSpecToPbSpec(sourceSpec))
	if err != nil {
		return err
	}
	if _, err := sourcePbClient.Init(ctx, &source.Init_Request{
		Spec: specBytes,
	}); err != nil {
		return err
	}
	tablesRes, err := sourcePbClient.GetDynamicTables(ctx, &source.GetDynamicTables_Request{})
	if err != nil {
		return err
	}
	for i := range destinationsClients {
		destSpecBytes, err := json.Marshal(CLIDestinationSpecToPbSpec(destinationSpecs[i]))
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Configure(ctx, &destination.Configure_Request{
			Config: destSpecBytes,
		}); err != nil {
			return err
		}
	}

	transformedSchemasBytes := make([][][]byte, 0, len(destinationsPbClients))
	for i := range destinationsPbClients {
		destinationSchemasBytes, err := transformSourceV2DestV3Schemas(tablesRes.Tables, destinationsTransformers[i])
		if err != nil {
			return err
		}
		transformedSchemasBytes = append(transformedSchemasBytes, destinationSchemasBytes)
	}

	if !noMigrate {
		migrateStart := time.Now().UTC()
		fmt.Printf("Starting migration for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)
		for i := range destinationsClients {
			if _, err := destinationsPbClients[i].Migrate(ctx, &destination.Migrate_Request{
				Tables: transformedSchemasBytes[i],
			}); err != nil {
				return err
			}
		}
		migrateTimeTook := time.Since(migrateStart)
		fmt.Printf("Migration completed successfully.\n")
		log.Info().
			Str("source", sourceSpec.VersionString()).
			Strs("destinations", destinationStrings).
			Float64("time_took", migrateTimeTook.Seconds()).
			Msg("End migration")
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting sync for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	syncClient, err := sourcePbClient.Sync(ctx, &source.Sync_Request{
		SyncTime: timestamppb.New(syncTime),
	})
	if err != nil {
		return err
	}
	writeClients := make([]destination.Destination_WriteClient, len(destinationsPbClients))
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
		if err := writeClients[i].Send(&destination.Write_Request{
			Source:    sourceSpec.Name,
			Tables:    transformedSchemasBytes[i],
			Timestamp: timestamppb.New(syncTime),
		}); err != nil {
			return err
		}
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
	for {
		r, err := syncClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		_ = bar.Add(1)
		for i := range destinationsPbClients {
			transformedResourceBytes, err := transformSourceV2DestV3Resource(r.Resource, destinationsTransformers[i])
			if err != nil {
				return err
			}
			if err := writeClients[i].Send(&destination.Write_Request{
				Resource: transformedResourceBytes,
			}); err != nil {
				return err
			}
		}
	}
	for i := range destinationsClients {
		if destinationSpecs[i].WriteMode == specs.WriteModeOverwriteDeleteStale {
			_, err := destinationsPbClients[i].DeleteStale(ctx, &destination.DeleteStale_Request{
				Tables:    transformedSchemasBytes[i],
				Source:    sourceSpec.Name,
				Timestamp: timestamppb.New(syncTime),
			})
			if err != nil {
				return err
			}
		}

		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Close(ctx, &destination.Close_Request{}); err != nil {
			return err
		}
	}

	getMetricsRes, err := sourcePbClient.GetMetrics(ctx, &source.GetMetrics_Request{})
	if err != nil {
		return err
	}
	if err := json.Unmarshal(getMetricsRes.Metrics, &mt); err != nil {
		return err
	}

	err = bar.Finish()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to finish progress bar")
	}
	syncTimeTook := time.Since(syncTime)
	exitReason = ExitReasonCompleted
	fmt.Printf("Sync completed successfully. Resources: %d, Errors: %d, Panics: %d, Time: %s\n", mt.TotalResources(), mt.TotalErrors(), mt.TotalPanics(), syncTimeTook.Truncate(time.Second).String())
	return nil
}
