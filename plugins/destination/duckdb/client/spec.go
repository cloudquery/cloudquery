package client

import (
	"time"
)

type Spec struct {
	ConnectionString string        `json:"connection_string,omitempty"`
	WaitAfterDelete  time.Duration `json:"wait_after_delete,omitempty"`
}

func (s *Spec) SetDefaults() {
	const minWaitAfterDelete = 100 * time.Millisecond
	if s.WaitAfterDelete < minWaitAfterDelete {
		s.WaitAfterDelete = minWaitAfterDelete
	}
}
