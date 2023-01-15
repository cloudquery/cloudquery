package cmd

import (
	"os/exec"
	"testing"
)

func TestSource(t *testing.T) {
	tmpDir := t.TempDir()
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"source", "test-org", "test", "--output", tmpDir})
	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
	goModTidy := exec.Command("go", "mod", "tidy")
	goModTidy.Dir = tmpDir
	if out, err := goModTidy.CombinedOutput(); err != nil {
		t.Error(string(out) + err.Error())
	}
	goBuild := exec.Command("go", "build", ".")
	goBuild.Dir = tmpDir
	if out, err := goBuild.CombinedOutput(); err != nil {
		t.Error(string(out) + err.Error())
	}
}
