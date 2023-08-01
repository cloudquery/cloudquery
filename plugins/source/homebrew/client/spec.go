package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	Concurrency int                `json:"concurrency"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency <= 0 {
		// use default concurrency recommended by plugin-SDK
		s.Concurrency = scheduler.DefaultConcurrency
	}
}

func (*Spec) Validate() error {
	return nil
}
