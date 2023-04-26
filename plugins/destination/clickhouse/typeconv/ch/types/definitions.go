package types

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
)

func FieldType(field arrow.Field) (string, error) {
	_type, err := dataType(field.Type)
	if err != nil {
		return "", err
	}

	// We allow nullable values in arrays, but arrays shouldn't be nullable themselves in ClickHouse
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return _type, nil
	}

	return "Nullable(" + _type + ")", nil
}

func FieldDefinition(field arrow.Field) (string, error) {
	_type, err := FieldType(field)
	if err != nil {
		return "", err
	}

	return util.SanitizeID(field.Name) + " " + _type, nil
}
