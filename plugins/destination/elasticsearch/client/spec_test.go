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
			Name: "spec with str addresses",
			Spec: `{"addresses": "address"}`,
			Err:  true,
		},
		{
			Name: "spec with valid addresses",
			Spec: `{"addresses": ["address"]}`,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"unknown": "test"}`,
			Err:  true,
		},
	})
}
