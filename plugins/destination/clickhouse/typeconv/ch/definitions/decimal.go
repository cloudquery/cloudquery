package definitions

import (
	"strconv"
)

func decimalType(precision, scale int32, minPrecision, maxPrecision int32) string {
	// https://clickhouse.com/docs/en/sql-reference/data-types/decimal

	if precision < scale {
		precision = scale
	}

	switch {
	case precision < minPrecision:
		precision = minPrecision
	case precision > maxPrecision:
		precision = maxPrecision
	}

	if scale > precision {
		scale = precision
	}

	return "Decimal(" + strconv.FormatInt(int64(precision), 10) + "," + strconv.FormatInt(int64(scale), 10) + ")"
}
