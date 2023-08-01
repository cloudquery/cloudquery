package client

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultConcurrency = 1000
)

type Spec struct {
	AdAccountId string             `json:"ad_account_id"`
	AccessToken string             `json:"access_token"`
	Concurrency int                `json:"concurrency"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) Validate() error {
	if s.AdAccountId == "" {
		return fmt.Errorf("ad_account_id is required")
	}

	if s.AccessToken == "" {
		return fmt.Errorf("access_token is required")
	}

	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}
