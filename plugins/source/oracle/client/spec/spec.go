package spec

import _ "embed"

type Spec struct {
	Concurrency int `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

//go:embed schema.json
var JSONSchema string
