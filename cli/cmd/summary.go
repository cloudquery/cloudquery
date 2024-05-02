package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

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
	CliVersion          string    `json:"cli_version"`
	DestinationErrors   uint64    `json:"destination_errors"`
	DestinationName     string    `json:"destination_name"`
	DestinationPath     string    `json:"destination_path"`
	DestinationVersion  string    `json:"destination_version"`
	DestinationWarnings uint64    `json:"destination_warnings"`
	Resources           uint64    `json:"resources"`
	SourceErrors        uint64    `json:"source_errors"`
	SourceName          string    `json:"source_name"`
	SourcePath          string    `json:"source_path"`
	SourceVersion       string    `json:"source_version"`
	SourceWarnings      uint64    `json:"source_warnings"`
	SyncID              string    `json:"sync_id"`
	SyncStartTime       time.Time `json:"sync_start_time"`
}

func persistSummary(filename string, summaries []syncSummary) error {
	// if filename is not specified then we don't need to persist the summary and can return
	if filename == "" {
		return nil
	}
	err := checkFilePath(filename)
	if err != nil {
		return fmt.Errorf("failed to validate summary file path: %w", err)
	}
	for _, summary := range summaries {
		dataBytes, err := json.Marshal(summary)
		if err != nil {
			return err
		}
		dataBytes = append(dataBytes, []byte("\n")...)
		err = appendToFile(filename, dataBytes)
		if err != nil {
			return fmt.Errorf("failed to append summary to file: %w", err)
		}
	}
	return nil
}

func appendToFile(fileName string, data []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func checkFilePath(filename string) error {
	dirPath := filepath.Dir(filename)
	return os.MkdirAll(dirPath, 0755)
}

func generateSummaryTable() *schema.Table {
	tableName := "cloudquery_sync_summary"
	t := schema.Tables{{
		Name: tableName,
		Transform: transformers.TransformWithStruct(
			&syncSummary{},
		),
		Columns: []schema.Column{},
	}}
	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}
	//for i := range t[0].Columns {
	//	t[0].Columns[i].NotNull = true
	//}
	return t[0]
}

func sendSummary(writeClients []plugin.Plugin_WriteClient, destinationSpecs []specs.Destination, destinationsClients []*managedplugin.Client, destinationTransformers []*transformer.RecordTransformer, summary *syncSummary, noMigrate bool) error {
	summaryTable := generateSummaryTable()
	summaryTableSchema := summaryTable.ToArrowSchema()

	defaultCaser := caser.New()
	for i := range destinationsClients {
		if !destinationSpecs[i].SyncSummary {
			continue
		}

		transformedSchema := destinationTransformers[i].TransformSchema(summaryTableSchema)
		transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
		if err != nil {
			return err
		}
		wr := &plugin.Write_Request{}

		// Respect the noMigrate flag
		if !noMigrate {
			wr.Message = &plugin.Write_Request_MigrateTable{
				MigrateTable: &plugin.Write_MessageMigrateTable{
					MigrateForce: destinationSpecs[i].MigrateMode == specs.MigrateModeForced,
					Table:        transformedSchemaBytes,
				},
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
		summary.CliVersion = Version

		resource := schema.NewResourceData(summaryTable, nil, nil)
		for _, col := range summaryTable.Columns {
			err = resource.Set(col.Name, funk.Get(summary, defaultCaser.ToPascal(col.Name), funk.WithAllowZero()))
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
