package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		// delete all records that are not in the source and are older than syncTime
		if _, err := c.client.Database(c.pluginSpec.Database).Collection(table.Name).DeleteMany(ctx, bson.M{"_cq_source_name": source, "_cq_sync_time": bson.M{"$lt": syncTime}}); err != nil {
			return err
		}
	}
	return nil
}
