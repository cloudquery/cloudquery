package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) SchemaTypeToSqlite(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "integer"
	case schema.TypeInt:
		return "integer"
	case schema.TypeFloat:
		return "real"
	case schema.TypeByteArray:
		return "blob"
	default:
		return "text"
	}
}

func (*Client) sqliteTypeToSchema(t string) schema.ValueType {
	switch t {
	case "integer":
		return schema.TypeInt
	case "real":
		return schema.TypeFloat
	case "text":
		return schema.TypeString
	case "blob":
		return schema.TypeByteArray
	case "timestamp":
		return schema.TypeTimestamp
	default:
		panic("unknown type")
	}
}
