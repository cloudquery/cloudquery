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

	// path patterns: should be a clean path
	cleanPath := &jsonschema.Schema{
		Title: "`path` is a clean path value",
		Not: &jsonschema.Schema{
			Title: "`path` is not a clean path value",
			AnyOf: []*jsonschema.Schema{
				{
					Title: "`path` contains `./`",
					Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
						properties := orderedmap.New[string, *jsonschema.Schema]()
						properties.Set("path", &jsonschema.Schema{
							Type:    "string",
							Pattern: `^.*\./.*$`,
						})
						return properties
					}(),
				},
				{
					Title: "`path` contains `//`",
					Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
						properties := orderedmap.New[string, *jsonschema.Schema]()
						properties.Set("path", &jsonschema.Schema{
							Type:    "string",
							Pattern: `^.*//.*$`,
						})
						return properties
					}(),
				},
			},
		},
	}

	pathWithUUID := &jsonschema.Schema{
		Title: "Require {{UUID}} to be present in path",
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			// we make the non-zero requirement, so we want to allow only null here
			properties := orderedmap.New[string, *jsonschema.Schema]()
			properties.Set("path", &jsonschema.Schema{
				Type:    "string",
				Pattern: `^.*\{\{UUID\}\}.*$`,
			})
			return properties
		}(),
	}
	// no_rotate:true -> no {{UUID}} should be present in path
	noRotateNoUUID := &jsonschema.Schema{
		Title: "Disallow {{UUID}} in path when using no_rotate",
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
			Not: pathWithUUID,
		},
	}

	sc.AllOf = append(sc.AllOf, noRotateNoBatch, cleanPath, noRotateNoUUID)
}

//go:embed schema.json
var JSONSchema string
