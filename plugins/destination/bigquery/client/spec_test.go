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
			Name: "spec with project_id",
			Spec: `{"project_id": "value"}`,
			Err:  true,
		},
		{
			Name: "spec with dataset_id",
			Spec: `{"dataset_id": "value"}`,
			Err:  true,
		},
		{
			Name: "spec with project_id and dataset_id",
			Spec: `{"project_id": "foo", "dataset_id": "bar"}`,
		},
		{
			Name: "spec with bool project_id",
			Spec: `{"project_id": true, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with null project_id",
			Spec: `{"project_id": null, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with int project_id",
			Spec: `{"project_id": 123, "dataset_id": "bar"}`,
			Err:  true,
		},
		{
			Name: "spec with bool batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":false}`,
			Err:  true,
		},
		{
			Name: "spec with null batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":null}`,
			Err:  true,
		},
		{
			Name: "spec with string batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":"str"}`,
			Err:  true,
		},
		{
			Name: "spec with array batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size":["abc"]}`,
			Err:  true,
		},
		{
			Name: "spec with proper batch_size",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "batch_size": 7}`,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"project_id": "foo", "dataset_id": "bar", "unknown": "test"}`,
			Err:  true,
		},
	})
}
