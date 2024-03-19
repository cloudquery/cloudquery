package client

import (
	_ "embed"
	"fmt"
	"time"
)

type Spec struct {
	// The number of items to fetch concurrently
	ItemConcurrency int `json:"item_concurrency" jsonschema:"minimum=1,default=100"`
	// RFC3339 formatted timestamp. Syncing will begin with posts after this date. If not specified, the plugin will fetch all items.
	StartTime string `json:"start_time" jsonschema:"format=date-time"`
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

//go:embed schema.json
var JSONSchema string
