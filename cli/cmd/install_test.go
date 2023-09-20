package cmd

import (
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstall(t *testing.T) {
	configs := []struct {
		name      string
		config    string
		wantFiles []string
	}{
		{
			name:   "sync_success_sourcev1_destv0",
			config: "sync-success-sourcev1-destv0.yml",
			wantFiles: []string{
				"cloudquery.log",
				"plugins/destination/cloudquery/test/v2.1.2/plugin",
				"plugins/destination/cloudquery/test/v2.1.2/plugin.zip",
				"plugins/source/cloudquery/test/v2.0.3/plugin",
				"plugins/source/cloudquery/test/v2.0.3/plugin.zip",
			},
		},
		{
			name:   "multiple_sources",
			config: "multiple-sources.yml",
			wantFiles: []string{
				"cloudquery.log",
				"plugins/destination/cloudquery/test/v2.2.7/plugin",
				"plugins/destination/cloudquery/test/v2.2.7/plugin.zip",
				"plugins/source/cloudquery/test/v3.1.7/plugin",
				"plugins/source/cloudquery/test/v3.1.7/plugin.zip",
			},
		},
		{
			name:   "multiple_destinations",
			config: "multiple-destinations.yml",
			wantFiles: []string{
				"cloudquery.log",
				"plugins/destination/cloudquery/test/v2.1.0/plugin",
				"plugins/destination/cloudquery/test/v2.1.0/plugin.zip",
				"plugins/destination/cloudquery/test/v2.2.7/plugin",
				"plugins/destination/cloudquery/test/v2.2.7/plugin.zip",
				"plugins/source/cloudquery/test/v2.0.3/plugin",
				"plugins/source/cloudquery/test/v2.0.3/plugin.zip",
			},
		},
		{
			name:   "multiple_sources_destinations",
			config: "multiple-sources-destinations.yml",
			wantFiles: []string{
				"cloudquery.log",
				"plugins/destination/cloudquery/test/v2.2.7/plugin",
				"plugins/destination/cloudquery/test/v2.2.7/plugin.zip",
				"plugins/source/cloudquery/test/v3.1.7/plugin",
				"plugins/source/cloudquery/test/v3.1.7/plugin.zip",
			},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cqDir := t.TempDir()
			t.Cleanup(func() {
				CloseLogFile()
				os.RemoveAll(cqDir)
			})
			testConfig := path.Join(currentDir, "testdata", tc.config)
			logFileName := path.Join(cqDir, "cloudquery.log")
			cmd := NewCmdRoot()
			cmd.SetArgs([]string{"install", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
			err := cmd.Execute()
			assert.NoError(t, err)

			// check if all files were created
			sort.Strings(tc.wantFiles)
			justFiles := readFiles(t, cqDir, "")
			assert.ElementsMatch(t, tc.wantFiles, justFiles)

			// check that log was written and contains some lines
			b, logFileError := os.ReadFile(path.Join(cqDir, "cloudquery.log"))
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
		})
	}
}

func readFiles(t *testing.T, basedir, prefix string) []string {
	files, err := os.ReadDir(basedir)
	assert.NoError(t, err)
	var justFiles []string
	for i := range files {
		name := files[i].Name()
		name = strings.TrimSuffix(name, ".exe")
		name = strings.ReplaceAll(name, ".exe.zip", ".zip")

		if !files[i].IsDir() {
			justFiles = append(justFiles, path.Join(prefix, name))
			continue
		}

		justFiles = append(justFiles, readFiles(t, path.Join(basedir, files[i].Name()), path.Join(prefix, name))...)
	}
	sort.Strings(justFiles)
	return justFiles
}
