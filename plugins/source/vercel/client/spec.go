package client

import "errors"

type Spec struct {
	AccessToken string   `json:"access_token"`
	TeamIDs     []string `json:"team_ids,omitempty"`

	EndpointURL string `json:"endpoint_url,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
	PageSize    int64  `json:"page_size,omitempty"`
	MaxRetries  int64  `json:"max_retries,omitempty"`
	MaxWait     int64  `json:"max_wait_secs,omitempty"`

	Concurrency int `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.EndpointURL == "" {
		s.EndpointURL = "https://api.vercel.com"
	}
	if s.Timeout < 1 {
		s.Timeout = 5
	}
	if s.PageSize < 1 {
		s.PageSize = 100
	}
	if s.MaxRetries < 1 {
		s.MaxRetries = 10
	}
	if s.MaxWait < 1 {
		s.MaxWait = 300
	}
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (s Spec) Validate() error {
	if s.AccessToken == "" {
		return errors.New("no access token provided")
	}

	return nil
}
