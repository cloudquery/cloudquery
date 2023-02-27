package client

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
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
	case schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet:
		return "nvarchar(255)"
	case schema.TypeString,
		schema.TypeStringArray,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray:
		return "text"
	case schema.TypeJSON:
		return "json"
	default:
		panic("unknown type " + t.String())
	}
}

func SchemaType(tableName string, columnName string, sqlType string) (schema.ValueType, error) {
	sqlTypeToSchemaType := map[string]schema.ValueType{
		"bool":          schema.TypeBool,
		"tinyint(1)":    schema.TypeBool,
		"bigint":        schema.TypeInt,
		"bigint(20)":    schema.TypeInt,
		"float":         schema.TypeFloat,
		"binary(16)":    schema.TypeUUID,
		"blob":          schema.TypeByteArray,
		"datetime":      schema.TypeTimestamp,
		"nvarchar(255)": schema.TypeInet,
		"varchar(255)":  schema.TypeInet,
		"text":          schema.TypeString,
		"json":          schema.TypeJSON,
	}

	if v, ok := sqlTypeToSchemaType[sqlType]; ok {
		return v, nil
	}

	return schema.TypeInvalid, fmt.Errorf("got unknown MySQL type %q for column %q of table %q while trying to convert it to CloudQuery internal schema type. Supported MySQL types are %q", sqlType, columnName, tableName, maps.Keys(sqlTypeToSchemaType))
}
