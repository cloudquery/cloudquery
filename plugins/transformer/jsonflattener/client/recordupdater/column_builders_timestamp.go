package recordupdater

import (
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/schemaupdater"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

type TimestampColumnsBuilder struct {
	i          int
	values     map[string][]*time.Time
	typeSchema map[string]string
}

func NewTimestampColumnsBuilder(typeSchema map[string]string, originalColumn *types.JSONArray) columnBuilder {
	b := &TimestampColumnsBuilder{i: -1, values: make(map[string][]*time.Time), typeSchema: typeSchema}
	for key, typ := range typeSchema {
		if typ != schemaupdater.TimestampType {
			continue
		}
		b.values[key] = make([]*time.Time, originalColumn.Len())
	}
	return b
}

func (b *TimestampColumnsBuilder) addRow(row map[string]any) {
	b.i++
	for key, typ := range b.typeSchema {
		if typ != schemaupdater.TimestampType {
			continue
		}
		value, exists := row[key]
		if !exists {
			b.values[key][b.i] = nil
			continue
		}
		if v, ok := value.(time.Time); ok {
			b.values[key][b.i] = &v
		}
	}
}

func (b *TimestampColumnsBuilder) build(key string) (arrow.Array, error) {
	if _, ok := b.values[key]; !ok {
		return nil, nil
	}
	return buildTimestampColumn(b.values[key])
}

func buildTimestampColumn(values []*time.Time) (arrow.Array, error) {
	bld := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
	defer bld.Release()
	for _, value := range values {
		if value == nil {
			bld.AppendNull()
			continue
		}
		ts, err := arrow.TimestampFromTime(*value, arrow.Microsecond)
		if err != nil {
			return nil, err
		}
		bld.Append(ts)
	}
	return bld.NewArray(), nil
}
