package client

import "fmt"

type Spec struct {
	ConnectionString string   `json:"connection_string,omitempty"`
	PgxLogLevel      LogLevel `json:"pgx_log_level,omitempty"`
	CDC              bool     `json:"cdc,omitempty"`
	CDCId            string   `json:"cdc_id,omitempty"`
}

func (*Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	if s.CDC && s.CDCId == "" {
		return fmt.Errorf("cdc_id must be set when cdc is enabled")
	}
	return nil
}
