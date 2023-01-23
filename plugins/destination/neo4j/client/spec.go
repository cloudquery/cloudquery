package client

import (
	"fmt"
)

type Spec struct {
	ConnectionString string `json:"connection_string"`
	Username         string `json:"username"`
	Password         string `json:"password"`
}

func (*Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	if s.ConnectionString == "" {
		return fmt.Errorf("connection_string is required")
	}
	return nil
}
