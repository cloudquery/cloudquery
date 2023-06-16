package types

import (
	"testing"

	"github.com/apache/arrow/go/v13/arrow"
)

func Test_decimalType(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: &arrow.Decimal128Type{Scale: 12}, expected: "Decimal(19,12)"},
		{dataType: &arrow.Decimal256Type{Scale: 12}, expected: "Decimal(39,12)"},
	} {
		ensureDefinition(t, tc)
	}
}
