package values

import (
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

func Record(sc *arrow.Schema, data []any) (arrow.Record, error) {
	if len(data) != len(sc.Fields()) {
		return nil, fmt.Errorf("mismatching field amount: have %d, want %d", len(data), len(sc.Fields()))
	}

	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, builder := range builder.Fields() {
		if err := buildValue(builder, data[i]); err != nil {
			return nil, err
		}
	}

	return builder.NewRecord(), nil
}

func unwrap(t schema.ValueType, v any) any {
	switch t {
	case schema.TypeByteArray:
		v := v.(**string)
		if *v == nil {
			return nil
		}
		return []byte(**v)

	case schema.TypeJSON:
		v := v.(**string)
		if *v == nil {
			return nil
		}
		return unescape(**v)

	case schema.TypeStringArray,
		schema.TypeMacAddrArray,
		schema.TypeCIDRArray,
		schema.TypeInetArray:
		return *v.(*[]*string)

	case schema.TypeIntArray:
		return *v.(*[]*int64)

	case schema.TypeUUIDArray:
		return *v.(*[]*uuid.UUID)

	default:
		return v
	}
}
