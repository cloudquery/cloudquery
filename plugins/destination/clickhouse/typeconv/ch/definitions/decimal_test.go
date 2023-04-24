package definitions

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
)

func Test_decimalType(t *testing.T) {
	type testCase struct {
		_type arrow.DecimalType
		exp   string
	}

	for _, tc := range []testCase{
		{_type: &arrow.Decimal128Type{Scale: 128}, exp: "String"},
		{_type: &arrow.Decimal128Type{Scale: 12}, exp: "Decimal(19,12)"},
		{_type: &arrow.Decimal256Type{Scale: 12}, exp: "Decimal(39,12)"},
	} {
		ensureDefinition(t, tc._type, tc.exp)
	}
}
