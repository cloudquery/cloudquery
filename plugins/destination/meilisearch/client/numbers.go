package client

import (
	"fmt"
	"math"
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

func toInt8(val any) (int8, error) {
	v, err := toInt64(val)
	if err != nil {
		return 0, err
	}
	if v < math.MinInt8 || v > math.MaxInt8 {
		return 0, fmt.Errorf("value %d out of range for int8", v)
	}
	return int8(v), nil
}

func toInt16(val any) (int16, error) {
	v, err := toInt64(val)
	if err != nil {
		return 0, err
	}
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, fmt.Errorf("value %d out of range for int16", v)
	}
	return int16(v), nil
}

func toInt32(val any) (int32, error) {
	v, err := toInt64(val)
	if err != nil {
		return 0, err
	}
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("value %d out of range for int32", v)
	}
	return int32(v), nil
}

func toUint8(val any) (uint8, error) {
	v, err := toUint64(val)
	if err != nil {
		return 0, err
	}
	if v > math.MaxUint8 {
		return 0, fmt.Errorf("value %d out of range for uint8", v)
	}
	return uint8(v), nil
}

func toUint16(val any) (uint16, error) {
	v, err := toUint64(val)
	if err != nil {
		return 0, err
	}
	if v > math.MaxUint16 {
		return 0, fmt.Errorf("value %d out of range for uint16", v)
	}
	return uint16(v), nil
}

func toUint32(val any) (uint32, error) {
	v, err := toUint64(val)
	if err != nil {
		return 0, err
	}
	if v > math.MaxUint32 {
		return 0, fmt.Errorf("value %d out of range for uint32", v)
	}
	return uint32(v), nil
}
