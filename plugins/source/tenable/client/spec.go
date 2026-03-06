package client

import "fmt"

type Spec struct {
	// Tenable.io API access key
	AccessKey string `json:"access_key" jsonschema:"required,minLength=1"`
	// Tenable.io API secret key
	SecretKey string `json:"secret_key" jsonschema:"required,minLength=1"`
	// Maximum number of concurrent requests
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10
	}
}

func (s *Spec) Validate() error {
	if s.AccessKey == "" {
		return fmt.Errorf("access_key is required")
	}
	if s.SecretKey == "" {
		return fmt.Errorf("secret_key is required")
	}
	return nil
}
