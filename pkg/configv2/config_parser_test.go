package configv2

import (
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
		name                     string
		configYaml               string
		expectedConfig           *Config
		expectedValidationErrors []string
		expectedError            bool
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
			nil,
			false,
		},
		{
			"violate mutual exclusive dsn and connection block",
			`
cloudquery:
  connection:
    username: "cq"
    dsn:  "postgres://postgres:pass@localhost:5432/postgres"
  providers:
  - name: "test"
    version: "1.1.0"
`,
			&Config{
				CloudQuery: &CloudQuery{
					Connection: &Connection{
						Username: "cq",
						DSN:      "postgres://postgres:pass@localhost:5432/postgres",
					},
					Providers: []RequiredProvider{{
						Name:    "test",
						Version: "1.1.0",
					}},
				},
			},
			nil,
			true,
		},
		{
			"Connection block only",
			`
cloudquery:
  connection:
    username: "postgres"
    password: "pass"
    host: "localhost"
    port: 15432
    database: "cq"
    sslmode: "disable"
  providers:
  - name: "test"
    version: "1.1.0"
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
					Providers: []RequiredProvider{{
						Name:    "test",
						Version: "1.1.0",
					}},
				},
			},
			nil,
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
  providers:
  - name: "test"
    version: "1.1.0"
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
					Providers: []RequiredProvider{{
						Name:    "test",
						Version: "1.1.0",
					}},
				},
			},
			nil,
			false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, result, err := UnmarshalConfig([]byte(tc.configYaml))
			if err != nil {
				if tc.expectedError {
					return
				}
				t.Fatalf("unexpected error: %s", err)
			}

			validationErrors := result.Errors()
			for i := range validationErrors {
				t.Log("validation errors:")
				desc := validationErrors[i].String()
				t.Log(desc)
			}
			if len(validationErrors) != len(tc.expectedValidationErrors) {
				t.Fatalf("expected %d validation errors, got %d", len(tc.expectedValidationErrors), len(validationErrors))
			} else {
				for i := range tc.expectedValidationErrors {
					if tc.expectedValidationErrors[i] != validationErrors[i].Description() {
						t.Errorf("expected validation error %s, got %s", tc.expectedValidationErrors[i], validationErrors[i].Description())
					}
				}
			}

			if diff := cmp.Diff(cfg, tc.expectedConfig); diff != "" {
				t.Errorf("Config mismatch (-want +got):\n%s", diff)
			}

		})
	}
}
