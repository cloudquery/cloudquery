package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
)

func Test_arrayType(t *testing.T) {
	for _, tc := range []testCase{
		{columnType: "Array(Nullable(String))", expected: arrow.ListOf(new(arrow.StringType))},
		{columnType: "Array(String)", expected: arrow.ListOfNonNullable(new(arrow.StringType))},
		{columnType: "Array(Tuple(`f1` Bool))", expected: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
		{columnType: "Nested(`f1` Bool)", expected: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
	} {
		ensureField(t, tc)
	}
}
