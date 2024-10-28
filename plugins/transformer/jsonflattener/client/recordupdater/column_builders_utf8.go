package recordupdater

import (
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type UTF8ColumnsBuilder struct {
	i          int
	values     map[string][]*string
	typeSchema map[string]string
}

func NewUTF8ColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &UTF8ColumnsBuilder{i: -1, values: make(map[string][]*string), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.UTF8Type {
			continue
		}
		b.values[key] = make([]*string, originalColumn.Len())
	}
	return b
}

func (b *UTF8ColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.UTF8Type {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(string); ok {
			b.values[key][b.i] = &v
		}
	}
}

func (b *UTF8ColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildUTF8Column(b.values[key]), nil
}

func buildUTF8Column(values []*string) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewStringArray()
}
