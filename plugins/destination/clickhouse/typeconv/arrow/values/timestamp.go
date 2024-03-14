package values

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
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
