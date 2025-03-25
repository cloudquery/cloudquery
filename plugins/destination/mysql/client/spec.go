package client

import (
	_ "embed"
	"errors"
)

const (
	batchSize      = 1000
	batchSizeBytes = 4 * 1024 * 1024 // 4 MB
)

type Spec struct {
	// Connection string to connect to the database. See the [Go driver documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for details.
	//
	// - `"user:password@tcp(127.0.0.1:3306)/dbname"` connect with TCP
	// - `"user:password@127.0.0.1:3306/dbname?charset=utf8mb4&parseTime=True&loc=Local"` connect and set charset, time parsing, and location
	// - `"user:password@localhost:3306/dbname?timeout=30s&readTimeout=1s&writeTimeout=1s"` connect and set various timeouts
	// - `"user:password@/dbname?loc=UTC&allowNativePasswords=true&tls=preferred"` connect and set location and native password allowance, and prefer TLS
	ConnectionString string `json:"connection_string" jsonschema:"required,minLength=1"`

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

func (s Spec) Validate() error {
	if s.ConnectionString == "" {
		return errors.New("connection_string is required")
	}
	return nil
}
