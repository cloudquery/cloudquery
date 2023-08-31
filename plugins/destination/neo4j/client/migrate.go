package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

// Migrate tables. Like mongo, neo4j does not have a schema, so this is a no-op.
func (*Client) MigrateTables(context.Context, message.WriteMigrateTables) error {
	return nil
}
