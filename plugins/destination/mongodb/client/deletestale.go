package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Client) DeleteStale(ctx context.Context, tables schema.Tables, source string, syncTime time.Time) error {
	for _, table := range tables {
		tableName := table.Name
		// delete all records that are not in the source and are older than syncTime
		if _, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).DeleteMany(ctx, bson.M{
			schema.CqSourceNameColumn.Name: source,
			schema.CqSyncTimeColumn.Name:   bson.M{"$lt": syncTime},
		}); err != nil {
			return err
		}
	}
	return nil
}
