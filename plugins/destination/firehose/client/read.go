package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	panic("not implemented")
}
