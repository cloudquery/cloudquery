package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.readCSV(ctx, table, sourceName, res)
	case FormatTypeJSON:
		return c.readJSON(ctx, table, sourceName, res)
	default:
		panic("unknown backend type " + c.csvSpec.Backend.String())
	}
}
