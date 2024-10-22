package recordupdater

import (
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const UTF8Type = "utf8"

type UTF8ColumnsBuilder struct {
	i          int
	values     map[string][]*string
	typeSchema map[string]string
}

func NewUTF8ColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) ColumnBuilder {
	b := &UTF8ColumnsBuilder{i: -1, values: make(map[string][]*string), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != UTF8Type {
			continue
		}
		b.values[key] = make([]*string, originalColumn.Len())
	}
	return b
}

func (b *UTF8ColumnsBuilder) AddRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != UTF8Type {
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

func (b *UTF8ColumnsBuilder) Build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildUTF8Column(b.values[key]), nil
}

func buildUTF8Column(values []*string) arrow.Array {
	bld := array.NewStringBuilder(memory.DefaultAllocator)
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		bld.Append(*value)
	}
	return bld.NewStringArray()
}
