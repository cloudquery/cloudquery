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
			Name: "spec with connection_string, username and password",
			Spec: `{"connection_string": "file", "username": "user", "password": "pass"}`,
		},
		{
			Name: "spec with bool connection_string",
			Spec: `{"connection_string": true, "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with null connection_string",
			Spec: `{"connection_string": null, "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with int connection_string",
			Spec: `{"connection_string": 123, "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"connection_string": "abc", "batch_size":false, "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"connection_string": "abc", "batch_size":null, "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"connection_string": "abc", "batch_size":"str", "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "batch_size":["abc"], "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "unknown": "test", "username": "user", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with missing username",
			Spec: `{"connection_string": "abc", "password": "pass"}`,
			Err:  true,
		},
		{
			Name: "spec with missing password",
			Spec: `{"connection_string": "abc", "username": "user"}`,
			Err:  true,
		},
	})
}
