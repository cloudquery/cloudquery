package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestKindFromString(t *testing.T) {
	k, err := KindFromString("source")
	require.NoError(t, err)
	require.Equal(t, KindSource, k)

	k, err = KindFromString("destination")
	require.NoError(t, err)
	require.Equal(t, KindDestination, k)

	k, err = KindFromString("Destination")
	require.Error(t, err)
	require.Equal(t, KindSource, k)

	k, err = KindFromString("")
	require.Error(t, err)
	require.Equal(t, KindSource, k)
}

func TestKind_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(Kind))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
			Err:  true,
			Spec: `""`,
		},
		{
			Name: "null",
			Err:  true,
			Spec: `null`,
		},
		{
			Name: "bad type",
			Err:  true,
			Spec: `123`,
		},
		{
			Name: "bad value",
			Err:  true,
			Spec: `"extra"`,
		},
		{
			Name: "source",
			Spec: `"source"`,
		},
		{
			Name: "destination",
			Spec: `"destination"`,
		},
	})
}
