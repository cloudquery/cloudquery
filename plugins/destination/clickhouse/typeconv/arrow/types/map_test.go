package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func Test_mapType(t *testing.T) {
	for _, tc := range []testCase{
		{
			columnType: "Map(String, Nullable(Bool))",
			expected:   arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)),
		},
		{
			columnType: "Map(String, Nullable(Decimal(19,0)))",
			expected:   arrow.MapOf(new(arrow.StringType), &arrow.Decimal128Type{Precision: 19}),
		},
		{
			columnType: "Map(String, Nullable(Decimal(22,0)))",
			expected:   arrow.MapOf(new(arrow.StringType), &arrow.Decimal128Type{Precision: 22}),
		},
		{
			columnType: "Map(String, Nullable(Decimal(42,0)))",
			expected:   arrow.MapOf(new(arrow.StringType), &arrow.Decimal256Type{Precision: 42}),
		},
		{
			columnType: "Map(String, Nullable(Bool))",
			expected:   arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)),
		},
		{
			columnType: "Map(Float64, Nullable(Bool))", // although it's impossible in CH
			expected:   arrow.MapOf(new(arrow.Float64Type), new(arrow.BooleanType)),
		},
		{
			columnType: "Map(String, Nullable(Tuple(`f` Bool, `f_nullable` Nullable(Bool))))", // although Nullable(Tuple) is impossible in CH
			expected: arrow.MapOf(
				new(arrow.StringType),
				arrow.StructOf(
					arrow.Field{Name: "f", Type: new(arrow.BooleanType)},
					arrow.Field{Name: "f_nullable", Type: new(arrow.BooleanType), Nullable: true},
				),
			),
		},
		{
			columnType: "Map(String, Nullable(Tuple(`bool_n` Nullable(Bool), `bool` Bool, `list` Array(Nullable(UUID)), `map` Map(Int32, Nullable(Float64)), `map_n` Map(Int32, Nullable(Float64)), `map_uuid` Map(UUID, Nullable(UUID)), `mapped_to_string` String)))",
			expected: arrow.MapOf(
				new(arrow.StringType),
				arrow.StructOf(
					arrow.Field{Name: "bool_n", Type: new(arrow.BooleanType), Nullable: true},
					arrow.Field{Name: "bool", Type: new(arrow.BooleanType)},
					arrow.Field{Name: "list", Type: arrow.ListOf(types.NewUUIDType())},
					arrow.Field{Name: "map", Type: arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type))},
					arrow.Field{Name: "map_n", Type: arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type))},
					arrow.Field{Name: "map_uuid", Type: arrow.MapOf(types.NewUUIDType(), types.NewUUIDType())},
					arrow.Field{Name: "mapped_to_string", Type: new(arrow.StringType)},
				),
			),
		},
	} {
		ensureField(t, tc)
	}
}
