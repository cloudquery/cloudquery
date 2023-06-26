package client

import "github.com/cloudquery/plugin-sdk/v4/scheduler"

type Spec struct {
	Scheduler string `json:"scheduler"`
}

func (s *Spec) SetDefaults() {
	if s.Scheduler == "" {
		s.Scheduler = scheduler.StrategyDFS.String()
	}
}

func (s *Spec) Validate() error {
	_, err := scheduler.StrategyForName(s.Scheduler)
	return err
}
