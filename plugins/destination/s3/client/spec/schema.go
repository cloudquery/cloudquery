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

	// batching enabled -> require {{UUID}} in path
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
		Then: pathWithUUID,
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"path": "the {{UUID}} placeholder must be present in the path",
				},
			},
		},
	}

	forceParquet := func(sc *jsonschema.Schema, field string) *jsonschema.Schema {
		val := *sc.Properties.Value(field)
		val.Enum = nil
		val.Const = "parquet"
		return &val
	}

	// write_empty_objects_for_empty_tables enabled -> require parquet format
	parquetEmptyObjects := &jsonschema.Schema{
		Title: "write_empty_objects_for_empty_tables requires parquet format",
		If: &jsonschema.Schema{
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				noRotate := *sc.Properties.Value("write_empty_objects_for_empty_tables")
				noRotate.Default = nil
				noRotate.Const = true
				noRotate.Description = ""
				properties := orderedmap.New[string, *jsonschema.Schema]()
				properties.Set("write_empty_objects_for_empty_tables", &noRotate)
				return properties
			}(),
			Required: []string{"write_empty_objects_for_empty_tables"},
		},
		Then: &jsonschema.Schema{
			// require properties not to be empty or null
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				properties := jsonschema.NewProperties()
				properties.Set("format", forceParquet(sc, "format"))
				return properties
			}(),
			Required: []string{"format"},
		},
		Extras: map[string]any{
			"errorMessage": map[string]any{
				"properties": map[string]any{
					"format": "when using `write_empty_objects_for_empty_tables` format must be set to `parquet`",
				},
			},
		},
	}

	sc.AllOf = append(sc.AllOf, cleanPath, noRotateNoUUID, noRotateNoBatch, uuidWhenBatching, parquetEmptyObjects)
}

//go:embed schema.json
var JSONSchema string
