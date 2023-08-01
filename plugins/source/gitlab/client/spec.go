package client

import (
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultConcurrency = 10000
)

type Spec struct {
	Token       string             `json:"access_token,omitempty"`
	BaseURL     string             `json:"base_url,omitempty"`
	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
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
