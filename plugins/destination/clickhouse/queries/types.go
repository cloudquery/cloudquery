package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func chType(column *schema.Column) string {
	// We allow nullable values in arrays, but arrays cannot be nullable themselves
	switch column.Type {
	case schema.TypeIntArray:
		return "Array(Nullable(Int64))"

	case schema.TypeStringArray,
		schema.TypeMacAddrArray,
		schema.TypeCIDRArray,
		schema.TypeInetArray:
		return "Array(Nullable(String))"

	case schema.TypeUUIDArray:
		return "Array(Nullable(UUID))"
	}

	// Now we need to account for the nullable option
	var typ string
	switch t := column.Type; t {
	case schema.TypeBool:
		typ = "Bool"

	case schema.TypeFloat:
		typ = "Float64"

	case schema.TypeInt:
		typ = "Int64"

	case schema.TypeString,
		schema.TypeByteArray,
		// ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment.
		// https://github.com/ClickHouse/ClickHouse/issues/46190
		schema.TypeJSON,
		schema.TypeMacAddr,
		schema.TypeCIDR,
		schema.TypeInet:
		typ = "String"

	case schema.TypeTimestamp:
		typ = "DateTime64(9)"

	case schema.TypeUUID:
		typ = "UUID"

	default:
		panic("unsupported type " + t.String())
	}

	if column.CreationOptions.NotNull {
		return typ
	}

	return "Nullable(" + typ + ")"
}
