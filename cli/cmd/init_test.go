package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	KeyArrowDown = "\033[B"
	KeyEnter     = "\r"
)

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
