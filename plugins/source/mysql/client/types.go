package client

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func SQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bool"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float"
	case schema.TypeUUID:
		return "binary(16)"
	case schema.TypeByteArray:
		return "blob"
	case schema.TypeTimestamp:
		return "datetime"
	case schema.TypeString,
		schema.TypeStringArray,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray,
		schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet:
		return "text"
	case schema.TypeJSON:
		return "json"
	default:
		panic("unknown type " + t.String())
	}
}

func SchemaType(tableName string, columnName string, dataType string, columnType string) (schema.ValueType, error) {
	if columnType == "binary(16)" {
		return schema.TypeUUID, nil
	}
	if columnType == "tinyint(1)" {
		return schema.TypeBool, nil
	}
	switch dataType {
	case "bool", "boolean":
		return schema.TypeBool, nil
	case "float", "decimal", "double":
		return schema.TypeFloat, nil
	case "datetime", "timestamp":
		return schema.TypeTimestamp, nil
	case "json":
		return schema.TypeJSON, nil
	case "binary":
		return schema.TypeByteArray, nil
	case "tinyint", "smallint", "int", "mediumint", "bigint":
		return schema.TypeInt, nil
	case "blob":
		return schema.TypeByteArray, nil
	}

	return schema.TypeString, nil
}
