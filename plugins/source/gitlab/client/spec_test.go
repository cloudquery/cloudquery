package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "valid",
			Spec: `{"access_token":"token"}`,
			Err:  false,
		},
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
			Name: "base url",
			Spec: `{"access_token":"token","base_url":"https://gitlab.com"}`,
			Err:  false,
		},
		{
			Name: "concurrency",
			Spec: `{"access_token":"token","concurrency":100}`,
			Err:  false,
		},
		{
			Name: "invalid field",
			Spec: `{"access_token":"token", "unknown_field":"value"}`,
			Err:  true,
		},
	})
}
