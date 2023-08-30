package client

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	CDCId            string   `json:"cdc_id,omitempty"`
	RowsPerRecord    int      `json:"rows_per_record,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.RowsPerRecord < 1 {
		s.RowsPerRecord = 1
	}
}
