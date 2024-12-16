package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func Test_listType(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: arrow.ListOf(new(arrow.StringType)), expected: "Array(Nullable(String))"},
		{dataType: arrow.ListOfNonNullable(new(arrow.StringType)), expected: "Array(String)"},
		{dataType: arrow.ListOf(new(types.UUIDType)), expected: "Array(Nullable(UUID))"},
		{dataType: arrow.ListOfNonNullable(new(types.UUIDType)), expected: "Array(UUID)"},
		{
			dataType: arrow.ListOf(arrow.MapOf(new(arrow.StringType), &arrow.Decimal128Type{Precision: 19})),
			expected: "Array(Map(String, Nullable(Decimal(19,0))))",
		},
	} {
		ensureDefinition(t, tc)
	}
}
