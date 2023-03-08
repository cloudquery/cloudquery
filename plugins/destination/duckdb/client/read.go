package client

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/marcboeker/go-duckdb"
)

const (
	readSQL = `SELECT %s FROM "%s" WHERE _cq_source_name = $1 order by _cq_sync_time asc`
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
			var r []byte
			results = append(results, &r)
		case schema.TypeString:
			var r *string
			results = append(results, &r)
		case schema.TypeByteArray:
			var r sql.RawBytes
			results = append(results, &r)
		case schema.TypeStringArray:
			var r duckdb.Composite[[]string]
			results = append(results, &r)
		case schema.TypeTimestamp:
			var r *string
			results = append(results, &r)
		case schema.TypeJSON:
			var r string
			results = append(results, &r)
		case schema.TypeUUIDArray:
			var r duckdb.Composite[[][]byte]
			results = append(results, &r)
		case schema.TypeCIDR:
			var r *string
			results = append(results, &r)
		case schema.TypeCIDRArray:
			var r duckdb.Composite[[]string]
			results = append(results, &r)
		case schema.TypeMacAddr:
			var r *string
			results = append(results, &r)
		case schema.TypeMacAddrArray:
			var r duckdb.Composite[[]string]
			results = append(results, &r)
		case schema.TypeInet:
			var r *string
			results = append(results, &r)
		case schema.TypeInetArray:
			var r duckdb.Composite[[]string]
			results = append(results, &r)
		case schema.TypeIntArray:
			var r duckdb.Composite[[]int]
			results = append(results, &r)
		}
	}
	return results
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	colNames := make([]string, 0, len(table.Columns))
	for _, col := range table.Columns {
		colNames = append(colNames, `"`+col.Name+`"`)
	}
	cols := strings.Join(colNames, ", ")
	rows, err := c.db.Query(fmt.Sprintf(readSQL, cols, table.Name), sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := c.createResultsArray(table)
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		for i := range values {
			switch v := values[i].(type) {
			case *duckdb.Composite[[]string]:
				values[i] = v.Get()
			case *duckdb.Composite[[][]byte]:
				values[i] = v.Get()
			case *duckdb.Composite[[]int]:
				values[i] = v.Get()
			}
		}
		res <- values
	}
	return nil
}
