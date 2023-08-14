package client

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	CDCId            string   `json:"cdc_id,omitempty"`
}
