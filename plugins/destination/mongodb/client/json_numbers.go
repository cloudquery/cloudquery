package client

import (
	"bytes"
	"strconv"

	"github.com/goccy/go-json"
)

// unmarshalNestedJSON decodes a nested value with UseNumber so int64/uint64 values that
// exceed float64 precision survive, then converts the json.Number values to the integer
// types the BSON encoder understands.
func unmarshalNestedJSON(data []byte) (any, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()

	var val any
	if err := dec.Decode(&val); err != nil {
		return nil, err
	}

	return convertJSONNumbers(val), nil
}

func convertJSONNumbers(val any) any {
	switch v := val.(type) {
	case json.Number:
		if i, err := strconv.ParseInt(v.String(), 10, 64); err == nil {
			return i
		}
		if u, err := strconv.ParseUint(v.String(), 10, 64); err == nil {
			return CustomUnit64(u)
		}
		f, err := v.Float64()
		if err != nil {
			return v.String()
		}
		return f
	case map[string]any:
		for k := range v {
			v[k] = convertJSONNumbers(v[k])
		}
		return v
	case []any:
		for i := range v {
			v[i] = convertJSONNumbers(v[i])
		}
		return v
	default:
		return val
	}
}
