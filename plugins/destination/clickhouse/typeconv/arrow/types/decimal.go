package types

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

func decimalType(name string, col *column.Decimal) *arrow.Field {
	precision, scale := int32(col.Precision()), int32(col.Scale())
	if precision <= 38 {
		return &arrow.Field{Name: name, Type: &arrow.Decimal128Type{Precision: precision, Scale: scale}}
	}
	return &arrow.Field{Name: name, Type: &arrow.Decimal256Type{Precision: precision, Scale: scale}}
}
