package values

import (
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
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
