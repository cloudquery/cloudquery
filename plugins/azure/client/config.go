package client

import "github.com/cloudquery/cq-provider-sdk/cqproto"

// Provider Configuration

type Config struct {
	Subscriptions []string `hcl:"subscriptions,optional"`

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
		//  Optional. if you not specified, cloudquery tries to access all subscriptions available to tenant
		//  subscriptions = ["<YOUR_SUBSCRIPTION_ID_HERE>"]
}`
	default:
		return `
Optional. if you not specified, cloudquery tries to access all subscriptions available to tenant
subscriptions:
  - <YOUR_SUBSCRIPTION_ID_HERE>
`
	}
}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
