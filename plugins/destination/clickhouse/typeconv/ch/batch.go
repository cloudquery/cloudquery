package ch

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
)

func BatchAddRecord(batch driver.Batch, record arrow.Record) error {
	for i, col := range record.Columns() {
		if err := addColumn(batch.Column(i), col); err != nil {
			return err
		}
	}
	return nil
}

func addColumn(column driver.BatchColumn, arr arrow.Array) error {
	for i := 0; i < arr.Len(); i++ {
		switch {
		case arr.IsNull(i), !arr.IsValid(i):
			if err := column.Append(nil); err != nil {
				return err
			}
		default:
			panic("impl")
		}
	}
	return nil
}
