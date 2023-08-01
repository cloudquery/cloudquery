package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

type Spec struct {
	NumClients  int                `json:"num_clients"`
	Concurrency int                `json:"concurrency,omitempty"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.NumClients <= 0 {
		s.NumClients = 1
	}
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (*Spec) Validate() error {
	return nil
}
