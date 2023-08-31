package client

import (
	"fmt"
	"time"
)

type Spec struct {
	ItemConcurrency int    `json:"item_concurrency"`
	StartTime       string `json:"start_time"`
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
}

func (s *Spec) Validate() error {
	if s.StartTime != "" {
		_, err := time.Parse(time.RFC3339, s.StartTime)
		if err != nil {
			return fmt.Errorf("could not parse start_time: %v", err)
		}
	}
	return nil
}
