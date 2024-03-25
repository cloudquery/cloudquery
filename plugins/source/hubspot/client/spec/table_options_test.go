package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestTableOptions_JSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(TableOptions))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "null",
			Spec: `null`,
		},
		{
			Name: "array",
			Spec: `[]`,
			Err:  true,
		},
		{
			Name: "integer",
			Spec: `123`,
			Err:  true,
		},
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "map",
			Spec: `{"a":{}}`,
		},
		{
			Name: "map with null value",
			Spec: `{"a": null}`,
		},
		{
			Name: "map with integer value",
			Spec: `{"a": 123}`,
			Err:  true,
		},
		{
			Name: "map with extra key in value",
			Spec: `{"a": {"extra":[""]}}`,
			Err:  true,
		},
	})
}

func TestTableOptionsSpec_JSONSchema(t *testing.T) {
	data, err := jsonschema.Generate(new(TableOptionsSpec))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "null",
			Spec: `null`,
			Err:  true,
		},
		{
			Name: "array",
			Spec: `[]`,
			Err:  true,
		},
		{
			Name: "integer",
			Spec: `123`,
			Err:  true,
		},
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra",
			Spec: `{"a":true}`,
			Err:  true,
		},
		{
			Name: "empty properties",
			Spec: `{"properties": []}`,
		},
		{
			Name: "null properties",
			Spec: `{"properties": null}`,
		},
		{
			Name: "integer properties",
			Spec: `{"properties": 123}`,
			Err:  true,
		},
		{
			Name: "empty properties value",
			Spec: `{"properties": [""]}`,
			Err:  true,
		},
		{
			Name: "null properties value",
			Spec: `{"properties": [null]}`,
			Err:  true,
		},
		{
			Name: "integer properties value",
			Spec: `{"properties": [123]}`,
			Err:  true,
		},
		{
			Name: "proper properties value",
			Spec: `{"properties": ["abc"]}`,
		},
		{
			Name: "empty associations",
			Spec: `{"associations": []}`,
		},
		{
			Name: "null associations",
			Spec: `{"associations": null}`,
		},
		{
			Name: "integer associations",
			Spec: `{"associations": 123}`,
			Err:  true,
		},
		{
			Name: "empty associations value",
			Spec: `{"associations": [""]}`,
			Err:  true,
		},
		{
			Name: "null associations value",
			Spec: `{"associations": [null]}`,
			Err:  true,
		},
		{
			Name: "integer associations value",
			Spec: `{"associations": [123]}`,
			Err:  true,
		},
		{
			Name: "proper associations value",
			Spec: `{"associations": ["abc"]}`,
		},
	})
}
