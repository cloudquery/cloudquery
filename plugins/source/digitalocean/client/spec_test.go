package client

import (
	"github.com/cloudquery/codegen/jsonschema"
	"testing"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec",
			Spec: `{}`,
		},
		{
			Name: "spec with token",
			Spec: `{"token": "abc"}`,
		},
		{
			Name: "spec with integer token",
			Spec: `{"token": 123}`,
			Err:  true,
		},
		{
			Name: "spec with spaces_access_key",
			Spec: `{"spaces_access_key": "abc"}`,
		},
		{
			Name: "spec with integer spaces_access_key",
			Spec: `{"spaces_access_key": 123}`,
			Err:  true,
		},
		{
			Name: "spec with spaces_access_key_id",
			Spec: `{"spaces_access_key_id": "abc"}`,
		},
		{
			Name: "spec with integer spaces_access_key_id",
			Spec: `{"spaces_access_key_id": 123}`,
			Err:  true,
		},
		{
			Name: "spec with empty spaces_regions",
			Spec: `{"spaces_regions":[]}`,
		},
		{
			Name: "spec with valid spaces_regions",
			Spec: `{"spaces_regions":["abc"]}`,
		},
		{
			Name: "spec with blank spaces_regions",
			Spec: `{"spaces_regions":[""]}`,
			Err:  true,
		},
		{
			Name: "spec with integer spaces_regions",
			Spec: `{"spaces_regions":[123]}`,
			Err:  true,
		},
	})
}
