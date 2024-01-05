package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpec(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra field",
			Err:  true,
			Spec: `{"extra": 0}`,
		},
		{
			Name: "zero concurrency",
			Err:  true,
			Spec: `{"concurrency": 0}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"concurrency": 0}`,
		},
		{
			Name: "bad concurrency",
			Err:  true,
			Spec: `{"concurrency": "abc"}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency": 123}`,
		},
		{
			Name: "null region_metadata",
			Spec: `{"region_metadata": null}`,
		},
		// other cases for region_metadata are tested separately
	})
}
