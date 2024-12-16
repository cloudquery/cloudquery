package types

import (
	"fmt"
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
func decimalType(decimalType arrow.DecimalType) (string, error) {
	precision, scale := decimalType.GetPrecision(), decimalType.GetScale()
	// sanity check
	switch {
	case precision < 1:
		return "", fmt.Errorf("precision (%d) must be in range [1, 76]", precision)
	case scale < 0, scale > precision:
		return "", fmt.Errorf("scale (%d) must be in range [0, %d]", scale, precision)
	}

	return "Decimal(" + strconv.FormatInt(int64(precision), 10) + "," + strconv.FormatInt(int64(scale), 10) + ")", nil
}
