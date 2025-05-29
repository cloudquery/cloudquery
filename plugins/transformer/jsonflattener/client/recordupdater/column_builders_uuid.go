package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type UUIDColumnsBuilder struct {
	i          int
	values     map[string][]*string
	typeSchema map[string]string
}

func NewUUIDColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &UUIDColumnsBuilder{i: -1, values: make(map[string][]*string), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.UUIDType {
			continue
		}
		b.values[key] = make([]*string, originalColumn.Len())
	}
	return b
}

func (b *UUIDColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.UUIDType {
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

func (b *UUIDColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildUUIDColumn(b.values[key]), nil
}

func buildUUIDColumn(values []*string) arrow.Array {
	bld := types.NewUUIDBuilder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		err := bld.AppendValueFromString(*value)
		if err != nil {
			bld.AppendNull()
		}
	}
	return bld.NewUUIDArray()
}
