package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
		},
		{
			Name: "app_token",
			Spec: `{"app_token": "token"}`,
		},
		{
			Name: "empty app_token",
			Spec: `{"app_token": ""}`,
			Err:  true,
		},
		{
			Name: "max_requests_per_second == -1 is invalid",
			Spec: `{"app_token": "token", "max_requests_per_second": -1}`,
			Err:  true,
		},
		{
			Name: "max_requests_per_second == 0 is invalid",
			Spec: `{"app_token": "token", "max_requests_per_second": 0}`,
			Err:  true,
		},
		{
			Name: "max_requests_per_second == 1 is valid",
			Spec: `{"app_token": "token", "max_requests_per_second": 1}`,
		},
		{
			Name: "concurrency == -1 is invalid",
			Spec: `{"app_token": "token", "concurrency": -1}`,
			Err:  true,
		},
		{
			Name: "concurrency == 0 is invalid",
			Spec: `{"app_token": "token", "concurrency": 0}`,
			Err:  true,
		},
		{
			Name: "concurrency == 1 is valid",
			Spec: `{"app_token": "token", "concurrency": 1}`,
		},
		{
			Name: "table_options == null is valid",
			Spec: `{"app_token": "token", "table_options": null}`,
		},
	})
}

func TestTableOptionsJSONSchema(t *testing.T) {
	schema, err := jsonschema.Generate(TableOptions{})

	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "empty table options",
			Spec: `{}`,
		},
		{
			Name: "hubspot_crm_companies = null is invalid",
			Spec: `{"hubspot_crm_companies": null}`,
		},
		{
			Name: "hubspot_crm_companiess = [] is invalid",
			Spec: `{"hubspot_crm_companies": []}`,
			Err:  true,
		},
		{
			Name: "hubspot_crm_companies = {} is valid",
			Spec: `{"hubspot_crm_companies": {}}`,
		},
	})
}
