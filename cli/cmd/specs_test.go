package cmd

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseJSONSchema(t *testing.T) {
	sc, err := parseJSONSchema(`{"type":"object","properties":{"name":{"$ref":"#/$defs/name"}},"$defs":{"name":{"type":"string"}}}`)
	require.NoError(t, err)
	require.NoError(t, sc.Validate(map[string]any{"name": "value"}))
	require.Error(t, sc.Validate(map[string]any{"name": 1}))
}

func TestParseJSONSchemaDoesNotLoadLocalFiles(t *testing.T) {
	localSchema := filepath.Join(t.TempDir(), "local.json")
	require.NoError(t, os.WriteFile(localSchema, []byte(`{"type":"string"}`), 0600))
	fileURL := url.URL{Scheme: "file", Path: filepath.ToSlash(localSchema)}

	_, err := parseJSONSchema(`{"type":"object","properties":{"name":{"$ref":"` + fileURL.String() + `"}}}`)
	require.Error(t, err)
}
