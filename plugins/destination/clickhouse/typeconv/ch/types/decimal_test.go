package types

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
)

func Test_decimalType(t *testing.T) {
	for _, tc := range []testCase{
		{_type: &arrow.Decimal128Type{Scale: 12}, expected: "Decimal(19,12)"},
		{_type: &arrow.Decimal256Type{Scale: 12}, expected: "Decimal(39,12)"},
	} {
		ensureDefinition(t, tc)
	}
}
