package client_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/client"
	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, client.JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string",
			Spec: `{"connection_string": "conn"}`,
		},
		{
			Name: "spec with connection_string and schema",
			Spec: `{"connection_string": "conn", "schema":"foo"}`,
		},
		{
			Name: "spec with connection_string and auth_mode azure",
			Spec: `{"connection_string": "conn", "auth_mode":"azure"}`,
		},
		{
			Name: "spec with connection_string and auth_mode ms",
			Spec: `{"connection_string": "conn", "auth_mode":"ms"}`,
		},
		{
			Name: "spec with connection_string and empty auth_mode",
			Spec: `{"connection_string": "conn", "auth_mode":""}`,
			Err:  true,
		},
		{
			Name: "spec with connection_string and invalid auth_mode",
			Spec: `{"connection_string": "conn", "auth_mode":"invalid"}`,
			Err:  true,
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
			Spec: `{"connection_string": "abc",  "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"connection_string": "abc", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with proper batch_size",
			Spec: `{"connection_string": "abc",  "batch_size":7}`,
		},
		{
			Name: "spec with bool batch_timeout",
			Spec: `{"connection_string": "abc", "batch_timeout":true}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"connection_string": "abc", "unknown": "test"}`,
			Err:  true,
		},
	})
}
