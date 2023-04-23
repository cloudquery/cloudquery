package ch

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/value"
)

func BatchAddRecord(batch driver.Batch, record arrow.Record) error {
	for i, col := range record.Columns() {
		data, err := value.FromArray(col)
		if err != nil {
			return err
		}
		if err := batch.Column(i).Append(data); err != nil {
			return err
		}
	}
	return nil
}

func mapColumnData[K comparable, V any](keys []*K, items []*V) any {
	if len(keys) != len(items) {
		panic("keys & items have different len")
	}

	if len(keys) == 0 {
		return nil
	}

	return nil
	//mapType := reflect.MapOf(
	//	reflect.TypeOf(*keys[0]),
	//	reflect.TypeOf(*items[0]),
	//)
}
