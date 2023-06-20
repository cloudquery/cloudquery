package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	configs := []struct {
		name   string
		config string
		err    string
	}{
		{
			name:   "sync_success_sourcev1_destv0",
			config: "sync-success-sourcev1-destv0.yml",
		},
		{
			name:   "multiple_sources",
			config: "multiple-sources.yml",
		},
		{
			name:   "multiple_destinations",
			config: "multiple-destinations.yml",
		},
		{
			name:   "multiple_sources_destinations",
			config: "multiple-sources-destinations.yml",
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
			cmd.SetArgs([]string{"sync", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
			err := cmd.Execute()
			if tc.err != "" {
				assert.Contains(t, err.Error(), tc.err)
			} else {
				assert.NoError(t, err)
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

func TestFindMaxSupportedVersion(t *testing.T) {
	cases := []struct {
		name          string
		giveVersions  []int
		giveSupported int
		want          int
	}{
		{name: "support_less", giveVersions: []int{1, 2, 3}, giveSupported: 2, want: 2},
		{name: "support_same", giveVersions: []int{1, 2, 3}, giveSupported: 3, want: 3},
		{name: "support_more", giveVersions: []int{1, 2, 3}, giveSupported: 4, want: 3},
		{name: "support_none", giveVersions: []int{3, 4, 5}, giveSupported: 2, want: -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMaxSupportedVersion(tc.giveVersions, tc.giveSupported)
			assert.Equal(t, tc.want, got)
		})
	}
}
