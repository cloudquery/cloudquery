package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type BoolColumnsBuilder struct {
	i          int
	values     map[string][]*bool
	typeSchema map[string]string
}

func NewBoolColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &BoolColumnsBuilder{i: -1, values: make(map[string][]*bool), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != BoolType {
			continue
		}
		b.values[key] = make([]*bool, originalColumn.Len())
	}
	return b
}

func (b *BoolColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != BoolType {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(bool); ok {
			b.values[key][b.i] = &v
		}
	}
}

func (b *BoolColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildBoolColumn(b.values[key]), nil
}

func buildBoolColumn(values []*bool) arrow.Array {
	bld := array.NewBooleanBuilder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewBooleanArray()
}
