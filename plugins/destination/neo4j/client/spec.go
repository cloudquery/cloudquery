package client

import (
	_ "embed"
	"errors"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4 // 4MB
)

type Spec struct {
	// Connection string to connect to the database. This can be a URL or a DSN, as per official [neo4j docs](https://neo4j.com/docs/browser-manual/current/operations/dbms-connection/#uri-scheme).
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1"`

	// Username to connect to the database.
	Username string `json:"username" jsonschema:"required,minLength=1"`

	// Password to connect to the database.
	Password string `json:"password" jsonschema:"required,minLength=1"`

	// Number of records to batch together before sending to the database.
	BatchSize int64 `json:"batch_size" jsonschema:"minimum=1,default=1000"`

	// Number of bytes (as Arrow buffer size) to batch together before sending to the database.
	BatchSizeBytes int64 `json:"batch_size_bytes" jsonschema:"minimum=1,default=4194304"`
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
		return errors.New("connection_string is required")
	}
	return nil
}
