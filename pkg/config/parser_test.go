package config

import (
	"testing"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/stretchr/testify/assert"
)

const testConfig = `cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
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

const testAliasProviderConfig = `cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
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
}

provider "aws" {
  alias = "another-aws"
  configuration {
	account "dev" {
		role_arn ="12312312"
	}
	account "ron" {}
  }
  resources = ["slow_resource"]
}`

const testMultipleProviderConfig = `cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
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
}

provider "aws" {
  configuration {
	account "dev" {
		role_arn ="12312312"
	}
	account "ron" {}
  }
  resources = ["slow_resource"]
}
`
const expectedDuplicateProviderError = "test.hcl:21,1-15: Provider Alias Required; Provider with name aws already exists, use alias in provider configuration block."

const testDuplicateAliasProviderConfig = `cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    source = "cloudquery"
    version = "v0.0.0"
  }
}

provider "aws" {
  alias = "same-aws"
  configuration {
	account "dev" {
		role_arn ="12312312"
	}
	account "ron" {}
  }
  resources = ["slow_resource"]
}

provider "aws" {
  alias = "same-aws"
  configuration {
	account "dev" {
		role_arn ="12312312"
	}
	account "ron" {}
  }
  resources = ["slow_resource"]
}
`
const expectedDuplicateAliasProviderError = "test.hcl:23,3-21: Duplicate Alias; Provider with alias same-aws for provider aws already exists, give it a different alias."

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
			Connection: &Connection{DSN: "postgres://postgres:pass@localhost:5432/postgres"},
			Providers: []*RequiredProvider{{
				Name:    "test",
				Source:  "cloudquery",
				Version: "v0.0.0",
			}},
		},
		Providers: []*Provider{
			{
				Name:          "aws",
				Alias:         "aws",
				Resources:     []string{"slow_resource"},
				Configuration: nil,
			},
		},
	}, cfg)
}

func TestParser_DuplicateProviderNaming(t *testing.T) {
	p := NewParser(nil)
	_, diags := p.LoadConfigFromSource("test.hcl", []byte(testMultipleProviderConfig))
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateProviderError, diags[0].Error())
}

func TestParser_AliasedProvider(t *testing.T) {
	p := NewParser(nil)
	cfg, diags := p.LoadConfigFromSource("test.hcl", []byte(testAliasProviderConfig))
	assert.Nil(t, diags)
	_, err := cfg.GetProvider("another-aws")
	assert.Nil(t, err)
	_, err = cfg.GetProvider("aws")
	assert.Nil(t, err)
}

func TestParser_DuplicateAliasedProvider(t *testing.T) {
	p := NewParser(nil)
	_, diags := p.LoadConfigFromSource("test.hcl", []byte(testDuplicateAliasProviderConfig))
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateAliasProviderError, diags[0].Error())
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
