package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestSync(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testDataDir := path.Join(currentDir, "testdata")
	tmpDir := t.TempDir()
	logFileName := path.Join(tmpDir, "cloudquery.log")
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"sync", testDataDir, "--cq-dir", tmpDir, "--log-file-name", logFileName})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	// check that log was written and contains some lines from the plugin
	b, err := os.ReadFile(logFileName)
	if err != nil {
		t.Fatal("failed to read cloudquery.log:", err)
	}
	content := string(b)
	if len(content) == 0 {
		t.Fatalf("cloudquery.log empty; expected some logs")
	}
}
