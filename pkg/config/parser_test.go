package config

import (
	"testing"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/stretchr/testify/assert"
)

const testConfig = `cloudquery {
  connection {
    dsn =  "host=localhost user=postgres password=pass DB.name=postgres port=5432"
  }
  provider "test" {
    source = "cloudquery"
    version = "v0.0.0"
  }
}

provider "aws" {
  configuration {
	account "dev" {
		role_arn ="12312312"
	}
	account "ron" {}
  }
  resources = ["slow_resource"]
}`

type Account struct {
	ID      string `hcl:",label"`
	RoleARN string `hcl:"role_arn,optional"`
}

type AwsConfig struct {
	Regions    []string  `hcl:"regions,optional"`
	Accounts   []Account `hcl:"account,block"`
	AWSDebug   bool      `hcl:"aws_debug,optional"`
	MaxRetries int       `hcl:"max_retries,optional" default:"5"`
	MaxBackoff int       `hcl:"max_backoff,optional" default:"30"`
}

func TestParser_LoadConfigFromSource(t *testing.T) {
	p := NewParser(nil)
	cfg, diags := p.LoadConfigFromSource("test.hcl", []byte(testConfig))
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].Configuration)
	cfg.Providers[0].Configuration = nil
	assert.Equal(t, &Config{
		CloudQuery: CloudQuery{
			Connection: &Connection{DSN: "host=localhost user=postgres password=pass DB.name=postgres port=5432"},
			Providers: []*RequiredProvider{{
				Name:    "test",
				Source:  "cloudquery",
				Version: "v0.0.0",
			}},
		},
		Providers: []*Provider{
			{
				Name:          "aws",
				Resources:     []string{"slow_resource"},
				Configuration: nil,
			},
		},
	}, cfg)
}

func TestProviderLoadConfiguration(t *testing.T) {
	p := NewParser(nil)
	cfg, diags := p.LoadConfigFromSource("test.hcl", []byte(testConfig))
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].Configuration)

	res, d := convert.Body(cfg.Providers[0].Configuration, convert.Options{Simplify: false})
	assert.Nil(t, d)
	c := AwsConfig{}
	errs := hclsimple.Decode("res.json", res, nil, &c)
	assert.Nil(t, errs)

}
