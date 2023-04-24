package values

import (
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildDate32Values(builder primitiveBuilder[arrow.Date32], value *time.Time) {
	switch {
	case value == nil, value == (*time.Time)(nil):
		builder.AppendNull()
	default:
		builder.Append(arrow.Date32FromTime(*value))
	}
}
func buildDate64Values(builder primitiveBuilder[arrow.Date64], value *time.Time) {
	switch {
	case value == nil, value == (*time.Time)(nil):
		builder.AppendNull()
	default:
		builder.Append(arrow.Date64FromTime(*value))
	}
}

func buildTimestampValues(builder array.TimestampBuilder, value *time.Time) {
	switch {
	case value == nil, value == (*time.Time)(nil):
		builder.AppendNull()
	default:
		_type := builder.Type().(*arrow.TimestampType)
		switch expr {

		}
		arrow.TimestampFromStringInLocation()
		builder.Append(arrow.Date64FromTime(*value))
	}
}

func timeToTimestamp(unit arrow.TimeUnit)
