package client

import (
	"github.com/cloudquery/codegen/jsonschema"
	"testing"
)

func TestJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty spec is valid",
			Spec: `{}`,
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
			Name: "region_codes == null is valid",
			Spec: `{"region_codes":null}`,
		},
		{
			Name: "region_codes with empty string is invalid",
			Spec: `{"region_codes": [""]}`,
			Err:  true,
		},
		{
			Name: "offer_codes == [] is valid",
			Spec: `{"offer_codes": []}`,
		},
		{
			Name: "offer_codes == null is valid",
			Spec: `{"offer_codes":null}`,
		},
		{
			Name: "offer_codes == [] is valid",
			Spec: `{"offer_codes": []}`,
		},
		{
			Name: "offer_codes with empty string is invalid",
			Spec: `{"offer_codes": [""]}`,
			Err:  true,
		},
	})
}
