package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec is valid",
			Spec: `{}`,
		},
		{
			Name: "invalid start_time",
			Spec: `{"start_time":"now"}`,
			Err:  true,
		},
		{
			Name: "valid start_time",
			Spec: `{"start_time":"2000-01-01T00:00:01Z"}`,
		},
		{
			Name: "concurrency == -1 is invalid",
			Spec: `{"item_concurrency": -1}`,
			Err:  true,
		},
		{
			Name: "concurrency == 0 is invalid",
			Spec: `{"item_concurrency": 0}`,
			Err:  true,
		},
		{
			Name: "concurrency == 1 is valid",
			Spec: `{"item_concurrency": 1}`,
		},
	})
}
