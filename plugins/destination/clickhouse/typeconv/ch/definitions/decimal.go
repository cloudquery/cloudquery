package definitions

import (
	"fmt"
	"strconv"

	"github.com/apache/arrow/go/v12/arrow"
	"golang.org/x/exp/constraints"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
func decimalType(_type arrow.DecimalType) string {
	precision, scale := _type.GetPrecision(), _type.GetScale()
	if scale > 76 {
		// don't support this, default to String
		return "String"
	}

	precision = max(precision, scale)

	switch _type.ID() {
	case arrow.DECIMAL128:
		precision = ensureBetween(precision, 19, 38)
	case arrow.DECIMAL256:
		precision = ensureBetween(precision, 39, 76)
	default:
		// don't support this, default to String
		return "String"
	}

	return "Decimal(" + strconv.FormatInt(int64(precision), 10) + "," + strconv.FormatInt(int64(scale), 10) + ")"
}

func ensureBetween[A constraints.Ordered](x, from, to A) A {
	if from > to {
		panic(fmt.Sprintf("%v > %v", from, to))
	}

	if x < from {
		x = from
	}
	if x > to {
		x = to
	}

	return x
}

func max[A constraints.Ordered](x, y A) A {
	if x < y {
		return y
	}
	return x
}
