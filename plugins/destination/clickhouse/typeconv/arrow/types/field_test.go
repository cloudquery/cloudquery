package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/stretchr/testify/require"
)

func TestField(t *testing.T) {
	type testCase struct {
		_type    string
		expected arrow.DataType
	}

	for _, tc := range []testCase{
		{_type: "Bool", expected: new(arrow.BooleanType)},
		{_type: "Int8", expected: new(arrow.Int8Type)},
		{_type: "Int16", expected: new(arrow.Int16Type)},
		{_type: "Int32", expected: new(arrow.Int32Type)},
		{_type: "Int64", expected: new(arrow.Int64Type)},
		{_type: "UInt8", expected: new(arrow.Uint8Type)},
		{_type: "UInt16", expected: new(arrow.Uint16Type)},
		{_type: "UInt32", expected: new(arrow.Uint32Type)},
		{_type: "UInt64", expected: new(arrow.Uint64Type)},
		{_type: "Float32", expected: new(arrow.Float32Type)},
		{_type: "Float64", expected: new(arrow.Float64Type)},
		{_type: "FixedString(125)", expected: &arrow.FixedSizeBinaryType{ByteWidth: 125}},
		{_type: "Date", expected: new(arrow.StringType)},
		{_type: "Date32", expected: new(arrow.Date32Type)},
		{_type: "UUID", expected: new(types.UUIDType)},
		{_type: "Map(String, UInt64)", expected: arrow.MapOf(new(arrow.StringType), new(arrow.Uint64Type))},
		{_type: "Map(String, Nullable(Bool))", expected: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType))},
		{
			_type: "Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID))))))",
			expected: arrow.MapOf(
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
			_type: "Array(Map(String, Nullable(Tuple(`f1_bool` Bool, `f2_map` Map(String, Nullable(Tuple(`f1_uint8_nullable` Nullable(UInt8), `f2_uuid` UUID)))))))",
			expected: arrow.ListOfField(
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
		ensureField(t, tc._type, tc.expected)
	}
}

func ensureField(t *testing.T, _type string, expected arrow.DataType) {
	// simple
	field, err := Field("field", _type)
	require.NoError(t, err)
	require.False(t, field.Nullable)
	require.Truef(t, arrow.TypeEqual(expected, field.Type), "expected type:\n%s\nactual:\n%s", expected.String(), field.Type.String())

	// nullable
	field, err = Field("field", "Nullable("+_type+")")
	require.NoError(t, err)
	require.True(t, field.Nullable)
	require.True(t, arrow.TypeEqual(expected, field.Type))
}
