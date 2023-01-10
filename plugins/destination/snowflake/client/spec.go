package client

import "fmt"

type Spec struct {
	ConnectionString string `json:"connection_string,omitempty"`
}

func (*Spec) SetDefaults() {
	// stub for any future defaults
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
