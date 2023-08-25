package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) getNormalizedTables(ctx context.Context, tables schema.Tables) (schema.Tables, error) {
	if c.pgTables != nil {
		return c.normalizeTables(tables, c.pgTables), nil
	}
	pgTables, err := c.refreshTables(ctx, tables.TableNames())
	if err != nil {
		return nil, fmt.Errorf("failed listing postgres tables: %w", err)

	}
	return c.normalizeTables(tables, pgTables), nil
}

func (c *Client) refreshTables(ctx context.Context, include []string) (schema.Tables, error) {
	var err error
	c.pgTables, err = c.listTables(ctx, include, nil)
	return c.pgTables, err
}
