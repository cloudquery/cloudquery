package pgarrow

import (
	"strconv"

	"github.com/apache/arrow/go/v13/arrow"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func ArrowToPg10(t arrow.DataType) string {
	switch dt := t.(type) {
	case *cqtypes.UUIDType:
		return "uuid"
	case *cqtypes.JSONType:
		return "jsonb"
	case *cqtypes.MACType:
		return "macaddr"
	case *cqtypes.InetType:
		return "inet"
	case *arrow.ListType:
		return ArrowToPg10(dt.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return ArrowToPg10(dt.Elem()) + "[]"
	case *arrow.LargeListType:
		return ArrowToPg10(dt.Elem()) + "[]"
	case *arrow.MapType:
		return "text"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type:
		return "smallint"
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Int32Type:
		return "int"
	case *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint8Type:
		return "smallint"
	case *arrow.Uint16Type:
		return "int"
	case *arrow.Uint32Type:
		return "bigint"
	case *arrow.Uint64Type:
		return "numeric(20,0)"
	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "double precision"
	case arrow.DecimalType:
		return "numeric(" + strconv.Itoa(int(dt.GetPrecision())) + "," + strconv.Itoa(int(dt.GetScale())) + ")"
	case *arrow.StringType:
		return "text"
	case *arrow.BinaryType:
		return "bytea"
	case *arrow.LargeBinaryType:
		return "bytea"
	case *arrow.TimestampType:
		return "timestamp without time zone"
	default:
		return "text"
	}
}

// ArrowToCockroach converts arrow data type to cockroach data type. CockroachDB lacks support for
// some data types like macaddr and has different aliases for ints.
// See: https://www.cockroachlabs.com/docs/stable/int.html
func ArrowToCockroach(t arrow.DataType) string {
	switch dt := t.(type) {
	case *cqtypes.UUIDType:
		return "uuid"
	case *cqtypes.JSONType:
		return "jsonb"
	case *cqtypes.MACType:
		return "text"
	case *cqtypes.InetType:
		return "inet"
	case *arrow.ListType:
		return ArrowToCockroach(dt.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return ArrowToCockroach(dt.Elem()) + "[]"
	case *arrow.LargeListType:
		return ArrowToCockroach(dt.Elem()) + "[]"
	case *arrow.MapType:
		return "text"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type:
		return "int2"
	case *arrow.Int16Type:
		return "int2"
	case *arrow.Int32Type:
		return "int8"
	case *arrow.Int64Type:
		return "int8"
	case *arrow.Uint8Type:
		return "int2"
	case *arrow.Uint16Type:
		return "int8"
	case *arrow.Uint32Type:
		return "int8"
	case *arrow.Uint64Type:
		return "numeric(20,0)"
	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "double precision"
	case arrow.DecimalType:
		return "numeric(" + strconv.Itoa(int(dt.GetPrecision())) + "," + strconv.Itoa(int(dt.GetScale())) + ")"
	case *arrow.StringType:
		return "text"
	case *arrow.BinaryType:
		return "bytea"
	case *arrow.LargeBinaryType:
		return "bytea"
	case *arrow.TimestampType:
		return "timestamp without time zone"
	default:
		return "text"
	}
}
