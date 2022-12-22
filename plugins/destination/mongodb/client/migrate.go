package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

// Migrate tables. It is the responsibility of the CLI of the client to lock before running migrations.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	return nil
}
