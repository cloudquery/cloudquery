package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteStale is not currently implemented for BigQuery, as it only supports "append" write mode.
func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		// c.client.Database(c.pluginSpec.Database).Collection(table.Name).DeleteMany(ctx)
		// delete all records that are not in the source and are older than syncTime
		if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).DeleteMany(ctx, bson.M{"_cq_source_name": source, "_cq_sync_time": bson.M{"$lt": syncTime}}); err != nil {
			return err
		}
	}
	return nil
}
