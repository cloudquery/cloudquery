package client

import (
	"github.com/apache/arrow-go/v18/arrow"
)

func (*Client) arrowTypeToSqliteStr(t arrow.DataType) string {
	switch t.ID() {
	case arrow.BINARY, arrow.LARGE_BINARY:
		return "blob"
	case arrow.STRING, arrow.LARGE_STRING:
		return "text"
	case arrow.INT8, arrow.INT16, arrow.INT32, arrow.INT64, arrow.UINT8, arrow.UINT16, arrow.UINT32, arrow.UINT64:
		return "integer"
	case arrow.FLOAT16, arrow.FLOAT32, arrow.FLOAT64:
		return "real"
	case arrow.BOOL:
		return "boolean"
	case arrow.TIMESTAMP:
		return "timestamp"
	default:
		return "text"
	}
}

func (*Client) arrowTypeToSqlite(t arrow.DataType) arrow.DataType {
	switch t.ID() {
	case arrow.BINARY, arrow.LARGE_BINARY:
		return arrow.BinaryTypes.LargeString
	case arrow.STRING, arrow.LARGE_STRING:
		return arrow.BinaryTypes.LargeString
	case arrow.INT8, arrow.INT16, arrow.INT32, arrow.INT64, arrow.UINT8, arrow.UINT16, arrow.UINT32, arrow.UINT64:
		return arrow.PrimitiveTypes.Int64
	case arrow.FLOAT16, arrow.FLOAT32, arrow.FLOAT64:
		return arrow.PrimitiveTypes.Float64
	case arrow.BOOL:
		return arrow.FixedWidthTypes.Boolean
	case arrow.TIMESTAMP:
		return arrow.FixedWidthTypes.Timestamp_us
	default:
		return arrow.BinaryTypes.LargeString
	}
}

func (*Client) sqliteTypeToArrowType(t string) arrow.DataType {
	switch t {
	case "integer":
		return arrow.PrimitiveTypes.Int64
	case "real":
		return arrow.PrimitiveTypes.Float64
	case "text":
		return arrow.BinaryTypes.LargeString
	case "blob":
		return arrow.BinaryTypes.LargeBinary
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "timestamp":
		return arrow.FixedWidthTypes.Timestamp_us
	default:
		panic("unknown type: " + t)
	}
}
