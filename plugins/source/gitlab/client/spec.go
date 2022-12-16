package client

type Spec struct {
	Token   string `json:"access_token,omitempty"`
	BaseURL string `json:"base_url,omitempty"`
}
