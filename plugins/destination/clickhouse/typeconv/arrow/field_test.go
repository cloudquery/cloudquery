package arrow

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/stretchr/testify/require"
)

func TestField(t *testing.T) {
	type testCase struct {
		typ string
		exp arrow.DataType
	}

	tz, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)

	for _, tc := range []testCase{
		{typ: "Bool", exp: new(arrow.BooleanType)},
		{typ: "Int8", exp: new(arrow.Int8Type)},
		{typ: "Int16", exp: new(arrow.Int16Type)},
		{typ: "Int32", exp: new(arrow.Int32Type)},
		{typ: "Int64", exp: new(arrow.Int64Type)},
		{typ: "UInt8", exp: new(arrow.Uint8Type)},
		{typ: "UInt16", exp: new(arrow.Uint16Type)},
		{typ: "UInt32", exp: new(arrow.Uint32Type)},
		{typ: "UInt64", exp: new(arrow.Uint64Type)},
		{typ: "Float32", exp: new(arrow.Float32Type)},
		{typ: "Float64", exp: new(arrow.Float64Type)},
		{typ: "FixedString(125)", exp: &arrow.FixedSizeBinaryType{ByteWidth: 125}},
		{typ: "Date", exp: new(arrow.StringType)},
		{typ: "Date32", exp: new(arrow.Date32Type)},
		{typ: "DateTime64(0)", exp: &arrow.TimestampType{Unit: arrow.Second}},
		{typ: "DateTime64(3)", exp: &arrow.TimestampType{Unit: arrow.Millisecond}},
		{typ: "DateTime64(6)", exp: &arrow.TimestampType{Unit: arrow.Microsecond}},
		{typ: "DateTime64(9)", exp: &arrow.TimestampType{Unit: arrow.Nanosecond}},
		{typ: "DateTime64(9, 'America/New_York')", exp: &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: tz.String()}},
		{typ: "Decimal(38, 35)", exp: &arrow.Decimal128Type{Precision: 38, Scale: 35}},
		{typ: "Decimal(42, 35)", exp: &arrow.Decimal256Type{Precision: 42, Scale: 35}},
		{typ: "UUID", exp: new(types.UUIDType)},
		{typ: "Tuple(`f1` Bool)", exp: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)})},
		{typ: "Tuple(`f1` Nullable(Bool))", exp: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType), Nullable: true})},
		{typ: "Map(String, UInt64)", exp: arrow.MapOf(new(arrow.StringType), new(arrow.Uint64Type))},
		{typ: "Map(String, Nullable(Bool))", exp: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType))},
		{typ: "Array(Nullable(String))", exp: arrow.ListOf(new(arrow.StringType))},
		{typ: "Array(String)", exp: arrow.ListOfNonNullable(new(arrow.StringType))},
		{typ: "Array(Tuple(`f1` Bool))", exp: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
		{typ: "Nested(`f1` Bool)", exp: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
		{
			typ: "Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID))))))",
			exp: arrow.MapOf(
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
		{
			typ: "Array(Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID)))))))",
			exp: arrow.ListOfField(
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
			),
		},
	} {
		// simple
		field, err := Field("field", tc.typ)
		require.NoError(t, err)

		switch field.Type.ID() {
		case arrow.LIST:

		}
		if field.Type.ID() == arrow.LIST {
			require.Equal(t, field.Type.(*arrow.ListType).ElemField().Nullable, field.Nullable)
		} else {
			require.False(t, field.Nullable)
		}

		require.Truef(t, arrow.TypeEqual(tc.exp, field.Type), "expected type: %s, actual: %s", tc.exp.String(), field.Type.String())

		// nullable
		field, err = Field("field", "Nullable("+tc.typ+")")
		require.NoError(t, err)
		require.True(t, field.Nullable)
		require.True(t, arrow.TypeEqual(tc.exp, field.Type))
	}
}
