package client

import (
	"fmt"
)

type Spec struct {
	ConnectionString string `json:"connection_string"`
	Username         string `json:"username"`
	Password         string `json:"password"`

	BatchSize      int `json:"batch_size"`
	BatchSizeBytes int `json:"batch_size_bytes"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize == 0 {
		s.BatchSize = 1000
	}
	if s.BatchSizeBytes == 0 {
		s.BatchSizeBytes = 1024 * 1024 * 4
	}
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
