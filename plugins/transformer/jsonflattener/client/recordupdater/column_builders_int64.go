package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type Int64ColumnsBuilder struct {
	i          int
	values     map[string][]*int64
	typeSchema map[string]string
}

func NewInt64ColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &Int64ColumnsBuilder{i: -1, values: make(map[string][]*int64), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.Int64Type {
			continue
		}
		b.values[key] = make([]*int64, originalColumn.Len())
	}
	return b
}

func (b *Int64ColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.Int64Type {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(float64); ok {
			int64Value := int64(v)
			b.values[key][b.i] = &int64Value
		}
	}
}

func (b *Int64ColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildInt64Column(b.values[key]), nil
}

func buildInt64Column(values []*int64) arrow.Array {
	bld := array.NewInt64Builder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewInt64Array()
}
