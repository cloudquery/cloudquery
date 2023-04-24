package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_listType(t *testing.T) {
	for _, tc := range []testCase{
		{_type: arrow.ListOf(new(arrow.StringType)), expected: "Array(Nullable(String))"},
		{_type: arrow.ListOfNonNullable(new(arrow.StringType)), expected: "Array(String)"},
		{_type: arrow.ListOf(new(types.UUIDType)), expected: "Array(Nullable(UUID))"},
		{_type: arrow.ListOfNonNullable(new(types.UUIDType)), expected: "Array(UUID)"},
		{_type: arrow.ListOfField(
			arrow.Field{
				Name:     "map",
				Type:     arrow.MapOf(new(arrow.StringType), new(arrow.Decimal128Type)),
				Nullable: true,
			},
		), expected: "Array(Nullable(String))"},
	} {
		ensureDefinition(t, tc)
	}
}
