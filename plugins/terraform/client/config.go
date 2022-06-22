package client

import (
	"github.com/cloudquery/cq-provider-sdk/cqproto"
)

type Config struct {
	Config []BackendConfigBlock `yaml:"config" hcl:"config,block"`

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

	// local backend
	config "mylocal" {
      backend = "local"
      path = "./examples/terraform.tfstate"
    }
	// s3 backend
    config "myremote" {
      backend = "s3"
      bucket = "tf-states"
      key    = "terraform.tfstate"
      region = "us-east-1"
      role_arn = ""
    }
}
`
	default:
		return `
config:
  - name: mylocal # local backend
    backend: local
    path: ./examples/terraform.tfstate
  - name: myremote # s3 backend
    backend: s3
    bucket: tf-states
    key: terraform.tfstate
    region: us-east-1
    role_arn: ""
`
	}
}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
