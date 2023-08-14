package sync

import (
	"context"
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/rs/zerolog"
)

const (
	defaultConcurrency = 10000
)

type Client struct {
	Jira *jira.Client
}

type Spec struct {
	Username    string `json:"username"`
	Token       string `json:"token"`
	BaseURL     string `json:"base_url"`
	Concurrency int    `json:"concurrency"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}

func (s *Spec) Validate() error {
	if s.BaseURL == "" {
		return fmt.Errorf("domain is required")
	}
	if s.Token == "" {
		return fmt.Errorf("token is required")
	}
	if s.Username == "" {
		return fmt.Errorf("username is required")
	}
	return nil
}

func New(ctx context.Context, logger zerolog.Logger, spec *Spec) (*Client, error) {
	c := &Client{}
	tp := jira.BasicAuthTransport{
		Username: spec.Username,
		Password: spec.Token,
	}
	jiraClient, err := jira.NewClient(tp.Client(), spec.BaseURL)
	if err != nil {
		return nil, err
	}
	_, _, err = jiraClient.User.GetSelfWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}
	c.Jira = jiraClient
	return c, nil
}

func (*Client) ID() string {
	return "jira"
}
