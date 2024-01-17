package specs

import (
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
)

func TestWriteModeFromString(t *testing.T) {
	m, err := WriteModeFromString("overwrite-delete-stale")
	require.NoError(t, err)
	require.Equal(t, WriteModeOverwriteDeleteStale, m)

	m, err = WriteModeFromString("overwrite")
	require.NoError(t, err)
	require.Equal(t, WriteModeOverwrite, m)

	m, err = WriteModeFromString("append")
	require.NoError(t, err)
	require.Equal(t, WriteModeAppend, m)

	m, err = WriteModeFromString("Append")
	require.Error(t, err)
	require.Equal(t, WriteModeOverwriteDeleteStale, m)

	m, err = WriteModeFromString("")
	require.Error(t, err)
	require.Equal(t, WriteModeOverwriteDeleteStale, m)
}

func TestWriteMode_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(WriteMode))
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
			Name: "overwrite-delete-stale",
			Spec: `"overwrite-delete-stale"`,
		},
		{
			Name: "overwrite",
			Spec: `"overwrite"`,
		},
		{
			Name: "append",
			Spec: `"append"`,
		},
	})
}

func TestWriteModeRoundTrip(t *testing.T) {
	for _, writeModeStr := range AllWriteModes {
		writeMode, err := WriteModeFromString(writeModeStr)
		if err != nil {
			t.Fatal(err)
		}
		if writeModeStr != writeMode.String() {
			t.Fatalf("expected:%s got:%s", writeModeStr, writeMode.String())
		}
	}
}
