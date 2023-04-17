package client

import (
	"context"
	"reflect"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
)

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	query, params := queries.Read(sourceName, table)

	rows, err := c.conn.Query(ctx, query, params...)
	if err != nil {
		return err
	}
	defer rows.Close()

	columnTypes := rows.ColumnTypes()

	for rows.Next() {
		resource := rowArr(columnTypes)

		if err := rows.Scan(resource...); err != nil {
			return err
		}

		res <- resource
	}

	return nil
}

func rowArr(columnTypes []driver.ColumnType) []any {
	res := make([]any, len(columnTypes))
	for i := range columnTypes {
		res[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}
	return res
}
