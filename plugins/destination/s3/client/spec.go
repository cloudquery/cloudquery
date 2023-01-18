package client

import (
	"fmt"
	"strings"
)

type FormatType string

const (
	FormatTypeCSV  = "csv"
	FormatTypeJSON = "json"
)

type Spec struct {
	Bucket   string     `json:"bucket,omitempty"`
	Path     string     `json:"path,omitempty"`
	Format   FormatType `json:"format,omitempty"`
	NoRotate bool       `json:"no_rotate,omitempty"`
}

func (*Spec) SetDefaults() {}

func (s *Spec) Validate() error {
	if s.Bucket == "" {
		return fmt.Errorf("bucket is required")
	}
	if s.Path == "" {
		return fmt.Errorf("path is required")
	}
	if s.NoRotate && strings.Contains(s.Path, PathVarUUID) {
		return fmt.Errorf("path should not contain %s when no_rotate = true", PathVarUUID)
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
