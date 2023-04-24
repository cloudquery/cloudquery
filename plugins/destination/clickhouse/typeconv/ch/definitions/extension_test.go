package definitions

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_extensionType(t *testing.T) {
	type testCase struct {
		_type arrow.ExtensionType
		exp   string
	}

	for _, tc := range []testCase{
		{_type: new(types.UUIDType), exp: "UUID"},
		{_type: new(types.InetType), exp: "String"},
		{_type: new(types.MacType), exp: "String"},
		{_type: new(types.JSONType), exp: "String"},
	} {
		ensureDefinition(t, tc._type, tc.exp)
	}
}
