package types

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
)

func FieldType(field arrow.Field) (string, error) {
	dt, err := columnType(field.Type)
	if err != nil {
		return "", err
	}

	// We allow nullable values in arrays, but arrays shouldn't be nullable themselves in ClickHouse
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return dt, nil
	}

	return "Nullable(" + dt + ")", nil
}

func FieldDefinition(field arrow.Field) (string, error) {
	fieldType, err := FieldType(field)
	if err != nil {
		return "", err
	}

	return util.SanitizeID(field.Name) + " " + fieldType, nil
}
