package client

type Spec struct {
	Token    string   `json:"api_token,omitempty"`
	ApiKey   string   `json:"api_key,omitempty"`
	ApiEmail string   `json:"api_email,omitempty"`
	Accounts []string `json:"accounts,omitempty"`
	Zones    []string `json:"zones,omitempty"`
}
