package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	configs := []struct {
		name    string
		config  string
		err     string
		summary []syncSummary
	}{
		{
			name:   "sync_success_sourcev1_destv0",
			config: "sync-success-sourcev1-destv0.yml",
		},
		{
			name:   "multiple_sources",
			config: "multiple-sources.yml",
			summary: []syncSummary{
				{
					CliVersion:         "development",
					DestinationErrors:  0,
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationErrors:  0,
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
			},
		},
		{
			name:   "multiple_destinations",
			config: "multiple-destinations.yml",
			summary: []syncSummary{

				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
			},
		},
		{
			name:   "multiple_sources_destinations",
			config: "multiple-sources-destinations.yml",
			summary: []syncSummary{
				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test-1",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test-1",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test-2",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test-2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
			},
		},
		{
			name:   "different_backend_from_destination",
			config: "different-backend-from-destination.yml",
			summary: []syncSummary{
				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test-1",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test-1",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test-2",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test-2",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
				{
					CliVersion:         "development",
					DestinationName:    "test1",
					DestinationPath:    "cloudquery/test",
					DestinationVersion: "v2.2.14",
					Resources:          12,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v3.1.15",
				},
			},
		},
		{
			name:   "should fail with missing path error when path is missing",
			config: "sync-missing-path-error.yml",
			err:    "Error: failed to validate destination test: path is required",
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	cqDir := t.TempDir()
	defer os.RemoveAll(cqDir)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			defer CloseLogFile()
			testConfig := path.Join(currentDir, "testdata", tc.config)
			logFileName := path.Join(cqDir, "cloudquery.log")
			cmd := NewCmdRoot()

			argList := []string{"sync", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName}
			summaryPath := ""
			if len(tc.summary) > 0 {
				summaryPath = path.Join(cqDir, "/test/cloudquery-summary.jsonl")
				argList = append(argList, "--summary-location", summaryPath)
			}

			cmd.SetArgs(argList)
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
			}

			if len(tc.summary) > 0 {
				summaries := readSummaries(t, summaryPath)
				// have to ignore SyncID because it's random
				diff := cmp.Diff(tc.summary, summaries, cmpopts.IgnoreFields(syncSummary{}, "SyncID"))
				if diff != "" {
					t.Errorf("unexpected summaries: %v", diff)
				}
			}

			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(path.Join(cqDir, "cloudquery.log"))
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
		})

		t.Run(tc.name+"_no_migrate", func(t *testing.T) {
			defer CloseLogFile()
			testConfig := path.Join(currentDir, "testdata", tc.config)
			logFileName := path.Join(cqDir, "cloudquery.log")

			cmd := NewCmdRoot()
			cmd.SetArgs([]string{"sync", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName, "--no-migrate"})
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSyncCqDir(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", "sync-success-sourcev1-destv0.yml")
	cqDir := t.TempDir()
	defer os.RemoveAll(cqDir)
	logFileName := path.Join(cqDir, "cloudquery.log")

	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"sync", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
	err := cmd.Execute()
	defer CloseLogFile()
	require.NoError(t, err)

	// check that destination plugin was downloaded to the cache using --cq-dir
	p := path.Join(cqDir, "plugins")
	files, err := os.ReadDir(p)
	if err != nil {
		t.Fatalf("failed to read cache directory %v: %v", p, err)
	}
	require.NotEmpty(t, files, "destination plugin not downloaded to cache")
}

func TestFindMaxCommonVersion(t *testing.T) {
	cases := []struct {
		name       string
		givePlugin []int
		giveCLI    []int
		want       int
	}{
		{name: "support_less", givePlugin: []int{1, 2, 3}, giveCLI: []int{1, 2}, want: 2},
		{name: "support_same", givePlugin: []int{1, 2, 3}, giveCLI: []int{1, 2, 3}, want: 3},
		{name: "support_more", givePlugin: []int{1, 2, 3}, giveCLI: []int{2, 3, 4}, want: 3},
		{name: "support_only_lower", givePlugin: []int{3, 4, 5}, giveCLI: []int{6, 7}, want: -1},
		{name: "support_only_higher", givePlugin: []int{3, 4, 5}, giveCLI: []int{1, 2}, want: -2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaxCommonVersion(tc.givePlugin, tc.giveCLI)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSync_IsolatedPluginEnvironmentsInCloud(t *testing.T) {
	configs := []struct {
		name   string
		config string
		err    string
	}{
		{
			name:   "source-with-env",
			config: "source-with-env.yml",
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	cqDir := t.TempDir()
	t.Setenv("CLOUDQUERY_API_KEY", "cqsr_123")
	t.Setenv("_CQ_TEAM_NAME", "test_team")
	t.Setenv("_CQ_SYNC_NAME", "test_sync")
	t.Setenv("_CQ_SYNC_RUN_ID", uuid.Must(uuid.NewUUID()).String())
	t.Setenv("__SOURCE_TEST__TEST_KEY", "test_value")
	t.Setenv("NOT_TEST_ENV", "should_not_be_visible_to_plugin")

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			cmd.SetArgs([]string{"sync", testConfig, "--cq-dir", cqDir})
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func readSummaries(t *testing.T, filename string) []syncSummary {
	p, err := os.ReadFile(filename)
	assert.NoError(t, err)

	lines := bytes.Split(p, []byte{'\n'})
	summaries := make([]syncSummary, len(lines))
	for i, line := range lines {
		if len(line) == 0 {
			summaries = slices.Delete(summaries, i, i+1)
			continue
		}
		var v syncSummary
		assert.NoError(t, json.Unmarshal(line, &v))
		summaries[i] = v
	}
	return summaries
}
