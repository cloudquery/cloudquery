package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/cloudquery/cli/v6/internal/transformer"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/cloudquery/plugin-pb-go/pb/plugin/v3"
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/thoas/go-funk"
)

type syncSummary struct {
	CLIVersion          string    `json:"cli_version"`
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
	SourceTables        []string  `json:"source_tables"`
	SyncID              string    `json:"sync_id"`
	SyncTime            time.Time `json:"sync_time"`
	SyncGroupID         *string   `json:"sync_group_id,omitempty"`
	ShardNum            *int      `json:"shard_num,omitempty"`
	ShardTotal          *int      `json:"shard_total,omitempty"`
}

func persistSummary(filename string, summary syncSummary) error {
	// if filename is not specified then we don't need to persist the summary and can return
	if filename == "" {
		return nil
	}
	err := checkFilePath(filename)
	if err != nil {
		return fmt.Errorf("failed to validate summary file path: %w", err)
	}
	dataBytes, err := json.Marshal(summary)
	if err != nil {
		return err
	}
	dataBytes = append(dataBytes, []byte("\n")...)
	err = appendToFile(filename, dataBytes)
	if err != nil {
		return fmt.Errorf("failed to append summary to file: %w", err)
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

func generateSummaryTable() (*schema.Table, error) {
	tableName := "cloudquery_sync_summaries"
	t := schema.Tables{{
		Name: tableName,
		Transform: transformers.TransformWithStruct(
			&syncSummary{},
			// Already added as `_cq` internal columns
			transformers.WithSkipFields("SyncTime", "SyncGroupID"),
		),
	}}
	if err := transformers.TransformTables(t); err != nil {
		return nil, err
	}
	return t[0], nil
}

func migrateSummaryTable(writeClient safeWriteClient, destTransformer *transformer.RecordTransformer, spec specs.Destination) error {
	if !spec.SyncSummary {
		return nil
	}

	summaryTable, err := generateSummaryTable()
	if err != nil {
		return err
	}
	summaryTableSchema := summaryTable.ToArrowSchema()
	transformedSchema := destTransformer.TransformSchema(summaryTableSchema)
	transformedSchemaBytes, err := plugin.SchemaToBytes(transformedSchema)
	if err != nil {
		return err
	}
	wr := &plugin.Write_Request{}
	wr.Message = &plugin.Write_Request_MigrateTable{
		MigrateTable: &plugin.Write_MessageMigrateTable{
			MigrateForce: spec.MigrateMode == specs.MigrateModeForced,
			Table:        transformedSchemaBytes,
		},
	}
	if err := writeClient.Send(wr); err != nil {
		return handleSendError(err, writeClient, "migrate sync summary table")
	}
	return nil
}

func sendSummary(writeClient safeWriteClient, destinationSpec specs.Destination, destinationsClient *managedplugin.Client, destinationTransformer *transformer.RecordTransformer, summary *syncSummary, noMigrate bool) error {
	summaryTable, err := generateSummaryTable()
	if err != nil {
		return err
	}

	csr := caser.New(caser.WithCustomInitialisms(map[string]bool{"CLI": true}), caser.WithCustomExceptions(map[string]string{"cli": "CLI"}))

	// Respect the noMigrate flag
	if !noMigrate {
		if err := migrateSummaryTable(writeClient, destinationTransformer, destinationSpec); err != nil {
			return fmt.Errorf("failed to migrate sync summary table: %w", err)
		}
	}

	// Get Information about the DestinationPlugin
	m := destinationsClient.Metrics()
	summary.DestinationErrors = m.Errors
	summary.DestinationWarnings = m.Warnings

	summary.DestinationName = destinationSpec.Name
	summary.DestinationVersion = destinationSpec.Version
	summary.DestinationPath = destinationSpec.Path

	resource := schema.NewResourceData(summaryTable, nil, nil)
	for _, col := range summaryTable.Columns {
		err := resource.Set(col.Name, funk.Get(summary, csr.ToPascal(col.Name), funk.WithAllowZero()))
		if err != nil {
			return fmt.Errorf("failed to set %s: %w", col.Name, err)
		}
	}

	vector := resource.GetValues()
	arrowRecord := vector.ToArrowRecord(resource.Table.ToArrowSchema())

	transformedRecord := destinationTransformer.Transform(arrowRecord)
	transformedRecordBytes, err := plugin.RecordToBytes(transformedRecord)
	if err != nil {
		return fmt.Errorf("failed to transform sync summary bytes: %w", err)
	}

	wr := &plugin.Write_Request{}
	wr.Message = &plugin.Write_Request_Insert{
		Insert: &plugin.Write_MessageInsert{
			Record: transformedRecordBytes,
		},
	}
	if err := writeClient.Send(wr); err != nil {
		return handleSendError(err, writeClient, "insert sync summary")
	}

	return nil
}
