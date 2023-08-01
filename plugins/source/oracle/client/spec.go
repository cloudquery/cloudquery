package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (*Spec) Validate() error {
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
