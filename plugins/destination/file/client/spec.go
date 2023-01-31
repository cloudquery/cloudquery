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
	Directory string `json:"directory,omitempty"`
	NoRotate  bool   `json:"no_rotate,omitempty"`
	*filetypes.FileSpec
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
