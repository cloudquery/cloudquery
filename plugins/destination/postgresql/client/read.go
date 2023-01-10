package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5"
)

const (
	readSQL = "SELECT %s FROM %s WHERE _cq_source_name = $1 order by _cq_sync_time asc"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	colNames := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		colNames = append(colNames, pgx.Identifier{col.Name}.Sanitize())
	}
	cols := strings.Join(colNames, ",")
	sql := fmt.Sprintf(readSQL, cols, pgx.Identifier{table.Name}.Sanitize())
	rows, err := c.conn.Query(ctx, sql, sourceName)
	if err != nil {
		return err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err
		}
		res <- values
	}
	rows.Close()
	return nil
}
