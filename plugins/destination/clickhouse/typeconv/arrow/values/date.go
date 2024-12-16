package values

import (
	"time"

	"github.com/apache/arrow-go/v18/arrow"
)

func buildDate32Values(builder primitiveBuilder[arrow.Date32], value any) {
	v, ok := unwrap[time.Time](value)
	if !ok {
		builder.AppendNull()
		return
	}

	if v.IsZero() {
		// work-around for empty values
		builder.AppendEmptyValue()
		return
	}

	builder.Append(arrow.Date32FromTime(v))
}

func buildDate64Values(builder primitiveBuilder[arrow.Date64], value any) {
	v, ok := unwrap[time.Time](value)
	if !ok {
		builder.AppendNull()
		return
	}

	if v.IsZero() {
		// work-around for empty values
		builder.AppendEmptyValue()
		return
	}

	builder.Append(arrow.Date64FromTime(v))
}
