package pgarrow

import (
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func ArrowToPg10(t arrow.DataType) string {
	switch dt := t.(type) {
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
	case *arrow.Time32Type, *arrow.Time64Type:
		return "time without time zone"
	case *arrow.Date32Type, *arrow.Date64Type:
		return "date"
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
	default:
		return "text"
	}
}

// ArrowToCockroach converts arrow data type to cockroach data type. CockroachDB lacks support for
// some data types like macaddr and has different aliases for ints.
// See: https://www.cockroachlabs.com/docs/stable/int.html
func ArrowToCockroach(t arrow.DataType) string {
	switch dt := t.(type) {
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
	case *arrow.Time32Type, *arrow.Time64Type:
		return "time without time zone"
	case *arrow.Date32Type, *arrow.Date64Type:
		return "date"
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
	default:
		return "text"
	}
}

func ArrowToCrateDB(t arrow.DataType) string {
	switch dt := t.(type) {
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
		// CrateDB does not support storing numeric types.
		return "text"
	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "double precision"
	case arrow.DecimalType:
		// CrateDB does not support storing numeric types
		return "text"
	case *arrow.StringType:
		return "text"
	case *arrow.BinaryType:
		// CrateDB does not support bytea
		return "text"
	case *arrow.LargeBinaryType:
		// CrateDB does not support bytea
		return "text"
	case *arrow.TimestampType:
		return "timestamp without time zone"
	case *arrow.Time32Type, *arrow.Time64Type:
		// CrateDB does not support "time without time zone", and does not support
		// storage of "time with time zone".
		return "text"
	case *arrow.Date32Type, *arrow.Date64Type:
		// CrateDB does not support storing "date" type
		return "timestamp without time zone"
	case *cqtypes.UUIDType:
		// CrateDB does not support UUID type
		return "text"
	case *cqtypes.JSONType:
		// CrateDB calls JSON type "object"
		return "object"
	case *cqtypes.MACType:
		// CrateDB does not support macaddr type
		return "text"
	case *cqtypes.InetType:
		// CrateDB does not support inet type, but we can use IP
		return "ip"
	case *arrow.ListType:
		return ArrowToCrateDB(dt.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return ArrowToCrateDB(dt.Elem()) + "[]"
	case *arrow.LargeListType:
		return ArrowToCrateDB(dt.Elem()) + "[]"
	case *arrow.MapType:
		return "text"
	default:
		return "text"
	}
}
