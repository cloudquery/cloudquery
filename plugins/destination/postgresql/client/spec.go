package client

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
}

const defaultBatchSize = 1000

func (s *Spec) SetDefaults() {
}
