package cmd

import (
	"os"
	"path"
	"testing"
)

var docFiles = []string{
	"cloudquery.md",
	"cloudquery_sync.md",
	"cloudquery_migrate.md",
	"cloudquery_tables.md",
	"cloudquery_install.md",
}

func TestDoc(t *testing.T) {
	defer CloseLogFile()
	cmd := NewCmdRoot()
	tmpDir := t.TempDir()
	cqTmpDir := t.TempDir()
	logFileName := path.Join(cqTmpDir, "cloudquery.log")
	cmd.SetArgs([]string{"doc", tmpDir, "--cq-dir", cqTmpDir, "--log-file-name", logFileName})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != len(docFiles) {
		t.Errorf("expected %d files, got %d", len(docFiles), len(files))
	}
}
