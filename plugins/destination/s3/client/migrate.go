package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/message"
)

func (*Client) MigrateTables(context.Context, []*message.MigrateTable) error {
	// migrate is not needed in append mode
	return nil
}
