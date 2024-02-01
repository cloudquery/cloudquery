package client

import (
	_ "embed"
	"fmt"
)

const (
	batchSize      = 1000
	batchSizeBytes = 4 * 1024 * 1024 // 4 MB
)

type Spec struct {
	// Connection string to connect to the database. See the [Go driver documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for details.
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1"`

	// Maximum amount of items that may be grouped together to be written in a single write.
	BatchSize int `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`
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

func (s Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
