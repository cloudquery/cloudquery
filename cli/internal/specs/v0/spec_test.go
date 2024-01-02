package specs

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

//go:embed schema.json
var schema string

func TestSpec_JSONSchemaExtend(t *testing.T) {
	jsonschema.TestJSONSchema(t, schema, []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `{}`,
		},
		{
			Name: "missing kind",
			Err:  true,
			Spec: `{"kind":"", "spec":{}}`,
		},
		{
			Name: "empty kind",
			Err:  true,
			Spec: `{"kind":"", "spec":{}}`,
		},
		{
			Name: "null kind",
			Err:  true,
			Spec: `{"kind":null, "spec":{}}`,
		},
		{
			Name: "bad kind",
			Err:  true,
			Spec: `{"kind":123, "spec":{}}`,
		},
		{
			Name: "bad kind value",
			Err:  true,
			Spec: `{"kind":"123", "spec":{}}`,
		},
		{
			Name: "kind:source",
			Spec: `{"kind":"source", "spec":{}}`,
		},
		{
			Name: "kind:destination",
			Spec: `{"kind":"source", "spec":{}}`,
		},
		{
			Name: "missing spec",
			Err:  true,
			Spec: `{"kind":"source"}`,
		},
		{
			Name: "empty spec",
			Spec: `{"kind":"source", "spec":{}}`,
		},
		{
			Name: "null spec",
			Err:  true,
			Spec: `{"kind":"source", "spec":null}`,
		},
		{
			Name: "bad spec",
			Err:  true,
			Spec: `{"kind":"source", "spec":123}`,
		},
		{
			Name: "kind:destination,spec:destination",
			Spec: `{"kind":"destination", "spec":{"write_mode":"append"}}`,
		},
		{
			Name: "kind:source,spec:destination",
			Err:  true,
			Spec: `{"kind":"source", "spec":{"write_mode":"append"}}`,
		},
		{
			Name: "kind:destination,spec:source",
			Err:  true,
			Spec: `{"kind":"destination", "spec":{"destinations":["a","b","c"]}}`,
		},
		{
			Name: "kind:source,spec:source",
			Spec: `{"kind":"source", "spec":{"destinations":["a","b","c"]}}`,
		},
	})
}
