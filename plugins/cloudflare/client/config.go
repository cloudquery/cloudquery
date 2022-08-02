package client

// Provider Configuration

type Config struct {
	Token    string   `yaml:"api_token,omitempty"`
	ApiKey   string   `yaml:"api_key,omitempty"`
	ApiEmail string   `yaml:"api_email,omitempty"`
	Accounts []string `yaml:"accounts,omitempty"`
	Zones    []string `yaml:"zones,omitempty"`
}

func (Config) Example() string {
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
