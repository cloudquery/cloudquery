package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) Migrate(ctx context.Context, tables schema.Tables) error {
	// migrate is not needed in migrate mode
	return nil
}
