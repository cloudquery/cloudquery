package spec

import _ "embed"

type Spec struct {
	Contexts    []string `yaml:"contexts,omitempty" json:"contexts" jsonschema:"minLength=1"`
	Concurrency int      `yaml:"concurrency,omitempty" json:"concurrency" jsonschema:"minimum=1,default=50000"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency <= 0 {
		const defaultConcurrency = 50000
		s.Concurrency = defaultConcurrency
	}
}

//go:embed schema.json
var JSONSchema string
