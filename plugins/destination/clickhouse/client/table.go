package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) getTableDefinitions(ctx context.Context, messages message.WriteMigrateTables) (schema.Tables, error) {
	// need proper description without flattened columns
	const flattenNested0 = "SET flatten_nested = 0"
	if err := c.conn.Exec(ctx, flattenNested0); err != nil {
		return nil, err
	}

	query, params := queries.GetTablesSchema(c.database)
	rows, err := c.conn.Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return queries.ScanTableSchemas(rows, messages)
}
