package cmd

import (
	"io"
	"os"
	"path"
	"path/filepath"
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
			cmd.SetArgs([]string{"plugin", "install", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
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

func TestInferRegistry(t *testing.T) {
	configs := []struct {
		name               string
		config             string
		outContains        string
		wantErrContains    string
		wantErrNotContains string
	}{
		{
			name: "infer registry success",
			config: `
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: v10.0.0
  table_concurrency: 10
  destinations: [dummydest]
  tables: [test]
---
kind: destination
spec:
  name: dummydest
  path: localhost:7777
  registry: grpc
`,
			outContains:        "Downloading https://storage.googleapis.com/cq-cloud-releases/cloudquery/source/gcp/v10.0.0/",
			wantErrContains:    "",
			wantErrNotContains: "",
		},
		{
			name: "infer registry fail: recommend github",
			config: `
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  version: v99.999.0
  table_concurrency: 10
  destinations: [dummydest]
  tables: [test]
---
kind: destination
spec:
  name: dummydest
  path: localhost:7777
  registry: grpc
`,
			wantErrContains:    "Hint",
			wantErrNotContains: "",
		},
		{
			name: "no infer registry fail: don't recommend github",
			config: `
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  registry: cloudquery
  version: v99.999.0
  table_concurrency: 10
  destinations: [dummydest]
  tables: [test]
---
kind: destination
spec:
  name: dummydest
  path: localhost:7777
  registry: grpc
`,
			wantErrContains:    "",
			wantErrNotContains: "Hint",
		},
		{
			name: "no infer registry fail: don't recommend github for github",
			config: `
kind: source
spec:
  name: gcp
  path: cloudquery/gcp
  registry: github
  version: v99.999.0
  table_concurrency: 10
  destinations: [dummydest]
  tables: [test]
---
kind: destination
spec:
  name: dummydest
  path: localhost:7777
  registry: grpc
`,
			wantErrContains:    "",
			wantErrNotContains: "Hint",
		},
	}

	for _, tc := range configs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cqDir := t.TempDir()
			t.Cleanup(func() {
				os.RemoveAll(cqDir)
			})
			testConfig := filepath.Join(cqDir, "config.yml")
			logFileName := path.Join(cqDir, "cloudquery.log")
			err := os.WriteFile(testConfig, []byte(tc.config), 0644)
			require.NoError(t, err)

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() {
				os.Stdout = oldStdout
			}()

			cmd := NewCmdRoot()
			cmd.SetArgs([]string{"plugin", "install", testConfig, "--cq-dir", cqDir, "--log-file-name", logFileName})
			err = cmd.Execute()

			_ = w.Close()
			out, _ := io.ReadAll(r)
			if tc.outContains != "" {
				require.Contains(t, string(out), tc.outContains)
			}

			if tc.wantErrContains == "" && tc.wantErrNotContains == "" {
				require.NoError(t, err)
			}
			if tc.wantErrContains != "" {
				require.ErrorContains(t, err, tc.wantErrContains)
			}
			if tc.wantErrNotContains != "" {
				require.Error(t, err)
				require.NotContains(t, err.Error(), tc.wantErrNotContains)
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
