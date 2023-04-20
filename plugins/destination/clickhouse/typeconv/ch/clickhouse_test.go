package ch

import (
	"strings"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/stretchr/testify/require"
)

func TestDefinitions(t *testing.T) {
	type testCase struct {
		data arrow.DataType
		exp  string
	}
	replacer := strings.NewReplacer(
		"@", "_",
		"(", "_",
		"<", "_",
		"[", "_",
		")", "_",
		">", "_",
		"]", "_",
		" ", "_",
		":", "_",
		",", "_",
	)

	for _, tc := range []testCase{
		{data: new(arrow.BooleanType), exp: "Bool"},

		{data: new(arrow.Int8Type), exp: "Int8"},
		{data: new(arrow.Int16Type), exp: "Int16"},
		{data: new(arrow.Int32Type), exp: "Int32"},
		{data: new(arrow.Int64Type), exp: "Int64"},
		{data: new(arrow.Uint8Type), exp: "UInt8"},
		{data: new(arrow.Uint16Type), exp: "UInt16"},
		{data: new(arrow.Uint32Type), exp: "UInt32"},
		{data: new(arrow.Uint64Type), exp: "UInt64"},
		{data: new(arrow.Float16Type), exp: "Float32"},
		{data: new(arrow.Float32Type), exp: "Float32"},
		{data: new(arrow.Float64Type), exp: "Float64"},
		{data: &arrow.FixedSizeBinaryType{ByteWidth: 125}, exp: "FixedString(125)"},
		{data: new(arrow.Date32Type), exp: "Date32"},
		{data: new(arrow.Date64Type), exp: "DateTime64(3)"},
		{data: &arrow.TimestampType{Unit: arrow.Second}, exp: "DateTime64(0)"},
		{data: &arrow.TimestampType{Unit: arrow.Millisecond}, exp: "DateTime64(3)"},
		{data: &arrow.TimestampType{Unit: arrow.Microsecond}, exp: "DateTime64(6)"},
		{data: &arrow.TimestampType{Unit: arrow.Nanosecond}, exp: "DateTime64(9)"},
		{data: &arrow.Decimal128Type{Scale: 35}, exp: "Decimal128(35)"},
		{data: &arrow.Decimal256Type{Scale: 35}, exp: "Decimal256(35)"},
		{data: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}), exp: "Tuple(`f1` Bool)"},
		{data: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType), Nullable: true}), exp: "Tuple(`f1` Nullable(Bool))"},
		{data: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)), exp: "Map(String, Nullable(Bool))"},
		{data: arrow.ListOf(new(arrow.StringType)), exp: "Array(Nullable(String))"},
		{data: arrow.ListOfNonNullable(new(arrow.StringType)), exp: "Array(String)"},
		{data: new(types.UUIDType), exp: "UUID"},
		{data: new(types.InetType), exp: "String"},
		{data: new(types.MacType), exp: "String"},
		{data: new(types.JSONType), exp: "String"},
		{data: arrow.MapOf(
			new(arrow.StringType),
			arrow.StructOf(
				arrow.Field{Name: "f1_bool", Type: new(arrow.BooleanType)},
				arrow.Field{Name: "f2_map", Type: arrow.MapOf(
					new(arrow.StringType),
					arrow.StructOf(
						arrow.Field{Name: "f1_uint8_nullable", Type: new(arrow.Uint8Type), Nullable: true},
						arrow.Field{Name: "f2_uuid", Type: new(types.UUIDType)},
					),
				),
				},
			),
		), exp: "Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID))))))"},
		{data: arrow.ListOfField(
			arrow.Field{
				Name: "map",
				Type: arrow.MapOf(
					new(arrow.StringType),
					arrow.StructOf(
						arrow.Field{Name: "f1_bool", Type: new(arrow.BooleanType)},
						arrow.Field{Name: "f2_map", Type: arrow.MapOf(
							new(arrow.StringType),
							arrow.StructOf(
								arrow.Field{Name: "f1_uint8_nullable", Type: new(arrow.Uint8Type), Nullable: true},
								arrow.Field{Name: "f2_uuid", Type: new(types.UUIDType)},
							),
						),
						},
					),
				),
			},
		), exp: "Array(Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID)))))))"},
	} {
		// non-nullable
		field := arrow.Field{
			Name: replacer.Replace(tc.data.String()),
			Type: tc.data,
		}
		defs, err := Definitions(field)
		require.NoError(t, err)
		require.Len(t, defs, 1)
		require.Equal(t, util.SanitizeID(field.Name)+" "+tc.exp, defs[0])

		if field.Type.ID() == arrow.LIST {
			// arrays cannot be marked nullable in ClickHouse
			continue
		}

		// nullable
		field.Nullable = true
		defs, err = Definitions(field)
		require.NoError(t, err)
		require.Len(t, defs, 1)
		require.Equal(t, util.SanitizeID(field.Name)+" Nullable("+tc.exp+")", defs[0])
	}
}
