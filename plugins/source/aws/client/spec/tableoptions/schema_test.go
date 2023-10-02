package tableoptions

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

type jsonSchemaTestCase struct {
	name string
	spec string
	err  bool
}

func testJSONSchema(t *testing.T, cases []jsonSchemaTestCase) {
	validator, err := plugin.JSONSchemaValidator(JSONSchema)
	require.NoError(t, err)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var v any
			require.NoError(t, json.Unmarshal([]byte(tc.spec), &v))
			err := validator.Validate(v)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
