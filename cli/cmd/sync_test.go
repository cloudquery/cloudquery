package cmd

import (
	"path"
	"runtime"
	"testing"
)

func TestSync(t *testing.T) {
	// this works but some funny stuff is going on
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)
	testDataDir := path.Join(currentDir, "testdata")
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"sync", testDataDir})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
}
