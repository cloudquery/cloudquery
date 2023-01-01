package client

import "fmt"

type FormatType string

const (
	FormatTypeCSV  = "csv"
	FormatTypeJSON = "json"
)

type Spec struct {
	Brokers      []string   `json:"brokers,omitempty"`
	Format       FormatType `json:"format,omitempty"`
	Verbose      bool       `json:"verbose,omitempty"`
	SaslUsername string     `json:"sasl_username,omitempty"`
	SaslPassword string     `json:"sasl_password,omitempty"`
}

func (*Spec) SetDefaults() {
}

func (s *Spec) Validate() error {
	if len(s.Brokers) == 0 {
		return fmt.Errorf("at least one broker is required")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
