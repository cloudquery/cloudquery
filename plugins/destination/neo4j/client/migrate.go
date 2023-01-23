package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

// Migrate tables. Like mongo, neo4j does not have a schema, so this is a no-op.
func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}
