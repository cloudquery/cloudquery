package spec

import (
	"fmt"

	"github.com/cloudquery/filetypes/v4"
)

type Spec struct {
	filetypes.FileSpec

	// List of brokers to connect to.
	Brokers []string `json:"brokers,omitempty" jsonschema:"required,minLength=1,minItems=1"`

	// If `true`, the plugin will log all underlying Kafka client messages to the log.
	Verbose bool `json:"verbose,omitempty"`

	// If connecting via SASL/PLAIN, the username to use.
	SASLUsername string `json:"sasl_username,omitempty"`

	// If connecting via SASL/PLAIN, the password to use.
	SASLPassword string `json:"sasl_password,omitempty"`

	// Client ID to be set for Kafka API calls.
	ClientID string `json:"client_id,omitempty" jsonschema:"default=cq-destination-kafka"`

	// Number of records to write before starting a new object.
	BatchSize int `json:"batch_size" jsonschema:"minimum=1,default=1000"`
}

func (s *Spec) SetDefaults() {
	s.FileSpec.SetDefaults()

	if s.ClientID == "" {
		s.ClientID = "cq-destination-kafka"
	}

	if s.BatchSize < 1 {
		s.BatchSize = 1000
	}
}

func (s *Spec) Validate() error {
	if len(s.Brokers) == 0 {
		return fmt.Errorf("at least one broker is required")
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}
