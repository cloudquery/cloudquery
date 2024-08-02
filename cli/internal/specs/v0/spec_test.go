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
			Spec: `{"spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "empty kind",
			Err:  true,
			Spec: `{"kind":"","spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "null kind",
			Err:  true,
			Spec: `{"kind":null,"spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "bad kind",
			Err:  true,
			Spec: `{"kind":123,"spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "bad kind value",
			Err:  true,
			Spec: `{"kind":"123","spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "kind:source",
			Spec: `{"kind":"source","spec":{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a","b","c"]}}`,
		},
		{
			Name: "kind:destination",
			Spec: `{"kind":"destination","spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "kind:transformer",
			Spec: `{"kind":"transformer","spec":{"name":"a","path":"b","registry":"local"}}`,
		},
		{
			Name: "kind:transformer with additional properties",
			Spec: `{"kind":"transformer","spec":{"name":"a","path":"b","registry":"local","additional":"property"}}`,
			Err:  true,
		},
		{
			Name: "missing spec",
			Err:  true,
			Spec: `{"kind":"source"}`,
		},
		{
			Name: "empty spec",
			Err:  true,
			Spec: `{"kind":"source","spec":{}}`,
		},
		{
			Name: "null spec",
			Err:  true,
			Spec: `{"kind":"source","spec":null}`,
		},
		{
			Name: "bad spec",
			Err:  true,
			Spec: `{"kind":"source","spec":123}`,
		},
		{
			Name: "kind:destination,spec:destination",
			Spec: `{"kind":"destination","spec":{"name":"a","path":"b","registry":"local","write_mode":"append"}}`,
		},
		{
			Name: "kind:source,spec:destination",
			Err:  true,
			Spec: `{"kind":"source","spec":{"name":"a","path":"b","registry":"local","write_mode":"append"}}`,
		},
		{
			Name: "kind:destination,spec:source",
			Err:  true,
			Spec: `{"kind":"destination","spec":{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a","b","c"]}}`,
		},
		{
			Name: "kind:source,spec:source",
			Spec: `{"kind":"source","spec":{"name":"a","path":"b","registry":"local","tables":["*"],"destinations":["a","b","c"]}}`,
		},
	})
}
