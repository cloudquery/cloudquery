package client

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
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

	BatchSize      *int64               `json:"batch_size"`
	BatchSizeBytes *int64               `json:"batch_size_bytes"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout"`
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
		if s.NoRotate {
			s.BatchSize = int64ptr(0)
		} else {
			s.BatchSize = int64ptr(10000)
		}
	}
	if s.BatchSizeBytes == nil {
		if s.NoRotate {
			s.BatchSizeBytes = int64ptr(0)
		} else {
			s.BatchSizeBytes = int64ptr(50 * 1024 * 1024) // 50 MiB
		}
	}
	if s.BatchTimeout == nil {
		if s.NoRotate {
			d := configtype.NewDuration(0)
			s.BatchTimeout = &d
		} else {
			d := configtype.NewDuration(30 * time.Second)
			s.BatchTimeout = &d
		}
	}
}

func (s *Spec) Validate() error {
	if s.Directory == "" && s.Path == "" {
		return fmt.Errorf("either `directory` or `path` must be set")
	}
	if s.Directory != "" && s.Path != "" {
		return fmt.Errorf("only one of `directory` or `path` is allowed")
	}
	if s.NoRotate && strings.Contains(s.Path, PathVarUUID) {
		return fmt.Errorf("`path` should not contain %s when `no_rotate` = true", PathVarUUID)
	}
	if !strings.Contains(s.Path, PathVarUUID) && s.batchingEnabled() {
		return fmt.Errorf("`path` should contain %s when using a non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`", PathVarUUID)
	}
	if s.Format == "" {
		return fmt.Errorf("`format` is required")
	}
	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)) {
		return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
	}

	return nil
}

func (s *Spec) batchingEnabled() bool {
	switch {
	case (s.BatchSize != nil && *s.BatchSize > 0) ||
		(s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) ||
		(s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0) ||
		(!s.NoRotate && s.BatchSize == nil) ||
		(!s.NoRotate && s.BatchSizeBytes == nil) ||
		(!s.NoRotate && s.BatchTimeout == nil):
		return true
	default:
		return false
	}
}

func int64ptr(i int64) *int64 {
	return &i
}
