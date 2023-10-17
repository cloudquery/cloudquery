package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
)

func TestSpecJSONSchema(t *testing.T) {
	jsonschema.TestJSONSchema(t, JSONSchema, []jsonschema.TestCase{})
	// Engine is tested separately

}
