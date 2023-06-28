package client

import (
	"fmt"
	"time"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
)

type Spec struct {
	ItemConcurrency int      `json:"item_concurrency"`
	StartTime       string   `json:"start_time"`
	Backend         *Backend `json:"backend"`
}

type Backend struct {
	Name     string                 `json:"name"`
	Path     string                 `json:"path"`
	Registry managedplugin.Registry `json:"registry"`
	Version  string                 `json:"version"`
	Table    string                 `json:"table"`
	Spec     map[string]interface{} `json:"spec"`
}

func (s *Spec) SetDefaults() {
	if s.ItemConcurrency <= 0 {
		// Default to loading 100 concurrent items
		s.ItemConcurrency = 100
	}
	if s.Backend != nil {
		s.Backend.SetDefaults()
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
		if s.Backend.Name == "" {
			return fmt.Errorf("backend name is required")
		}
		if s.Backend.Version == "" {
			return fmt.Errorf("backend version is required")
		}
		if s.Backend.Path == "" {
			return fmt.Errorf("backend path is required")
		}
	}
	return nil
}

func (b *Backend) SetDefaults() {
	if b.Name == "" {
		b.Name = "cq_state"
	}
}
