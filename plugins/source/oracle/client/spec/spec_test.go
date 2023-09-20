package spec

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/stretchr/testify/require"
)

func getValidator(t *testing.T) *jsonschema.Schema {
	t.Helper()

	c := jsonschema.NewCompiler()
	c.Draft = jsonschema.Draft2020

	require.NoError(t, c.AddResource("schema.json", strings.NewReader(JSONSchema)))

	validator, err := c.Compile("schema.json")
	require.NoError(t, err)

	return validator
}

func TestSpec(t *testing.T) {
	validator := getValidator(t)

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
