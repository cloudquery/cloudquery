package client

import (
	"errors"
)

type Spec struct {
	// Required
	AccessToken string `json:"access_token,omitempty"`

	// Optional
	Timeout int64 `json:"timeout_secs,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Timeout < 1 {
		s.Timeout = 30
	}
}

func (s *Spec) Validate() error {
	if s.AccessToken == "" {
		return errors.New("no credentials provided")
	}

	return nil
}
