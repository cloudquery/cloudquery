package client

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	CDC              bool     `json:"cdc,omitempty"`
}

func (*Spec) SetDefaults() {
}

func (*Spec) Validate() error {
	return nil
}
