package spec

import (
	"errors"

	"github.com/cloudquery/filetypes/v4"
)

type topicDetails struct {
	// Number of partitions to create for the topic.
	NumPartitions int `json:"num_partitions,omitempty" jsonschema:"minimum=1,default=1"`
	// Replication factor for the topic.
	ReplicationFactor int `json:"replication_factor,omitempty" jsonschema:"minimum=1,default=1"`
}

type TlsDetails struct {
	// Path to the certificate file for client authentication
	CertFile *string `json:"cert_file,omitempty"`
	// Path to the key file for client authentication
	KeyFile *string `json:"key_file,omitempty"`
	// Path to the certificate authority file for TLS client authentication
	CaFile *string `json:"ca_file,omitempty"`
}

type Spec struct {
	filetypes.FileSpec

	// List of brokers to connect to.
	//
	// Example broker address:
	//
	// - `"localhost:9092"` default url for a local Kafka broker
	Brokers []string `json:"brokers,omitempty" jsonschema:"required,minLength=1,minItems=1"`

	// If `true`, the plugin will log all underlying Kafka client messages to the log.
	Verbose bool `json:"verbose,omitempty"`

	// If connecting via SASL/PLAIN, the username to use.
	SASLUsername string `json:"sasl_username,omitempty"`

	// If connecting via SASL/PLAIN, the password to use.
	SASLPassword string `json:"sasl_password,omitempty"`

	// Enforce TLS Verification when configuring a username to connect to Kafka.
	EnforceTLSVerification bool `json:"enforce_tls_verification,omitempty"`

	// TLS details for client authentication.
	TlsDetails *TlsDetails `json:"tls_details,omitempty"`

	// Number of records to write before starting a new object.
	BatchSize int64 `json:"batch_size" jsonschema:"minimum=1,default=1000"`

	// Topic details, such as number of partitions and replication factor.
	TopicDetails topicDetails `json:"topic_details"`
}

func (s *Spec) SetDefaults() {
	s.FileSpec.SetDefaults()

	if s.BatchSize < 1 {
		s.BatchSize = 1000
	}
	if s.TopicDetails.NumPartitions < 1 {
		s.TopicDetails.NumPartitions = 1
	}
	if s.TopicDetails.ReplicationFactor < 1 {
		s.TopicDetails.ReplicationFactor = 1
	}
}

func (s *Spec) Validate() error {
	if len(s.Brokers) == 0 {
		return errors.New("at least one broker is required")
	}

	// required for s.FileSpec.Validate call
	err := s.FileSpec.UnmarshalSpec()
	if err != nil {
		return err
	}
	s.FileSpec.SetDefaults()

	return s.FileSpec.Validate()
}
