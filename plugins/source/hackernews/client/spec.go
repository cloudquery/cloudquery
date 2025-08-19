package client

import (
	_ "embed"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	// The number of items to fetch concurrently
	ItemConcurrency int `json:"item_concurrency" jsonschema:"minimum=1,default=100"`
	// RFC3339 formatted timestamp. Syncing will begin with posts after this date.
	// Relative values like "3 days ago" are also supported.
	// If not specified, the plugin will default to 24 hours ago.
	StartTime configtype.Time `json:"start_time" jsonschema:"format=date-time"`
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
	if s.StartTime.IsZero() {
		// Default to 24 hours ago
		s.StartTime, _ = configtype.ParseTime("24 hours ago")
	}
}

func (*Spec) Validate() error {
	// validation for configtype.Time is done on unmarshalling
	return nil
}

//go:embed schema.json
var JSONSchema string
