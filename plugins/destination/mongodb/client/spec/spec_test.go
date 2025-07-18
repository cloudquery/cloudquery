package spec

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
			Spec: `{"connection_string": "conn"}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string and database",
			Spec: `{"connection_string": "conn", "database":"foo"}`,
		},
		{
			Name: "spec with bool connection_string",
			Spec: `{"connection_string": true, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with null connection_string",
			Spec: `{"connection_string": null, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with int connection_string",
			Spec: `{"connection_string": 123, "database":"foo"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "database":"foo", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "database":"foo", "unknown": "test"}`,
			Err:  true,
		},
	})
}
