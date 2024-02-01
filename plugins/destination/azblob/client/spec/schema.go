package spec

import (
	_ "embed"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

func (s Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	s.FileSpec.JSONSchemaExtend(sc) // need to call manually

	batchTimeout := sc.Properties.Value("batch_timeout").OneOf[0] // 0 - val, 1 - null
	batchTimeout.Default = "30s"

	// no_rotate:true -> only nulls for batch options
	noRotateNoBatch := &jsonschema.Schema{
		Title: "Disallow batching when using no_rotate",
		If: &jsonschema.Schema{
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				noRotate := *sc.Properties.Value("no_rotate")
				noRotate.Default = nil
				noRotate.Const = true
				noRotate.Description = ""
				properties := orderedmap.New[string, *jsonschema.Schema]()
				properties.Set("no_rotate", &noRotate)
				return properties
			}(),
			Required: []string{"no_rotate"},
		},
		Then: &jsonschema.Schema{
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				// we make the non-zero requirement, so we want to allow only null here
				null := &jsonschema.Schema{Type: "null"}
				properties := orderedmap.New[string, *jsonschema.Schema]()
				properties.Set("batch_size", null)
				properties.Set("batch_size_bytes", null)
				properties.Set("batch_timeout", null)
				return properties
			}(),
		},
	}

	sc.AllOf = append(sc.AllOf, noRotateNoBatch)
}

//go:embed schema.json
var JSONSchema string
