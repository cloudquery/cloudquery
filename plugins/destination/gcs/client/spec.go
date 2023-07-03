package client

import (
	"fmt"
	"time"

	"github.com/cloudquery/filetypes/v4"
)

type Spec struct {
	Bucket   string `json:"bucket,omitempty"`
	Path     string `json:"path,omitempty"`
	NoRotate bool   `json:"no_rotate,omitempty"`
	*filetypes.FileSpec

	BatchSize      *int64 `json:"batch_size"`
	BatchSizeBytes *int64 `json:"batch_size_bytes"`
	BatchTimeoutMs *int64 `json:"batch_timeout_ms"`
}

func (s *Spec) SetDefaults() {
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
	if s.BatchTimeoutMs == nil {
		if s.NoRotate {
			s.BatchTimeoutMs = int64ptr(0)
		} else {
			s.BatchTimeoutMs = int64ptr(int64(30 * time.Second / time.Millisecond)) // 30 seconds
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
	if s.Format == "" {
		return fmt.Errorf("`format` is required")
	}

	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeoutMs != nil && *s.BatchTimeoutMs > 0)) {
		return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
	}

	return nil
}

func int64ptr(i int64) *int64 {
	return &i
}
