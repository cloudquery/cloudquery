package client

import (
	"errors"
	"strings"
)

type Spec struct {
	APIKey string `json:"api_key"`

	MaxRetries  int64 `json:"max_retries"`
	RateLimit   int64 `json:"rate_limit"`
	StripeDebug bool  `json:"stripe_debug"`
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
}

func (s Spec) Validate() error {
	if s.APIKey == "" {
		return errors.New("no api_key provided")
	}
	return nil
}
