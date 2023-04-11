package queries

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"golang.org/x/exp/maps"
)

func SQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bit"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float" // = float(53)
	case schema.TypeUUID:
		return "uniqueidentifier"
	case schema.TypeByteArray:
		return "varbinary(max)"
	case schema.TypeTimestamp:
		return "datetime2"
	case schema.TypeString,
		schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet:
		return "nvarchar(4000)" // feasible to see these as PK, so need to limit the value
	case schema.TypeStringArray,
		schema.TypeJSON,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray:
		return "nvarchar(max)"
	default:
		panic("unknown type " + t.String())
	}
}

func SchemaType(tableName string, columnName string, sqlType string) (schema.ValueType, error) {
	sqlToSchema := map[string]schema.ValueType{
		"bit":              schema.TypeBool,
		"bigint":           schema.TypeInt,
		"float":            schema.TypeFloat,
		"uniqueidentifier": schema.TypeUUID,
		"varbinary(max)":   schema.TypeByteArray,
		"datetime2":        schema.TypeTimestamp,
		"nvarchar(4000)":   schema.TypeString,
		"nvarchar(max)":    schema.TypeStringArray,
	}

	if v, ok := sqlToSchema[sqlType]; ok {
		return v, nil
	}

	return schema.TypeInvalid, fmt.Errorf("got unknown MSSQL type %q of column %q for table %q while trying to convert it to CloudQuery internal schema type. Supported MSSQL types are %q", sqlType, columnName, tableName, maps.Keys(sqlToSchema))
}

// columnGoType has to be in sync with SQLType
func columnGoType(t schema.ValueType) reflect.Type {
	switch t {
	case schema.TypeBool:
		return reflect.TypeOf(true)
	case schema.TypeInt:
		return reflect.TypeOf(int64(0))
	case schema.TypeFloat:
		return reflect.TypeOf(float64(0))
	case schema.TypeUUID, schema.TypeByteArray:
		return reflect.TypeOf([]byte{})
	case schema.TypeTimestamp:
		return reflect.TypeOf(time.Time{})
	case schema.TypeString,
		schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet,
		schema.TypeStringArray,
		schema.TypeJSON,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray:
		return reflect.TypeOf("")
	default:
		panic("unknown type " + t.String())
	}
}
