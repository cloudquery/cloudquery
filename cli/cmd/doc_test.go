package cmd

import (
	"os"
	"testing"
)

var docFiles = []string{
	"cloudquery.md",
	"cloudquery_completion.md",
	"cloudquery_completion_bash.md",
	"cloudquery_completion_fish.md",
	"cloudquery_completion_powershell.md",
	"cloudquery_completion_zsh.md",
	"cloudquery_generate.md",
	"cloudquery_sync.md",
}

func TestDoc(t *testing.T) {
	tmpdir, tmpErr := os.MkdirTemp("", "docs_*")
	if tmpErr != nil {
		t.Fatalf("failed to create temporary directory: %v", tmpErr)
	}
	defer os.RemoveAll(tmpdir)
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"doc", tmpdir})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	files, err := os.ReadDir(tmpdir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != len(docFiles) {
		t.Errorf("expected %d files, got %d", len(docFiles), len(files))
	}
}
