package client_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/client"
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
			Name: "spec with host",
			Spec: `{"host": "conn"}`,
			Err:  true,
		},
		{
			Name: "spec with api_key",
			Spec: `{"api_key": "foo"}`,
			Err:  true,
		},
		{
			Name: "spec with host and api_key",
			Spec: `{"host": "conn", "api_key":"foo"}`,
		},
		{
			Name: "spec with bool host",
			Spec: `{"host": true, "api_key": "foo"}`,
			Err:  true,
		},
		{
			Name: "spec with null host",
			Spec: `{"host": null, "api_key": "foo"}`,
			Err:  true,
		},
		{
			Name: "spec with int host",
			Spec: `{"host": 123, "api_key": "foo"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"host": "abc", "api_key": "foo", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"host": "abc", "api_key": "foo", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"host": "abc", "api_key": "foo", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"host": "abc", "api_key": "foo", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with proper batch_size",
			Spec: `{"host": "abc", "api_key": "foo", "batch_size":7}`,
		},
		{
			Name: "spec with bool batch_timeout",
			Spec: `{"host": "abc", "api_key": "foo", "batch_timeout":true}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"host": "abc", "api_key": "foo", "unknown": "test"}`,
			Err:  true,
		},
	})
}
