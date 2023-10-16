package client

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestMetric_JSONSchemaExtend(t *testing.T) {
	sc, err := jsonschema.Generate(Metric{})
	require.NoError(t, err)

	jsonschema.TestJSONSchema(t, string(sc), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true, // missing "name"
			Spec: `{}`,
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
			Name: "proper name",
			Spec: `{"name":"abc"}`,
		},
	})
}
