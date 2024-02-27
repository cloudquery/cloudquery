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
			Name: "empty accounts",
			Spec: `{"accounts": []}`,
			Err:  true,
		},
		{
			Name: "spec with accounts",
			Spec: `{"accounts": [{"name": "test", "api_key": "test", "app_key": "test"}]}`,
		},
		{
			Name: "spec with accounts with unknown field",
			Spec: `{"accounts": [{"name": "test", "api_key": "test", "app_key": "test", "unknown": "test"}]}`,
			Err:  true,
		},
		{
			Name: "spec with accounts with missing field",
			Spec: `{"accounts": [{"name": "test", "api_key": "test"}]}`,
			Err:  true,
		},
		{
			Name: "spec with accounts with empty field",
			Spec: `{"accounts": [{"name": "", "api_key": "", "app_key": ""}]}`,
			Err:  true,
		},
		{
			Name: "spec with accounts with invalid field",
			Spec: `{"accounts": [{"name": 1, "api_key": 1, "app_key": 1}]}`,
			Err:  true,
		},
		{
			Name: "spec with concurrency",
			Spec: `{"accounts": [{"name": "test", "api_key": "test", "app_key": "test"}], "concurrency": 10000}`,
		},
		{
			Name: "spec with site",
			Spec: `{"accounts": [{"name": "test", "api_key": "test", "app_key": "test"}], "site": "datadoghq.com"}`,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"unknown": "test"}`,
			Err:  true,
		},
	})
}
