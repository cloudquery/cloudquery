package client

import (
	"fmt"
)

type FormatType string

type BackendType string

const (
	FormatTypeCSV  = "csv"
	FormatTypeJSON = "json"
)

type Spec struct {
	Directory string     `json:"directory,omitempty"`
	Format    FormatType `json:"format,omitempty"`
	// This is used for debugging purposes only
	NoRotate bool `json:"no_rotate,omitempty"`
}

func (*Spec) SetDefaults() {}

func (s *Spec) Validate() error {
	if s.Directory == "" {
		return fmt.Errorf("directory is required")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
