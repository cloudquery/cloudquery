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
}

func (s *Spec) Validate() error {
	if len(s.APIKey) == 0 {
		return fmt.Errorf("missing API key")
	}
	return nil
}
