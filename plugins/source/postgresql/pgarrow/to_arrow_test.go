package pgarrow

import (
	"testing"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func TestPg10ToArrow(t *testing.T) {
	cases := []struct {
		pgType string
		want   arrow.DataType
	}{
		{"boolean", arrow.FixedWidthTypes.Boolean},
		{"bigint", arrow.PrimitiveTypes.Int64},
		{"double precision", arrow.PrimitiveTypes.Float64},
		{"uuid", types.ExtensionTypes.UUID},
		{"bytea", arrow.BinaryTypes.Binary},
		{"text[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"json", types.ExtensionTypes.JSON},
		{"jsonb", types.ExtensionTypes.JSON},
		{"uuid[]", arrow.ListOf(types.ExtensionTypes.UUID)},
		{"cidr", types.ExtensionTypes.Inet},
		{"cidr[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"macaddr", types.ExtensionTypes.MAC},
		{"macaddr8", types.ExtensionTypes.MAC},
		{"macaddr[]", arrow.ListOf(types.ExtensionTypes.MAC)},
		{"macaddr8[]", arrow.ListOf(types.ExtensionTypes.MAC)},
		{"inet", types.ExtensionTypes.Inet},
		{"inet[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"timestamp", arrow.FixedWidthTypes.Timestamp_us},
		{"text[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"varchar(50)", arrow.BinaryTypes.String},
		{"varchar(50)[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"character(10)", arrow.BinaryTypes.String},
		{"integer[]", arrow.ListOf(arrow.PrimitiveTypes.Int32)},
		{"timestamp", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamptz", arrow.FixedWidthTypes.Timestamp_us},
		{"TIMESTAMPTZ USING timestamp(0)", arrow.FixedWidthTypes.Timestamp_s},
		{"TIMESTAMPTZ USING timestamp(3)", arrow.FixedWidthTypes.Timestamp_ms},
		{"TIMESTAMPTZ USING timestamp(6)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0)", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0) without time zone", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0) with time zone", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"date", arrow.FixedWidthTypes.Date32},
		{"date[]", arrow.ListOf(arrow.FixedWidthTypes.Date32)},
		{"time", arrow.FixedWidthTypes.Time64us},
		{"time[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time with time zone", arrow.FixedWidthTypes.Time64us},
		{"time with time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time without time zone", arrow.FixedWidthTypes.Time64us},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time(0)", arrow.FixedWidthTypes.Time32s},
		{"time(1)", arrow.FixedWidthTypes.Time32ms},
		{"time(2)", arrow.FixedWidthTypes.Time32ms},
		{"time(3)", arrow.FixedWidthTypes.Time32ms},
		{"time(4)", arrow.FixedWidthTypes.Time64us},
		{"time(5)", arrow.FixedWidthTypes.Time64us},
		{"time(6)", arrow.FixedWidthTypes.Time64us},
		{"time(0) without time zone", arrow.FixedWidthTypes.Time32s},
		{"time(1) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(2) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(3) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(4) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(5) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(6) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(0) with time zone", arrow.FixedWidthTypes.Time32s},
		{"time(1) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(2) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(3) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(4) with time zone", arrow.FixedWidthTypes.Time64us},
		{"time(5) with time zone", arrow.FixedWidthTypes.Time64us},
		{"time(6) with time zone", arrow.FixedWidthTypes.Time64us},
		// special case for uint64
		{"numeric(20,0)", arrow.PrimitiveTypes.Uint64},

		// types that are converted to string for now - more specific support for these types
		// may be added in the future
		{"numeric", arrow.BinaryTypes.String},
		{"numeric (1, 0)", arrow.BinaryTypes.String},
		{"numeric (1000, 1000)", arrow.BinaryTypes.String},
		{"interval", arrow.BinaryTypes.String},
		{"interval YEAR", arrow.BinaryTypes.String},
		{"interval YEAR TO MONTH", arrow.BinaryTypes.String},
		{"interval SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND (0)", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND (3)", arrow.BinaryTypes.String},
		{"interval SECOND (6)", arrow.BinaryTypes.String},
		{"interval HOUR TO MINUTE", arrow.BinaryTypes.String},
		{"interval YEAR", arrow.BinaryTypes.String},
		{"interval MONTH", arrow.BinaryTypes.String},
		{"interval DAY", arrow.BinaryTypes.String},
		{"interval HOUR", arrow.BinaryTypes.String},
		{"interval MINUTE", arrow.BinaryTypes.String},
		{"interval SECOND", arrow.BinaryTypes.String},
		{"interval YEAR TO MONTH", arrow.BinaryTypes.String},
		{"interval DAY TO HOUR", arrow.BinaryTypes.String},
		{"interval DAY TO MINUTE", arrow.BinaryTypes.String},
		{"interval DAY TO SECOND", arrow.BinaryTypes.String},
		{"interval HOUR TO MINUTE", arrow.BinaryTypes.String},
		{"interval HOUR TO SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND", arrow.BinaryTypes.String},
		{"money", arrow.BinaryTypes.String},
		{"box", arrow.BinaryTypes.String},
		{"bit", arrow.BinaryTypes.String},
		{"bit varying(10)", arrow.BinaryTypes.String},
		{"circle", arrow.BinaryTypes.String},
		{"line", arrow.BinaryTypes.String},
		{"point", arrow.BinaryTypes.String},
		{"path", arrow.BinaryTypes.String},
		{"polygon", arrow.BinaryTypes.String},
	}

	for _, c := range cases {
		t.Run(c.pgType, func(t *testing.T) {
			got := Pg10ToArrow(c.pgType)
			if !arrow.TypeEqual(got, c.want) {
				t.Errorf("Pg10ToArrow(%q) = %v, want %v", c.pgType, got, c.want)
			}
		})
	}
}
func TestCockroachToArrow(t *testing.T) {
	cases := []struct {
		pgType string
		want   arrow.DataType
	}{
		{"boolean", arrow.FixedWidthTypes.Boolean},
		{"bigint", arrow.PrimitiveTypes.Int64},
		{"double precision", arrow.PrimitiveTypes.Float64},
		{"uuid", types.ExtensionTypes.UUID},
		{"bytea", arrow.BinaryTypes.Binary},
		{"text[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"json", types.ExtensionTypes.JSON},
		{"jsonb", types.ExtensionTypes.JSON},
		{"uuid[]", arrow.ListOf(types.ExtensionTypes.UUID)},
		{"cidr", types.ExtensionTypes.Inet},
		{"cidr[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"inet", types.ExtensionTypes.Inet},
		{"inet[]", arrow.ListOf(types.ExtensionTypes.Inet)},
		{"timestamp", arrow.FixedWidthTypes.Timestamp_us},
		{"text[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"varchar(50)", arrow.BinaryTypes.String},
		{"varchar(50)[][]", arrow.ListOf(arrow.ListOf(arrow.BinaryTypes.String))},
		{"character(10)", arrow.BinaryTypes.String},
		{"integer[]", arrow.ListOf(arrow.PrimitiveTypes.Int64)},
		{"timestamp", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamptz", arrow.FixedWidthTypes.Timestamp_us},
		{"TIMESTAMPTZ USING timestamp(0)", arrow.FixedWidthTypes.Timestamp_s},
		{"TIMESTAMPTZ USING timestamp(3)", arrow.FixedWidthTypes.Timestamp_ms},
		{"TIMESTAMPTZ USING timestamp(6)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0)", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3)", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6)", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0) without time zone", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3) without time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6) without time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(0) with time zone", arrow.FixedWidthTypes.Timestamp_s},
		{"timestamp(1) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(2) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(3) with time zone", arrow.FixedWidthTypes.Timestamp_ms},
		{"timestamp(4) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(5) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"timestamp(6) with time zone", arrow.FixedWidthTypes.Timestamp_us},
		{"date", arrow.FixedWidthTypes.Date32},
		{"date[]", arrow.ListOf(arrow.FixedWidthTypes.Date32)},
		{"time", arrow.FixedWidthTypes.Time64us},
		{"time[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time with time zone", arrow.FixedWidthTypes.Time64us},
		{"time with time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time without time zone", arrow.FixedWidthTypes.Time64us},
		{"time without time zone[]", arrow.ListOf(arrow.FixedWidthTypes.Time64us)},
		{"time(0)", arrow.FixedWidthTypes.Time32s},
		{"time(1)", arrow.FixedWidthTypes.Time32ms},
		{"time(2)", arrow.FixedWidthTypes.Time32ms},
		{"time(3)", arrow.FixedWidthTypes.Time32ms},
		{"time(4)", arrow.FixedWidthTypes.Time64us},
		{"time(5)", arrow.FixedWidthTypes.Time64us},
		{"time(6)", arrow.FixedWidthTypes.Time64us},
		{"time(0) without time zone", arrow.FixedWidthTypes.Time32s},
		{"time(1) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(2) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(3) without time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(4) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(5) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(6) without time zone", arrow.FixedWidthTypes.Time64us},
		{"time(0) with time zone", arrow.FixedWidthTypes.Time32s},
		{"time(1) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(2) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(3) with time zone", arrow.FixedWidthTypes.Time32ms},
		{"time(4) with time zone", arrow.FixedWidthTypes.Time64us},
		{"time(5) with time zone", arrow.FixedWidthTypes.Time64us},
		{"time(6) with time zone", arrow.FixedWidthTypes.Time64us},
		// special case for uint64
		{"numeric(20,0)", arrow.PrimitiveTypes.Uint64},

		// types that are converted to string for now - more specific support for these types
		// may be added in the future
		{"macaddr", arrow.BinaryTypes.String},
		{"macaddr8", arrow.BinaryTypes.String},
		{"macaddr[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"macaddr8[]", arrow.ListOf(arrow.BinaryTypes.String)},
		{"numeric", arrow.BinaryTypes.String},
		{"numeric (1, 0)", arrow.BinaryTypes.String},
		{"numeric (1000, 1000)", arrow.BinaryTypes.String},
		{"interval", arrow.BinaryTypes.String},
		{"interval YEAR", arrow.BinaryTypes.String},
		{"interval YEAR TO MONTH", arrow.BinaryTypes.String},
		{"interval SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND (0)", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND (3)", arrow.BinaryTypes.String},
		{"interval SECOND (6)", arrow.BinaryTypes.String},
		{"interval HOUR TO MINUTE", arrow.BinaryTypes.String},
		{"interval YEAR", arrow.BinaryTypes.String},
		{"interval MONTH", arrow.BinaryTypes.String},
		{"interval DAY", arrow.BinaryTypes.String},
		{"interval HOUR", arrow.BinaryTypes.String},
		{"interval MINUTE", arrow.BinaryTypes.String},
		{"interval SECOND", arrow.BinaryTypes.String},
		{"interval YEAR TO MONTH", arrow.BinaryTypes.String},
		{"interval DAY TO HOUR", arrow.BinaryTypes.String},
		{"interval DAY TO MINUTE", arrow.BinaryTypes.String},
		{"interval DAY TO SECOND", arrow.BinaryTypes.String},
		{"interval HOUR TO MINUTE", arrow.BinaryTypes.String},
		{"interval HOUR TO SECOND", arrow.BinaryTypes.String},
		{"interval MINUTE TO SECOND", arrow.BinaryTypes.String},
		{"money", arrow.BinaryTypes.String},
		{"box", arrow.BinaryTypes.String},
		{"bit", arrow.BinaryTypes.String},
		{"bit varying(10)", arrow.BinaryTypes.String},
		{"circle", arrow.BinaryTypes.String},
		{"line", arrow.BinaryTypes.String},
		{"point", arrow.BinaryTypes.String},
		{"path", arrow.BinaryTypes.String},
		{"polygon", arrow.BinaryTypes.String},
	}

	for _, c := range cases {
		t.Run(c.pgType, func(t *testing.T) {
			got := CockroachToArrow(c.pgType)
			if !arrow.TypeEqual(got, c.want) {
				t.Errorf("CockroachToArrow(%q) = %v, want %v", c.pgType, got, c.want)
			}
		})
	}
}
