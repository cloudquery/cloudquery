package client

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	*filetypes.FileSpec
	NoRotate  bool   `json:"no_rotate,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	Region    string `json:"region,omitempty"`
	Path      string `json:"path,omitempty"`
	Athena    bool   `json:"athena,omitempty"`
	TestWrite *bool  `json:"test_write,omitempty"`
	Endpoint  string `json:"endpoint,omitempty"`

	BatchSize      *int64               `json:"batch_size"`
	BatchSizeBytes *int64               `json:"batch_size_bytes"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout"`
}

func (s *Spec) SetDefaults() {
	if !strings.Contains(s.Path, PathVarTable) {
		// for backwards-compatibility, default to given path plus /{{TABLE}}.[format].{{UUID}} if
		// no {{TABLE}} value is found in the path string
		s.Path += fmt.Sprintf("/%s.%s", PathVarTable, s.Format)
		if !s.NoRotate {
			s.Path += "." + PathVarUUID
		}
	}
	if s.TestWrite == nil {
		b := true
		s.TestWrite = &b
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
	if s.Bucket == "" {
		return fmt.Errorf("`bucket` is required")
	}
	if s.Path == "" {
		return fmt.Errorf("`path` is required")
	}
	if s.Region == "" {
		return fmt.Errorf("`region` is required")
	}
	if s.NoRotate && strings.Contains(s.Path, PathVarUUID) {
		return fmt.Errorf("`path` should not contain %s when `no_rotate` = true", PathVarUUID)
	}
	if !strings.Contains(s.Path, PathVarUUID) && s.batchingEnabled() {
		return fmt.Errorf("`path` should contain %s when using a non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`", PathVarUUID)
	}
	if path.IsAbs(s.Path) {
		return fmt.Errorf("`path` should not start with a \"/\"")
	}
	if s.Path != path.Clean(s.Path) {
		return fmt.Errorf("`path` should not contain relative paths or duplicate slashes")
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
