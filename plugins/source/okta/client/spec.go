package client

import (
	"errors"
	"time"

	"github.com/rs/zerolog"
)

type (
	Spec struct {
		Token       string     `json:"token,omitempty"`
		Domain      string     `json:"domain,omitempty"`
		RateLimit   *RateLimit `json:"rate_limit,omitempty"`
		Debug       bool       `json:"debug,omitempty"`
		Concurrency int        `json:"concurrency,omitempty"`
	}
	RateLimit struct {
		MaxBackoff time.Duration `json:"max_backoff,omitempty"`
		MaxRetries int32         `json:"max_retries,omitempty"`
	}
)

func (s *Spec) SetDefaults(logger *zerolog.Logger) {
	const (
		minRetries = int32(2)
		minBackOff = 30 * time.Second
	)

	if s.RateLimit == nil {
		s.RateLimit = new(RateLimit)
	}

	if s.RateLimit.MaxRetries < minRetries {
		s.RateLimit.MaxRetries = minRetries
	}

	if s.RateLimit.MaxBackoff < minBackOff {
		s.RateLimit.MaxBackoff = minBackOff
	}

	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (s Spec) Validate() error {
	if len(s.Token) == 0 {
		return errors.New("missing \"token\" in plugin configuration")
	}

	const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
	switch s.Domain {
	case "", exampleDomain:
		return errors.New("missing \"domain\" in plugin configuration")
	}

	return nil
}
