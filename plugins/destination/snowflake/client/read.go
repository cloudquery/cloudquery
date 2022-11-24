package client

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	readSQL = "SELECT * FROM %s WHERE \"_cq_source_name\" = ?"
)

func (*Client) createResultsArray(table *schema.Table) []interface{} {
	results := make([]interface{}, 0, len(table.Columns))
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

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	stmt := fmt.Sprintf(readSQL, table.Name)
	rows, err := c.db.Query(stmt, sourceName)
	if err != nil {
		return err
	}
	defer rows.Close()
	values := c.createResultsArray(table)
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("failed to read from table %s: %w", table.Name, err)
		}
		res <- values
	}
	rows.Close()
	return nil
}
