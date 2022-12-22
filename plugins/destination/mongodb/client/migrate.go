package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

// MongoDB. No migrations needed :)
func (*Client) Migrate(context.Context, schema.Tables) error {
	return nil
}
