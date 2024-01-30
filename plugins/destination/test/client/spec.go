package client

import _ "embed"

type Spec struct {
	// If true, will return an error on write rather than consume from the channel
	ErrorOnWrite bool `json:"error_on_write,omitempty" jsonschema:"default=false"`
}

//go:embed schema.json
var JSONSchema string
