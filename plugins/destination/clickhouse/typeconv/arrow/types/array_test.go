package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
)

func Test_arrayType(t *testing.T) {
	for _, tc := range []testCase{
		{_type: "Array(Nullable(String))", expected: arrow.ListOf(new(arrow.StringType))},
		{_type: "Array(String)", expected: arrow.ListOfNonNullable(new(arrow.StringType))},
		{_type: "Nullable(Array(String))", expected: arrow.ListOfNonNullable(new(arrow.StringType))},
		{_type: "Array(Tuple(`f1` Bool))", expected: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
		{_type: "Nested(`f1` Bool)", expected: arrow.ListOfNonNullable(arrow.StructOf(arrow.Field{Name: "f1", Type: new(arrow.BooleanType)}))},
	} {
		ensureField(t, tc)
	}
}
