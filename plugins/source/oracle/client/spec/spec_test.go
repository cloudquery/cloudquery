package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestSpec(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra field",
			Err:  true,
			Spec: `{"extra": 0}`,
		},
		{
			Name: "zero concurrency",
			Err:  true,
			Spec: `{"concurrency": 0}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"concurrency": 0}`,
		},
		{
			Name: "bad concurrency",
			Err:  true,
			Spec: `{"concurrency": "abc"}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency": 123}`,
		},
	})
}

func TestEnsureJSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(Spec))
	require.NoError(t, err)
	require.JSONEqf(t, string(data), JSONSchema, "new schema should be:\n%s\n", string(data))
}
