package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

var (
	TypeMapping = map[arrow.DataType]bigquery.FieldType{
		arrow.FixedWidthTypes.Boolean:      bigquery.BooleanFieldType,
		arrow.PrimitiveTypes.Int8:          bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Int16:         bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Int32:         bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Int64:         bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Uint8:         bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Uint16:        bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Uint32:        bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Uint64:        bigquery.IntegerFieldType,
		arrow.PrimitiveTypes.Float32:       bigquery.FloatFieldType,
		arrow.PrimitiveTypes.Float64:       bigquery.FloatFieldType,
		arrow.BinaryTypes.String:           bigquery.StringFieldType,
		arrow.BinaryTypes.LargeString:      bigquery.StringFieldType,
		arrow.BinaryTypes.Binary:           bigquery.BytesFieldType,
		arrow.BinaryTypes.LargeBinary:      bigquery.BytesFieldType,
		arrow.FixedWidthTypes.Date32:       bigquery.DateFieldType,
		arrow.FixedWidthTypes.Date64:       bigquery.DateFieldType,
		arrow.FixedWidthTypes.Timestamp_s:  bigquery.TimestampFieldType,
		arrow.FixedWidthTypes.Timestamp_ms: bigquery.TimestampFieldType,
		arrow.FixedWidthTypes.Timestamp_us: bigquery.TimestampFieldType,
		arrow.FixedWidthTypes.Timestamp_ns: bigquery.TimestampFieldType,
		types.ExtensionTypes.Inet:          bigquery.StringFieldType,
		types.ExtensionTypes.Mac:           bigquery.StringFieldType,
		types.ExtensionTypes.UUID:          bigquery.StringFieldType,
		types.ExtensionTypes.JSON:          bigquery.JSONFieldType,
	}
)

func (c *Client) SchemaTypeToBigQueryType(dt arrow.DataType) (ft bigquery.FieldType, repeated bool) {
	if arrow.IsListLike(dt.ID()) {
		ift, isRepeated := c.SchemaTypeToBigQueryType(dt.(*arrow.ListType).Elem())
		if isRepeated {
			// store nested arrays as JSON
			return bigquery.JSONFieldType, false
		}
		return ift, true
	}
	for k, v := range TypeMapping {
		if arrow.TypeEqual(k, dt) {
			return v, false
		}
	}
	panic("supported type: " + dt.Name())
}
