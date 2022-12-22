package client

import (
	"context"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	batchSize    = 1000
	writeTimeout = 5 * time.Minute
)

type item struct {
	cols map[string]bigquery.Value
}

func (i *item) Save() (map[string]bigquery.Value, string, error) {
	// we're not doing de-dup at the moment
	return i.cols, bigquery.NoDedupeID, nil
}



func (c *Client) WriteTableBatch(ctx context.Context, table *schema.Table, resources [][]any) error {
	operations := make([]mongo.WriteModel, len(resources))
	for i, resource := range resources {
		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		filter := make(bson.M, len(table.PrimaryKeys()))
		for _, pk := range table.PrimaryKeys() {
			filter[pk] = resource[table.Columns.Index(pk)]
		}
		operation.SetFilter(filter)
		update := make(bson.M, len(table.Columns))
		for i, col := range table.Columns {
			update[col.Name] = resource[i]
		}
		// operation.SetUpdate(update)
		operation.SetUpdate(bson.M{"$set": update})
		operations[i] = operation
	}
	if len(operations) == 0 {
		panic("no operations")
	}
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}
