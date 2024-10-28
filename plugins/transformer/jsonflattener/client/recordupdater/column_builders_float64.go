package recordupdater

import (
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type Float64ColumnsBuilder struct {
	i          int
	values     map[string][]*float64
	typeSchema map[string]string
}

func NewFloat64ColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &Float64ColumnsBuilder{i: -1, values: make(map[string][]*float64), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.Float64Type {
			continue
		}
		b.values[key] = make([]*float64, originalColumn.Len())
	}
	return b
}

func (b *Float64ColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.Float64Type {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(float64); ok {
			b.values[key][b.i] = &v
		}
	}
}

func (b *Float64ColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildFloat64Column(b.values[key]), nil
}

func buildFloat64Column(values []*float64) arrow.Array {
	bld := array.NewFloat64Builder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewFloat64Array()
}
