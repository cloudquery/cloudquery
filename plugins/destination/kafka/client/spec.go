package client

import (
	"fmt"

	"github.com/cloudquery/filetypes"
)

type Spec struct {
	Brokers      []string `json:"brokers,omitempty"`
	Verbose      bool     `json:"verbose,omitempty"`
	SaslUsername string   `json:"sasl_username,omitempty"`
	SaslPassword string   `json:"sasl_password,omitempty"`
	// This is currently only used for testing to wait for
	// kafka cluster to be ready in GitHub actions.
	MaxMetadataRetries int `json:"max_metadata_retries,omitempty"`
	*filetypes.FileSpec
}

func (s *Spec) SetDefaults() {
	if s.FileSpec == nil {
		s.FileSpec = &filetypes.FileSpec{}
	}
	s.FileSpec.SetDefaults()
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
