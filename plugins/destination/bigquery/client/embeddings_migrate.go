package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	pluginSchema "github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *ConcreteEmbeddingsClient) MigrateTables(ctx context.Context) error {
	for _, tableConfig := range c.spec.Tables {
		schema, err := c.buildSchemaForEmbeddingsTable(ctx, tableConfig)
		if err != nil {
			return err
		}
		tm := bigquery.TableMetadata{
			Name:             tableConfig.TargetTableName,
			Location:         "",
			Description:      fmt.Sprintf("Embeddings for table %s", tableConfig.SourceTableName),
			Schema:           schema,
			TimePartitioning: c.client.timePartitioning(),
		}

		if ok, err := c.client.doesTableExist(ctx, c.client.client, tableConfig.TargetTableName); err != nil {
			return err
		} else if ok {
			continue
		}

		if err := c.client.client.DatasetInProject(c.ProjectID, c.DatasetID).Table(tableConfig.TargetTableName).Create(ctx, &tm); err != nil {
			return err
		}
		if err := c.client.waitForTableToExist(ctx, c.client.client, &pluginSchema.Table{Name: tableConfig.TargetTableName}); err != nil {
			return err
		}
	}
	return nil
}

func (c *ConcreteEmbeddingsClient) buildSchemaForEmbeddingsTable(ctx context.Context, tableConfig TableConfig) ([]*bigquery.FieldSchema, error) {
	md, err := c.client.client.DatasetInProject(c.ProjectID, c.DatasetID).Table(tableConfig.SourceTableName).Metadata(ctx)
	if err != nil {
		return nil, err
	}

	metadataColumnSet := sliceToSet(tableConfig.MetadataColumns)
	columns := baseColumns()
	for _, column := range md.Schema {
		if _, ok := metadataColumnSet[column.Name]; ok {
			columns = append(columns, column)
		}
	}

	return columns, nil
}

func baseColumns() []*bigquery.FieldSchema {
	return []*bigquery.FieldSchema{
		{Name: "chunk_id", Type: bigquery.IntegerFieldType, Required: true},
		{Name: "chunk_text", Type: bigquery.StringFieldType, Required: true},
		{Name: "embedding", Type: bigquery.FloatFieldType, Repeated: true, Required: true},
	}
}
