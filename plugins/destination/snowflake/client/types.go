package client

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func (*Client) SchemaTypeToSnowflake(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.ListType, *arrow.FixedSizeListType:
		return "array"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type, *arrow.Uint8Type, *arrow.Int16Type, *arrow.Uint16Type,
		*arrow.Int32Type, *arrow.Uint32Type, *arrow.Int64Type, *arrow.Uint64Type:
		return "number"
	case *arrow.Float32Type, *arrow.Float64Type:
		return "float"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "binary"
	case *arrow.TimestampType:
		return "timestamp_tz"
	case *types.JSONType, *arrow.StructType:
		return "variant"
	default:
		return "text"
	}
}
