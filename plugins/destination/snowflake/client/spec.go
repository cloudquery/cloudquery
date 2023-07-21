package client

import "fmt"

const (
	defaultBatchSize          = 1000
	defaultBatchSizeBytes     = 4 * 1024 * 1024
	defaultMigrateConcurrency = 1
)

type Spec struct {
	ConnectionString   string `json:"connection_string,omitempty"`
	BatchSize          int    `json:"batch_size,omitempty"`
	BatchSizeBytes     int    `json:"batch_size_bytes,omitempty"`
	MigrateConcurrency int    `json:"migrate_concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	// stub for any future defaults
	if s.BatchSize == 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.MigrateConcurrency == 0 {
		s.MigrateConcurrency = defaultMigrateConcurrency
	}
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
