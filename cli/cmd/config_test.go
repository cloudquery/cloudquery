package cmd

import (
	"bytes"
	"io"
	"os"
	"path"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/assert"
)

func TestConfigSetGet(t *testing.T) {
	configDir := t.TempDir()
	t.Cleanup(func() {
		CloseLogFile()
		os.RemoveAll(configDir)
	})

	t.Setenv("XDG_CONFIG_HOME", configDir)
	xdg.Reload()

	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"config", "set", "team", "my-team"})
	err := cmd.Execute()
	assert.NoError(t, err)

	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"config", "get", "team"})
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	err = cmd.Execute()
	out, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "my-team\n", string(out))

	// check that the config file was created in the temporary directory,
	// not somewhere else
	_, err = os.Stat(path.Join(configDir, configPath))
	assert.NoError(t, err)
}
