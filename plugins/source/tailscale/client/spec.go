package client

import (
	"fmt"
)

type Spec struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	APIKey       string `json:"api_key,omitempty"`
	Tailnet      string `json:"tailnet,omitempty"`
	EndpointURL  string `json:"endpoint_url,omitempty"`
}

func (s *Spec) Validate() error {
	if len(s.Tailnet) == 0 {
		return fmt.Errorf("missing tailnet")
	}

	if s.APIKey == "" {
		if len(s.ClientID) == 0 {
			return fmt.Errorf("missing client_id")
		}
		if len(s.ClientSecret) == 0 {
			return fmt.Errorf("missing client_secret")
		}
	}

	return nil
}
