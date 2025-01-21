package recordupdater

import (
	"net"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type InetColumnsBuilder struct {
	i          int
	values     map[string][]*any
	typeSchema map[string]string
}

func NewInetColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &InetColumnsBuilder{i: -1, values: make(map[string][]*any), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != InetType {
			continue
		}
		b.values[key] = make([]*any, originalColumn.Len())
	}
	return b
}

func (b *InetColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != InetType {
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

func (b *InetColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildInetColumn(b.values[key]), nil
}

func buildInetColumn(values []*any) arrow.Array {
	bld := types.NewInetBuilder(memory.DefaultAllocator)
	defer bld.Release()
	for range values {
		bld.AppendNull()
	}
	return bld.NewInetArray()
}

func buildInetListColumn(values [][]net.IPNet) arrow.Array {
	bld := array.NewListBuilder(memory.DefaultAllocator, types.NewInetType())
	defer bld.Release()
	for range values {
		bld.AppendNull()
	}
	return bld.NewListArray()
}
