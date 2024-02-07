package client

import _ "embed"

const (
	defaultConcurrency = 10000
)

type Spec struct {
	Token       string   `json:"api_token,omitempty" jsonschema_extras:"x-cq-auth=true"`
	ApiKey      string   `json:"api_key,omitempty" jsonschema_extras:"x-cq-auth=true"`
	ApiEmail    string   `json:"api_email,omitempty" jsonschema_extras:"x-cq-auth=true"`
	Accounts    []string `json:"accounts,omitempty"`
	Zones       []string `json:"zones,omitempty"`
	Concurrency int      `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}

//go:embed schema.json
var JSONSchema string
