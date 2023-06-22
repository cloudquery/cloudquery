package client

import "github.com/cloudquery/plugin-pb-go/specs"

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	BatchSize        int      `json:"batch_size,omitempty"`
	BatchSizeBytes   int      `json:"batch_size_bytes,omitempty"`
	MigrateMode      string   `json:"migrate_mode,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.MigrateMode == "" {
		s.MigrateMode = specs.MigrateModeSafe.String()
	}
}
