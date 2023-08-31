package client

import "fmt"

const (
	batchSize      = 1000
	batchSizeBytes = 4 * 1024 * 1024
)

type Spec struct {
	BatchSize        int    `json:"batch_size,omitempty"`
	BatchSizeBytes   int    `json:"batch_size_bytes,omitempty"`
	ConnectionString string `json:"connection_string,omitempty"`
}

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
