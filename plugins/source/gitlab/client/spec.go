package client

import (
	_ "embed"

	"errors"
)

const (
	defaultConcurrency = 10000
)

// Spec is the (nested) spec used by the GitLab source plugin:
type Spec struct {
	// An access token for your GitLab server. Instructions on how to generate an access token here.
	Token string `json:"access_token,omitempty" jsonschema:"required,minLength=1"`
	//	URL for your self hosted GitLab server. Leave empty for GitLab SaaS. Not all tables are supported for GitLab SaaS.
	BaseURL string `json:"base_url,omitempty"`
	// A best effort maximum number of Go routines to use. Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) Validate() error {
	gitlabToken := s.Token
	if gitlabToken == "" {
		return errors.New("missing GitLab API token in configuration file")
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}

//go:embed schema.json
var JSONSchema string
