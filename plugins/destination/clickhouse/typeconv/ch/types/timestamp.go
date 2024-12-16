package types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
func timestampType(tsType *arrow.TimestampType) (string, error) {
	loc, err := tsType.GetZone()
	if err != nil {
		return "", err
	}

	if loc.String() == time.UTC.String() {
		// default is UTC, so we force it to nil here
		loc = nil
	}

	return timeType(tsType.Unit, loc)
}

func timeType(unit arrow.TimeUnit, loc *time.Location) (string, error) {
	var precision int
	switch unit {
	case arrow.Second:
		precision = 0
	case arrow.Millisecond:
		precision = 3
	case arrow.Microsecond:
		precision = 6
	case arrow.Nanosecond:
		precision = 9
	default:
		return "", fmt.Errorf("unsupported Apache Arrow time unit: %s", unit.String())
	}

	if loc != nil {
		return "DateTime64(" + strconv.Itoa(precision) + ", '" + loc.String() + "')", nil
	}

	return "DateTime64(" + strconv.Itoa(precision) + ")", nil
}
