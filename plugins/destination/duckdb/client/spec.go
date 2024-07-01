package client

import _ "embed"

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4 // 4MB
)

type Spec struct {
	// Absolute or relative path to a file, such as `./example.duckdb`
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`

	// Enables debug logging
	Debug bool `json:"debug,omitempty" jsonschema:"default=false"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
}
