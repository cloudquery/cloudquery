package spec

import _ "embed"

// CloudQuery Oracle source plugin configuration spec.
type Spec struct {
	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

//go:embed schema.json
var JSONSchema string
