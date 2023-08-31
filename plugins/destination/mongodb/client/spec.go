package client

import (
	"fmt"
)

const (
	defaultBatchSize      = 1000
	defaultBatchSizeBytes = 1024 * 1024 * 4
)

type Spec struct {
	ConnectionString string `json:"connection_string"`
	Database         string `json:"database"`
	BatchSize        int    `json:"batch_size"`
	BatchSizeBytes   int    `json:"batch_size_bytes"`
}

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
