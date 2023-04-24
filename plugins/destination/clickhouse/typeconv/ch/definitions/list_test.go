package definitions

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_listType(t *testing.T) {
	type testCase struct {
		_type listDataType
		exp   string
	}

	for _, tc := range []testCase{
		{_type: arrow.ListOf(new(arrow.StringType)), exp: "Array(Nullable(String))"},
		{_type: arrow.ListOfNonNullable(new(arrow.StringType)), exp: "Array(String)"},
		{_type: arrow.ListOf(new(types.UUIDType)), exp: "Array(Nullable(UUID))"},
		{_type: arrow.ListOfNonNullable(new(types.UUIDType)), exp: "Array(UUID)"},
		{_type: arrow.ListOfField(
			arrow.Field{
				Name:     "map",
				Type:     arrow.MapOf(new(arrow.StringType), new(arrow.Decimal128Type)),
				Nullable: true,
			},
		), exp: "Array(Nullable(String))"},
	} {
		ensureDefinition(t, tc._type, tc.exp)
	}
}
