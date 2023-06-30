package client

import (
	"fmt"
	"time"
)

type Spec struct {
	ItemConcurrency int      `json:"item_concurrency"`
	StartTime       string   `json:"start_time"`
	Backend         *Backend `json:"backend"`
}

type Backend struct {
	Table      string `json:"table"`
	Connection string `json:"connection"`
}

func (s *Spec) SetDefaults() {
	if s.ItemConcurrency <= 0 {
		// Default to loading 100 concurrent items
		s.ItemConcurrency = 100
	}
	if s.Backend != nil {
		if s.Backend.Table == "" {
			s.Backend.Table = "cloudquery_hackernews_state"
		}
	}
}

func (s *Spec) Validate() error {
	if s.StartTime != "" {
		_, err := time.Parse(time.RFC3339, s.StartTime)
		if err != nil {
			return fmt.Errorf("could not parse start_time: %v", err)
		}
	}
	if s.Backend != nil {
		if s.Backend.Connection == "" {
			return fmt.Errorf("backend connection is required")
		}
		if s.Backend.Table == "" {
			return fmt.Errorf("backend table is required")
		}
	}
	return nil
}
