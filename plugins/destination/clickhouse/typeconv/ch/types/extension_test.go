package types

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_extensionType(t *testing.T) {
	for _, tc := range []testCase{
		{_type: types.NewUUIDType(), expected: "UUID"},
		{_type: types.NewInetType(), expected: "String"},
		{_type: types.NewMacType(), expected: "String"},
		{_type: types.NewJSONType(), expected: "String"},
	} {
		ensureDefinition(t, tc)
	}
}
