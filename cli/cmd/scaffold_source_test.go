package cmd

import (
	"os"
	"path"
	"testing"
)

func TestScaffoldSource(t *testing.T) {
	outputDir := t.TempDir()
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"scaffold", "source", "cloudquery", "test", "--output", outputDir})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	for _, filePath := range destinationTemplates {
		if _, err := os.Stat(path.Join(outputDir, filePath)); err != nil {
			t.Fatal("file not found", err)
		}
	}
}