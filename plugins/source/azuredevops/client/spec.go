package client

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	PersonalAccessToken string `json:"personal_access_token"`
	OrganizationURL     string `json:"organization_url"`

	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
func (s Spec) Validate() error {
	if s.PersonalAccessToken == "" {
		return fmt.Errorf("missing personal access token in configuration")
	}
	if s.OrganizationURL == "" {
		return fmt.Errorf("missing organization url in configuration")
	}
	return nil
}
