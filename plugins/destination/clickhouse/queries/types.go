package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func chType(dataType arrow.DataType) string {

}

func chFieldType(field arrow.Field) string {
	typ := chType(field.Type)

	// We allow nullable values in arrays, but arrays cannot be nullable themselves
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return typ
	}

	return "Nullable(" + typ + ")"
}

func cqCol(name, typ string) schema.Column {
	res := schema.Column{Name: name}

	if strings.HasPrefix(typ, "Array(") {
		switch typ {
		case "Array(Nullable(Int64))":
			res.Type = schema.TypeIntArray
		case "Array(Nullable(String))":
			res.Type = schema.TypeStringArray
		case "Array(Nullable(UUID))":
			res.Type = schema.TypeUUIDArray
		default:
			panic("unsupported type " + typ)
		}
		return res
	}

	trimmed := strings.TrimPrefix(typ, "Nullable(")
	nullable := trimmed != typ
	if nullable {
		typ = strings.TrimSuffix(trimmed, ")")
	}

	switch typ {
	case "Bool":
		res.Type = schema.TypeBool
	case "Float64":
		res.Type = schema.TypeFloat
	case "Int64":
		res.Type = schema.TypeInt
	case "String":
		res.Type = schema.TypeString
	case "DateTime64(9)":
		res.Type = schema.TypeTimestamp
	case "UUID":
		res.Type = schema.TypeUUID
	default:
		panic("unsupported type " + typ)
	}

	res.CreationOptions.NotNull = !nullable

	return res
}

func normalizeColumn(column schema.Column) schema.Column {
	res := schema.Column{Name: column.Name}

	switch column.Type {
	case schema.TypeIntArray:
		res.Type = schema.TypeIntArray
	case schema.TypeStringArray,
		schema.TypeMacAddrArray,
		schema.TypeCIDRArray,
		schema.TypeInetArray:
		res.Type = schema.TypeStringArray
	case schema.TypeUUIDArray:
		res.Type = schema.TypeUUIDArray
	}

	if res.Type != schema.TypeInvalid {
		// arrays cannot be nullable in CH, so we don't bother with this check at all
		return res
	}

	switch t := column.Type; t {
	case schema.TypeBool:
		res.Type = schema.TypeBool

	case schema.TypeFloat:
		res.Type = schema.TypeFloat

	case schema.TypeInt:
		res.Type = schema.TypeInt

	case schema.TypeString,
		schema.TypeByteArray,
		// ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment.
		// https://github.com/ClickHouse/ClickHouse/issues/46190
		schema.TypeJSON,
		schema.TypeMacAddr,
		schema.TypeCIDR,
		schema.TypeInet:
		res.Type = schema.TypeString

	case schema.TypeTimestamp:
		res.Type = schema.TypeTimestamp

	case schema.TypeUUID:
		res.Type = schema.TypeUUID

	default:
		panic("unsupported type " + t.String())
	}

	res.CreationOptions.NotNull = column.CreationOptions.NotNull || column.CreationOptions.PrimaryKey

	return res
}
