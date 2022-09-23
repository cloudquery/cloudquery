package convert

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func SchemaTypeToPg(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "boolean", nil
	case schema.TypeInt:
		return "bigint", nil
	case schema.TypeFloat:
		return "real", nil
	case schema.TypeUUID:
		return "uuid", nil
	case schema.TypeString:
		return "text", nil
	case schema.TypeStringArray:
		return "text[]", nil
	case schema.TypeTimestamp:
		return "timestamp without time zone", nil
	case schema.TypeJSON:
		return "jsonb", nil
	case schema.TypeUUIDArray:
		return "uuid[]", nil
	case schema.TypeInetArray:
		return "inet[]", nil
	case schema.TypeCIDR:
		return "cidr", nil
	case schema.TypeCIDRArray:
		return "cidr[]", nil
	case schema.TypeMacAddr:
		return "macaddr", nil
	case schema.TypeMacAddrArray:
		return "macaddr[]", nil
	case schema.TypeInet:
		return "inet", nil
	case schema.TypeIntArray:
		return "bigint[]", nil
	default:
		return "", fmt.Errorf("unknown type %s", t)
	}
}

// ValueTypeFromString this function is mainly used by https://github.com/cloudquery/cq-gen
func ValueTypeFromString(s string) schema.ValueType {
	switch strings.TrimPrefix(strings.ToLower(s), "type") {
	case "bool":
		return schema.TypeBool
	case "int", "bigint", "smallint":
		return schema.TypeInt
	case "float":
		return schema.TypeFloat
	case "uuid":
		return schema.TypeUUID
	case "string":
		return schema.TypeString
	case "json":
		return schema.TypeJSON
	case "intarray":
		return schema.TypeIntArray
	case "stringarray":
		return schema.TypeStringArray
	case "bytearray":
		return schema.TypeByteArray
	case "timestamp":
		return schema.TypeTimestamp
	case "uuidarray":
		return schema.TypeUUIDArray
	case "inet":
		return schema.TypeInet
	case "inetarray":
		return schema.TypeInetArray
	case "macaddr":
		return schema.TypeMacAddr
	case "macaddrarray":
		return schema.TypeMacAddrArray
	case "cidr":
		return schema.TypeCIDR
	case "cidrarray":
		return schema.TypeCIDRArray
	case "invalid":
		return schema.TypeInvalid
	default:
		return schema.TypeInvalid
	}
}
