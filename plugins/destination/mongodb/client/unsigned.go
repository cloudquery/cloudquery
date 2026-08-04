package client

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// reinterpretUnsigned restores uint64 values that were stored as int64 bits by the
// CustomUnit64 encoder. Values above math.MaxInt64 come back from MongoDB as negative
// int64, which a uint64 builder rejects once the value is decoded exactly.
func reinterpretUnsigned(dt arrow.DataType, val any) any {
	if val == nil {
		return nil
	}

	switch t := dt.(type) {
	case *arrow.Uint64Type:
		if i, ok := val.(int64); ok {
			return uint64(i)
		}
		return val
	case *arrow.StructType:
		m, ok := val.(map[string]any)
		if !ok {
			return val
		}
		for _, f := range t.Fields() {
			if v, ok := m[f.Name]; ok {
				m[f.Name] = reinterpretUnsigned(f.Type, v)
			}
		}
		return m
	case arrow.ListLikeType:
		s, ok := val.([]any)
		if !ok {
			return val
		}
		for i := range s {
			s[i] = reinterpretUnsigned(t.Elem(), s[i])
		}
		return s
	default:
		return val
	}
}
