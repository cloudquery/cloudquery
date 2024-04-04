package cmd

import (
	"path"
	"testing"
)

func testCommandArgs(t *testing.T) []string {
	cqDir := t.TempDir()
	logFileName := path.Join(cqDir, "cloudquery.log")
	t.Cleanup(func() {
		CloseLogFile()
	})

	return []string{"--cq-dir", cqDir, "--log-file-name", logFileName}
}
