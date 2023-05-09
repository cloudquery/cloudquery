package types

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/types"
)

func Test_extensionType(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: types.NewUUIDType(), expected: "UUID"},
		{dataType: types.NewInetType(), expected: "String"},
		{dataType: types.NewMacType(), expected: "String"},
		{dataType: types.NewJSONType(), expected: "String"},
	} {
		ensureDefinition(t, tc)
	}
}
