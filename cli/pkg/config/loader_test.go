package config

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/cloudquery/cloudquery/cli/internal/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

type Account struct {
	ID      string `yaml:"id"`
	RoleARN string `yaml:"role_arn,omitempty"`
}

type AwsConfig struct {
	Regions    []string  `yaml:"regions,omitempty"`
	Accounts   []Account `yaml:"accounts"`
	AWSDebug   bool      `yaml:"aws_debug,omitempty"`
	MaxRetries int       `yaml:"max_retries,omitempty" default:"5"`
	MaxBackoff int       `yaml:"max_backoff,omitempty" default:"30"`
}

const testConfig = `
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: test
      source: cloudquery
      version: v0.0.0

providers:
  - name: aws
    configuration:
      accounts:
        - id: "dev"
          role_arn: "12312312"
        - id: "ron"
    resources: ["slow_resource"]
`

const expectedDuplicateProviderError = "provider with name aws already exists, use alias in provider configuration block"
const expectedDuplicateAliasProviderError = "provider with alias same-aws for provider aws-2 already exists, give it a different alias"

const bucketName = "myBucket"
const defaultPermissions = 0644

func TestLoader_LoadValidConfigFromFile(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/valid_config.yml")
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].ConfigBytes)
	cfg.Providers[0].ConfigBytes = nil
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
				Name:        "aws",
				Alias:       "aws",
				Resources:   []string{"slow_resource"},
				ConfigBytes: nil,
			},
		},
	}, cfg)
}

func TestLoader_BadVersion(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/bad_version.yml")
	assert.NotNil(t, diags)
	assert.Equal(t, "Provider \"test\" version \"invalid\" is invalid. Please set to 'latest' a or valid semantic version", diags[0].Description().Summary)
}

func TestLoader_DuplicateProviderNaming(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/duplicate_provider_name.yml")
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateProviderError, diags[0].Error())
}

func TestLoader_AliasedProvider(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/config_with_alias.yml")
	assert.Nil(t, diags)
	_, err := cfg.GetProvider("another-aws")
	assert.Nil(t, err)
	_, err = cfg.GetProvider("aws")
	assert.Nil(t, err)
}

func TestLoader_DuplicateAliasedProvider(t *testing.T) {
	p := NewParser()
	_, diags := p.LoadConfigFile("fixtures/duplicate_provider_alias.yml")
	assert.NotNil(t, diags)
	assert.Equal(t, expectedDuplicateAliasProviderError, diags[0].Error())
}

func TestProviderLoadConfiguration(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/valid_config.yml")
	assert.Nil(t, diags)
	assert.NotNil(t, cfg.Providers[0].ConfigBytes)

	c := AwsConfig{}
	errs := yaml.Unmarshal(cfg.Providers[0].ConfigBytes, &c)
	assert.Nil(t, errs)
}

func TestConfigEnvVariableSubstitution(t *testing.T) {
	p := NewParser(WithEnvironmentVariables(EnvVarPrefix, []string{
		"CQ_VAR_DSN=postgres://postgres:pass@localhost:5432/postgres",
		"CQ_VAR_ROLE_ARN=12312313",
	}))
	cfg, diags := p.LoadConfigFile("fixtures/env_vars.yml")
	if diags != nil {
		for _, d := range diags {
			t.Error(d.Error())
		}
		return
	}
	assert.Equal(t, "postgres://postgres:pass@localhost:5432/postgres", cfg.CloudQuery.Connection.DSN)

	c := AwsConfig{}
	errs := yaml.Unmarshal(cfg.Providers[0].ConfigBytes, &c)
	assert.Nil(t, errs)

	assert.Equal(t, "12312313", c.Accounts[0].RoleARN)
}

func TestLoader_LoadConfigNoSourceField(t *testing.T) {
	p := NewParser()
	cfg, diags := p.LoadConfigFile("fixtures/no_source.yml")
	assert.Nil(t, diags)
	// Check configuration was added, we will nil it after it to check the whole structure
	assert.NotNil(t, cfg.Providers[0].ConfigBytes)
	cfg.Providers[0].ConfigBytes = nil
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
				Name:        "test",
				Alias:       "test",
				Resources:   []string{"slow_resource"},
				ConfigBytes: nil,
			},
		},
	}, cfg)
	assert.Equal(t, cfg.CloudQuery.Providers[0].String(), "cq-provider-test@v0.0.0")
}

func TestLoader_LoadConfigFromSourceConnectionOptionality(t *testing.T) {
	cases := []struct {
		cfg           string
		expectedDSN   string
		expectedError bool
	}{
		{
			`
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
  providers:
    - name: aws
      version: latest
`,
			"postgres://postgres:pass@localhost:5432/postgres",
			false,
		},
		{
			`
cloudquery:
  connection:
    dsn: "postgres://postgres:pass@localhost:5432/postgres"
    database: "cq"
  providers:
    - name: aws
      version: latest
`,
			"",
			true,
		},
		{
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
    - name: aws
      version: latest
`,
			"postgres://postgres:pass@localhost:15432/cq?search_path=public&sslmode=disable",
			false,
		},
		{
			`
cloudquery:
  connection:
    username: "postgres"
    password: "pass"
    type: "postgres"
    host: "localhost"
    port: 15432
    database: "cq"
    sslmode: "disable"
    extras: [ "search_path=myschema" ]
  providers:
    - name: aws
      version: latest
`,
			"postgres://postgres:pass@localhost:15432/cq?search_path=myschema&sslmode=disable",
			false,
		},
	}
	for i := range cases {
		tc := cases[i]
		t.Run("case #"+strconv.Itoa(i+1), func(t *testing.T) {
			p := NewParser()
			parsedCfg, diags := p.LoadConfigFromSource([]byte(tc.cfg))
			if tc.expectedError {
				assert.True(t, diags.HasErrors())
			} else {
				assert.Len(t, diags.BySeverity(diag.ERROR), 0)
				if t.Failed() {
					return
				}
				assert.Equal(t, tc.expectedDSN, parsedCfg.CloudQuery.Connection.DSN)
			}
		})
	}
}

func Test_LoadFile(t *testing.T) {
	srvURL, backend := setupTestS3Bucket(t)
	cases := []struct {
		Path        string
		Name        string
		Configs     string
		Type        string
		SetupFile   bool
		ExpectError bool
	}{
		{
			Name:      "Success-S3Object",
			Type:      "s3",
			Path:      fmt.Sprintf("s3://%s/cloudquery.yml?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=%s", bucketName, srvURL.Host),
			Configs:   testConfig,
			SetupFile: true,
		},
		{
			Name:        "Failure-S3Object",
			Type:        "s3",
			Path:        fmt.Sprintf("s3://%s/cloudquery2.yml?region=us-east-1&disableSSL=true&s3ForcePathStyle=true&endpoint=%s", bucketName, srvURL.Host),
			Configs:     testConfig,
			SetupFile:   false,
			ExpectError: true,
		},
		{
			Name:      "Success-RelativePath",
			Type:      "file",
			Path:      "./asdf/asdf/asfd/teddstcdonfig.yml",
			Configs:   testConfig,
			SetupFile: true,
		},
		{
			Name:        "Failure-RelativePath-NotExists",
			Type:        "file",
			Path:        "./asdf/asdf/asfd/teddstcdonfig.yml",
			Configs:     testConfig,
			SetupFile:   false,
			ExpectError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var p *Parser
			switch tc.Type {
			case "s3":
				p = NewParser()
				os.Setenv("AWS_ANON", "true")
				defer os.Unsetenv("AWS_ANON")
				if tc.SetupFile {
					assert.NoError(t, putFile(backend, tc.Path, "text/yaml", tc.Configs))
				}

			case "file":
				appFS := afero.NewMemMapFs()
				p = NewParser(func(p *Parser) {
					p.fs = afero.Afero{Fs: appFS}
				})
				if tc.SetupFile {
					p.fs.WriteFile(tc.Path, []byte(tc.Configs), defaultPermissions)
				}
			}
			body, diags := p.LoadFile(tc.Path)
			if !tc.ExpectError {
				assert.Equal(t, 0, len(diags))
				cfg := []byte(tc.Configs)
				assert.Nil(t, diags)
				assert.Equal(t, cfg, body)
			} else {
				assert.Equal(t, 1, len(diags))
				assert.Equal(t, "Failed to read file: file does not exist. Hint: Try `cloudquery init <provider>`", diags[0].Description().Summary)
				assert.Equal(t, fmt.Sprintf("The file %q could not be read", tc.Path), diags[0].Description().Detail)
			}
		})
	}
}

func setupTestS3Bucket(t *testing.T) (*url.URL, *s3mem.Backend) {
	backend := s3mem.New()
	faker := gofakes3.New(backend)

	srv := httptest.NewServer(faker.Server())

	t.Cleanup(srv.Close)

	assert.NoError(t, backend.CreateBucket(bucketName))
	u, err := url.Parse(srv.URL)
	assert.NoError(t, err)
	return u, backend
}

func putFile(backend gofakes3.Backend, path, mime, content string) error {
	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	_, err = backend.PutObject(
		bucketName,
		strings.TrimPrefix(u.Path, "/"),
		map[string]string{"Content-Type": mime},
		bytes.NewBufferString(content),
		int64(len(content)),
	)

	return err
}
