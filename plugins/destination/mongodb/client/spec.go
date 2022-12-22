package client

import (
	"fmt"
)

type Spec struct {
	ConnectionString string `json:"connection_string"`
	Database         string `json:"database"`
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
