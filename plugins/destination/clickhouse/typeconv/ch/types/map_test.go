package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
)

func Test_mapType(t *testing.T) {
	for _, tc := range []testCase{
		{
			dataType: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)),
			expected: "Map(String, Nullable(Bool))",
		},
		{
			dataType: arrow.MapOf(new(arrow.StringType), new(arrow.Decimal128Type)),
			expected: "Map(String, Nullable(Decimal(19,0)))",
		},
		{
			dataType: arrow.MapOf(new(arrow.Float64Type), new(arrow.Decimal128Type)),
			expected: "String",
		},
	} {
		ensureDefinition(t, tc)
	}
}
