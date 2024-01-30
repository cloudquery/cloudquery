package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
		},
		{
			Name: "spec with err_on_write",
			Spec: `{"error_on_write": true}`,
		},
		{
			Name: "spec with err_on_write false",
			Spec: `{"error_on_write": false}`,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"unknown": "test"}`,
			Err:  true,
		},
	})
}
