package config

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
					Providers: RequiredProviders{{
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
