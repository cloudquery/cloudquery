package typeconv

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	_arrow "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/arrow"
	_clickhouse "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch"
)

func ArrowRecord(sc *arrow.Schema, data []any) (arrow.Record, error) {
	return _arrow.Record(sc, data)
}

func ClickHouseBatchAddRecord(batch driver.Batch, record arrow.Record) error {
	return _clickhouse.BatchAddRecord(batch, record)
}
