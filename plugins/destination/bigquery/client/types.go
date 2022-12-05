package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) SchemaTypeToBigQuery(t schema.ValueType) (ft bigquery.FieldType, repeated bool) {
	switch t {
	case schema.TypeBool:
		return bigquery.BooleanFieldType, false
	case schema.TypeInt:
		return bigquery.IntegerFieldType, false
	case schema.TypeFloat:
		return bigquery.FloatFieldType, false
	case schema.TypeUUID:
		return bigquery.StringFieldType, false
	case schema.TypeString:
		return bigquery.StringFieldType, false
	case schema.TypeByteArray:
		return bigquery.BytesFieldType, false
	case schema.TypeStringArray:
		return bigquery.StringFieldType, true
	case schema.TypeTimestamp:
		return bigquery.TimestampFieldType, false
	case schema.TypeJSON:
		return bigquery.JSONFieldType, false
	case schema.TypeUUIDArray:
		return bigquery.StringFieldType, true
	case schema.TypeCIDR:
		return bigquery.StringFieldType, false
	case schema.TypeCIDRArray:
		return bigquery.StringFieldType, true
	case schema.TypeMacAddr:
		return bigquery.StringFieldType, false
	case schema.TypeMacAddrArray:
		return bigquery.StringFieldType, true
	case schema.TypeInet:
		return bigquery.StringFieldType, false
	case schema.TypeInetArray:
		return bigquery.StringFieldType, true
	case schema.TypeIntArray:
		return bigquery.IntegerFieldType, true
	default:
		panic("unknown type")
	}
}
