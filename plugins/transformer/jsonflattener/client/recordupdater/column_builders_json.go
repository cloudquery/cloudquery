package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type JSONColumnsBuilder struct {
	i          int
	values     map[string][]*any
	typeSchema map[string]string
}

func NewJSONColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &JSONColumnsBuilder{i: -1, values: make(map[string][]*any), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.JSONType {
			continue
		}
		b.values[key] = make([]*any, originalColumn.Len())
	}
	return b
}

func (b *JSONColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.JSONType {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		b.values[key][b.i] = &value
	}
}

func (b *JSONColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildJSONColumn(b.values[key]), nil
}

func buildJSONColumn(values []*any) arrow.Array {
	bld := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewJSONType()))
	defer bld.Release()
	for _, value := range values {
		bld.Append(value)
	}
	return bld.NewJSONArray()
}
