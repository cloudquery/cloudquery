package client

import (
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

// Documents are decoded with UseNumber so int64/uint64 values that exceed float64
// precision survive as json.Number, which these helpers parse without rounding.

func toInt64(val any) (int64, error) {
	switch v := val.(type) {
	case json.Number:
		return strconv.ParseInt(v.String(), 10, 64)
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("unsupported numeric type %T", val)
	}
}

func toUint64(val any) (uint64, error) {
	switch v := val.(type) {
	case json.Number:
		return strconv.ParseUint(v.String(), 10, 64)
	case float64:
		return uint64(v), nil
	default:
		return 0, fmt.Errorf("unsupported numeric type %T", val)
	}
}

func toFloat64(val any) (float64, error) {
	switch v := val.(type) {
	case json.Number:
		return v.Float64()
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("unsupported numeric type %T", val)
	}
}
