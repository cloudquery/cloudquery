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
		},
		{
			Name: "null max_requests_per_second",
			Spec: `{"max_requests_per_second": null}`,
		},
		{
			Name: "max_requests_per_second == -1 is invalid",
			Spec: `{"max_requests_per_second": -1}`,
			Err:  true,
		},
		{
			Name: "max_requests_per_second == 0 is invalid",
			Spec: `{"max_requests_per_second": 0}`,
			Err:  true,
		},
		{
			Name: "max_requests_per_second == 1 is valid",
			Spec: `{"max_requests_per_second": 1}`,
		},
		{
			Name: "concurrency == -1 is invalid",
			Spec: `{"concurrency": -1}`,
			Err:  true,
		},
		{
			Name: "concurrency == 0 is invalid",
			Spec: `{"concurrency": 0}`,
			Err:  true,
		},
		{
			Name: "concurrency == 1 is valid",
			Spec: `{"concurrency": 1}`,
		},
		{
			Name: "table_options == null is valid",
			Spec: `{"table_options": null}`,
		},
		{
			Name: "table_options == {} is valid",
			Spec: `{"table_options": {}}`,
		},
		{
			Name: "spec with table options = null is invalid",
			Spec: `{"table_options": {"hubspot_crm_companies": null}}`,
		},
		{
			Name: "spec with table options = [] is invalid",
			Spec: `{"table_options": {"hubspot_crm_companies": []}}`,
			Err:  true,
		},
		{
			Name: "spec with table options = {} is valid",
			Spec: `{"table_options": {"hubspot_crm_companies": {}}}`,
		},
	})
}
