package client

import "github.com/cloudquery/cq-provider-sdk/cqproto"

// Provider Configuration

type Config struct {
	Token  string `hcl:"token,optional"`
	Domain string `hcl:"domain"`

	requestedFormat cqproto.ConfigFormat
}

func NewConfig(f cqproto.ConfigFormat) *Config {
	return &Config{
		requestedFormat: f,
	}
}

func (c Config) Example() string {
	switch c.requestedFormat {
	case cqproto.ConfigHCL:
		return `configuration {
	// Optional. Okta Token to access API, you can set this with OKTA_API_TOKEN env variable
    // token = <YOUR_OKTA_TOKEN>
	// Required. You okta domain name
    // domain =  https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com
}
`
	default:
		return `
Optional. Okta Token to access API, you can set this with OKTA_API_TOKEN env variable
token = <YOUR_OKTA_TOKEN>
Required. You okta domain name
domain =  https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com
`
	}
}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
