package client

import (
	_ "embed"
	"errors"
)

const (
	defaultConcurrency = 10000
)

type Spec struct {
	Token       string   `json:"api_token,omitempty"`
	ApiKey      string   `json:"api_key,omitempty"`
	ApiEmail    string   `json:"api_email,omitempty"`
	Accounts    []string `json:"accounts,omitempty"`
	Zones       []string `json:"zones,omitempty"`
	Concurrency int      `json:"concurrency,omitempty"`
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

//go:embed schema.json
var JSONSchema string
