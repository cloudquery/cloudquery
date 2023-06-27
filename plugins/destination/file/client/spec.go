package client

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/filetypes/v4"
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

	BatchSize      *int64 `json:"batch_size"`
	BatchSizeBytes *int64 `json:"batch_size_bytes"`
}

func (s *Spec) SetDefaults() {
	if s.Directory != "" {
		s.Path = path.Join(s.Directory, fmt.Sprintf("%s.%s", PathVarTable, s.Format))
		if !s.NoRotate {
			s.Path += "." + PathVarUUID
		}
	}
	if !strings.Contains(s.Path, PathVarTable) {
		s.Path = path.Join(s.Path, fmt.Sprintf("%s.%s", PathVarTable, s.Format))
	}
	if s.BatchSize == nil {
		i := int64(10000)
		s.BatchSize = &i
	}
	if s.BatchSizeBytes == nil {
		i := int64(50 * 1024 * 1024) // 50 MiB
		s.BatchSizeBytes = &i
	}
}

func (s *Spec) Validate() error {
	if s.Directory == "" && s.Path == "" {
		return fmt.Errorf("either `directory` or `path` must be set")
	}
	if s.Directory != "" && s.Path != "" {
		return fmt.Errorf("only one of `directory` or `path` is allowed")
	}
	if s.Format == "" {
		return fmt.Errorf("`format` is required")
	}
	return nil
}
