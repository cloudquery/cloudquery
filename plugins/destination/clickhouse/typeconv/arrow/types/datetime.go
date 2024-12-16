package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/timezone"
	"github.com/apache/arrow-go/v18/arrow"
)

func dateTimeType(name string, col *column.DateTime) (*arrow.Field, error) {
	timeZone, err := getTimeZone(params(col.Type()))
	if err != nil {
		return nil, err
	}

	return &arrow.Field{Name: name, Type: &arrow.TimestampType{Unit: arrow.Second, TimeZone: timeZone}}, nil
}

func dateTime64Type(name string, col *column.DateTime64) (*arrow.Field, error) {
	// need to parse
	params := strings.SplitN(params(col.Type()), ",", 2)
	precision, err := strconv.Atoi(params[0])
	if err != nil {
		return nil, err
	}

	var unit arrow.TimeUnit
	switch {
	case precision == 0:
		unit = arrow.Second
	case precision <= 3:
		// This is the same as arrow.DATE64, so we need to canonize the schema
		unit = arrow.Millisecond
	case precision <= 6:
		unit = arrow.Microsecond
	case precision <= 9:
		unit = arrow.Nanosecond
	default:
		return nil, fmt.Errorf("unsupported DateTime64 precision: %d (supported values: 0-9)", precision)
	}

	var timeZone string
	if len(params) > 1 {
		timeZone, err = getTimeZone(params[1])
		if err != nil {
			return nil, err
		}
	}

	return &arrow.Field{Name: name, Type: &arrow.TimestampType{Unit: unit, TimeZone: timeZone}}, nil
}

func getTimeZone(param string) (string, error) {
	tzName := strings.Trim(strings.TrimSpace(param), "'")
	if len(tzName) == 0 {
		return "", nil
	}

	loc, err := timezone.Load(tzName)
	if err != nil {
		return "", err
	}

	return loc.String(), nil
}
