package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	documents := make([]any, len(resources))
	for i, resource := range resources {
		document := make(bson.M, len(table.Columns))
		for l, col := range table.Columns {
			document[col.Name] = resource[l]
		}
		documents[i] = document
	}
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).InsertMany(ctx, documents); err != nil {
		return err
	}
	return nil
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	operations := make([]mongo.WriteModel, len(resources))
	pks := table.PrimaryKeys()
	for i, resource := range resources {
		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		if len(pks) == 0 {
			// If no primary keys are defined, use all columns as a filter
			pks = table.Columns.Names()
		}
		filter := make(bson.M, len(pks))
		for _, pk := range pks {
			filter[pk] = resource[table.Columns.Index(pk)]
		}
		operation.SetFilter(filter)
		update := make(bson.M, len(table.Columns))
		for l, col := range table.Columns {
			update[col.Name] = resource[l]
		}
		operation.SetUpdate(bson.M{"$set": update})
		operations[i] = operation
	}
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return c.appendTableBatch(ctx, table, resources)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		return c.overwriteTableBatch(ctx, table, resources)
	default:
		panic("unsupported write mode " + c.spec.WriteMode.String())
	}
}
