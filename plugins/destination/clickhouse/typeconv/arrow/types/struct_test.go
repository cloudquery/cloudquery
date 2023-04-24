package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_structType(t *testing.T) {
	for _, tc := range []testCase{
		{
			_type:    "Tuple(`f1` Bool)",
			expected: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}),
		},
		{
			_type:    "Tuple(`f1` Nullable(Bool))",
			expected: arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType), Nullable: true}),
		},
		{
			_type: "Tuple(`bool_list` Array(Bool), `bool_list_nullable` Array(Nullable(Bool)))",
			expected: arrow.StructOf(
				arrow.Field{
					Name: "bool_list",
					Type: arrow.ListOfNonNullable(new(arrow.BooleanType)),
				},
				arrow.Field{
					Name:     "bool_list_nullable",
					Type:     arrow.ListOf(new(arrow.BooleanType)),
					Nullable: true,
				},
			),
		},
		{
			_type: "Tuple(`uuid_list` Array(Nullable(UUID)), `struct` Tuple(`bool_list` Array(Bool), `bool_list_nullable` Array(Nullable(Bool))))",
			expected: arrow.StructOf(
				arrow.Field{
					Name:     "uuid_list",
					Type:     arrow.ListOf(types.NewUUIDType()),
					Nullable: true,
				},
				arrow.Field{
					Name: "struct",
					Type: arrow.StructOf(
						arrow.Field{
							Name: "bool_list",
							Type: arrow.ListOfNonNullable(new(arrow.BooleanType)),
						},
						arrow.Field{
							Name:     "bool_list_nullable",
							Type:     arrow.ListOf(new(arrow.BooleanType)),
							Nullable: true,
						},
					),
				},
			),
		},
	} {
		ensureField(t, tc)
	}
}
