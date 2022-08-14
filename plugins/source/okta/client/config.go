package client

// Provider Configuration

type Config struct {
	Token  string `yaml:"token,omitempty"`
	Domain string `yaml:"domain"`
}

func (Config) Example() string {
	return `
# Optional. Okta Token to access API, you can set this with OKTA_API_TOKEN env variable
# token: "<YOUR_OKTA_TOKEN>"
# Required. You okta domain name
# domain: "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
`
}
