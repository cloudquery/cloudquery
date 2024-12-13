package types

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

func nullableType(name string, col *column.Nullable) (*arrow.Field, error) {
	base, err := fieldFromColumn(col.Base())
	if err != nil {
		return nil, err
	}
	return &arrow.Field{Name: name, Type: base.Type, Nullable: true}, nil
}
