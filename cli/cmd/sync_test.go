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
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	configs := []struct {
		name    string
		config  string
		shard   string
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
					CLIVersion:        "development",
					DestinationErrors: 0,
					DestinationName:   "test",
					DestinationPath:   "cloudquery/test",
					Resources:         13,
					SourceName:        "test",
					SourcePath:        "cloudquery/test",
					SourceTables:      []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
				{
					CLIVersion:        "development",
					DestinationErrors: 0,
					DestinationName:   "test",
					DestinationPath:   "cloudquery/test",
					Resources:         13,
					SourceName:        "test2",
					SourcePath:        "cloudquery/test",
					SourceTables:      []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
			},
		},
		{
			name:   "multiple_destinations",
			config: "multiple-destinations.yml",
		},
		{
			name:   "multiple_sources_destinations",
			config: "multiple-sources-destinations.yml",
			summary: []syncSummary{
				{
					CLIVersion:      "development",
					DestinationName: "test-1",
					DestinationPath: "cloudquery/test",
					Resources:       13,
					SourceName:      "test-1",
					SourcePath:      "cloudquery/test",
					SourceTables:    []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
				{
					CLIVersion:      "development",
					DestinationName: "test-2",
					DestinationPath: "cloudquery/test",
					Resources:       13,
					SourceName:      "test-2",
					SourcePath:      "cloudquery/test",
					SourceTables:    []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
			},
		},
		{
			name:   "different_backend_from_destination",
			config: "different-backend-from-destination.yml",
			summary: []syncSummary{
				{
					CLIVersion:      "development",
					DestinationName: "test1",
					DestinationPath: "cloudquery/test",
					Resources:       13,
					SourceName:      "test",
					SourcePath:      "cloudquery/test",
					SourceTables:    []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
			},
		},
		{
			name:   "with_sync_group_id",
			config: "with-sync-group-id.yml",
			summary: []syncSummary{
				{
					CLIVersion:      "development",
					DestinationName: "test1",
					DestinationPath: "cloudquery/test",
					Resources:       13,
					SourceName:      "test",
					SourcePath:      "cloudquery/test",
					SourceTables:    []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
					SyncGroupID:     lo.ToPtr("sync_group_id_test"),
				},
			},
		},
		{
			name:   "with_sync_group_id_and_shard",
			config: "with-sync-group-id.yml",
			shard:  "1/2",
			summary: []syncSummary{
				{
					CLIVersion:      "development",
					DestinationName: "test1",
					DestinationPath: "cloudquery/test",
					// Less resources due to sharding
					Resources:    11,
					SourceName:   "test",
					SourcePath:   "cloudquery/test",
					SourceTables: []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
					SyncGroupID:  lo.ToPtr("sync_group_id_test"),
					ShardNum:     lo.ToPtr(1),
					ShardTotal:   lo.ToPtr(2),
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

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()

			baseArgs := testCommandArgs(t)
			argList := append([]string{"sync", testConfig}, baseArgs...)
			summaryPath := ""
			if len(tc.summary) > 0 {
				tmp := t.TempDir()
				summaryPath = path.Join(tmp, "/test/cloudquery-summary.jsonl")
				argList = append(argList, "--summary-location", summaryPath)
			}

			if tc.shard != "" {
				argList = append(argList, "--shard", tc.shard)
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
				// Ignore random fields or fields that are updated over time
				diff := cmp.Diff(tc.summary, summaries, cmpopts.IgnoreFields(syncSummary{}, "SyncID", "DestinationVersion", "SourceVersion", "SyncTime"))
				for _, s := range summaries {
					assert.NotEmpty(t, s.SyncID)
					assert.NotEmpty(t, s.SyncTime)
					assert.NotEmpty(t, s.DestinationVersion)
					assert.NotEmpty(t, s.SourceVersion)
				}
				require.Empty(t, diff, "unexpected summaries: %v", diff)
			}

			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(baseArgs[3])
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
		})

		t.Run(tc.name+"_no_migrate", func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)

			cmd := NewCmdRoot()
			cmd.SetArgs(append([]string{"sync", testConfig, "--no-migrate"}, testCommandArgs(t)...))
			err := cmd.Execute()
			if tc.err != "" {
				require.Contains(t, err.Error(), tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

type syncSummaryTable struct {
	CLIVersion          string   `json:"cli_version"`
	DestinationErrors   uint64   `json:"destination_errors"`
	DestinationName     string   `json:"destination_name"`
	DestinationPath     string   `json:"destination_path"`
	DestinationVersion  string   `json:"destination_version"`
	DestinationWarnings uint64   `json:"destination_warnings"`
	Resources           uint64   `json:"resources"`
	SourceErrors        uint64   `json:"source_errors"`
	SourcePath          string   `json:"source_path"`
	SourceVersion       string   `json:"source_version"`
	SourceWarnings      uint64   `json:"source_warnings"`
	SourceTables        []string `json:"source_tables"`
	SyncID              string   `json:"sync_id"`
	ShardNum            *int     `json:"shard_num,omitempty"`
	ShardTotal          *int     `json:"shard_total,omitempty"`
	// Internal columns are prefixed with _cq_ in the destination schema (hence in the file destination JSON)
	SyncGroupID *string `json:"_cq_sync_group_id,omitempty"`
	SyncTime    string  `json:"_cq_sync_time"`
	SourceName  string  `json:"_cq_source_name"`
}

func TestSyncWithSummaryTable(t *testing.T) {
	configs := []struct {
		name         string
		config       string
		shard        string
		err          string
		summaryTable []syncSummaryTable
	}{
		{
			name:   "with-destination-summary",
			config: "with-destination-summary.yml",
			summaryTable: []syncSummaryTable{
				{
					CLIVersion:         "development",
					DestinationErrors:  0,
					DestinationName:    "test",
					DestinationPath:    "cloudquery/file",
					DestinationVersion: "v5.2.5",
					Resources:          13,
					SourceName:         "test",
					SourcePath:         "cloudquery/test",
					SourceVersion:      "v4.5.1",
					SourceTables:       []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
				},
			},
		},
		{
			name:   "with-destination-summary-with-sync-group-id-and-shard",
			config: "with-destination-summary-with-sync-group-id-and-shard.yml",
			shard:  "1/2",
			summaryTable: []syncSummaryTable{
				{
					CLIVersion:         "development",
					DestinationErrors:  0,
					DestinationName:    "test",
					DestinationPath:    "cloudquery/file",
					DestinationVersion: "v5.2.5",
					// Less resources due to sharding
					Resources:     11,
					SourceName:    "test_1_2",
					SourcePath:    "cloudquery/test",
					SourceVersion: "v4.5.1",
					SourceTables:  []string{"test_some_table", "test_sub_table", "test_testdata_table", "test_paid_table"},
					SyncGroupID:   lo.ToPtr("sync_group_id_test"),
					ShardNum:      lo.ToPtr(1),
					ShardTotal:    lo.ToPtr(2),
				},
			},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			baseArgs := testCommandArgs(t)
			argList := append([]string{"sync", testConfig}, baseArgs...)

			summaryTablePath := ""
			if len(tc.summaryTable) > 0 {
				datadir := t.TempDir()
				summaryTablePath = path.Join(datadir, "/data/cloudquery_sync_summaries")
				// this is the only way to inject the dynamic output path
				os.Setenv("CQ_FILE_DESTINATION", path.Join(datadir, "/data/{{TABLE}}/{{UUID}}.{{FORMAT}}"))
			}
			if tc.shard != "" {
				argList = append(argList, "--shard", tc.shard)
			}
			cmd.SetArgs(argList)
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
			}
			summaries := []syncSummaryTable{}
			// find all json files in the data directory
			files, err := os.ReadDir(summaryTablePath)
			if err != nil {
				t.Fatalf("failed to read directory %v: %v", summaryTablePath, err)
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				b, err := os.ReadFile(path.Join(summaryTablePath, file.Name()))
				if err != nil {
					t.Fatalf("failed to read file %v: %v", file.Name(), err)
				}
				var v syncSummaryTable
				assert.NoError(t, json.Unmarshal(b, &v))
				summaries = append(summaries, v)
			}

			// Ignore random fields or fields that are updated over time
			diff := cmp.Diff(tc.summaryTable, summaries, cmpopts.IgnoreFields(syncSummaryTable{}, "SyncID", "SyncTime", "DestinationVersion", "SourceVersion"))
			for _, s := range summaries {
				assert.NotEmpty(t, s.SyncID)
				assert.NotEmpty(t, s.SyncTime)
				assert.NotEmpty(t, s.DestinationVersion)
				assert.NotEmpty(t, s.SourceVersion)
			}

			require.Empty(t, diff, "unexpected summaries: %v", diff)

			// have to ignore SyncID because it's random and plugin versions since we update those frequently using an automated process
			// also ignore SyncTime because it's a timestamp
			for _, s := range summaries {
				assert.NotEmpty(t, s.SyncID)
			}
		})
	}
}

func TestSyncCqDir(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", "sync-success-sourcev1-destv0.yml")

	cmd := NewCmdRoot()
	baseArgs := testCommandArgs(t)
	cmd.SetArgs(append([]string{"sync", testConfig}, baseArgs...))
	err := cmd.Execute()
	require.NoError(t, err)

	// check that destination plugin was downloaded to the cache using --cq-dir
	p := path.Join(baseArgs[1], "plugins")
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

	t.Setenv("CQ_CLOUD", "1")
	t.Setenv("_CQ_TEAM_NAME", "cloudquery")
	t.Setenv("_CQ_SYNC_NAME", "test_sync")
	t.Setenv("_CQ_SYNC_RUN_ID", uuid.Must(uuid.NewUUID()).String())
	t.Setenv("__SOURCE_TEST__TEST_KEY", "test_value")
	t.Setenv("NOT_TEST_ENV", "should_not_be_visible_to_plugin")

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			testConfig := path.Join(currentDir, "testdata", tc.config)
			cmd := NewCmdRoot()
			cmd.SetArgs(append([]string{"sync", testConfig}, testCommandArgs(t)...))
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
