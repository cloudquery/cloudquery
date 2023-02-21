package queries

import (
	"reflect"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
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

func SchemaType(t string) schema.ValueType {
	switch t {
	case "bit":
		return schema.TypeBool
	case "bigint":
		return schema.TypeInt
	case "float":
		return schema.TypeFloat
	case "uniqueidentifier":
		return schema.TypeUUID
	case "varbinary(max)":
		return schema.TypeByteArray
	case "datetime2":
		return schema.TypeTimestamp
	case "nvarchar(4000)":
		return schema.TypeString
	case "nvarchar(max)":
		return schema.TypeStringArray
	default:
		panic("unknown type " + t)
	}
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
