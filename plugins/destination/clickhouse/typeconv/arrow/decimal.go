package arrow

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
)

func decimalType(name string, col *column.Decimal) (*arrow.Field, error) {
	var decimal arrow.DecimalType
	if precision := col.Precision(); precision <= 38 {
		decimal = &arrow.Decimal128Type{Precision: int32(precision), Scale: int32(col.Scale())}
	} else {
		decimal = &arrow.Decimal256Type{Precision: int32(precision), Scale: int32(col.Scale())}
	}
	return &arrow.Field{Name: name, Type: decimal}, nil
}
