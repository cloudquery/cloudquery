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
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"no_rotate":        "batching options must not be present when no_rotate is enabled",
					"batch_size":       "batching options must not be present when no_rotate is enabled",
					"batch_size_bytes": "batching options must not be present when no_rotate is enabled",
					"batch_timeout":    "batching options must not be present when no_rotate is enabled",
				},
			},
		},
	}

	// path patterns: should be a clean path
	cleanPath := &jsonschema.Schema{
		Title: "`path` is a clean path value",
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"path": "value must not contain ./ or //",
				},
			},
		},
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			// we make the non-zero requirement, so we want to allow only null here
			properties := orderedmap.New[string, *jsonschema.Schema]()
			properties.Set("path", &jsonschema.Schema{
				Type: "string",
				Not: &jsonschema.Schema{
					AnyOf: []*jsonschema.Schema{
						{
							Pattern: `^.*\./.*$`,
						},
						{
							Pattern: `^.*//.*$`,
						},
					},
				},
			})
			return properties
		}(),
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

	pathNotWithUUID := &jsonschema.Schema{
		Title: "Disallow {{UUID}} in path",
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			// we make the non-zero requirement, so we want to allow only null here
			properties := orderedmap.New[string, *jsonschema.Schema]()
			properties.Set("path", &jsonschema.Schema{
				Type: "string",
				Not: &jsonschema.Schema{
					Pattern: `^.*\{\{UUID\}\}.*$`,
				},
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
		Then: pathNotWithUUID,
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"path": "the {{UUID}} placeholder must not be present in the path when no_rotate is enabled",
				},
			},
		},
	}

	// batching enabled -> require {{UUID}} in path or require no path variables in path,
	// since we will use UUID by default if batch
	uuidWhenBatching := &jsonschema.Schema{
		Title: "Require {{UUID}} in path when batching",
		If: &jsonschema.Schema{
			// It's enough to disallow setting no_rotate to true
			// As otherwise we're requiring the positive batch size (& bytes) values
			Title: "Disallow setting no_rotate to true",
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				noRotate := *sc.Properties.Value("no_rotate")
				noRotate.Default = nil
				noRotate.Const = false
				noRotate.Description = ""
				properties := orderedmap.New[string, *jsonschema.Schema]()
				properties.Set("no_rotate", &noRotate)
				return properties
			}(),
		},
		Then: &jsonschema.Schema{
			AnyOf: []*jsonschema.Schema{
				pathWithUUID,
				{
					Title: "`path` does not contain path variables",
					Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
						properties := orderedmap.New[string, *jsonschema.Schema]()
						properties.Set("path", &jsonschema.Schema{
							Not: &jsonschema.Schema{
								Type:    "string",
								Pattern: `^.*\{\{.*\}\}.*$`,
							},
						})
						return properties
					}(),
				},
			},
		},
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"path": "value must contain the {{UUID}} placeholder or no path variables at all",
				},
			},
		},
	}

	sc.AllOf = append(sc.AllOf, noRotateNoBatch, cleanPath, noRotateNoUUID, uuidWhenBatching)
}

//go:embed schema.json
var JSONSchema string
