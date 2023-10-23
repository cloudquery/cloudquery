package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec_JSONSchemaExtend(t *testing.T) {
	// reports & oauth are tested in depth separately
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing property_id
			Spec: `{}`,
		},
		{
			Name: "empty property_id",
			Err:  true,
			Spec: `{"property_id":""}`,
		},
		{
			Name: "proper",
			Spec: `{"property_id":"abc"}`,
		},
		{
			Name: "empty start_date",
			Err:  true,
			Spec: `{"property_id":"abc","start_date":""}`,
		},
		{
			Name: "bad start_date",
			Err:  true,
			Spec: `{"property_id":"abc","start_date":"2023.10.21"}`,
		},
		{
			Name: "proper start_date",
			Spec: `{"property_id":"abc","start_date":"2023-10-21"}`,
		},
		{
			Name: "null oauth",
			Spec: `{"property_id":"abc","oauth":null}`,
		},
		{
			Name: "empty oauth",
			Err:  true,
			Spec: `{"property_id":"abc","oauth":{}}`,
		},
		{
			Name: "null reports",
			Spec: `{"property_id":"abc","reports":null}`,
		},
		{
			Name: "empty reports",
			Spec: `{"property_id":"abc","reports":[]}`,
		},
		{
			Name: "bad reports",
			Err:  true,
			Spec: `{"property_id":"abc","reports":123}`,
		},
		{
			Name: "concurrency < 1",
			Err:  true,
			Spec: `{"property_id":"abc","concurrency":0}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"property_id":"abc","concurrency":null}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"property_id":"abc","concurrency":123}`,
		},
	})
}
