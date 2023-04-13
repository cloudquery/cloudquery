package client

import (
	"fmt"
)

type Spec struct {
	// APIKey required to access Snyk API
	APIKey string `json:"api_key,omitempty"`

	// Organizations is a list of organizations to fetch information from.
	// By default, will fetch from all organizations available for user.
	Organizations []string `json:"organizations,omitempty"`

	// EndpointURL is an optional parameter to override the API URL for snyk.Client.
	// It defaults to https://api.snyk.io/api/
	EndpointURL string `json:"endpoint_url,omitempty"`

	// Retries is an optional parameter to override the default number of retries for retryable requests.
	Retries int `json:"retries,omitempty"`

	// RetryDelaySeconds is an optional parameter to override the default backoff time for retryable requests.
	RetryDelaySeconds int `json:"retry_delay_seconds,omitempty"`
}

func (s *Spec) Validate() error {
	if len(s.APIKey) == 0 {
		return fmt.Errorf("missing API key")
	}
	return nil
}
