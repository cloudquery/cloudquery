package values

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
)

func BatchAddRecord(batch driver.Batch, record arrow.Record) error {
	for i, col := range record.Columns() {
		data, err := FromArray(col)
		if err != nil {
			return err
		}
		if err := batch.Column(i).Append(data); err != nil {
			return err
		}
	}
	return nil
}
