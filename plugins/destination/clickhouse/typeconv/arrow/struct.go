package arrow

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
)

func structType(tuple *column.Tuple) (*arrow.StructType, error) {
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

	return arrow.StructOf(fields...), nil
}
