package client

import (
	"cloud.google.com/go/elasticsearch"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) SchemaTypeToElasticsearch(t schema.ValueType) (ft elasticsearch.FieldType, repeated bool) {
	switch t {
	case schema.TypeBool:
		return elasticsearch.BooleanFieldType, false
	case schema.TypeInt:
		return elasticsearch.IntegerFieldType, false
	case schema.TypeFloat:
		return elasticsearch.FloatFieldType, false
	case schema.TypeUUID:
		return elasticsearch.StringFieldType, false
	case schema.TypeString:
		return elasticsearch.StringFieldType, false
	case schema.TypeByteArray:
		return elasticsearch.BytesFieldType, false
	case schema.TypeStringArray:
		return elasticsearch.StringFieldType, true
	case schema.TypeTimestamp:
		return elasticsearch.TimestampFieldType, false
	case schema.TypeJSON:
		return elasticsearch.JSONFieldType, false
	case schema.TypeUUIDArray:
		return elasticsearch.StringFieldType, true
	case schema.TypeCIDR:
		return elasticsearch.StringFieldType, false
	case schema.TypeCIDRArray:
		return elasticsearch.StringFieldType, true
	case schema.TypeMacAddr:
		return elasticsearch.StringFieldType, false
	case schema.TypeMacAddrArray:
		return elasticsearch.StringFieldType, true
	case schema.TypeInet:
		return elasticsearch.StringFieldType, false
	case schema.TypeInetArray:
		return elasticsearch.StringFieldType, true
	case schema.TypeIntArray:
		return elasticsearch.IntegerFieldType, true
	default:
		panic("unknown type")
	}
}
