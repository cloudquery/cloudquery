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
			Err:  true,
		},
		{
			Name: "missing token",
			Spec: `{"access_token":""}`,
			Err:  true,
		},
		{
			Name: "valid",
			Spec: `{"access_token":"token"}`,
			Err:  false,
		},
	})
}
