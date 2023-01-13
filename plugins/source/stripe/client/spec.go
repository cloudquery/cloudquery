package client

type Spec struct {
	APIKey string `json:"api_key"`

	MaxRetries  int64 `json:"max_retries"`
	RateLimit   int64 `json:"rate_limit"`
	StripeDebug bool  `json:"stripe_debug"`
}
