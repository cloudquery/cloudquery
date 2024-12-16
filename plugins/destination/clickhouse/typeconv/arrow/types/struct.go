package types

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

func structType(name string, tuple *column.Tuple) (*arrow.Field, error) {
	columns, err := parseTupleType(tuple.Type(), time.UTC)
	if err != nil {
		return nil, err
	}

	fields := make([]arrow.Field, len(columns))
	for i, col := range columns {
		field, err := fieldFromColumn(col)
		if err != nil {
			return nil, err
		}
		fields[i] = *field
	}

	return &arrow.Field{Name: name, Type: arrow.StructOf(fields...)}, nil
}
