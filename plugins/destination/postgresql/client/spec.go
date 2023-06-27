package client

import (
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v4/writers"
)

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	MigrateMode      string   `json:"migrate_mode,omitempty"`
	BatchSize        int      `json:"batch_size,omitempty"`
	BatchSizeBytes   int      `json:"batch_size_bytes,omitempty"`
	BatchTimeoutMs   int      `json:"batch_timeout_ms,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.MigrateMode == "" {
		s.MigrateMode = specs.MigrateModeSafe.String()
	}
	if s.BatchSize <= 0 {
		s.BatchSize = writers.DefaultBatchSize
	}
	if s.BatchSizeBytes <= 0 {
		s.BatchSizeBytes = writers.DefaultBatchSizeBytes
	}
	if s.BatchTimeoutMs <= 0 {
		s.BatchTimeoutMs = writers.DefaultBatchTimeoutSeconds * 1000
	}
}
