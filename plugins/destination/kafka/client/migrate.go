package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// For Kafka migrate is a no-op
func (*Client) Migrate(context.Context, schema.Tables) error {
	return nil
}
