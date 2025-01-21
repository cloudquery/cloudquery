package recordupdater

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
)

type UUIDColumnsBuilder struct {
	i          int
	values     map[string][]*any
	typeSchema map[string]string
}

func NewUUIDColumnsBuilder(typeSchema map[string]string, originalColumn *types.UUIDArray) columnBuilder {
	b := &UUIDColumnsBuilder{i: -1, values: make(map[string][]*any), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != UUIDType {
			continue
		}
		b.values[key] = make([]*any, originalColumn.Len())
	}
	return b
}

func (b *UUIDColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != UUIDType {
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

func (b *UUIDColumnsBuilder) build(key string) (arrow.Array, error) {
	return nil, nil
}

func buildUUIDColumn(values []uuid.UUID) arrow.Array {
	bld := types.NewUUIDBuilder(memory.DefaultAllocator)
	defer bld.Release()
	for _, value := range values {
		bld.Append(value)
	}
	return bld.NewUUIDArray()
}

func buildUUIDListColumn(values [][]uuid.UUID) (arrow.Array, error) {
	bld := array.NewListBuilder(memory.DefaultAllocator, types.NewUUIDType())
	defer bld.Release()
	for range values {
		bld.AppendNull()
	}
	return bld.NewArray(), nil
}
