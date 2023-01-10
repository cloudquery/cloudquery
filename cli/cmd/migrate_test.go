package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestMigrate(t *testing.T) {
	defer CloseLogFile()
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testConfig := path.Join(currentDir, "testdata", "sync-success.yml")
	cmd := NewCmdRoot()
	tmpDir := t.TempDir()
	logFileName := path.Join(tmpDir, "cloudquery.log")
	cmd.SetArgs([]string{"migrate", testConfig, "--cq-dir", tmpDir, "--log-file-name", logFileName})
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
