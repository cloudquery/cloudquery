package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"
)

func TestMigrate(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testDataDir := path.Join(currentDir, "testdata")
	cmd := NewCmdRoot()
	tmpDir := t.TempDir()
	logFileName := path.Join(tmpDir, "cloudquery.log")
	cmd.SetArgs([]string{"migrate", testDataDir, "--cq-dir", tmpDir, "--log-file-name", logFileName})
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
