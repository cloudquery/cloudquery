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
			Name: "spec with connection_string",
			Spec: `{"connection_string": "file"}`,
		},
		{
			Name: "spec with bool connection_string",
			Spec: `{"connection_string": true}`,
			Err:  true,
		},
		{
			Name: "spec with null connection_string",
			Spec: `{"connection_string": null}`,
			Err:  true,
		},
		{
			Name: "spec with int connection_string",
			Spec: `{"connection_string": 123}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"connection_string": "abc", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"connection_string": "abc", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"connection_string": "abc", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "unknown": "test"}`,
			Err:  true,
		},
	})
}
