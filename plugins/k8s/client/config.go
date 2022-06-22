package client

import "github.com/cloudquery/cq-provider-sdk/cqproto"

type Config struct {
	Contexts []string `hcl:"contexts,optional"`

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
  // Optional. Set contexts that you want to fetch. If it is not given then all contexts from config are iterated over.
  // contexts = ["YOUR_CONTEXT_NAME1", "YOUR_CONTEXT_NAME2"]
}`
	default:
		return `
Optional. Set contexts that you want to fetch. If it is not given then all contexts from config are iterated over.
contexts:
  - YOUR_CONTEXT_NAME1
  - YOUR_CONTEXT_NAME2
`
	}
}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
