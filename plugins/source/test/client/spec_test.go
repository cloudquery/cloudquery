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
			Name: "spec with valid num_clients",
			Spec: `{"num_clients": 10}`,
		},
		{
			Name: "spec with null num_clients",
			Spec: `{"num_clients": null}`,
			Err:  true,
		},
		{
			Name: "spec with bad num_clients",
			Spec: `{"num_clients": 0}`,
			Err:  true,
		},
		{
			Name: "spec with bool num_clients",
			Spec: `{"num_clients": true}`,
			Err:  true,
		},
		{
			Name: "spec with str num_clients",
			Spec: `{"num_clients": "abc"}`,
			Err:  true,
		},
		{
			Name: "spec with array num_clients",
			Spec: `{"num_clients": [0]}`,
			Err:  true,
		},
		{
			Name: "spec with num_rows",
			Spec: `{"num_rows":7}`,
		},
		{
			Name: "spec with negative num_rows",
			Spec: `{"num_rows":-7}`,
			Err:  true,
		},
		{
			Name: "spec with str num_rows",
			Spec: `{"num_rows":"abc"}`,
			Err:  true,
		},
		{
			Name: "spec with null num_rows",
			Spec: `{"num_rows":null}`,
		},
		{
			Name: "spec with unknown field",
			Spec: `{"unknown": "test"}`,
			Err:  true,
		},
	})
}
