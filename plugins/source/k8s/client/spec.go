package client

import (
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
)

const (
	defaultConcurrency = 50000
)

type Spec struct {
	Contexts    []string           `yaml:"contexts,omitempty" json:"contexts"`
	Concurrency int                `yaml:"concurrency,omitempty" json:"concurrency"`
	Scheduler   scheduler.Strategy `json:"scheduler,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = defaultConcurrency
	}
}
