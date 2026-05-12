package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/stretchr/testify/require"
)

func TestPluginTypeFromKind(t *testing.T) {
	src, err := pluginTypeFromKind("source")
	require.NoError(t, err)
	require.Equal(t, managedplugin.PluginSource, src)

	dst, err := pluginTypeFromKind("destination")
	require.NoError(t, err)
	require.Equal(t, managedplugin.PluginDestination, dst)

	_, err = pluginTypeFromKind("transformer")
	require.Error(t, err)
}

func TestSchemaFileName(t *testing.T) {
	require.Equal(t, "aws@v33.0.0.json", schemaFileName("aws", "v33.0.0"))
	require.Equal(t, "aws.json", schemaFileName("aws", ""))
}

func TestWriteSchemaOutput(t *testing.T) {
	const schema = `{"type":"object"}`

	t.Run("to schemas dir with versioned name", func(t *testing.T) {
		dir := t.TempDir()
		sub := path.Join(dir, "nested")
		require.NoError(t, writeSchemaOutput(schema, "aws", "v33.0.0", sub))
		got, err := os.ReadFile(path.Join(sub, "aws@v33.0.0.json"))
		require.NoError(t, err)
		require.Equal(t, schema, string(got))
	})

	t.Run("to schemas dir without version falls back to unversioned name", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, writeSchemaOutput(schema, "aws", "", dir))
		got, err := os.ReadFile(path.Join(dir, "aws.json"))
		require.NoError(t, err)
		require.Equal(t, schema, string(got))
	})
}
