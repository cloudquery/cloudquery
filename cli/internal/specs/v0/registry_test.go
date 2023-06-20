package specs

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestRegistryJsonMarshalUnmarshal(t *testing.T) {
	b, err := json.Marshal(RegistryGrpc)
	if err != nil {
		t.Fatal("failed to marshal registry:", err)
	}
	var registry Registry
	if err := json.Unmarshal(b, &registry); err != nil {
		t.Fatal("failed to unmarshal registry:", err)
	}
	if registry != RegistryGrpc {
		t.Fatal("expected registry to be github, but got:", registry)
	}
}

func TestRegistryYamlMarshalUnmarsahl(t *testing.T) {
	b, err := yaml.Marshal(RegistryGrpc)
	if err != nil {
		t.Fatal("failed to marshal registry:", err)
	}
	var registry Registry
	if err := yaml.Unmarshal(b, &registry); err != nil {
		t.Fatal("failed to unmarshal registry:", err)
	}
	if registry != RegistryGrpc {
		t.Fatal("expected registry to be github, but got:", registry)
	}
}
