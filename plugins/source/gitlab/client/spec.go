package client

import "errors"

const (
	defaultConcurrency = 10000
)

type Spec struct {
	Token       string `json:"access_token,omitempty"`
	BaseURL     string `json:"base_url,omitempty"`
	Concurrency int    `json:"concurrency,omitempty"`
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
