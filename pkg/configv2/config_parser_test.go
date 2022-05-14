package configv2

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testConfig = `
cloudquery:
  connection: {}
#    dsn: "postgres://postgres:pass@localhost:5432/postgres"
#  providers:
#    - name: "test"
#      source: "cloudquery"
#      version: "v0.0.0"

providers:
- name: "aws"
  configuration:
    accounts:
    - name: "dev"
      role_arn: "12312312"
    - name: "ron"
  resources:
  - "slow_resource"
`

const testNoSource = `
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: "test"
    version: "v0.0.0"

providers:
  - name: "aws"
  configuration:
    accounts:
    - name: "dev"
      role_arn: "12312312"
    - name: "ron"
  resources:
    - name: "slow_resource"
`
const testAliasProviderConfig = `
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: "test"
      version: "v0.0.0"

providers:
  - name: "aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
  - name: "aws"
    alias = "another-aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
`

const testMultipleProviderConfig = `
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: "test"
      version: "v0.0.0"

providers:
  - name: "aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
  - name: "aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
`

const expectedDuplicateProviderError = "test.hcl:21,1-15: Provider Alias Required; Provider with name aws already exists, use alias in provider configuration block."

const testDuplicateAliasProviderConfig = `
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: "test"
      version: "v0.0.0"

providers:
  - name: "aws"
    alias = "another-aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
  - name: "aws"
    alias = "another-aws"
    configuration:
      accounts:
      - name: "dev"
        role_arn: "12312312"
      - name: "ron"
    resources:
      - name: "slow_resource"
`

const expectedDuplicateAliasProviderError = "test.hcl:23,3-21: Duplicate Alias; Provider with alias same-aws for provider aws already exists, give it a different alias."

const testBadVersion = `cloudquery {
  connection {
    dsn =  "postgres://postgres:pass@localhost:5432/postgres"
  }
  provider "test" {
    source = "cloudquery"
    version = "0.0.0"
  }
}`

type Account struct {
	ID      string `yaml:",label"`
	RoleARN string `yaml:"role_arn,optional"`
}

// func TestParser_LoadConfigFromSource(t *testing.T) {
// 	p := NewParser()
// 	cfg, diags := p.LoadConfigFromSource([]byte(testConfig))
// 	assert.Nil(t, diags)
// 	// Check configuration was added, we will nil it after it to check the whole structure
// 	assert.NotNil(t, cfg.Providers[0].Configuration)
// 	cfg.Providers[0].Configuration = yaml.Node{}
// 	// source := "cloudquery"
// 	// assert.Equal(t, &Config{
// 	// 	CloudQuery: CloudQuery{
// 	// 		Connection: &Connection{DSN: "postgres://postgres:pass@localhost:5432/postgres"},
// 	// 		Providers: []RequiredProvider{{
// 	// 			Name:    "test",
// 	// 			Source:  &source,
// 	// 			Version: "v0.0.0",
// 	// 		}},
// 	// 		Logger: &logging.Config{},
// 	// 	},
// 	// 	Providers: []Provider{
// 	// 		{
// 	// 			Name:          "aws",
// 	// 			Alias:         "aws",
// 	// 			Resources:     []string{"slow_resource"},
// 	// 			Configuration: yaml.Node{},
// 	// 		},
// 	// 	},
// 	// }, cfg)
// }

// func TestParser_BadVersion(t *testing.T) {
// 	p := NewParser()
// 	_, diags := p.LoadConfigFromSource([]byte(testBadVersion))
// 	assert.NotNil(t, diags)
// 	assert.Equal(t, "test.hcl:1,1-11: Provider test version 0.0.0 is invalid; Please set to 'latest' version or valid semantic versioning starting with vX.Y.Z", diags[0].Error())
// }

// func TestParser_DuplicateProviderNaming(t *testing.T) {
// 	p := NewParser()
// 	_, diags := p.LoadConfigFromSource([]byte(testMultipleProviderConfig))
// 	assert.NotNil(t, diags)
// 	assert.Equal(t, expectedDuplicateProviderError, diags[0].Error())
// }

// func TestParser_AliasedProvider(t *testing.T) {
// 	p := NewParser()
// 	cfg, diags := p.LoadConfigFromSource([]byte(testAliasProviderConfig))
// 	assert.Nil(t, diags)
// 	_, err := cfg.GetProvider("another-aws")
// 	assert.Nil(t, err)
// 	_, err = cfg.GetProvider("aws")
// 	assert.Nil(t, err)
// }

// func TestParser_DuplicateAliasedProvider(t *testing.T) {
// 	p := NewParser()
// 	_, diags := p.LoadConfigFromSource([]byte(testDuplicateAliasProviderConfig))
// 	assert.NotNil(t, diags)
// 	assert.Equal(t, expectedDuplicateAliasProviderError, diags[0].Error())
// }

// func TestProviderLoadConfiguration(t *testing.T) {
// 	p := NewParser()
// 	cfg, diags := p.LoadConfigFromSource([]byte(testConfig))
// 	assert.Nil(t, diags)
// 	assert.NotNil(t, cfg.Providers[0].Configuration)
// }

const testEnvVarConfig = `
cloudquery:
  connection:
    dsn: ${DSN}"
  providers:
    - name: "test"
      source: "cloudquery"
      version: "v0.0.0"

providers:
  - name: "aws"
    configuration:
      accounts:
        - name: "dev"
          role_arn: "${ROLE_ARN}"
        - name: "ron"
    resources:
      - name: "slow_resource"
`

// func TestConfigEnvVariableSubstitution(t *testing.T) {
// 	p := NewParser(WithEnvironmentVariables([]string{
// 		"DSN=postgres://postgres:pass@localhost:5432/postgres",
// 		"ROLE_ARN=12312312",
// 	}))
// 	cfg, diags := p.LoadConfigFromSource([]byte(testEnvVarConfig))
// 	if diags != nil {
// 		for _, d := range diags {
// 			t.Error(d.Summary)
// 		}
// 		return
// 	}
// 	assert.Equal(t, "postgres://postgres:pass@localhost:5432/postgres", cfg.CloudQuery.Connection.DSN)
// }

// func TestParser_LoadConfigNoSourceField(t *testing.T) {
// 	p := NewParser()
// 	cfg, diags := p.LoadConfigFromSource([]byte(testNoSource))
// 	assert.Nil(t, diags)
// 	// Check configuration was added, we will nil it after it to check the whole structure
// 	// assert.NotNil(t, cfg.Providers[0].Configuration)
// 	// cfg.Providers[0].Configuration = nil
// 	// assert.Equal(t, &Config{
// 	// 	CloudQuery: CloudQuery{
// 	// 		Connection: &Connection{DSN: "postgres://postgres:pass@localhost:5432/postgres"},
// 	// 		Providers: []RequiredProvider{{
// 	// 			Name:    "test",
// 	// 			Source:  nil,
// 	// 			Version: "v0.0.0",
// 	// 		}},
// 	// 		Logger: &logging.Config{},
// 	// 	},
// 	// 	Providers: []Provider{
// 	// 		{
// 	// 			Name:          "test",
// 	// 			Alias:         "test",
// 	// 			Resources:     []string{"slow_resource"},
// 	// 			Configuration: nil,
// 	// 		},
// 	// 	},
// 	// }, cfg)
// 	assert.Equal(t, cfg.CloudQuery.Providers[0].String(), "cq-provider-test@v0.0.0")
// }

func TestParser_LoadConfigFromSourceConnectionOptionality(t *testing.T) {
	cases := []struct {
		Name           string
		cfg            string
		expectedConfig *Config
		expectedError  bool
	}{
		{
			"DSN only",
			`
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
  - name: "test"
    version: "1.1.0"
`,
			&Config{
				CloudQuery: &CloudQuery{
					Connection: &Connection{
						DSN: "postgres://postgres:pass@localhost:5432/postgres",
					},
					Providers: []RequiredProvider{{
						Name:    "test",
						Version: "1.1.0",
					}},
				},
			},
			false,
		},
		{
			"violate mutual exclusive dsn and connection block",
			`
cloudquery:
  dsn:  "postgres://postgres:pass@localhost:5432/postgres"
  connection:
    database: "cq"
`,
			nil,
			true,
		},
		{
			"Connection block only",
			`
cloudquery;
  connection:
    username: "postgres"
    password: "pass"
    host: "localhost"
    port: 15432
    database: "cq"
    sslmode: "disable"
`,
			&Config{
				CloudQuery: &CloudQuery{
					Connection: &Connection{
						Username: "postgres",
						Password: "pass",
						Host:     "localhost",
						Port:     15432,
						Database: "cq",
						SSLMode:  "disable",
					},
				},
			},
			false,
		},
		{
			"Connection block with extras",
			`
cloudquery:
  connection:
    username: "postgres"
    password: "pass"
    host: "localhost"
    port: 15432
    database: "cq"
    sslmode: "disable"
    extras:
      - "search_path=myschema"
`,
			&Config{
				CloudQuery: &CloudQuery{
					Connection: &Connection{
						Username: "postgres",
						Password: "pass",
						Host:     "localhost",
						Port:     15432,
						Database: "cq",
						SSLMode:  "disable",
						Extras:   []string{"search_path=myschema"},
					},
				},
			},
			false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			cfg, result, err := UnmarshalConfig([]byte(tc.cfg))
			if err != nil {
				if tc.expectedError {
					return
				}
				t.Fatalf("unexpected error: %s", err)
			}
			fmt.Println(result)
			if diff := cmp.Diff(cfg, tc.expectedConfig); diff != "" {
				t.Errorf("Config mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
