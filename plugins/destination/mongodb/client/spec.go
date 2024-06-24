package client

import (
	_ "embed"
	"fmt"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4
)

type Spec struct {
	// MongoDB URI as described in the official MongoDB [documentation](https://www.mongodb.com/docs/manual/reference/connection-string/).
	//
	// Example connection strings:
	// - `"mongodb://username:password@hostname:port/database"` basic connection
	// - `"mongodb+srv://username:password@cluster.example.com/database"` connecting to a MongoDB Atlas cluster
	// - `"mongodb://localhost:27017/myDatabase?authSource=admin"` specify authentication source
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1"`

	// Database to sync the data to.
	Database string `json:"database" jsonschema:"required,minLength=1"`

	// Maximum number of items that may be grouped together to be written in a single write.
	BatchSize int64 `json:"batch_size,omitempty" jsonschema:"minimum=1,default=1000"`

	// Maximum size of items that may be grouped together to be written in a single write.
	BatchSizeBytes int64 `json:"batch_size_bytes,omitempty" jsonschema:"minimum=1,default=4194304"`
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

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	if s.Database == "" {
		return fmt.Errorf("database is required")
	}
	return nil
}
