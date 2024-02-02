package spec

import (
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

const (
	defaultBatchSize      = 10000
	defaultBatchSizeBytes = 100000000
	defaultBatchTimeout   = 60 * time.Second
)

type Spec struct {
	ConnectionString string              `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel            `json:"pgx_log_level,omitempty"`
	BatchSize        int                 `json:"batch_size,omitempty"`
	BatchSizeBytes   int                 `json:"batch_size_bytes,omitempty"`
	BatchTimeout     configtype.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.BatchSize <= 0 {
		s.BatchSize = defaultBatchSize
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = defaultBatchSizeBytes
	}
	if s.BatchTimeout.Duration() <= 0 {
		s.BatchTimeout = configtype.NewDuration(defaultBatchTimeout)
	}
}

func (s *Spec) Validate() error {
	if len(s.ConnectionString) == 0 {
		return errors.New("`connection_string` is required")
	}
	return nil
}
