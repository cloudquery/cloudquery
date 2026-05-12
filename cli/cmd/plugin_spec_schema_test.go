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

func TestWriteSchemaOutput(t *testing.T) {
	const schema = `{"type":"object"}`

	t.Run("to explicit output file", func(t *testing.T) {
		dir := t.TempDir()
		out := path.Join(dir, "custom.json")
		require.NoError(t, writeSchemaOutput(schema, "aws", out, ""))
		got, err := os.ReadFile(out)
		require.NoError(t, err)
		require.Equal(t, schema, string(got))
	})

	t.Run("to schemas dir by plugin name", func(t *testing.T) {
		dir := t.TempDir()
		sub := path.Join(dir, "nested")
		require.NoError(t, writeSchemaOutput(schema, "aws", "", sub))
		got, err := os.ReadFile(path.Join(sub, "aws.json"))
		require.NoError(t, err)
		require.Equal(t, schema, string(got))
	})
}
