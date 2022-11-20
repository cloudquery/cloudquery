package client

import "fmt"

type Spec struct {
	DSN string `json:"dsn,omitempty"`
}

func (*Spec) SetDefaults() {
	// stub for any future defaults
}

func (s *Spec) Validate() error {
	if s.DSN == "" {
		return fmt.Errorf("dsn is required")
	}
	return nil
}
