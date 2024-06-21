package client

import _ "embed"

const (
	batchSize      = 10000
	batchSizeBytes = 10 * 1024 * 1024 // 10 MB
)

type Spec struct {
	// Path to a file, such as `./mydb.sql`
	ConnectionString string `json:"connection_string,omitempty" jsonschema:"required,minLength=1"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`
}

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults() {
	if s.BatchSize == 0 {
		s.BatchSize = batchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = batchSizeBytes
	}
}
