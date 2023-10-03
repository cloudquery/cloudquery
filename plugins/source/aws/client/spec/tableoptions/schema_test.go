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
			require.NoErrorf(t, json.Unmarshal([]byte(tc.spec), &v), "failed input:\n%s\n", tc.spec)
			err := validator.Validate(v)
			if tc.err {
				require.Errorf(t, err, "failed input:\n%s\n", tc.spec)
			} else {
				require.NoErrorf(t, err, "failed input:\n%s\n", tc.spec)
			}
		})
	}
}

func jsonWithRemovedKeys(t *testing.T, val any, keys ...string) string {
	data, err := json.Marshal(val)
	require.NoError(t, err)

	var m any
	require.NoError(t, json.Unmarshal(data, &m))

	switch m := m.(type) {
	case map[string]any:
		for _, k := range keys {
			delete(m, k)
		}
	default:
		t.Fatalf("failed to remove JSON keys from value of type %T", m)
	}

	data, err = json.MarshalIndent(m, "", "  ")
	require.NoError(t, err)
	return string(data)
}
