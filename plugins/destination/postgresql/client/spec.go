package client

import "github.com/cloudquery/plugin-pb-go/specs/v0"

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	MigrateMode      string   `json:"migrate_mode,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.MigrateMode == "" {
		s.MigrateMode = specs.MigrateModeSafe.String()
	}
}
