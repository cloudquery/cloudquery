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
			Name: "spec with unknown field",
			Spec: `{"unknown": "field"}`,
			Err:  true,
		},
		{
			Name: "spec with team_ids",
			Spec: `{"team_ids": ["team1", "team2"]}`,
		},
		{
			Name: "spec with max_requests_per_second",
			Spec: `{"max_requests_per_second": 5}`,
		},
		{
			Name: "spec with negative max_requests_per_second",
			Spec: `{"max_requests_per_second": -5}`,
			Err:  true,
		},
		{
			Name: "spec with concurrency",
			Spec: `{"concurrency": 5}`,
		},
		{
			Name: "spec with negative concurrency",
			Spec: `{"concurrency": -5}`,
			Err:  true,
		},
		{
			Name: "spec with all fields",
			Spec: `{"team_ids": ["team1", "team2"], "max_requests_per_second": 5, "concurrency": 5}`,
		},
	})
}
