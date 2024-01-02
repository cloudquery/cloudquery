package specs

import (
	"encoding/json"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestBackendJsonMarshalUnmarshal(t *testing.T) {
	b, err := json.Marshal(BackendLocal)
	if err != nil {
		t.Fatal("failed to marshal backend:", err)
	}
	var backend Backend
	if err := json.Unmarshal(b, &backend); err != nil {
		t.Fatal("failed to unmarshal backend:", err)
	}
	if backend != BackendLocal {
		t.Fatal("expected backend to be local, but got:", backend)
	}
}

func TestBackendYamlMarshalUnmarshal(t *testing.T) {
	b, err := yaml.Marshal(BackendLocal)
	if err != nil {
		t.Fatal("failed to marshal backend:", err)
	}
	var backend Backend
	if err := yaml.Unmarshal(b, &backend); err != nil {
		t.Fatal("failed to unmarshal backend:", err)
	}
	if backend != BackendLocal {
		t.Fatal("expected backend to be local, but got:", backend)
	}
}

func TestBackendFromString(t *testing.T) {
	b, err := BackendFromString("none")
	require.NoError(t, err)
	require.Equal(t, BackendNone, b)

	b, err = BackendFromString("local")
	require.NoError(t, err)
	require.Equal(t, BackendLocal, b)

	b, err = BackendFromString("Local")
	require.Error(t, err)
	require.Equal(t, BackendNone, b)

	b, err = BackendFromString("")
	require.Error(t, err)
	require.Equal(t, BackendNone, b)
}

func TestBackend_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(Backend))
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
			Name: "none",
			Spec: `"none"`,
		},
		{
			Name: "local",
			Spec: `"local"`,
		},
	})
}
