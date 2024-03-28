//go:build !windows

package cmd

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInferRegistry(t *testing.T) {
	t.Skip()
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
			logFileName := path.Join(cqDir, "cloudquery.log")
			t.Cleanup(func() {
				CloseLogFile()
			})
			testConfig := filepath.Join(cqDir, "config.yml")

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
