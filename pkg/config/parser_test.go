package config

import (
	"strconv"
	"testing"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
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

const expectedDuplicateProviderError = "fixtures/duplicate_provider_name.hcl:21,1-15: Provider Alias Required; Provider with name aws already exists, use alias in provider configuration block."

const expectedDuplicateAliasProviderError = "fixtures/duplicate_provider_alias.hcl:23,3-21: Duplicate Alias; Provider with alias same-aws for provider aws already exists, give it a different alias."
const expectedDuplicateAliasProviderErrorYaml = "<empty>:0,0-0: Duplicate Alias; Provider with alias same-aws for provider aws-2 already exists, give it a different alias."

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
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/valid_config.hcl")
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].Configuration)
	cfg.Providers[0].Configuration = nil
	source := "cloudquery"
	assert.Equal(t, &Config{
		CloudQuery: CloudQuery{
			Connection: &Connection{DSN: "postgres://postgres:pass@localhost:5432/postgres"},
			Providers: []*RequiredProvider{{
				Name:    "test",
				Source:  &source,
				Version: "v0.0.0",
			}},
			Logger: &logging.Config{},
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

func TestParser_BadVersion(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/bad_version.hcl")
	assert.NotNil(t, diags)
	assert.Equal(t, "Provider test version 0.0.0 is invalid", diags[0].Error())
}

func TestParser_DuplicateProviderNaming(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/duplicate_provider_name.hcl")
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateProviderError, diags[0].Error())
}

func TestParser_AliasedProvider(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/config_with_alias.hcl")
	assert.Nil(t, diags)
	_, err := cfg.GetProvider("another-aws")
	assert.Nil(t, err)
	_, err = cfg.GetProvider("aws")
	assert.Nil(t, err)
}

func TestParser_DuplicateAliasedProvider(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/duplicate_provider_alias.hcl")
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateAliasProviderError, diags[0].Error())
}

func TestParser_DuplicateAliasedProviderYaml(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/duplicate_provider_alias.yaml")
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateAliasProviderErrorYaml, diags[0].Error())
}

func TestProviderLoadConfiguration(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/valid_config.hcl")
	assert.Nil(t, diags)
	assert.NotNil(t, cfg.Providers[0].Configuration)

	c := AwsConfig{}
	errs := hclsimple.Decode("res.hcl", cfg.Providers[0].Configuration, nil, &c)
	assert.Nil(t, errs)
}

func TestWithEnvironmentVariables(t *testing.T) {
	const prefix = "PREFIX_"
	p := NewParser(WithEnvironmentVariables(prefix, []string{prefix + "VAR1=value1", prefix + "Var2=value 2"}))
	assert.Equal(t, "value1", p.HCLContext.Variables["VAR1"].AsString())
	assert.Equal(t, "value 2", p.HCLContext.Variables["Var2"].AsString())
}

func TestConfigEnvVariableSubstitution(t *testing.T) {
	p := NewParser(WithEnvironmentVariables(EnvVarPrefix, []string{
		"CQ_VAR_DSN=postgres://postgres:pass@localhost:5432/postgres",
		"CQ_VAR_ROLE_ARN=12312312",
	}))
	cfg, diags := p.LoadConfigFile("fixtures/env_vars.hcl")
	if diags != nil {
		for _, d := range diags {
			t.Error(d.Error())
		}
		return
	}
	assert.Equal(t, "postgres://postgres:pass@localhost:5432/postgres", cfg.CloudQuery.Connection.DSN)

	c := AwsConfig{}
	errs := hclsimple.Decode("res.hcl", cfg.Providers[0].Configuration, nil, &c)
	assert.Nil(t, errs)

	assert.Equal(t, "12312312", c.Accounts[0].RoleARN)
}

func TestParser_LoadConfigNoSourceField(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/no_source.hcl")
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].Configuration)
	cfg.Providers[0].Configuration = nil
	assert.Equal(t, &Config{
		CloudQuery: CloudQuery{
			Connection: &Connection{DSN: "postgres://postgres:pass@localhost:5432/postgres"},
			Providers: []*RequiredProvider{{
				Name:    "test",
				Source:  nil,
				Version: "v0.0.0",
			}},
			Logger: &logging.Config{},
		},
		Providers: []*Provider{
			{
				Name:          "test",
				Alias:         "test",
				Resources:     []string{"slow_resource"},
				Configuration: nil,
			},
		},
	}, cfg)
	assert.Equal(t, cfg.CloudQuery.Providers[0].String(), "cq-provider-test@v0.0.0")
}

func TestParser_LoadConfigFromSourceConnectionOptionality(t *testing.T) {
	cases := []struct {
		cfg           string
		expectedDSN   string
		expectedError bool
	}{
		{
			`
cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
  }
}
`,
			"postgres://postgres:pass@localhost:5432/postgres",
			false,
		},
		{
			`
cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
    database = "cq"
  }
}
`,
			"",
			true,
		},
		{
			`
cloudquery {
  connection {
    username = "postgres"
    password = "pass"
    host = "localhost"
    port = 15432
    database = "cq"
    sslmode = "disable"
  }
}
`,
			"postgres://postgres:pass@localhost:15432/cq?sslmode=disable",
			false,
		},
		{
			`
cloudquery {
  connection {
    username = "postgres"
    password = "pass"
    type = "postgres"
    host = "localhost"
    port = 15432
    database = "cq"
    sslmode = "disable"
	extras = [ "search_path=myschema" ]
  }
}
`,
			"postgres://postgres:pass@localhost:15432/cq?search_path=myschema&sslmode=disable",
			false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run("case #"+strconv.Itoa(i+1), func(t *testing.T) {
			p := NewParser()
			parsedCfg, diags := p.LoadConfigFromSource("test.hcl", []byte(tc.cfg))
			if tc.expectedError {
				assert.True(t, diags.HasErrors())
			} else {
				assert.Len(t, diags.BySeverity(diag.ERROR), 0)
				assert.Equal(t, tc.expectedDSN, parsedCfg.CloudQuery.Connection.DSN)
			}
		})
	}
}
