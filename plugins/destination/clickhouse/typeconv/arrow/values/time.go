package values

import (
	"fmt"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func buildTimestampValues(builder *array.TimestampBuilder, value any) error {
	v, ok := unwrap[time.Time](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	t, err := timeToTimestamp(v, builder.Type().(*arrow.TimestampType))
	if err != nil {
		return err
	}

	if v.IsZero() {
		// work-around for empty values
		builder.AppendEmptyValue()
		return nil
	}

	builder.Append(t)
	return nil
}

func timeToTimestamp(value time.Time, tsType *arrow.TimestampType) (arrow.Timestamp, error) {
	loc, err := tsType.GetZone()
	if err != nil {
		return arrow.Timestamp(0), err
	}
	if loc != nil {
		value = value.In(loc)
	}

	return arrow.TimestampFromTime(value, tsType.Unit)
}

func buildTime32Values(builder primitiveBuilder[arrow.Time32], value any, dt *arrow.Time32Type) error {
	v, ok := unwrap[time.Time](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	if v.IsZero() {
		// work-around for empty values
		builder.AppendEmptyValue()
		return nil
	}

	t := v.Sub(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	switch dt.Unit {
	case arrow.Second:
		builder.Append(arrow.Time32(t.Seconds()))
	case arrow.Millisecond:
		builder.Append(arrow.Time32(t.Milliseconds()))
	default:
		return fmt.Errorf("unsupported unit %q for time32", dt.Unit.String())
	}
	return nil
}

func buildTime64Values(builder primitiveBuilder[arrow.Time64], value any, dt *arrow.Time64Type) error {
	v, ok := unwrap[time.Time](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	if v.IsZero() {
		// work-around for empty values
		builder.AppendEmptyValue()
		return nil
	}

	t := v.Sub(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	switch dt.Unit {
	case arrow.Microsecond:
		builder.Append(arrow.Time64(t.Microseconds()))
	case arrow.Nanosecond:
		builder.Append(arrow.Time64(t.Nanoseconds()))
	default:
		return fmt.Errorf("unsupported unit %q for time64", dt.Unit.String())
	}
	return nil
}
