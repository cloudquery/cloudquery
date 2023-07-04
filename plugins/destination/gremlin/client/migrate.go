package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

// Migrate tables. Like neo4j, gremlin does not have a schema, so this is a no-op.
func (*Client) MigrateTables(context.Context, message.WriteMigrateTables) error {
	return nil
}
