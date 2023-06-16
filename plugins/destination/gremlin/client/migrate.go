package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
)

// Migrate tables. Like neo4j, gremlin does not have a schema, so this is a no-op.
func (*Client) Migrate(_ context.Context, _ schema.Tables) error {
	return nil
}
