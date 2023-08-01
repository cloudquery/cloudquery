package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultConcurrency = 10000
)

type Spec struct {
	FastlyAPIKey string             `json:"fastly_api_key"`
	Services     []string           `json:"services"`
	Concurrency  int                `json:"concurrency"`
	Scheduler    scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}
