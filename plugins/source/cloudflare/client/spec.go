package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultConccurency = 10000
)

type Spec struct {
	Token       string             `json:"api_token,omitempty"`
	ApiKey      string             `json:"api_key,omitempty"`
	ApiEmail    string             `json:"api_email,omitempty"`
	Accounts    []string           `json:"accounts,omitempty"`
	Zones       []string           `json:"zones,omitempty"`
	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConccurency
	}
}
