package client

type Spec struct {
	AccessToken string `json:"access_token"`

	MaxRetries int64 `json:"max_retries"`
}
