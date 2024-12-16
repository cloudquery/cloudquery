package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
)

func Test_decimalType(t *testing.T) {
	for _, tc := range []testCase{
		{columnType: "Decimal(5, 2)", expected: &arrow.Decimal128Type{Precision: 5, Scale: 2}},
		{columnType: "Decimal(19, 10)", expected: &arrow.Decimal128Type{Precision: 19, Scale: 10}},
		{columnType: "Decimal(38, 35)", expected: &arrow.Decimal128Type{Precision: 38, Scale: 35}},
		{columnType: "Decimal(42, 35)", expected: &arrow.Decimal256Type{Precision: 42, Scale: 35}},
	} {
		ensureField(t, tc)
	}
}
