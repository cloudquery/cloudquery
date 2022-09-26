package convert

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

// SchemaTypeToPg copied from cli/internal/destinations/postgresql
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

// ValueTypeFromString uses plugin-sdk, but with fix for typo in inetarray
func ValueTypeFromString(s string) schema.ValueType {
	if strings.TrimPrefix(strings.ToLower(s), "type") == "inetarray" {
		return schema.TypeInetArray
	} else if strings.TrimPrefix(strings.ToLower(s), "type") == "timestamp without timezone" {
		return schema.TypeTimestamp
	}
	return schema.ValueTypeFromString(s)
}
