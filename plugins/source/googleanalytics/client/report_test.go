package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestReport_JSONSchemaExtend(t *testing.T) {
	sc, err := jsonschema.Generate(Report{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(sc), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing name & metrics
			Spec: `{}`,
		},
		{
			Name: "proper",
			Spec: `{"name":"1","metrics":[{"name":"2"}]}`,
		},
		{
			Name: "null name",
			Err:  true,
			Spec: `{"name":null,"metrics":[{"name":"2"}]}`,
		},
		{
			Name: "bad name type",
			Err:  true,
			Spec: `{"name":123,"metrics":[{"name":"2"}]}`,
		},
		{
			Name: "empty metrics",
			Err:  true,
			Spec: `{"name":"1","metrics":[]}`,
		},
		{
			Name: "null metrics",
			Err:  true,
			Spec: `{"name":"1","metrics":null}`,
		},
		{
			Name: "bad metrics",
			Err:  true,
			Spec: `{"name":"1","metrics":123}`,
		},
		{
			Name: "empty metrics entry",
			Err:  true,
			Spec: `{"name":"1","metrics":[{}]}`,
		},
		{
			Name: "null metrics entry",
			Err:  true,
			Spec: `{"name":"1","metrics":[null]}`,
		},
		{
			Name: "bad metrics entry",
			Err:  true,
			Spec: `{"name":"1","metrics":[123]}`,
		},

		{
			Name: "empty dimensions",
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":[]}`,
		},
		{
			Name: "null dimensions",
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":null}`,
		},
		{
			Name: "bad dimensions",
			Err:  true,
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":123}`,
		},
		{
			Name: "empty dimensions entry",
			Err:  true,
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":[""]}`,
		},
		{
			Name: "null dimensions entry",
			Err:  true,
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":[null]}`,
		},
		{
			Name: "proper dimensions entry",
			Spec: `{"name":"1","metrics":[{"name":"2"}],"dimensions":["123"]}`,
		},
	})
}
