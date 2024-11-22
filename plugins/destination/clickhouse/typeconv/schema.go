package typeconv

import (
	"github.com/apache/arrow/go/v17/arrow"
	arrow_types "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/arrow/types"
	ch_types "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func CanonizedTables(messages message.WriteMigrateTables) (schema.Tables, error) {
	canonized := make(schema.Tables, len(messages))
	var err error
	for i, msg := range messages {
		canonized[i], err = CanonizedTable(msg.Table)
		if err != nil {
			return nil, err
		}
	}
	return canonized, nil
}

func CanonizedTable(table *schema.Table) (*schema.Table, error) {
	columns := make(schema.ColumnList, len(table.Columns))
	for i, col := range table.Columns {
		canonized, err := CanonizedField(col.ToArrowField())
		if err != nil {
			return nil, err
		}
		columns[i] = schema.NewColumnFromArrowField(*canonized)
	}

	return &schema.Table{Name: table.Name, Columns: columns}, nil
}

// CanonizedField allows to know what type we associate the Apache Arrow type with.
// Several different Apache Arrow types will produce the same canonical type & that'll be the type we'll use in the database.
func CanonizedField(field arrow.Field) (*arrow.Field, error) {
	// 1 - convert to the ClickHouse
	fieldType, err := ch_types.FieldType(field)
	if err != nil {
		return nil, err
	}
	// 2 - convert back to Apache Arrow
	temp, err := arrow_types.Field(field.Name, fieldType)
	if err != nil {
		return nil, err
	}
	clone := field
	clone.Type = temp.Type
	return &clone, nil
}
