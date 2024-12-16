package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func Test_structType(t *testing.T) {
	for _, tc := range []testCase{
		{
			dataType: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}),
			expected: "Tuple(`f1` Bool)",
		},
		{
			dataType: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType), Nullable: true}),
			expected: "Tuple(`f1` Nullable(Bool))",
		},
		{
			dataType: arrow.StructOf(
				arrow.Field{Name: "bool_list", Type: arrow.ListOfNonNullable(new(arrow.BooleanType))},
				arrow.Field{Name: "bool_list_nullable", Type: arrow.ListOf(new(arrow.BooleanType))},
			),
			expected: "Tuple(`bool_list` Array(Bool), `bool_list_nullable` Array(Nullable(Bool)))",
		},
		{
			dataType: arrow.StructOf(
				arrow.Field{Name: "uuid_list", Type: arrow.ListOf(types.NewUUIDType())},
				arrow.Field{
					Name: "struct",
					Type: arrow.StructOf(
						arrow.Field{Name: "bool_list", Type: arrow.ListOfNonNullable(new(arrow.BooleanType))},
						arrow.Field{Name: "bool_list_nullable", Type: arrow.ListOf(new(arrow.BooleanType))},
					),
				},
			),
			expected: "Tuple(`uuid_list` Array(Nullable(UUID)), `struct` Tuple(`bool_list` Array(Bool), `bool_list_nullable` Array(Nullable(Bool))))",
		},
	} {
		ensureDefinition(t, tc)
	}
}
