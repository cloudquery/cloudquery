package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (*Client) Migrate(ctx context.Context, schemas schema.Schemas) error {
	// migrate is not needed in append mode
	return nil
}
