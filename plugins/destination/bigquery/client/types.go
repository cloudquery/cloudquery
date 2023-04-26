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

func GetValueForBigQuery(arr arrow.Array, i int) any {
	return arr.GetOneForMarshal(i) // maybe this works for now?

	//dt := arr.DataType()
	//switch {
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Boolean):
	//	return arr.(*array.Boolean).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int8):
	//	return arr.(*array.Int8).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int16):
	//	return arr.(*array.Int16).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int32):
	//	return arr.(*array.Int32).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Int64):
	//	return arr.(*array.Int64).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint8):
	//	return arr.(*array.Uint8).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint16):
	//	return arr.(*array.Uint16).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint32):
	//	return arr.(*array.Uint32).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Uint64):
	//	return arr.(*array.Uint64).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float32):
	//	return arr.(*array.Float32).Value(i)
	//case arrow.TypeEqual(dt, arrow.PrimitiveTypes.Float64):
	//	return arr.(*array.Float64).Value(i)
	//case arrow.TypeEqual(dt, arrow.BinaryTypes.String):
	//	return arr.(*array.String).Value(i)
	//case arrow.TypeEqual(dt, arrow.ListOf(arrow.BinaryTypes.String)):
	//	m := arr.(*array.List).GetOneForMarshal(i)
	//	j, _ := json.Marshal(m)
	//	var v []string
	//	json.Unmarshal(j, &v)
	//	return v
	//case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeString):
	//	return arr.(*array.LargeString).Value(i)
	//case arrow.TypeEqual(dt, arrow.BinaryTypes.Binary):
	//	return arr.(*array.Binary).Value(i)
	//case arrow.TypeEqual(dt, arrow.BinaryTypes.LargeBinary):
	//	return arr.(*array.LargeBinary).Value(i)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date32):
	//	return arr.(*array.Date32).Value(i)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Date64):
	//	return arr.(*array.Date64).Value(i)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_s):
	//	return arr.(*array.Timestamp).Value(i).ToTime(arrow.Second)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ms):
	//	return arr.(*array.Timestamp).Value(i).ToTime(arrow.Millisecond)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_us):
	//	return arr.(*array.Timestamp).Value(i).ToTime(arrow.Microsecond)
	//case arrow.TypeEqual(dt, arrow.FixedWidthTypes.Timestamp_ns):
	//	return arr.(*array.Timestamp).Value(i).ToTime(arrow.Nanosecond)
	//case arrow.TypeEqual(dt, types.ExtensionTypes.Inet):
	//	return arr.(*types.InetArray).ValueStr(i)
	//case arrow.TypeEqual(dt, types.ExtensionTypes.Mac):
	//	return arr.(*types.MacArray).ValueStr(i)
	//case arrow.TypeEqual(dt, types.ExtensionTypes.UUID):
	//	return arr.(*types.UUIDArray).ValueStr(i)
	//case arrow.TypeEqual(dt, types.ExtensionTypes.JSON):
	//	return arr.(*types.JSONArray).ValueStr(i)
	//case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Inet)):
	//	// not sure if there is a better way to do this available right now
	//	m := arr.(*array.List).GetOneForMarshal(i)
	//	j, _ := json.Marshal(m)
	//	var v []string
	//	json.Unmarshal(j, &v)
	//	return v
	//case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.Mac)):
	//	// not sure if there is a better way to do this available right now
	//	m := arr.(*array.List).GetOneForMarshal(i)
	//	j, _ := json.Marshal(m)
	//	var v []string
	//	json.Unmarshal(j, &v)
	//	return v
	//case arrow.TypeEqual(dt, arrow.ListOf(types.ExtensionTypes.UUID)):
	//	// not sure if there is a better way to do this available right now
	//	m := arr.(*array.List).GetOneForMarshal(i)
	//	j, _ := json.Marshal(m)
	//	var v []string
	//	json.Unmarshal(j, &v)
	//	return v
	//}
	//panic("unhandled type: " + dt.String())
}
