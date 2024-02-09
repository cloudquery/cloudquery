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
			Name: "spec with bearer_token",
			Spec: `{"bearer_token": "tok"}`,
		},
		{
			Name: "spec with bearer_token and empty notion_version",
			Spec: `{"bearer_token": "tok", "notion_version": ""}`,
			Err:  true,
		},
		{
			Name: "spec with bearer_token and valid notion_version",
			Spec: `{"bearer_token": "tok", "notion_version": "2021-01-01"}`,
		},
		{
			Name: "spec with bearer_token and bool notion_version",
			Spec: `{"bearer_token": "tok", "notion_version": true}`,
			Err:  true,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"bearer_token": "tok", "unknown": "test"}`,
			Err:  true,
		},
	})
}
