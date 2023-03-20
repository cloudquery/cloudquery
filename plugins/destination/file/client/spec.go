package client

import (
	"fmt"

	"github.com/cloudquery/filetypes"
)

type Spec struct {
	*filetypes.FileSpec
	Directory string `json:"directory,omitempty"`
	NoRotate  bool   `json:"no_rotate,omitempty"`
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
