package cmd

import (
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestSync(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll("cloudquery.log")
	})

	// this works but some funny stuff is going on
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testDataDir := path.Join(currentDir, "testdata")
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"sync", testDataDir})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	// check that log was written and contains some lines from the plugin
	b, err := os.ReadFile("cloudquery.log")
	if err != nil {
		t.Fatal("failed to read cloudquery.log:", err)
	}
	content := string(b)
	if len(content) == 0 {
		t.Fatalf("cloudquery.log empty, but expected some logs")
	}
	if !strings.Contains(content, "plugin=") {
		t.Errorf("cloudquery.log is expected to contain lines with string plugin=, but none found. Logs were: \n%v", content)
	}
}
