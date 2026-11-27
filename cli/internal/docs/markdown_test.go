package docs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestMarkdownFilePath(t *testing.T) {
	dir := t.TempDir()

	got, err := MarkdownFilePath(dir, "aws_ec2_instances")
	require.NoError(t, err)
	require.Equal(t, filepath.Join(dir, "aws_ec2_instances.md"), got)

	for _, name := range []string{"../escaped", "nested/table", "/tmp/absolute"} {
		t.Run(name, func(t *testing.T) {
			_, err := MarkdownFilePath(dir, name)
			require.Error(t, err)
		})
	}
}

func TestGenerateRejectsTableNameOutsideOutputDir(t *testing.T) {
	parent := t.TempDir()
	dir := filepath.Join(parent, "docs")
	require.NoError(t, os.Mkdir(dir, 0755))

	g := NewGenerator("test", schema.Tables{{Name: "../escaped"}})
	require.Error(t, g.Generate(dir, FormatMarkdown))

	_, err := os.Stat(filepath.Join(parent, "escaped.md"))
	require.ErrorIs(t, err, os.ErrNotExist)
}
