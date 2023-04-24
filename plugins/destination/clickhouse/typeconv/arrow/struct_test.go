package arrow

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_structType(t *testing.T) {
	type testCase struct {
		_type    string
		expected arrow.DataType
	}

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
					Type: arrow.ListOfField(arrow.Field{Name: "bool_list", Type: new(arrow.BooleanType)}),
				},
				arrow.Field{
					Name: "bool_list_nullable",
					Type: arrow.ListOfField(
						arrow.Field{
							Name:     "bool_list_nullable",
							Type:     new(arrow.BooleanType),
							Nullable: true,
						},
					),
					Nullable: true,
				},
			),
		},
		{
			_type: "Tuple(`uuid_list` Array(Nullable(UUID)), `struct` Tuple(`bool_list` Array(Bool), `bool_list_nullable` Array(Nullable(Bool))))",
			expected: arrow.StructOf(
				arrow.Field{
					Name: "uuid_list",
					Type: arrow.ListOfField(
						arrow.Field{
							Name:     "uuid_list",
							Type:     types.NewUUIDType(),
							Nullable: true,
						},
					),
					Nullable: true,
				},
				arrow.Field{
					Name: "struct",
					Type: arrow.StructOf(
						arrow.Field{
							Name: "bool_list",
							Type: arrow.ListOfField(arrow.Field{Name: "bool_list", Type: new(arrow.BooleanType)}),
						},
						arrow.Field{
							Name: "bool_list_nullable",
							Type: arrow.ListOfField(
								arrow.Field{
									Name:     "bool_list_nullable",
									Type:     new(arrow.BooleanType),
									Nullable: true,
								},
							),
							Nullable: true,
						},
					),
				},
			),
		},
	} {
		ensureField(t, tc._type, tc.expected)
	}
}
