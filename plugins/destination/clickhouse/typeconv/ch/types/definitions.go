package types

import (
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/util"
)

func FieldType(field arrow.Field) (string, error) {
	dt, err := ColumnType(field.Type)
	if err != nil {
		return "", err
	}

	if CanBeNullable(field.Type) && field.Nullable {
		return "Nullable(" + dt + ")", nil
	}

	return dt, nil
}

func FieldDefinition(field arrow.Field) (string, error) {
	fieldType, err := FieldType(field)
	if err != nil {
		return "", err
	}

	return util.SanitizeID(field.Name) + " " + fieldType, nil
}
