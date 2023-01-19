package client

import (
	"fmt"

	"github.com/cloudquery/filetypes"
)

type Spec struct {
	Bucket string `json:"bucket,omitempty"`
	Path   string `json:"path,omitempty"`
	*filetypes.FileSpec
}

func (s *Spec) SetDefaults() {
	s.FileSpec.SetDefaults()
}

func (s *Spec) Validate() error {
	if s.Bucket == "" {
		return fmt.Errorf("bucket is required")
	}
	if s.Path == "" {
		return fmt.Errorf("path is required")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
