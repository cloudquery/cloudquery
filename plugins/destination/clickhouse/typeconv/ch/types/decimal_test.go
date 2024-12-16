package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
)

func Test_decimalType(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: &arrow.Decimal128Type{Precision: 1}, expected: "Decimal(1,0)"},
		{dataType: &arrow.Decimal128Type{Precision: 12, Scale: 10}, expected: "Decimal(12,10)"},
		{dataType: &arrow.Decimal128Type{Precision: 19, Scale: 12}, expected: "Decimal(19,12)"},
		{dataType: &arrow.Decimal256Type{Precision: 39, Scale: 12}, expected: "Decimal(39,12)"},
	} {
		ensureDefinition(t, tc)
	}
}
