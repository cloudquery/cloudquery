package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
)

func Test_decimalType(t *testing.T) {
	type testCase struct {
		_type    string
		expected arrow.DataType
	}

	for _, tc := range []testCase{
		{_type: "Decimal(5, 2)", expected: &arrow.Decimal128Type{Precision: 5, Scale: 2}},
		{_type: "Decimal(19, 10)", expected: &arrow.Decimal128Type{Precision: 19, Scale: 10}},
		{_type: "Decimal(38, 35)", expected: &arrow.Decimal128Type{Precision: 38, Scale: 35}},
		{_type: "Decimal(42, 35)", expected: &arrow.Decimal256Type{Precision: 42, Scale: 35}},
	} {
		ensureField(t, tc._type, tc.expected)
	}
}
