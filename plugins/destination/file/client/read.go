package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	if !c.testMode && c.csvSpec.Backend != BackendTypeLocal {
		return fmt.Errorf("read is only supported in test mode with non-local backend")
	}
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.readCSV(ctx, table, sourceName, res)
	case FormatTypeJSON:
		return c.readJSON(ctx, table, sourceName, res)
	default:
		panic("unknown backend type " + c.csvSpec.Backend)
	}
}
