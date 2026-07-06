package cmd

import (
	_ "embed"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"testing"

	cqapi "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	KeyArrowDown = "\033[B"
	KeyEnter     = "\r"
)

//go:embed testdata/init/aws_example.md
var awsExample string

//go:embed testdata/init/postgresql_example.md
var postgresqlExample string

func TestInit(t *testing.T) {
	configs := []struct {
		name         string
		source       string
		destination  string
		yes          bool
		expectedPath string
		expectedSpec func(*testing.T, *specs.SpecReader)
	}{
		{
			name:         "accept defaults with --yes flag is set",
			yes:          true,
			expectedPath: "aws_to_postgresql.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "aws", specs.Sources[0].Name)
				require.Equal(t, "cloudquery/aws", specs.Sources[0].Path)
				require.Equal(t, "postgresql", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/postgresql", specs.Destinations[0].Path)
			},
		},
		{
			name:         "uses source and destination flags",
			source:       "cloudquery/gcp",
			destination:  "cloudquery/sqlite",
			expectedPath: "gcp_to_sqlite.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "gcp", specs.Sources[0].Name)
				require.Equal(t, "cloudquery/gcp", specs.Sources[0].Path)
				require.Equal(t, "sqlite", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/sqlite", specs.Destinations[0].Path)
			},
		},
		{
			name:         "automatically prepends cloudquery/ to source and destination",
			source:       "azure",
			destination:  "bigquery",
			expectedPath: "azure_to_bigquery.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "azure", specs.Sources[0].Name)
				require.Equal(t, "cloudquery/azure", specs.Sources[0].Path)
				require.Equal(t, "bigquery", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/bigquery", specs.Destinations[0].Path)
			},
		},
		{
			name:         "can generate spec file for community plugins",
			source:       "hermanschaaf/chess-com",
			destination:  "bigquery",
			expectedPath: "chess-com_to_bigquery.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "chess-com", specs.Sources[0].Name)
				require.Equal(t, "hermanschaaf/chess-com", specs.Sources[0].Path)
				require.Equal(t, "bigquery", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/bigquery", specs.Destinations[0].Path)
			},
		},
		{
			name:         "generate spec file from source list prompt",
			destination:  "bigquery",
			expectedPath: "azure_to_bigquery.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "azure", specs.Sources[0].Name)
				require.Equal(t, "cloudquery/azure", specs.Sources[0].Path)
				require.Equal(t, "bigquery", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/bigquery", specs.Destinations[0].Path)
			},
		},
		{
			name:         "generate spec file from destination list prompt",
			source:       "gcp",
			expectedPath: "gcp_to_s3.yaml",
			expectedSpec: func(t *testing.T, specs *specs.SpecReader) {
				require.Len(t, specs.Sources, 1)
				require.Len(t, specs.Destinations, 1)
				require.Equal(t, "gcp", specs.Sources[0].Name)
				require.Equal(t, "cloudquery/gcp", specs.Sources[0].Path)
				require.Equal(t, "s3", specs.Destinations[0].Name)
				require.Equal(t, "cloudquery/s3", specs.Destinations[0].Path)
			},
		},
	}

	// Set environment variables so the spec reader doesn't fail
	t.Setenv("POSTGRESQL_CONNECTION_STRING", "test")
	t.Setenv("PROJECT_ID", "test")
	t.Setenv("DATASET_ID", "test")
	for _, tc := range configs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cmd := NewCmdRoot()
			baseArgs := testCommandArgs(t)
			initArgs := []string{"init"}
			cqDir := baseArgs[1]
			expectedSpecPath := path.Join(cqDir, tc.expectedPath)
			initArgs = append(initArgs, "--spec-path", expectedSpecPath)
			if tc.source != "" {
				initArgs = append(initArgs, "--source", tc.source)
			}
			if tc.destination != "" {
				initArgs = append(initArgs, "--destination", tc.destination)
			}
			if tc.yes {
				initArgs = append(initArgs, "--yes")
			}
			// Add --disable-ai flag to prevent AI command from being triggered
			initArgs = append(initArgs, "--disable-ai")
			cmd.SetArgs(append(initArgs, baseArgs...))

			if !tc.yes && (tc.source == "" || tc.destination == "") {
				// skip the prompt tests on Windows
				if runtime.GOOS == "windows" {
					t.Skip("Skipping prompt tests on Windows")
					return
				}
				oldStdin := os.Stdin
				r, w, err := os.Pipe()
				require.NoError(t, err)
				defer func() {
					r.Close()
					w.Close()
					os.Stdin = oldStdin
				}()
				if tc.source == "" {
					// Select second source on the list
					_, err = w.WriteString(KeyArrowDown + KeyEnter)
					require.NoError(t, err)
				}
				if tc.destination == "" {
					// Select third destination on the list
					_, err = w.WriteString(KeyArrowDown + KeyArrowDown + KeyEnter)
					require.NoError(t, err)
				}
				os.Stdin = r
			}
			err := cmd.Execute()
			assert.NoError(t, err)
			require.FileExists(t, expectedSpecPath)

			specReader, err := specs.NewSpecReader([]string{expectedSpecPath})
			require.NoError(t, err)
			tc.expectedSpec(t, specReader)
		})
	}
}

func Test_configForSourcePlugin(t *testing.T) {
	testCases := []struct {
		name         string
		source       cqapi.ListPlugin
		version      *cqapi.PluginVersionDetails
		expectedSpec func(*testing.T, *specs.SpecReader)
	}{
		{
			name: "without example config",
			source: cqapi.ListPlugin{
				Name:          "aws",
				TeamName:      "cloudquery",
				Kind:          cqapi.PluginKindSource,
				LatestVersion: lo.ToPtr("v27.0.0"),
			},
			version: &cqapi.PluginVersionDetails{
				Name: "v27.0.0",
			},
			expectedSpec: func(t *testing.T, sr *specs.SpecReader) {
				require.Len(t, sr.Sources, 1)
				require.Equal(t, "aws", sr.Sources[0].Name)
				require.Equal(t, "cloudquery/aws", sr.Sources[0].Path)
				require.Equal(t, "v27.0.0", sr.Sources[0].Version)
				require.Equal(t, []string{"*"}, sr.Sources[0].Tables)
				require.Equal(t, []string{"DESTINATION_NAME"}, sr.Sources[0].Destinations)
				require.Len(t, sr.Sources[0].Spec, 0)
			},
		},
		{
			name: "with example config",
			source: cqapi.ListPlugin{
				Name:          "aws",
				TeamName:      "cloudquery",
				Kind:          cqapi.PluginKindSource,
				LatestVersion: lo.ToPtr("v27.0.0"),
			},
			version: &cqapi.PluginVersionDetails{
				Name:          "v27.0.0",
				ExampleConfig: awsExample,
			},
			expectedSpec: func(t *testing.T, sr *specs.SpecReader) {
				require.Len(t, sr.Sources, 1)
				require.Equal(t, "aws", sr.Sources[0].Name)
				require.Equal(t, "cloudquery/aws", sr.Sources[0].Path)
				require.Equal(t, "v27.0.0", sr.Sources[0].Version)
				require.Equal(t, []string{"aws_ec2_instances"}, sr.Sources[0].Tables)
				require.Equal(t, []string{"DESTINATION_NAME"}, sr.Sources[0].Destinations)
				require.Len(t, sr.Sources[0].Spec, 1)
				require.Equal(t, json.Number("100"), sr.Sources[0].Spec["concurrency"])
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			specPath := path.Join(tempDir, "spec.yaml")
			spec := configForSourcePlugin(tc.source, tc.version)
			require.NoError(t, os.WriteFile(specPath, []byte(spec), 0644))

			specReader, err := specs.NewRelaxedSpecReader([]string{specPath})
			require.NoError(t, err)
			tc.expectedSpec(t, specReader)
		})
	}
}

func Test_configForDestinationPlugin(t *testing.T) {
	testCases := []struct {
		name         string
		destination  cqapi.ListPlugin
		version      *cqapi.PluginVersionDetails
		expectedSpec func(*testing.T, *specs.SpecReader)
	}{
		{
			name: "without example config",
			destination: cqapi.ListPlugin{
				Name:          "postgresql",
				TeamName:      "cloudquery",
				Kind:          cqapi.PluginKindDestination,
				LatestVersion: lo.ToPtr("v8.0.0"),
			},
			version: &cqapi.PluginVersionDetails{
				Name: "v8.0.0",
			},
			expectedSpec: func(t *testing.T, sr *specs.SpecReader) {
				require.Len(t, sr.Destinations, 1)
				require.Equal(t, "postgresql", sr.Destinations[0].Name)
				require.Equal(t, "cloudquery/postgresql", sr.Destinations[0].Path)
				require.Equal(t, "v8.0.0", sr.Destinations[0].Version)
				require.Len(t, sr.Destinations[0].Spec, 0)
			},
		},
		{
			name: "with example config",
			destination: cqapi.ListPlugin{
				Name:          "postgresql",
				TeamName:      "cloudquery",
				Kind:          cqapi.PluginKindDestination,
				LatestVersion: lo.ToPtr("v8.0.0"),
			},
			version: &cqapi.PluginVersionDetails{
				Name:          "v8.0.0",
				ExampleConfig: postgresqlExample,
			},
			expectedSpec: func(t *testing.T, sr *specs.SpecReader) {
				require.Len(t, sr.Destinations, 1)
				require.Equal(t, "postgresql", sr.Destinations[0].Name)
				require.Equal(t, "cloudquery/postgresql", sr.Destinations[0].Path)
				require.Equal(t, "v8.0.0", sr.Destinations[0].Version)
				require.Len(t, sr.Destinations[0].Spec, 1)
				require.Equal(t, "test", sr.Destinations[0].Spec["connection_string"])
			},
		},
	}

	// Set environment variables so the spec reader doesn't fail
	t.Setenv("POSTGRESQL_CONNECTION_STRING", "test")
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			specPath := path.Join(tempDir, "spec.yaml")
			spec := configForDestinationPlugin(tc.destination, tc.version)
			require.NoError(t, os.WriteFile(specPath, []byte(spec), 0644))

			specReader, err := specs.NewRelaxedSpecReader([]string{specPath})
			require.NoError(t, err)
			tc.expectedSpec(t, specReader)
		})
	}
}

func Test_withRecommendedTables(t *testing.T) {
	rec := []string{"aws_ec2_instances", "aws_s3_buckets"}
	base := "kind: source\nspec:\n  name: aws\n  path: cloudquery/aws\n  version: v1.0.0\n  destinations: [platform]\n"

	// Assert on the parsed result rather than exact whitespace — the yaml.Node
	// re-marshal may reformat, but the tables selection must be exactly `rec`.
	parseTables := func(t *testing.T, spec string) []string {
		t.Helper()
		p := path.Join(t.TempDir(), "spec.yaml")
		require.NoError(t, os.WriteFile(p, []byte(spec), 0644))
		sr, err := specs.NewRelaxedSpecReader([]string{p})
		require.NoError(t, err)
		require.Len(t, sr.Sources, 1)
		return sr.Sources[0].Tables
	}

	t.Run("overrides the wildcard", func(t *testing.T) {
		require.Equal(t, rec, parseTables(t, withRecommendedTables(base+"  tables: ['*']\n", rec)))
	})

	t.Run("overrides a specific inline example list (e.g. AWS hub example)", func(t *testing.T) {
		got := withRecommendedTables(base+`  tables: ["aws_s3_buckets"]`+"\n", rec)
		require.Equal(t, rec, parseTables(t, got))
		require.NotContains(t, got, `["aws_s3_buckets"]`)
	})

	t.Run("overrides a block-style list", func(t *testing.T) {
		got := withRecommendedTables(base+"  tables:\n    - aws_s3_buckets\n    - aws_iam_users\n", rec)
		require.Equal(t, rec, parseTables(t, got))
		require.NotContains(t, got, "aws_iam_users", "old block items removed")
	})

	t.Run("preserves comments and other fields", func(t *testing.T) {
		in := "kind: source\nspec:\n  name: aws\n  # fill in your credentials below\n  path: cloudquery/aws\n  version: v1.0.0\n" +
			`  tables: ["aws_s3_buckets"]` + "\n  destinations: [platform]\n"
		got := withRecommendedTables(in, rec)
		require.Contains(t, got, "# fill in your credentials below", "comments are preserved")
		require.Equal(t, rec, parseTables(t, got))
	})

	t.Run("adds tables when the example config has none", func(t *testing.T) {
		// base has no `tables:` key; the recommended set should be added.
		got := withRecommendedTables(base, rec)
		require.Equal(t, rec, parseTables(t, got))
	})

	t.Run("no-op on empty recommendation", func(t *testing.T) {
		spec := base + `  tables: ["aws_s3_buckets"]` + "\n"
		require.Equal(t, spec, withRecommendedTables(spec, nil))
	})

	t.Run("no source spec mapping leaves the spec unchanged", func(t *testing.T) {
		spec := "kind: source\n"
		require.Equal(t, spec, withRecommendedTables(spec, rec))
	})

	t.Run("multi-doc: edits the source, keeps the other doc", func(t *testing.T) {
		// A source doc + a companion destination doc must both survive; only the
		// source's tables change, the destination doc is left intact.
		in := base + `  tables: ["aws_s3_buckets"]` + "\n---\nkind: destination\nspec:\n  name: postgresql\n  path: cloudquery/postgresql\n  version: v8.0.0\n"
		got := withRecommendedTables(in, rec)
		require.Contains(t, got, "kind: destination", "the destination doc is not truncated")
		require.Contains(t, got, "postgresql")
		require.NotContains(t, got, `["aws_s3_buckets"]`)

		p := path.Join(t.TempDir(), "spec.yaml")
		require.NoError(t, os.WriteFile(p, []byte(got), 0644))
		sr, err := specs.NewRelaxedSpecReader([]string{p})
		require.NoError(t, err)
		require.Len(t, sr.Sources, 1)
		require.Len(t, sr.Destinations, 1, "destination doc preserved")
		require.Equal(t, rec, sr.Sources[0].Tables)
	})
}
