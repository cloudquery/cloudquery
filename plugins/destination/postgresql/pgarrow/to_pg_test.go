package pgarrow

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/assert"
)

func TestArrowToPg10(t *testing.T) {
	cases := []struct {
		want     string
		dataType arrow.DataType
	}{
		{"boolean", arrow.FixedWidthTypes.Boolean},
		{"bigint", arrow.PrimitiveTypes.Int64},
		{"double precision", arrow.PrimitiveTypes.Float64},
		{"uuid", types.ExtensionTypes.UUID},
		{"bytea", arrow.BinaryTypes.Binary},
		{"text[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"jsonb", types.ExtensionTypes.JSON},
		{"uuid[]", arrow.ListOf(types.ExtensionTypes.UUID)},
		{"macaddr", types.ExtensionTypes.MAC},
		{"macaddr[]", arrow.ListOf(types.ExtensionTypes.MAC)},
		{"inet", types.ExtensionTypes.Inet},
		{"inet[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"timestamp without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"text[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"int[]", arrow.ListOf(arrow.PrimitiveTypes.Int32)},
		{"date", arrow.FixedWidthTypes.Date32},
		{"date[]", arrow.ListOf(arrow.FixedWidthTypes.Date32)},
		{"time without time zone", arrow.FixedWidthTypes.Time32s},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time32s)},
		{"time without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time32ms)},
		{"time without time zone", arrow.FixedWidthTypes.Time64us},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time without time zone", arrow.FixedWidthTypes.Time64ns},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64ns)},
		// special case for uint64
		{"numeric(20,0)", arrow.PrimitiveTypes.Uint64},
		{"numeric(38,0)", &arrow.Decimal128Type{Precision: 38, Scale: 0}},
		{"numeric(1,0)", &arrow.Decimal128Type{Precision: 1, Scale: 0}},
		{"numeric(38,15)", &arrow.Decimal128Type{Precision: 38, Scale: 15}},
		{"numeric(50,25)", &arrow.Decimal256Type{Precision: 50, Scale: 25}},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			assert.Equal(t, c.want, ArrowToPg10(c.dataType))
		})
	}
}
func TestArrowToCockroach(t *testing.T) {
	cases := []struct {
		want     string
		dataType arrow.DataType
	}{
		{"boolean", arrow.FixedWidthTypes.Boolean},
		{"int8", arrow.PrimitiveTypes.Int64},
		{"double precision", arrow.PrimitiveTypes.Float64},
		{"uuid", types.ExtensionTypes.UUID},
		{"bytea", arrow.BinaryTypes.Binary},
		{"text[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"jsonb", types.ExtensionTypes.JSON},
		{"uuid[]", arrow.ListOf(types.ExtensionTypes.UUID)},
		{"inet", types.ExtensionTypes.Inet},
		{"inet[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"timestamp without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"text[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"int8[]", arrow.ListOf(arrow.PrimitiveTypes.Int32)},
		{"date", arrow.FixedWidthTypes.Date32},
		{"date[]", arrow.ListOf(arrow.FixedWidthTypes.Date32)},
		{"time without time zone", arrow.FixedWidthTypes.Time32s},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time32s)},
		{"time without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time32ms)},
		{"time without time zone", arrow.FixedWidthTypes.Time64us},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time without time zone", arrow.FixedWidthTypes.Time64ns},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64ns)},
		// special case for uint64
		{"numeric(20,0)", arrow.PrimitiveTypes.Uint64},
		{"numeric(38,0)", &arrow.Decimal128Type{Precision: 38, Scale: 0}},
		{"numeric(1,0)", &arrow.Decimal128Type{Precision: 1, Scale: 0}},
		{"numeric(38,15)", &arrow.Decimal128Type{Precision: 38, Scale: 15}},
		{"numeric(50,25)", &arrow.Decimal256Type{Precision: 50, Scale: 25}},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			assert.Equal(t, c.want, ArrowToCockroach(c.dataType))
		})
	}
}
