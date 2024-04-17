package client

import (
	_ "embed"
	"errors"

	"github.com/invopop/jsonschema"
)

const (
	defaultConcurrency = 10000
)

type Spec struct {
	Token       string   `json:"api_token,omitempty" jsonschema:"minLength=1,example=${CLOUDFLARE_API_TOKEN}"`
	ApiKey      string   `json:"api_key,omitempty" jsonschema:"minLength=1,example=${CLOUDFLARE_API_KEY}"`
	ApiEmail    string   `json:"api_email,omitempty" jsonschema:"minLength=1,example=${CLOUDFLARE_EMAIL}"`
	Accounts    []string `json:"accounts,omitempty"`
	Zones       []string `json:"zones,omitempty"`
	Concurrency int      `json:"concurrency,omitempty" jsonschema:"default=10000"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}

func (s *Spec) Validate() error {
	if s.Token == "" && s.ApiKey == "" && s.ApiEmail == "" {
		return errors.New("either api_token or api_key/api_email are required")
	}

	if s.Token != "" {
		if s.ApiKey != "" || s.ApiEmail != "" {
			return errors.New("api_token and api_key/api_email are mutually exclusive")
		}
		return nil
	}

	if s.ApiKey == "" {
		return errors.New("api_key is required when api_email is provided")
	}

	if s.ApiEmail == "" {
		return errors.New("api_email is required when api_key is provided")
	}

	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.OneOf = []*jsonschema.Schema{
		{Required: []string{"api_token"}},
		{Required: []string{"api_key", "api_email"}},
	}
}

//go:embed schema.json
var JSONSchema string
