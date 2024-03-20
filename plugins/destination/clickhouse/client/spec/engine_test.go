package spec

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestEngineJSONSchema(t *testing.T) {
	schema, err := jsonschema.Generate(Engine{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(schema), []jsonschema.TestCase{
		{
			Name: "null",
			Err:  true,
			Spec: `null`,
		},
		{
			Name: "bad",
			Err:  true,
			Spec: `123`,
		},
		{
			Name: "empty",
			Spec: `{}`,
		},
		{
			Name: "extra keyword",
			Err:  true,
			Spec: `{"extra":true}`,
		},
		{
			Name: "empty name",
			Err:  true,
			Spec: `{"name":""}`,
		},
		{
			Name: "null name",
			Err:  true,
			Spec: `{"name":null}`,
		},
		{
			Name: "bad name type",
			Err:  true,
			Spec: `{"name":123}`,
		},
		{
			Name: "bad name format",
			Err:  true,
			Spec: `{"name":"SomeEngine"}`,
		},
		{
			Name: "name: MergeTree",
			Spec: `{"name":"MergeTree"}`,
		},
		{
			Name: "name: SomeRubbishMergeTree",
			Spec: `{"name":"SomeRubbishMergeTree"}`,
		},
		{
			Name: "empty parameters",
			Spec: `{"parameters":[]}`,
		},
		{
			Name: "null parameters",
			Spec: `{"parameters":null}`,
		},
		{
			Name: "bad parameters",
			Err:  true, // it's either null or array
			Spec: `{"parameters":123}`,
		},
		{
			Name: "parameters empty string entry",
			Spec: `{"parameters":[""]}`,
		},
		{
			Name: "parameters string entry",
			Spec: `{"parameters":["123"]}`,
		},
		{
			Name: "parameters zero float entry",
			Spec: `{"parameters":[0.0]}`,
		},
		{
			Name: "parameters positive float entry",
			Spec: `{"parameters":[5.7]}`,
		},
		{
			Name: "parameters negative float entry",
			Spec: `{"parameters":[-5.7]}`,
		},
		{
			Name: "parameters zero integer entry",
			Spec: `{"parameters":[0]}`,
		},
		{
			Name: "parameters positive integer entry",
			Spec: `{"parameters":[123]}`,
		},
		{
			Name: "parameters negative integer entry",
			Spec: `{"parameters":[-123]}`,
		},
		{
			Name: "parameters empty object entry",
			Spec: `{"parameters":[{}]}`,
		},
		{
			Name: "parameters non-empty object entry",
			Spec: `{"parameters":[{"a":[{"b":123}]}]}`,
		},
		{
			Name: "parameters empty array entry",
			Spec: `{"parameters":[[]]}`,
		},
		{
			Name: "parameters non-empty array entry",
			Spec: `{"parameters":[[{"a":[{"b":123}]}]]}`,
		},
		{
			Name: "null parameters entry",
			Spec: `{"parameters":[null]}`,
			Err:  true,
		},
	})
}
