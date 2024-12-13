package recordupdater

import (
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
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
		b.values[key][b.i] = nil
		if !exists {
			continue
		}
		switch v := value.(type) {
		case string:
			formats := []string{
				"2006-01-02T15:04:05.000Z",
				time.RFC3339,
				time.RFC3339Nano,
				"2006-01-02T15:04:05Z",
				"2006-01-02 15:04:05",
				"2006-01-02",
				time.RFC822,
				time.RFC850,
				time.RFC1123,
				time.RFC1123Z,
				time.UnixDate,
				time.RubyDate,
				time.ANSIC,
			}

			for _, format := range formats {
				if parsed, err := time.Parse(format, v); err == nil {
					b.values[key][b.i] = &parsed
					break
				}
			}
		case time.Time:
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
