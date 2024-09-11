package client

import (
	_ "embed"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	// If true, will return an error on any write message rather than consume from the channel
	ErrorOnWrite bool `json:"error_on_write,omitempty" jsonschema:"default=false"`

	// If true, will return an error on migrate table messages rather than consume from the channel
	ErrorOnMigrate bool `json:"error_on_migrate,omitempty" jsonschema:"default=false"`

	// Whether to use a BatchWriter or not.
	BatchWriter bool `json:"batch_writer" jsonschema:"default=false"`

	// Maximum number of items that may be grouped together to be written in a single write.
	//
	// Defaults to `10000`.
	BatchSize *int64 `json:"batch_size" jsonschema:"minimum=1,default=10000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	//
	// Defaults to `52428800` (50 MiB).
	BatchSizeBytes *int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=52428800"`

	// Maximum interval between batch writes.
	//
	// Defaults to `30s`.
	BatchTimeout *configtype.Duration `json:"batch_timeout" jsonschema:"default=30s"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.BatchSize == nil {
		s.BatchSize = ptr(int64(10000))
	}
	if s.BatchSizeBytes == nil {
		s.BatchSizeBytes = ptr(int64(50 * 1024 * 1024)) // 50 MiB
	}
	if s.BatchTimeout == nil {
		s.BatchTimeout = ptr(configtype.NewDuration(30 * time.Second))
	}
}

func ptr[A any](a A) *A {
	return &a
}
