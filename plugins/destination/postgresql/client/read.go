package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any, opts destination.ReadOptions) error {
	readSQL := []string{"SELECT"}

	// add columns
	if len(opts.Columns) == 0 {
		readSQL = append(readSQL, "*")
	} else {
		for i, col := range opts.Columns {
			if i > 0 {
				readSQL = append(readSQL, ",")
			}
			readSQL = append(readSQL, pgx.Identifier{col}.Sanitize())
		}
	}

	// add table name
	readSQL = append(readSQL, fmt.Sprintf("FROM %s", pgx.Identifier{table.Name}.Sanitize()))

	// add source name
	readSQL = append(readSQL, fmt.Sprintf("WHERE %s = $1", pgx.Identifier{schema.CqSourceNameColumn.Name}.Sanitize()))

	// add order by
	if len(opts.OrderBy) == 0 {
		readSQL = append(readSQL, "ORDER BY _cq_sync_time ASC")
	} else {
		readSQL = append(readSQL, "ORDER BY")
		for i, col := range opts.OrderBy {
			if i > 0 {
				readSQL = append(readSQL, ",")
			}
			order := "ASC"
			if col.Desc {
				order = "DESC"
			}
			readSQL = append(readSQL, fmt.Sprintf("%s %s", pgx.Identifier{col.Name}.Sanitize(), order))
		}
	}

	// add limit
	if opts.Limit > 0 {
		readSQL = append(readSQL, fmt.Sprintf("LIMIT %d", opts.Limit))
	}

	sql := strings.Join(readSQL, " ")
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
