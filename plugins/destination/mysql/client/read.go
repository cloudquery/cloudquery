package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) createResultsArray(table *schema.Table) []any {
	results := make([]any, 0, len(table.Columns))
	for _, col := range table.Columns {
		switch col.Type {
		case schema.TypeBool:
			var r *bool
			results = append(results, &r)
		case schema.TypeInt:
			var r *int
			results = append(results, &r)
		case schema.TypeFloat:
			var r *float64
			results = append(results, &r)
		case schema.TypeUUID:
			var r *[]byte
			results = append(results, &r)
		case schema.TypeString:
			var r *string
			results = append(results, &r)
		case schema.TypeByteArray:
			var r *[]byte
			results = append(results, &r)
		case schema.TypeStringArray:
			var r string
			results = append(results, &r)
		case schema.TypeTimestamp:
			var r *time.Time
			results = append(results, &r)
		case schema.TypeJSON:
			var r string
			results = append(results, &r)
		case schema.TypeUUIDArray:
			var r string
			results = append(results, &r)
		case schema.TypeCIDR:
			var r *string
			results = append(results, &r)
		case schema.TypeCIDRArray:
			var r string
			results = append(results, &r)
		case schema.TypeMacAddr:
			var r *string
			results = append(results, &r)
		case schema.TypeMacAddrArray:
			var r string
			results = append(results, &r)
		case schema.TypeInet:
			var r *string
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

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	builder := strings.Builder{}
	builder.WriteString("SELECT")
	for i, col := range table.Columns {
		builder.WriteString(" " + identifier(col.Name))
		if i != len(table.Columns)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("FROM " + identifier(table.Name) + " WHERE _cq_source_name = ? ORDER BY _cq_sync_time ASC")
	rows, err := c.db.QueryContext(ctx, builder.String(), sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		res <- values
	}
	return nil
}
