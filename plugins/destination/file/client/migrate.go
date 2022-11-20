package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func isSameColumns(record []string, columns []string) bool {
	if len(record) != len(columns) {
		return false
	}
	for i, v := range record {
		if v != columns[i] {
			return false
		}
	}
	return true
}

func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	// migrate is not needed in migrate mode
	return nil
}
