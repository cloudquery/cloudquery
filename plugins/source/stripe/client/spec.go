package client

import (
	"errors"
	"strings"
)

type Spec struct {
	APIKey string `json:"api_key"`

	MaxRetries  int64 `json:"max_retries,omitempty"`
	RateLimit   int64 `json:"rate_limit,omitempty"`
	StripeDebug bool  `json:"stripe_debug,omitempty"`

	Concurrency int `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.RateLimit < 1 {
		// https://stripe.com/docs/rate-limits
		if strings.HasPrefix(s.APIKey, "sk_live") {
			s.RateLimit = 90
		} else {
			s.RateLimit = 20
		}
	}

	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (s Spec) Validate() error {
	if s.APIKey == "" {
		return errors.New("no api_key provided")
	}
	return nil
}
