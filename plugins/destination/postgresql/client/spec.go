package client

import (
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
)

const (
	defaultBatchSize      = 10000
	defaultBatchSizeBytes = 1000000
	defaultBatchTimeout   = 10 * time.Second
)

type Spec struct {
	ConnectionString string              `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel            `json:"pgx_log_level,omitempty"`
	MigrateMode      string              `json:"migrate_mode,omitempty"`
	BatchSize        int                 `json:"batch_size,omitempty"`
	BatchSizeBytes   int                 `json:"batch_size_bytes,omitempty"`
	BatchTimeout     configtype.Duration `json:"batch_timeout,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.MigrateMode == "" {
		s.MigrateMode = specs.MigrateModeSafe.String()
	}
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
