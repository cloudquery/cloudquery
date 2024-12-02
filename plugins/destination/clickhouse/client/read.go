package client

import (
	"context"
	"reflect"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/queries"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/typeconv/arrow/values"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	rows, err := c.conn.Query(ctx, queries.Read(table))
	if err != nil {
		return err
	}
	defer rows.Close()

	row := rowArr(rows.ColumnTypes())
	builder := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	for rows.Next() {
		if err = rows.Scan(row...); err != nil {
			return err
		}

		if err = values.AppendToRecordBuilder(builder, row); err != nil {
			return err
		}
	}

	res <- builder.NewRecord()
	return nil
}

func rowArr(columnTypes []driver.ColumnType) []any {
	res := make([]any, len(columnTypes))
	for i := range columnTypes {
		res[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}
	return res
}
