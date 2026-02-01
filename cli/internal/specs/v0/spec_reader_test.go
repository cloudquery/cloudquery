package specs

import (
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
	"time"

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

func getexpectedCertificate() string {
	cert := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAupzbCx3FoPinoH2ULfQLLk47vkKkjY7na97xFbtAvY+/Aomc
YLpxEl3GkNGRMlMeaDP2qxhpBeH85xtR350CorBgN4lnRmhPCgIZ0wxiMbwzQ7HO
vVa+Ha9RrN89vx3/SlW73WJyPd0WXw3WQCnjThVtRicOAIRUnkwPvLtHVmteo6LI
IW0LpKxZWi0pJrET3ka1gQDNRh5/b1NGfP2Ecv2jqtT3YG4YkPx3ll21ZhW4BurF
1drQXSwrowAk4tiLH0g+IYakTNEWmobtvXq+DdNFQzHPYa2Nm9K1ExjI2N12U4Jf
JRfPSGqfQIOTiaZPOQl/H0mFZRElN+2eBxbMDwIDAQABAoIBADdqW/ulmCnwnSqi
EA5DYcya68/Yj3AAB0X3uuTmqdeA58pzne31f51iHpSjvvfQSf/MqovtYEagcM8a
REpgoEc6lB/53CLC1/HTZOLQ0xoM1rZcB1Yfe65qARmSY44s9MIYyoR39w/a5wlM
HRsJtVfbMgt6joRlx5EIakXz4T/OeWkEvefHRn93ykiNPT3SSvyrFZ4B2OIGla/G
LxA9mRsGNy0HSymhW671aXZDDZ97zFzrLYvf2hOhKa5xWLJKs13xWycI6ofpnwQ9
WXPtbBclqVz0IIrYipq5XzJ8smEoGu0ZOu1dO8DoypPNITvlZYUjX9xmoE9CAiAp
IL+ryAkCgYEA68nF1VZx+hH66oEbb5APhde/0Qpdrbnaq6s/hhsVBseTI10vi5V7
dyCliNhnCU3JAIVi0v+kTXPJhs3wLmDhSrFLWxkPt1HzHw+m1eRP1rQH+M7CSrVF
yQV5KYrs3mtIO4AYoFo+bO0tmn2DVKg6FO6leiy+nLAuPk1uMWvAfZcCgYEAypv1
IHUU4vv7XoufrKTWTji/piGTgQxft8Tk+IaFTBphY6gpFCXB4jm97kriUVL0htN7
JLkpzGrdhP9ufKwAWH/5Jlna6OQpNAVIhziYoQjBKBoL/LtjCHdHqbj881QrVLyk
ITeecgO6x36zpf09/Dw8+k1VOVs62IXq5tF8ZEkCgYAFpC43jHntocB/G9esM6Yr
bZ4JQlY7cdbphI9ghgVaxCuhDPm2PT1W/FD5lTPh5RqKCKb0pWko8TxBHWxBr8+0
GcnTxCW8HRnUBGvZcjz2xhfqvAeqAexJgvgDJm/EYoy337i3HXGg6YvNxnL984hw
N8V9xtRIq25vzypzxEA2wQKBgQC+4HrpHySUS2yfv86oaYr0moYDT3KVi0DJ8pb8
hE8kSV4i8xPwRToJlPiYfLgGga6ZLre++yqjyLH1UGeY0LpqpfXl6ZVQ/1LKDYgs
zGcOnx7KVu+gJDHCkg1TmlHENDG2XRoLqUh+hYD73SQGZzR0Y5PXA/AcXxRrVI7e
8dDM8QKBgQCbH5XBA5R8FhLxY3di8AjqgnfYlrTemouRcDratVDb47pWq453OkZQ
5JSK6iUOQP+ZA4HXitBmErQ3D1ALCenE6Yk3BXBi7PaZHpYA4h4xnE+3qitGC5fY
mq67p1yluTqgCsU8qshVPBpynF9CANc5AOGoh/MAff6U2y2KXUbuUw==
-----END RSA PRIVATE KEY-----`
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(cert, "\n", "\r\n")
	}
	return cert
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

var boolTrue = true
var boolFalse = false
var specLoaderTestCases = []specLoaderTestCase{
	{
		name: "success",
		path: []string{getPath("gcp.yml"), getPath("dir")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:     "gcp",
					Path:     "cloudquery/gcp",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresqlv2"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:     "aws",
					Path:     "cloudquery/aws",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:     "postgresqlv2",
					Path:     "cloudquery/postgresql",
					Version:  "v1.0.0",
					Registry: RegistryGRPC,
				},
				WriteMode: WriteModeOverwrite,
				Spec:      map[string]any{"credentials": "mytestcreds"},
			},
			{
				Metadata: Metadata{
					Name:     "postgresql",
					Path:     "cloudquery/postgresql",
					Version:  "v1.0.0",
					Registry: RegistryGRPC,
				},
				WriteMode: WriteModeOverwrite,
			},
		},
	},
	{
		name: "success_yaml_extension",
		path: []string{getPath("gcp.yml"), getPath("dir_yaml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:     "gcp",
					Path:     "cloudquery/gcp",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresqlv2"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:     "aws",
					Path:     "cloudquery/aws",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:     "postgresqlv2",
					Path:     "cloudquery/postgresql",
					Version:  "v1.0.0",
					Registry: RegistryGRPC,
				},
				WriteMode: WriteModeOverwrite,
				Spec:      map[string]any{"credentials": "mytestcreds"},
			},
			{
				Metadata: Metadata{
					Name:     "postgresql",
					Path:     "cloudquery/postgresql",
					Version:  "v1.0.0",
					Registry: RegistryGRPC,
				},
				WriteMode: WriteModeOverwrite,
			},
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
		name: "sync_group_id_append",
		path: []string{getPath("sync_group_id_append.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:     "gcp",
					Path:     "cloudquery/gcp",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:    "postgresql",
					Path:    "cloudquery/postgresql",
					Version: "v1.0.0",
				},
				WriteMode:   WriteModeAppend,
				SyncGroupId: "{{YEAR}}-{{MONTH}}-{{DAY}}-{{HOUR}}-{{MINUTE}}",
			},
		},
	},
	{
		name: "sync_group_id_default",
		path: []string{getPath("sync_group_id_default.yml")},
		err: func() string {
			return "destination postgresql: sync_group_id is not supported with write_mode: overwrite-delete-stale"
		},
	},
	{
		name: "sync_group_id_overwrite_delete_stale",
		path: []string{getPath("sync_group_id_overwrite_delete_stale.yml")},
		err: func() string {
			return "destination postgresql: sync_group_id is not supported with write_mode: overwrite-delete-stale"
		},
	},
	{
		name: "sync_group_id_overwrite",
		path: []string{getPath("sync_group_id_overwrite.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:     "gcp",
					Path:     "cloudquery/gcp",
					Version:  "v1.0.0",
					Registry: RegistryLocal,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:    "postgresql",
					Path:    "cloudquery/postgresql",
					Version: "v1.0.0",
				},
				WriteMode:   WriteModeOverwrite,
				SyncGroupId: "{{YEAR}}-{{MONTH}}-{{DAY}}-{{HOUR}}-{{MINUTE}}",
			},
		},
	},
	{
		name: "multiple sources success",
		path: []string{getPath("multiple_sources.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "aws",
					Path:             "cloudquery/aws",
					Version:          "v4.6.1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:             "azure",
					Path:             "cloudquery/azure",
					Version:          "v1.3.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.6.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"connection_string": "postgresql://postgres:pass@localhost:5432/postgres"},
			},
		},
	},
	{
		name: "environment variables",
		path: []string{getPath("env_variables.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "aws",
					Path:             "cloudquery/aws",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:             "azure",
					Path:             "cloudquery/azure",
					Version:          "v1.3.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql", "postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.6.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"connection_string": "postgresql://localhost:5432/cloudquery?sslmode=disable", "version": "#v1"},
			},
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
			{
				Metadata: Metadata{
					Name:     "aws",
					Path:     "cloudquery/aws",
					Version:  "v1",
					Registry: RegistryCloudQuery,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:     "azure",
					Path:     "cloudquery/azure",
					Version:  "v1.3.3",
					Registry: RegistryCloudQuery,
				},
				Destinations: []string{"postgresql", "postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:     "postgresql",
					Path:     "cloudquery/postgresql",
					Version:  "v1.6.3",
					Registry: RegistryCloudQuery,
				},
				Spec: map[string]any{},
			},
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
			{
				Metadata: Metadata{
					Name:             "test",
					Path:             "cloudquery/test",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"custom_version": "#v1"},
			},
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
			{
				Metadata: Metadata{
					Name:     "test",
					Path:     "cloudquery/test",
					Version:  "v1",
					Registry: RegistryCloudQuery,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:     "postgresql",
					Path:     "cloudquery/postgresql",
					Version:  "v1",
					Registry: RegistryCloudQuery,
				},
				Spec: map[string]any{},
			},
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
			{
				Metadata: Metadata{
					Name:             "0123456789",
					Path:             "cloudquery/aws",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"0987654321"},
				Tables:       []string{"test"},
			},
			{
				Metadata: Metadata{
					Name:             "012345",
					Path:             "cloudquery/aws",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"0987654321"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "0987654321",
					Path:             "cloudquery/postgresql",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"connection_string": "postgresql://localhost:5432/cloudquery?sslmode=disable"},
			},
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
			{
				Metadata: Metadata{
					Name:             "gcp",
					Path:             "cloudquery/gcp",
					Registry:         RegistryCloudQuery,
					Version:          "v1.0.0",
					registryInferred: true,
				},
				Destinations: []string{"bigquery"},
				Tables:       []string{"*"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "bigquery",
					Path:             "cloudquery/bigquery",
					Registry:         RegistryCloudQuery,
					Version:          "v3.1.0",
					registryInferred: true,
				},
				Spec: map[string]any{"service_account_key_json": getExpectedApplicationDefaultCredentials()},
			},
		},
	},
	{
		name: "success importing JSON file using YAML pipe operator",
		path: []string{getPath("json_file_yaml_pipe.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "gcp",
					Path:             "cloudquery/gcp",
					Registry:         RegistryCloudQuery,
					Version:          "v1.0.0",
					registryInferred: true,
				},
				Destinations: []string{"bigquery"},
				Tables:       []string{"*"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "bigquery",
					Path:             "cloudquery/bigquery",
					Registry:         RegistryCloudQuery,
					Version:          "v3.1.0",
					registryInferred: true,
				},
				Spec: map[string]any{"service_account_key_json": getExpectedApplicationDefaultCredentials()},
			},
		},
	},
	{
		name: "skip_dependent_tables defaults to true",
		path: []string{getPath("skip_dependent_tables/default.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "gcp",
					Path:             "cloudquery/gcp",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				SkipDependentTables: &boolTrue,
				Destinations:        []string{"postgresql"},
				Tables:              []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				WriteMode: WriteModeOverwriteDeleteStale,
			},
		},
	},
	{
		name: "skip_dependent_tables when set to true",
		path: []string{getPath("skip_dependent_tables/true.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "gcp",
					Path:             "cloudquery/gcp",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				SkipDependentTables: &boolTrue,
				Destinations:        []string{"postgresql"},
				Tables:              []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				WriteMode: WriteModeOverwriteDeleteStale,
			},
		},
	},
	{
		name: "skip_dependent_tables when set to false",
		path: []string{getPath("skip_dependent_tables/false.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "gcp",
					Path:             "cloudquery/gcp",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				SkipDependentTables: &boolFalse,
				Destinations:        []string{"postgresql"},
				Tables:              []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				WriteMode: WriteModeOverwriteDeleteStale,
			},
		},
	},
	{
		name: "Time substitution",
		path: []string{getPath("time_substitution.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "aws",
					Path:             "cloudquery/aws",
					Version:          "v1.3.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
				Spec:         map[string]any{"table_options": map[string]any{"aws_cloudtrail_events": map[string]any{"lookup_events": []any{map[string]any{"start_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)}}}}},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.6.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{},
			},
		},
	},
	{
		name: "Time substitution with formatting",
		path: []string{getPath("time_substitution_with_format.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "aws",
					Path:             "cloudquery/aws",
					Version:          "v1.3.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
				Spec:         map[string]any{"table_options": map[string]any{"aws_cloudtrail_events": map[string]any{"lookup_events": []any{map[string]any{"start_time": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.DateOnly)}}}}},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1.6.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{},
			},
		},
	},
	{
		name: "Time substitution with error",
		path: []string{getPath("time_substitution_with_error.yml")},
		err: func() string {
			return "failed to expand time variable in file testdata/time_substitution_with_error.yml (section 1): failed to substitute time\ninvalid time format: brkn123"
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "aws",
					Path:             "cloudquery/aws",
					Version:          "v1.3.3",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"postgresql"},
				Tables:       []string{"test"},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "postgresql",
					Path:             "cloudquery/postgresql",
					Version:          "v1",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"custom_version": "#v1"},
			},
		},
	},
	{
		name: "success importing PEM file",
		path: []string{getPath("pem_file.yml")},
		err: func() string {
			return ""
		},
		sources: []*Source{
			{
				Metadata: Metadata{
					Name:             "test-source",
					Path:             "cloudquery/test",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Destinations: []string{"test-destination"},
				Tables:       []string{"*"},
				Spec:         map[string]any{"certificate": getexpectedCertificate()},
			},
		},
		destinations: []*Destination{
			{
				Metadata: Metadata{
					Name:             "test-destination",
					Path:             "cloudquery/postgresql",
					Version:          "v1.0.0",
					Registry:         RegistryCloudQuery,
					registryInferred: true,
				},
				Spec: map[string]any{"connection_string": "postgresql://localhost:5432/test"},
			},
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
				d.SetDefaults()
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
		t.Fatal("expected source with account id 0123456789")
	}
	if specReader.GetSourceByName("0123456789").Name != "0123456789" {
		t.Fatalf("got: %s expected: %s", specReader.GetSourceByName("0123456789").Name, "0123456789")
	}
	if specReader.GetDestinationByName("0987654321") == nil {
		t.Fatal("expected destination with account id 0987654321")
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

func TestShouldEscapeFileContent(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected bool
	}{
		{
			name:     "single line content should not be escaped",
			content:  []byte("single line"),
			expected: false,
		},
		{
			name:     "valid JSON should be escaped",
			content:  []byte(`{"key": "value"}`),
			expected: false,
		},
		{
			name:     "multiline JSON should be escaped",
			content:  []byte("{\n  \"key\": \"value\"\n}"),
			expected: true,
		},
		{
			name: "valid PEM certificate should NOT be escaped",
			content: []byte(`-----BEGIN CERTIFICATE-----
		MIIDXTCCAkWgAwIBAgIJAKJ5cIv5zJ5OMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
		BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
		aWRnaXRzIFB0eSBMdGQwHhcNMTkwNjEyMDAwMDAwWhcNMjkwNjA5MDAwMDAwWjBF
		MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
		ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
		CgKCAQEAw8rL0rFt1pFvL1pKvL0qL6pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL
		1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFv
		L1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pF
		vL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p
		FvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1
		pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL
		1wIDAQABo1AwTjAdBgNVHQ4EFgQUj1pFvL1pFvL1pFvL1pFvL1pFvL0wHwYDVR0j
		BBgwFoAUj1pFvL1pFvL1pFvL1pFvL1pFvL0wDAYDVR0TBAUwAwEB/zANBgkqhkiG
		9w0BAQsFAAOCAQEAq1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p
		FvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1
		pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL
		1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFv
		L1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pF
		vL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p
		FvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFg==
		-----END CERTIFICATE-----
		`),
			expected: false,
		},
		{
			name:     "valid PEM private key should be escaped",
			content:  []byte(`-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWgAwIBAgIJAKJ5cIv5zJ5OMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTkwNjEyMDAwMDAwWhcNMjkwNjA5MDAwMDAwWjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAw8rL0rFt1pFvL1pKvL0qL6pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL\n1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFv\nL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pF\nvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p\nFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1\npFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL\n1wIDAQABo1AwTjAdBgNVHQ4EFgQUj1pFvL1pFvL1pFvL1pFvL1pFvL0wHwYDVR0j\nBBgwFoAUj1pFvL1pFvL1pFvL1pFvL1pFvL0wDAYDVR0TBAUwAwEB/zANBgkqhkiG\n9w0BAQsFAAOCAQEAq1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p\nFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1\npFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL\n1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFv\nL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pF\nvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1p\nFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFvL1pFg==\n-----END CERTIFICATE-----\n`),
			expected: false,
		},
		{
			name:     "multiline non-JSON non-PEM should be escaped",
			content:  []byte("line1\nline2\nline3"),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := shouldEscapeFileContent(tc.content)
			if result != tc.expected {
				t.Errorf("shouldEscapeFileContent() = %v, expected %v", result, tc.expected)
			}
		})
	}
}
