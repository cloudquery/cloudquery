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
		name            string
		config          string
		wantSourceFiles int
		wantDestFiles   int
	}{
		{
			name:            "sync_success_sourcev1_destv0",
			config:          "sync-success-sourcev1-destv0.yml",
			wantSourceFiles: 2,
			wantDestFiles:   2,
		},
		{
			name:            "multiple_sources",
			config:          "multiple-sources.yml",
			wantSourceFiles: 2,
			wantDestFiles:   2,
		},
		{
			name:            "multiple_destinations",
			config:          "multiple-destinations.yml",
			wantSourceFiles: 2,
			wantDestFiles:   4,
		},
		{
			name:            "multiple_sources_destinations",
			config:          "multiple-sources-destinations.yml",
			wantSourceFiles: 2,
			wantDestFiles:   2,
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
			justFiles := readFiles(t, cqDir, "")

			sourceFiles, destFiles := 0, 0
			for _, file := range justFiles {
				if strings.HasPrefix(file, "plugins/source") {
					sourceFiles++
				} else if strings.HasPrefix(file, "plugins/destination") {
					destFiles++
				}
			}
			assert.Equalf(t, tc.wantSourceFiles, sourceFiles, "expected %d source files, got %d", tc.wantSourceFiles, sourceFiles)
			assert.Equalf(t, tc.wantDestFiles, destFiles, "expected %d destination files, got %d", tc.wantDestFiles, destFiles)
			if t.Failed() {
				t.Logf("files found: %v", justFiles)
				t.FailNow()
			}

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

		if !files[i].IsDir() {
			justFiles = append(justFiles, path.Join(prefix, name))
			continue
		}

		justFiles = append(justFiles, readFiles(t, path.Join(basedir, files[i].Name()), path.Join(prefix, name))...)
	}
	sort.Strings(justFiles)
	return justFiles
}
