package definitions

import (
	"strconv"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
func timestampType(_type *arrow.TimestampType) string {
	loc, err := _type.GetZone()
	if err != nil {
		// This will also result in the _type.GetToTimeFunc error.
		// Just default to String
		// TODO: consider panic/error
		return "String"
	}
	if loc.String() == time.UTC.String() {
		loc = nil
	}

	var precision int
	switch _type.TimeUnit() {
	case arrow.Second:
		precision = 0
	case arrow.Millisecond:
		precision = 3
	case arrow.Microsecond:
		precision = 6
	case arrow.Nanosecond:
		precision = 9
	default:
		// this will also result in the _type.GetToTimeFunc error.
		// Just default to String
		// TODO: consider panic/error
		return "String"
	}

	if loc != nil {
		return "DateTime64(" + strconv.Itoa(precision) + ", '" + loc.String() + "')"
	}

	return "DateTime64(" + strconv.Itoa(precision) + ")"
}
