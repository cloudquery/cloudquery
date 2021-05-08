package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	    id = 1
		regions = ["us-east1"]
		resources = ["ec2"]
	}
	account "ron" {
	    id = 1
		regions = ["us-east1"]
		resources = ["ec2"]
	}
  }
  resources = ["slow_resource"]
}`


func TestParser_LoadConfigFromSource(t *testing.T) {
	p := NewParser(nil)
	cfg, diags := p.LoadConfigFromSource("test.hcl", []byte(testConfig))
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].Configuration)
	cfg.Providers[0].Configuration = nil
	assert.Equal(t, &Config{
		CloudQuery: CloudQuery{
			Connection: Connection{DSN: "host=localhost user=postgres password=pass DB.name=postgres port=5432"},
			Providers: []*RequiredProvider{{
				Name:    "test",
				Source:  "cloudquery",
				Version: "v0.0.0",
			}},
		},
		Providers:  []*Provider{
			{
				Name:          "aws",
				Resources:     []string{"slow_resource"},
				Configuration: nil,
			},
		},
	}, cfg)

}
