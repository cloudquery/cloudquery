package types

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow-go/v18/arrow"
)

func arrayType(name string, col *column.Array) (*arrow.Field, error) {
	base, err := fieldFromColumn(col.Base())
	if err != nil {
		return nil, err
	}

	var dataType arrow.DataType
	_, nullable := col.Base().(*column.Nullable)
	if nullable {
		dataType = arrow.ListOf(base.Type)
	} else {
		dataType = arrow.ListOfNonNullable(base.Type)
	}

	return &arrow.Field{Name: name, Type: dataType}, nil
}
