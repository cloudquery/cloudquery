package types

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
)

func arrayType(name string, col *column.Array) (*arrow.Field, error) {
	base, err := fieldFromColumn(col.Base())
	if err != nil {
		return nil, err
	}

	var _type arrow.DataType
	if _, nullable := col.Base().(*column.Nullable); nullable {
		_type = arrow.ListOf(base.Type)
	} else {
		_type = arrow.ListOfNonNullable(base.Type)
	}

	// ClockHouse arrays are always non-nullable
	return &arrow.Field{Name: name, Type: _type, Nullable: false}, nil
}
