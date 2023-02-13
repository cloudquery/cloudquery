package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) getTableDefinitions(ctx context.Context) (schema.Tables, error) {
	query, params := queries.GetTablesSchema(c.database)
	rows, err := c.conn.Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return queries.ScanTableSchemas(rows)
}
