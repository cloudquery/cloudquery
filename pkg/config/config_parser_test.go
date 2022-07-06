package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DecodeConfig_Extras(t *testing.T) {
	cases := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{
			name: "should fail if extra key in root",
			input: `
cloudquery:
    providers:
        - name: aws
          version: latest
    connection:
        type: postgres
        username: postgres
        password: pass
        host: localhost
        port: 5432
        database: postgres
        sslmode: disable
blurb:
	blah: blah
`,
			expectedError: true,
		},
		{
			name: "should fail if extra key in cq",
			input: `
cloudquery:
    providers:
        - name: aws
          version: latest
    connection:
        type: postgres
        username: postgres
        password: pass
        host: localhost
        port: 5432
        database: postgres
        sslmode: disable
	some-extra: blah
`,
			expectedError: true,
		},
		{
			name: "should not fail if no extra key in provider block",
			input: `
cloudquery:
    providers:
        - name: aws
          version: latest
    connection:
        type: postgres
        username: postgres
        password: pass
        host: localhost
        port: 5432
        database: postgres
        sslmode: disable
providers:
    # provider configurations
    - name: aws
      configuration:
        regions:
           - us-east-1
           - us-west-2
      resources:
        - accessanalyzer.analyzers
`,
			expectedError: false,
		},
		{
			name: "should fail if extra key in provider block",
			input: `
cloudquery:
    providers:
        - name: aws
          version: latest
    connection:
        type: postgres
        username: postgres
        password: pass
        host: localhost
        port: 5432
        database: postgres
        sslmode: disable
providers:
    # provider configurations
    - name: aws
      configuration:
      regions:
        - us-east-1
        - us-west-2
      resources:
        - accessanalyzer.analyzers
`,
			expectedError: true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			c, diags := decodeConfig(strings.NewReader(tc.input))
			assert.Equal(t, tc.expectedError, diags.HasErrors())
			if diags.HasErrors() {
				assert.Nil(t, c)
			} else {
				assert.NotNil(t, c)
			}
		})
	}
}

func Test_ProcessConfig_Connection(t *testing.T) {
	cases := []struct {
		name           string
		input          *Connection
		expectedResult string
		expectedError  bool
	}{
		{
			"should use the default port if none is specified",
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: `postgres`,
			},
			"postgres://user:pass@localhost:5432/postgres",
			false,
		},
		{
			"should use the provided port if specified",
			&Connection{
				Username: `user`,
				Type:     `postgres`,
				Host:     `localhost`,
				Port:     15432,
				Database: `postgres`,
			},
			"postgres://user@localhost:15432/postgres",
			false,
		},
		{
			"should append extras as query arguments if specified",
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: `postdb`,
				SSLMode:  `disable`,
				Extras:   []string{"a=b", "c=d", "e", "sslmode=enable"},
			},
			"postgres://user:pass@localhost:5432/postdb?a=b&c=d&e=&sslmode=disable",
			false,
		},
		{
			"should error if host is missing",
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     ``,
				Database: `postgres`,
			},
			"",
			true,
		},
		{
			"should error if database is missing",
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: ``,
			},
			"",
			true,
		},
		{
			"should error if dsn is set from config",
			&Connection{
				Username: `user`,
				Password: `pass`,
				Host:     `localhost`,
				Database: `postgres`,
				DSN:      "dsn",
			},
			"dsn",
			true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			config := Config{CloudQuery: CloudQuery{Connection: tc.input}}
			diags := ProcessConfig(&config)
			assert.Equal(t, tc.expectedError, diags.HasErrors())
			assert.Equal(t, tc.expectedResult, config.CloudQuery.Connection.DSN)
		})
	}
}

func TestHandle_ProcessConfigProviderVersion(t *testing.T) {
	cases := []struct {
		name            string
		providerVersion string
		expectedResult  string
		expectedError   bool
	}{
		{
			"should allow loose version",
			"v0.10",
			"v0.10.0",
			false,
		},
		{
			"should allow version without 'v' prefix",
			"0.10",
			"v0.10.0",
			false,
		},
		{
			"should allow 'latest' version",
			"latest",
			"latest",
			false,
		},
		{
			"should error if invalid semver",
			"invalid",
			"invalid",
			true,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			provider := &RequiredProvider{Name: "aws", Version: tc.providerVersion}
			config := Config{CloudQuery: CloudQuery{Providers: RequiredProviders{provider}, Connection: &Connection{DSN: "postgres://user:pass@localhost:5432/postgres"}}}
			diags := ProcessConfig(&config)
			assert.Equal(t, tc.expectedError, diags.HasErrors())
			assert.Equal(t, tc.expectedResult, config.CloudQuery.Providers[0].Version)
		})
	}
}
