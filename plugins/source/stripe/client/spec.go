package client

type Spec struct {
	APIKey string `json:"api_key"`

	MaxRetries  int64 `json:"max_retries"`
	StripeDebug bool  `json:"stripe_debug"`
}
