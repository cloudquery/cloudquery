package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpecJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "empty contexts",
			Spec: `{"contexts":[]}`,
		},
		{
			Name: "null contexts",
			Spec: `{"contexts":null}`,
		},
		{
			Name: "bad contexts",
			Err:  true,
			Spec: `{"contexts":123}`,
		},
		{
			Name: "empty contexts entry",
			Err:  true,
			Spec: `{"contexts":[""]}`,
		},
		{
			Name: "null contexts entry",
			Err:  true,
			Spec: `{"contexts":[null]}`,
		},
		{
			Name: "bad contexts entry",
			Err:  true,
			Spec: `{"contexts":[123]}`,
		},
		{
			Name: "proper contexts entry",
			Spec: `{"contexts":["some-ctx"]}`,
		},
		{
			Name: "zero concurrency",
			Err:  true,
			Spec: `{"concurrency":0}`,
		},
		{
			Name: "null concurrency",
			Err:  true,
			Spec: `{"concurrency":null}`,
		},
		{
			Name: "bad concurrency",
			Err:  true,
			Spec: `{"concurrency":3.5}`,
		},
		{
			Name: "proper concurrency",
			Spec: `{"concurrency":5}`,
		},
	})
}
