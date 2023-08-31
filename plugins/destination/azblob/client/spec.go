package client

import (
	"fmt"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

type Spec struct {
	StorageAccount string `json:"storage_account,omitempty"`
	Container      string `json:"container,omitempty"`
	Path           string `json:"path,omitempty"`
	NoRotate       bool   `json:"no_rotate,omitempty"`
	*filetypes.FileSpec

	BatchSize      *int64               `json:"batch_size"`
	BatchSizeBytes *int64               `json:"batch_size_bytes"`
	BatchTimeout   *configtype.Duration `json:"batch_timeout"`
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
	if s.StorageAccount == "" {
		return fmt.Errorf("`storage_account` is required")
	}
	if s.Container == "" {
		return fmt.Errorf("`container` is required")
	}
	if s.Path == "" {
		return fmt.Errorf("`path` is required")
	}
	if s.Format == "" {
		return fmt.Errorf("`format` is required")
	}
	if s.NoRotate && ((s.BatchSize != nil && *s.BatchSize > 0) || (s.BatchSizeBytes != nil && *s.BatchSizeBytes > 0) || (s.BatchTimeout != nil && s.BatchTimeout.Duration() > 0)) {
		return fmt.Errorf("`no_rotate` cannot be used with non-zero `batch_size`, `batch_size_bytes` or `batch_timeout_ms`")
	}

	return nil
}

func int64ptr(i int64) *int64 {
	return &i
}
