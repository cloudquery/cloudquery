package client

type Spec struct {
	Token    string   `json:"api_token,omitempty"`
	ApiKey   string   `json:"api_key,omitempty"`
	ApiEmail string   `json:"api_email,omitempty"`
	Accounts []string `json:"accounts,omitempty"`
	Zones    []string `json:"zones,omitempty"`
}

func (Spec) Example() string {
	return `
// Use can use either the API token or the API key
// API token is preferred

// API token to access Cloudflare resources, also can be set with the CLOUDFLARE_API_TOKEN environment variable
api_token: "<YOUR_API_TOKEN_HERE>"
// API key to access Cloudflare resources, also can be set with the CLOUDFLARE_API_KEY environment variable
api_key: "<YOUR_API_KEY_HERE>"
// API email to access Cloudflare resources, also can be set with the CLOUDFLARE_API_EMAIL environment variable
api_email: "<YOUR_API_EMAIL_HERE>"

// List of accounts to target, if empty, all accounts will be targeted
//accounts:
// - "<YOUR_ACCOUNT_ID>"

// List of accounts to target, if empty, all available zones will be targeted
//zones:
// - "<YOUR_ZONE_ID>"
`
}
