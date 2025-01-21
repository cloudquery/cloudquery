package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type Int32ColumnsBuilder struct {
	i          int
	values     map[string][]*int32
	typeSchema map[string]string
}

func NewInt32ColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &Int32ColumnsBuilder{i: -1, values: make(map[string][]*int32), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != Int32Type {
			continue
		}
		b.values[key] = make([]*int32, originalColumn.Len())
	}
	return b
}

func (b *Int32ColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != Int32Type {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(float64); ok {
			int32Value := int32(v)
			b.values[key][b.i] = &int32Value
		}
	}
}

func (b *Int32ColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildInt32Column(b.values[key]), nil
}

func buildInt32Column(values []*int32) arrow.Array {
	bld := array.NewInt32Builder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewInt32Array()
}

func buildInt32ListColumn(values [][]int32) (arrow.Array, error) {
	bld := array.NewListBuilder(memory.DefaultAllocator, arrow.PrimitiveTypes.Int32)
	defer bld.Release()
	for range values {
		bld.AppendNull()
	}
	return bld.NewArray(), nil
}
