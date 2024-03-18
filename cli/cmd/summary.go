package cmd

import (
	"fmt"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/thoas/go-funk"
)

type syncSummary struct {
	CliVersion          string
	DestinationErrors   uint64
	DestinationName     string
	DestinationPath     string
	DestinationVersion  string
	DestinationWarnings uint64
	Resources           uint64
	SourceErrors        uint64
	SourceName          string
	SourcePath          string
	SourceVersion       string
	SourceWarnings      uint64
	SyncID              string
}

func generateSummaryTable() *schema.Table {
	tableName := "cloudquery_sync_summary"
	t := schema.Tables{{
		Name:      tableName,
		Transform: transformers.TransformWithStruct(&syncSummary{}),
		Columns:   []schema.Column{},
	}}

	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}

	return t[0]
}

func generateSummaryMigrateMessage(destTransformer *transformer.RecordTransformer, destSpec specs.Destination) (*plugin.Write_Request, error) {
	// generate schema of the summary table
	summaryTable := generateSummaryTable()
	summaryTableSchema := summaryTable.ToArrowSchema()
	transformedSchema := destTransformer.TransformSchema(summaryTableSchema)
	transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
	if err != nil {
		return nil, err
	}

	return &plugin.Write_Request{
		Message: &plugin.Write_Request_MigrateTable{
			MigrateTable: &plugin.Write_MessageMigrateTable{
				MigrateForce: destSpec.MigrateMode == specs.MigrateModeForced,
				Table:        transformedSchemaBytes,
			},
		},
	}, nil
}

func sendSummary(writeClients []plugin.Plugin_WriteClient, destinationSpecs []specs.Destination, destinationsClients []*managedplugin.Client, destinationTransformers []*transformer.RecordTransformer, summary *syncSummary, noMigrate bool) error {
	summaryTable := generateSummaryTable()
	defaultCaser := caser.New()
	for i := range destinationsClients {
		if !destinationSpecs[i].SyncSummary {
			continue
		}
		wr := &plugin.Write_Request{}
		// Respect the noMigrate flag
		if !noMigrate {
			wr, err := generateSummaryMigrateMessage(destinationTransformers[i], destinationSpecs[i])
			if err != nil {
				return err
			}
			if err := writeClients[i].Send(wr); err != nil {
				return handleSendError(err, writeClients[i], "migrate sync summary table")
			}
		}

		// Get Information about the DestinationPlugin
		m := destinationsClients[i].Metrics()
		summary.DestinationErrors = m.Errors
		summary.DestinationWarnings = m.Warnings

		summary.DestinationName = destinationSpecs[i].Name
		summary.DestinationVersion = destinationSpecs[i].Version
		summary.DestinationPath = destinationSpecs[i].Path

		resource := schema.NewResourceData(summaryTable, nil, nil)
		for _, col := range summaryTable.Columns {
			err := resource.Set(col.Name, funk.Get(summary, defaultCaser.ToPascal(col.Name)))
			if err != nil {
				return fmt.Errorf("failed to set %s: %w", col.Name, err)
			}
		}

		vector := resource.GetValues()
		arrowRecord := vector.ToArrowRecord(resource.Table.ToArrowSchema())

		transformedRecord := destinationTransformers[i].Transform(arrowRecord)
		transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
		if err != nil {
			return fmt.Errorf("failed to transform sync summary bytes: %w", err)
		}
		wr = &plugin.Write_Request{}
		wr.Message = &plugin.Write_Request_Insert{
			Insert: &plugin.Write_MessageInsert{
				Record: transformedRecordBytes,
			},
		}
		if err := writeClients[i].Send(wr); err != nil {
			return handleSendError(err, writeClients[i], "insert sync summary")
		}
	}
	return nil
}
