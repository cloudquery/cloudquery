package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
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
		types.ExtensionTypes.MAC:           bigquery.StringFieldType,
		types.ExtensionTypes.UUID:          bigquery.StringFieldType,
		types.ExtensionTypes.JSON:          bigquery.JSONFieldType,
	}
)

func (c *Client) ArrowTypeToBigQuery(dt arrow.DataType) (ft bigquery.FieldType, schema []*bigquery.FieldSchema, repeated bool) {
	switch v := dt.(type) {
	case *arrow.ListType:
		ift, sch, isRepeated := c.ArrowTypeToBigQuery(v.Elem())
		if isRepeated {
			// store nested arrays as JSON
			return bigquery.JSONFieldType, nil, false
		}
		return ift, sch, true
	case *arrow.MapType:
		// store maps as JSON
		return bigquery.JSONFieldType, nil, false
	}

	switch dt.ID() {
	case arrow.STRUCT:
		return bigquery.RecordFieldType, nil, false
	case arrow.INTERVAL_MONTH_DAY_NANO:
		return bigquery.RecordFieldType, nil, false
	}
	for k, v := range TypeMapping {
		if arrow.TypeEqual(k, dt) {
			return v, nil, false
		}
	}
	// default to string
	return bigquery.StringFieldType, nil, false
}
