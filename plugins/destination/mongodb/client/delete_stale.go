package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	for _, msg := range msgs {
		tableName := msg.TableName
		// delete all records that are not in the source and are older than syncTime
		err := retryWrite(ctx, c.logger, c.spec.WriteRetry, tableName, func() error {
			return c.runWrite(ctx, func(ctx context.Context) error {
				_, err := c.client.Database(c.spec.Database).Collection(tableName).DeleteMany(ctx, bson.M{
					schema.CqSourceNameColumn.Name: msg.SourceName,
					schema.CqSyncTimeColumn.Name:   bson.M{"$lt": msg.SyncTime},
				})
				return err
			})
		})
		if err != nil {
			return err
		}

	}
	return nil
}
