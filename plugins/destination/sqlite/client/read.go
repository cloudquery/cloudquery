package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) createResultsArray(table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		switch col.Type {
		case schema.TypeBool:
			var r bool
			results = append(results, &r)
		case schema.TypeInt:
			var r int
			results = append(results, &r)
		case schema.TypeFloat:
			var r float64
			results = append(results, &r)
		case schema.TypeUUID:
			var r string
			results = append(results, &r)
		case schema.TypeString:
			var r string
			results = append(results, &r)
		case schema.TypeByteArray:
			var r sql.RawBytes
			results = append(results, &r)
		case schema.TypeStringArray:
			var r string
			results = append(results, &r)
		case schema.TypeTimestamp:
			var r string
			results = append(results, &r)
		case schema.TypeJSON:
			var r string
			results = append(results, &r)
		case schema.TypeUUIDArray:
			var r string
			results = append(results, &r)
		case schema.TypeCIDR:
			var r string
			results = append(results, &r)
		case schema.TypeCIDRArray:
			var r string
			results = append(results, &r)
		case schema.TypeMacAddr:
			var r string
			results = append(results, &r)
		case schema.TypeMacAddrArray:
			var r string
			results = append(results, &r)
		case schema.TypeInet:
			var r string
			results = append(results, &r)
		case schema.TypeInetArray:
			var r string
			results = append(results, &r)
		case schema.TypeIntArray:
			var r string
			results = append(results, &r)
		}
	}
	return results
}

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
			readSQL = append(readSQL, col)
		}
	}

	// add table name
	readSQL = append(readSQL, fmt.Sprintf("FROM %s", sanitizeName(table.Name)))

	// add source name
	readSQL = append(readSQL, fmt.Sprintf("WHERE %s = $1", sanitizeName(schema.CqSourceNameColumn.Name)))

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
			readSQL = append(readSQL, fmt.Sprintf("%s %s", sanitizeName(col.Name), order))
		}
	}

	// add limit
	if opts.Limit > 0 {
		readSQL = append(readSQL, fmt.Sprintf("LIMIT %d", opts.Limit))
	}

	rows, err := c.db.Query(strings.Join(readSQL, " "), sourceName)
	if err != nil {
		return err
	}
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		res <- values
	}
	rows.Close()
	return nil
}

func sanitizeName(s string) string {
	s = strings.ReplaceAll(s, string([]byte{0}), "")
	s = `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
	return s
}
