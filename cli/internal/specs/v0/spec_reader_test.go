package specs

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const expectedApplicationDefaultCredentials = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"project_id\",\n  \"private_key_id\": \"private_key_id\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----privatekey\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"client_email\",\n  \"client_id\": \"client_id\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"auth_provider_x509_cert_url\",\n  \"client_x509_cert_url\": \"client_x509_cert_url\",\n  \"universe_domain\": \"googleapis.com\"\n}\n"

var expectedApplicationDefaultCredentialsWindows = strings.ReplaceAll(expectedApplicationDefaultCredentials, "\n", "\r\n")

func getExpectedApplicationDefaultCredentials() string {
	if runtime.GOOS == "windows" {
		return expectedApplicationDefaultCredentialsWindows
	}
	return expectedApplicationDefaultCredentials
}

type specLoaderTestCase struct {
	name         string
	path         []string
	err          func() string
	sources      []*Source
	destinations []*Destination
	envVariables map[string]string
}

func getPath(pathParts ...string) string {
	return path.Join("testdata", path.Join(pathParts...))
}

var specLoaderTestCases = []specLoaderTestCase{
	{
		name: "success",
		path: []string{getPath("gcp.yml"), getPath("dir")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "gcp", Path: "cloudquery/gcp", Version: "v1.0.0", Registry: RegistryLocal, Destinations: []string{"postgresqlv2"}, TableConcurrency: 10, Tables: []string{"test"}},
			{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: RegistryLocal, Destinations: []string{"postgresql"}, TableConcurrency: 10, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresqlv2", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: RegistryGrpc, WriteMode: WriteModeOverwrite, Spec: map[string]any{"credentials": "mytestcreds"}},
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: RegistryGrpc, WriteMode: WriteModeOverwrite},
		},
	},
	{
		name: "success_yaml_extension",
		path: []string{getPath("gcp.yml"), getPath("dir_yaml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "gcp", Path: "cloudquery/gcp", Version: "v1.0.0", Registry: RegistryLocal, Destinations: []string{"postgresqlv2"}, TableConcurrency: 10, Tables: []string{"test"}},
			{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: RegistryLocal, Destinations: []string{"postgresql"}, TableConcurrency: 10, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresqlv2", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: RegistryGrpc, WriteMode: WriteModeOverwrite, Spec: map[string]any{"credentials": "mytestcreds"}},
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: RegistryGrpc, WriteMode: WriteModeOverwrite},
		},
	},
	{
		name: "duplicate_source",
		path: []string{getPath("gcp.yml"), getPath("gcp.yml")},
		err: func() string {
			return "duplicate source name gcp"
		},
	},
	{
		name: "no_such_file",
		path: []string{getPath("dir", "no_such_file.yml"), getPath("dir", "postgresql.yml")},
		err: func() string {
			if runtime.GOOS == "windows" {
				return "open testdata/dir/no_such_file.yml: The system cannot find the file specified."
			}
			return "open testdata/dir/no_such_file.yml: no such file or directory"
		},
	},
	{
		name: "duplicate_destination",
		path: []string{getPath("dir", "postgresql.yml"), getPath("dir", "postgresql.yml")},
		err: func() string {
			return "duplicate destination name postgresql"
		},
	},
	{
		name: "different_versions_for_destinations",
		path: []string{getPath("gcp.yml"), getPath("gcpv2.yml")},
		err: func() string {
			return "destination postgresqlv2 is used by multiple sources cloudquery/gcp with different versions"
		},
	},
	{
		name: "multiple sources success",
		path: []string{getPath("multiple_sources.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "aws", Path: "cloudquery/aws", Version: "v4.6.1", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
			{Name: "azure", Path: "cloudquery/azure", Version: "v1.3.3", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1.6.3", Registry: RegistryGithub, Spec: map[string]any{"connection_string": "postgresql://postgres:pass@localhost:5432/postgres"}},
		},
	},
	{
		name: "environment variables",
		path: []string{getPath("env_variables.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "aws", Path: "cloudquery/aws", Version: "v1", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
			{Name: "azure", Path: "cloudquery/azure", Version: "v1.3.3", Registry: RegistryGithub, Destinations: []string{"postgresql", "postgresql"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1.6.3", Registry: RegistryGithub, Spec: map[string]any{"connection_string": "postgresql://localhost:5432/cloudquery?sslmode=disable", "version": "#v1"}},
		},
		envVariables: map[string]string{
			"VERSION":           "v1",
			"DESTINATIONS":      "postgresql",
			"CONNECTION_STRING": "postgresql://localhost:5432/cloudquery?sslmode=disable",
		},
	},
	{
		name: "environment variables with error",
		path: []string{getPath("env_variables.yml")},
		err: func() string {
			return "failed to expand environment variable in file testdata/env_variables.yml (section 3): env variable CONNECTION_STRING not found"
		},
		sources: []*Source{
			{Name: "aws", Path: "cloudquery/aws", Version: "v1", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
			{Name: "azure", Path: "cloudquery/azure", Version: "v1.3.3", Registry: RegistryGithub, Destinations: []string{"postgresql", "postgresql"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1.6.3", Registry: RegistryGithub, Spec: map[string]any{}},
		},
		envVariables: map[string]string{
			"VERSION":      "v1",
			"DESTINATIONS": "postgresql",
		},
	},
	{
		name: "environment variables in string without error",
		path: []string{getPath("env_variable_in_string.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "test", Path: "cloudquery/test", Version: "v1", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1", Registry: RegistryGithub, Spec: map[string]any{"custom_version": "#v1"}},
		},
		envVariables: map[string]string{
			"VERSION": "v1",
		},
	},
	{
		name: "environment variables in string with error",
		path: []string{getPath("env_variable_in_string.yml")},
		err: func() string {
			return "failed to expand environment variable in file testdata/env_variable_in_string.yml (section 2): env variable VERSION not found"
		},
		sources: []*Source{
			{Name: "test", Path: "cloudquery/test", Version: "v1", Registry: RegistryGithub, Destinations: []string{"postgresql"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "postgresql", Path: "cloudquery/postgresql", Version: "v1", Registry: RegistryGithub, Spec: map[string]any{}},
		},
		envVariables: map[string]string{},
	},
	{
		name: "number in name field",
		path: []string{getPath("numbers.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "0123456789", Path: "cloudquery/aws", Version: "v1", Registry: RegistryGithub, Destinations: []string{"0987654321"}, Tables: []string{"test"}},
			{Name: "012345", Path: "cloudquery/aws", Version: "v1", Registry: RegistryGithub, Destinations: []string{"0987654321"}, Tables: []string{"test"}},
		},
		destinations: []*Destination{
			{Name: "0987654321", Path: "cloudquery/postgresql", Version: "v1", Registry: RegistryGithub, Spec: map[string]any{"connection_string": "postgresql://localhost:5432/cloudquery?sslmode=disable"}},
		},
		envVariables: map[string]string{
			"ACCOUNT_ID": "0123456789",
		},
	},
	{
		name: "success importing JSON file",
		path: []string{getPath("json_file.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "gcp", Path: "cloudquery/gcp", Version: "v1.0.0", Destinations: []string{"bigquery"}, Tables: []string{"*"}},
		},
		destinations: []*Destination{
			{Name: "bigquery", Path: "cloudquery/bigquery", Version: "v3.1.0", Spec: map[string]any{"service_account_key_json": getExpectedApplicationDefaultCredentials()}},
		},
	},
	{
		name: "success importing JSON file using YAML pipe operator",
		path: []string{getPath("json_file_yaml_pipe.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{Name: "gcp", Path: "cloudquery/gcp", Version: "v1.0.0", Destinations: []string{"bigquery"}, Tables: []string{"*"}},
		},
		destinations: []*Destination{
			{Name: "bigquery", Path: "cloudquery/bigquery", Version: "v3.1.0", Spec: map[string]any{"service_account_key_json": getExpectedApplicationDefaultCredentials()}},
		},
	},
}

func TestLoadSpecs(t *testing.T) {
	for _, tc := range specLoaderTestCases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.envVariables {
				t.Setenv(k, v)
			}
			specReader, err := NewSpecReader(tc.path)
			expectedErr := tc.err()
			if err != nil {
				if err.Error() != expectedErr {
					t.Fatalf("expected error: '%s', got: '%s'", expectedErr, err)
				}
				return
			}
			if expectedErr != "" {
				t.Fatalf("expected error: %s, got nil", expectedErr)
			}

			for _, s := range tc.sources {
				s.SetDefaults()
			}

			for _, d := range tc.destinations {
				d.SetDefaults(0, 0)
			}

			require.Equal(t, tc.sources, specReader.Sources)
			require.Equal(t, tc.destinations, specReader.Destinations)
		})
	}
}

func TestLoadSpecWithAccountNumbers(t *testing.T) {
	t.Setenv("ACCOUNT_ID", "0123456789")
	specReader, err := NewSpecReader([]string{getPath("numbers.yml")})
	if err != nil {
		t.Fatal(err)
	}
	if len(specReader.Sources) != 2 {
		t.Fatalf("got: %d expected: %d", len(specReader.Sources), 2)
	}
	if len(specReader.Destinations) != 1 {
		t.Fatalf("got: %d expected: %d", len(specReader.Destinations), 1)
	}
	if specReader.GetSourceByName("0123456789") == nil {
		t.Fatalf("expected source with account id 0123456789")
	}
	if specReader.GetSourceByName("0123456789").Name != "0123456789" {
		t.Fatalf("got: %s expected: %s", specReader.GetSourceByName("0123456789").Name, "0123456789")
	}
	if specReader.GetDestinationByName("0987654321") == nil {
		t.Fatalf("expected destination with account id 0987654321")
	}
	if specReader.GetDestinationByName("0987654321").Name != "0987654321" {
		t.Fatalf("got: %s expected: %s", specReader.GetDestinationByName("0987654321").Name, "0987654321")
	}
}

func TestExpandFile(t *testing.T) {
	cfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: ${file:./testdata/creds.txt}
		otherstuff: 2
		credentials1: [${file:./testdata/creds.txt}, ${file:./testdata/creds1.txt}]
	`)
	expectedCfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: mytestcreds
		otherstuff: 2
		credentials1: [mytestcreds, anothercredtest]
	`)
	expandedCfg, err := expandFileConfig(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expandedCfg, expectedCfg) {
		t.Fatalf("got: %s expected: %s", expandedCfg, expectedCfg)
	}

	badCfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: ${file:./testdata/creds2.txt}
		otherstuff: 2
	`)
	_, err = expandFileConfig(badCfg)
	if !os.IsNotExist(err) {
		t.Fatalf("expected error: %s, got: %s", os.ErrNotExist, err)
	}
}

func TestExpandEnv(t *testing.T) {
	os.Setenv("TEST_ENV_CREDS", "mytestcreds")
	os.Setenv("TEST_ENV_CREDS2", "anothercredtest")
	cfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: ${TEST_ENV_CREDS}
		otherstuff: 2
		credentials1: [${TEST_ENV_CREDS}, ${TEST_ENV_CREDS2}]
	`)
	expectedCfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: mytestcreds
		otherstuff: 2
		credentials1: [mytestcreds, anothercredtest]
	`)
	expandedCfg, err := expandEnv(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expandedCfg, expectedCfg) {
		t.Fatalf("got: %s expected: %s", expandedCfg, expectedCfg)
	}

	badCfg := []byte(`
kind: source
spec:
	name: test
	version: v1.0.0
	spec:
		credentials: ${TEST_ENV_CREDS1}
		otherstuff: 2
	`)
	_, err = expandEnv(badCfg)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
