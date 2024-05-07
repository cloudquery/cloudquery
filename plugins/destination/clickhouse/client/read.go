package client

import (
	"context"
	"reflect"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/arrow/values"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	sc := table.ToArrowSchema()
	query := queries.Read(table)

	rows, err := c.conn.Query(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	columnTypes := rows.ColumnTypes()

	for rows.Next() {
		row := rowArr(columnTypes)

		if err := rows.Scan(row...); err != nil {
			return err
		}

		record, err := values.Record(sc, row)
		if err != nil {
			return err
		}

		res <- record
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
