package client

import (
	"fmt"

	"github.com/cloudquery/filetypes"
)

type FormatType string

const (
	FormatTypeCSV  = "csv"
	FormatTypeJSON = "json"
)

type Spec struct {
	StorageAccount string `json:"storage_account,omitempty"`
	Container      string `json:"container,omitempty"`
	Path           string `json:"path,omitempty"`
	*filetypes.FileSpec
}

func (s *Spec) SetDefaults() {
	s.FileSpec.SetDefaults()
}

func (s *Spec) Validate() error {
	if s.StorageAccount == "" {
		return fmt.Errorf("bucket is required")
	}
	if s.Container == "" {
		return fmt.Errorf("container is required")
	}
	if s.Path == "" {
		return fmt.Errorf("path is required")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
