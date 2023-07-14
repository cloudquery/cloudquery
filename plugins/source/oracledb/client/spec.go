package client

import "fmt"

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
	Concurrency      int    `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 100
	}
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
