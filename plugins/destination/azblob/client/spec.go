package client

import (
	"fmt"

	"github.com/cloudquery/filetypes/v4"
)

type Spec struct {
	StorageAccount string `json:"storage_account,omitempty"`
	Container      string `json:"container,omitempty"`
	Path           string `json:"path,omitempty"`
	NoRotate       bool   `json:"no_rotate,omitempty"`
	*filetypes.FileSpec

	BatchSize      int `json:"batch_size"`
	BatchSizeBytes int `json:"batch_size_bytes"`
}

func (*Spec) SetDefaults() {}

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
