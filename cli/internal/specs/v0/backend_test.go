package specs

import (
	"encoding/json"
	"testing"

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
