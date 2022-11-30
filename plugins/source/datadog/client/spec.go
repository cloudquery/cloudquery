package client

type Spec struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Name   string `json:"name"`
	APIKey string `json:"api_key"`
	AppKey string `json:"app_key"`
	APIUrl string `json:"api_url,omitempty"`
}
