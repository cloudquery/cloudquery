package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/rs/zerolog/log"
)

// nolint:dupl
func migrateConnectionV3(ctx context.Context, sourceClient *managedplugin.Client, destinationsClients managedplugin.Clients, sourceSpec specs.Source, destinationSpecs []specs.Destination) error {
	destinationStrings := make([]string, len(destinationSpecs))
	for i := range destinationSpecs {
		destinationStrings[i] = destinationSpecs[i].VersionString()
	}
	migrateStart := time.Now().UTC()
	log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("Start migration")
	defer log.Info().Str("source", sourceSpec.Name).Strs("destinations", destinationStrings).Time("migrate_time", migrateStart).Msg("End migration")

	sourcePbClient := plugin.NewPluginClient(sourceClient.Conn)
	destinationsPbClients := make([]plugin.PluginClient, len(destinationsClients))
	destinationTransformers := make([]*transformer.RecordTransformer, len(destinationsClients))
	for i := range destinationsClients {
		destinationsPbClients[i] = plugin.NewPluginClient(destinationsClients[i].Conn)
		opts := []transformer.RecordTransformerOption{
			transformer.WithSourceNameColumn(sourceSpec.Name),
			transformer.WithSyncTimeColumn(migrateStart),
		}
		if destinationSpecs[i].WriteMode == specs.WriteModeAppend {
			opts = append(opts, transformer.WithRemovePKs())
		} else if destinationSpecs[i].PKMode == specs.PKModeCQID {
			opts = append(opts, transformer.WithRemovePKs())
			opts = append(opts, transformer.WithCQIDPrimaryKey())
		}
		destinationTransformers[i] = transformer.NewRecordTransformer(opts...)
	}

	// initialize destinations first, so that their connections may be used as backends by the source
	for i := range destinationsClients {
		destSpec := destinationSpecs[i]
		destSpecBytes, err := json.Marshal(destSpec.Spec)
		if err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Init(ctx, &plugin.Init_Request{
			Spec: destSpecBytes,
		}); err != nil {
			return err
		}
	}

	specBytes, err := json.Marshal(sourceSpec.Spec)
	if err != nil {
		return err
	}
	if _, err := sourcePbClient.Init(ctx, &plugin.Init_Request{
		Spec: specBytes,
	}); err != nil {
		return err
	}

	writeClients := make([]plugin.Plugin_WriteClient, len(destinationsPbClients))
	for i := range destinationsPbClients {
		writeClients[i], err = destinationsPbClients[i].Write(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Str("source", sourceSpec.VersionString()).Strs("destinations", destinationStrings).Msg("Start fetching resources")
	fmt.Printf("Starting migration for: %s -> %s\n", sourceSpec.VersionString(), destinationStrings)

	getTablesRes, err := sourcePbClient.GetTables(ctx, &plugin.GetTables_Request{
		Tables:     sourceSpec.Tables,
		SkipTables: sourceSpec.SkipTables,
	})
	if err != nil {
		return err
	}

	schemas, err := plugin.NewSchemasFromBytes(getTablesRes.Tables)
	if err != nil {
		return err
	}

	for i := range destinationsClients {
		for _, sc := range schemas {
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
		if _, err := writeClients[i].CloseAndRecv(); err != nil {
			return err
		}
		if _, err := destinationsPbClients[i].Close(ctx, &plugin.Close_Request{}); err != nil {
			return err
		}
	}

	migrateTimeTook := time.Since(migrateStart)
	fmt.Println("Migration completed successfully.")
	log.Info().Str("source", sourceSpec.Name).
		Strs("destinations", sourceSpec.Destinations).
		Float64("time_took", migrateTimeTook.Seconds()).
		Msg("Migration completed successfully")
	return nil
}
