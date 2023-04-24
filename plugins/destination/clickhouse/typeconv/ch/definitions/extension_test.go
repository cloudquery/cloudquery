package definitions

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_extensionType(t *testing.T) {
	type testCase struct {
		_type    arrow.ExtensionType
		expected string
	}

	for _, tc := range []testCase{
		{_type: types.NewUUIDType(), expected: "UUID"},
		{_type: types.NewInetType(), expected: "String"},
		{_type: types.NewMacType(), expected: "String"},
		{_type: types.NewJSONType(), expected: "String"},
	} {
		ensureDefinition(t, tc._type, tc.expected)
	}
}
