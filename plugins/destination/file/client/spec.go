package client

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/filetypes/v2"
)

const (
	PathVarFormat = "{{FORMAT}}"
	PathVarTable  = "{{TABLE}}"
	PathVarUUID   = "{{UUID}}"
	YearVar       = "{{YEAR}}"
	MonthVar      = "{{MONTH}}"
	DayVar        = "{{DAY}}"
	HourVar       = "{{HOUR}}"
	MinuteVar     = "{{MINUTE}}"
)

type Spec struct {
	*filetypes.FileSpec
	Directory string `json:"directory,omitempty"`
	NoRotate  bool   `json:"no_rotate,omitempty"`
	Path      string `json:"path,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Directory != "" {
		s.Path = path.Join(s.Directory, fmt.Sprintf("%s.%s", PathVarTable, s.Format))
		if !s.NoRotate {
			s.Path += "." + PathVarUUID
		}
	} else {
		s.Directory = path.Dir(s.Path)
	}
	if !strings.Contains(s.Path, PathVarTable) {
		s.Path = path.Join(s.Path, fmt.Sprintf("%s.%s", PathVarTable, s.Format))
	}
}

func (s *Spec) Validate() error {
	if s.Directory == "" && s.Path == "" {
		return fmt.Errorf("either directory or path must be set")
	}
	if s.Directory != "" && s.Path != "" {
		return fmt.Errorf("only one of directory or path is allowed")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}
	return nil
}
