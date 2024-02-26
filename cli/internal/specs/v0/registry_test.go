package specs

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/cloudquery/codegen/jsonschema"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestRegistryJsonMarshalUnmarshal(t *testing.T) {
	b, err := json.Marshal(RegistryGRPC)
	if err != nil {
		t.Fatal("failed to marshal registry:", err)
	}
	var registry Registry
	if err := json.Unmarshal(b, &registry); err != nil {
		t.Fatal("failed to unmarshal registry:", err)
	}
	if registry != RegistryGRPC {
		t.Fatal("expected registry to be grpc, but got:", registry)
	}
}

func TestRegistryYamlMarshalUnmarshal(t *testing.T) {
	b, err := yaml.Marshal(RegistryGRPC)
	if err != nil {
		t.Fatal("failed to marshal registry:", err)
	}
	var registry Registry
	if err := yaml.Unmarshal(b, &registry); err != nil {
		t.Fatal("failed to unmarshal registry:", err)
	}
	if registry != RegistryGRPC {
		t.Fatal("expected registry to be github, but got:", registry)
	}
}

func TestRegistryFromString(t *testing.T) {
	for idx, name := range AllRegistries {
		t.Run(name, func(t *testing.T) {
			r, err := RegistryFromString(name)
			require.NoError(t, err)
			require.Equal(t, Registry(idx), r)

			if len(name) > 0 {
				name = strings.ToUpper(name)
				t.Run(name, func(t *testing.T) {
					r, err := RegistryFromString(strings.ToUpper(name))
					require.Error(t, err)
					require.Equal(t, RegistryUnset, r)
				})
			}
		})
	}
}

func TestRegistry_JSONSchemaExtend(t *testing.T) {
	data, err := jsonschema.Generate(new(Registry))
	require.NoError(t, err)
	jsonschema.TestJSONSchema(t, string(data), []jsonschema.TestCase{
		{
			Name: "empty",
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
			Name: "unset",
			Spec: `""`,
		},
		{
			Name: "github",
			Spec: `"github"`,
		},
		{
			Name: "local",
			Spec: `"local"`,
		},
		{
			Name: "grpc",
			Spec: `"grpc"`,
		},
		{
			Name: "docker",
			Spec: `"docker"`,
		},
		{
			Name: "cloudquery",
			Spec: `"cloudquery"`,
		},
	})
}
