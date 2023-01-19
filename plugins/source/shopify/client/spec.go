package client

type Spec struct {
	// Either
	AccessToken string `json:"access_token,omitempty"`
	// or
	APIKey    string `json:"api_key,omitempty"`
	APISecret string `json:"api_secret,omitempty"`

	// Required
	ShopURL string `json:"shop_url"`

	// Optional
	Timeout    int64 `json:"timeout_secs,omitempty"`
	MaxRetries int64 `json:"max_retries,omitempty"`
	PageSize   int64 `json:"page_size,omitempty"`
}
