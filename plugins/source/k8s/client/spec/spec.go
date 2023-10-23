package spec

import _ "embed"

// CloudQuery Kubernetes source plugin config spec.
type Spec struct {
	// Specify K8s contexts to connect to.
	// Specifying `*` will connect to all contexts available in the K8s config file (usually `~/.kube/config`).
	//
	// Default (empty or `null`) value results in using the default context from K8s's config file.
	Contexts []string `yaml:"contexts,omitempty" json:"contexts" jsonschema:"minLength=1"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `yaml:"concurrency,omitempty" json:"concurrency" jsonschema:"minimum=1,default=50000"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency <= 0 {
		const defaultConcurrency = 50000
		s.Concurrency = defaultConcurrency
	}
}

//go:embed schema.json
var JSONSchema string
