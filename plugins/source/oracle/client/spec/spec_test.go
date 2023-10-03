package spec

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

func TestSpec(t *testing.T) {
	validator, err := plugin.JSONSchemaValidator(JSONSchema)
	require.NoError(t, err)

	type testCase struct {
		name string
		spec string
		err  bool
	}

	for _, tc := range []testCase{
		{
			name: "empty",
			spec: `{}`,
		},
		{
			name: "proper",
			spec: `{"concurrency": 2}`,
		},
		{
			name: "bad",
			spec: `{"concurrency": 0}`,
			err:  true,
		},
	} {
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

func TestEnsureJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(Spec))
	require.NoError(t, err)
	require.JSONEqf(t, string(data), JSONSchema, "new schema should be:\n%s\n", string(data))
}
