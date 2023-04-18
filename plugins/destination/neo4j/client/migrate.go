package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// Migrate tables. Like mongo, neo4j does not have a schema, so this is a no-op.
func (*Client) Migrate(ctx context.Context, tables schema.Schemas) error {
	return nil
}
