package client

import (
	"encoding/json"
)

const (
	StateVersion = 4
)

// Hashicorp terraform state v4
// https://github.com/hashicorp/terraform/blob/main/internal/states/statefile/version4.go

type TerraformData struct {
	State State
}

type State struct {
	Version          uint64                 `json:"version"`
	TerraformVersion string                 `json:"terraform_version"`
	Serial           uint64                 `json:"serial"`
	Lineage          string                 `json:"lineage"`
	RootOutputs      map[string]OutputState `json:"outputs"`
	Resources        []Resource             `json:"resources"`
}

type OutputState struct {
	ValueRaw     json.RawMessage `json:"value"`
	ValueTypeRaw json.RawMessage `json:"type"`
	Sensitive    bool            `json:"sensitive,omitempty"`
}

type Resource struct {
	Module         string     `json:"module,omitempty"`
	Mode           string     `json:"mode"`
	Type           string     `json:"type"`
	Name           string     `json:"name"`
	EachMode       string     `json:"each,omitempty"`
	ProviderConfig string     `json:"provider"`
	Instances      []Instance `json:"instances"`
}

type Instance struct {
	IndexKey interface{} `json:"index_key,omitempty"`
	Status   string      `json:"status,omitempty"`
	Deposed  string      `json:"deposed,omitempty"`

	SchemaVersion           uint64            `json:"schema_version"`
	AttributesRaw           json.RawMessage   `json:"attributes,omitempty"`
	AttributesFlat          map[string]string `json:"attributes_flat,omitempty"`
	AttributeSensitivePaths json.RawMessage   `json:"sensitive_attributes,omitempty,"`

	PrivateRaw []byte `json:"private,omitempty"`

	Dependencies []string `json:"dependencies,omitempty"`

	CreateBeforeDestroy bool `json:"create_before_destroy,omitempty"`
}
