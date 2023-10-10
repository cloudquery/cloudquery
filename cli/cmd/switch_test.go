package cmd

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/adrg/xdg"
	"github.com/cloudquery/cloudquery/cli/internal/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSwitch(t *testing.T) {
	configDir := t.TempDir()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/teams":
			// write json response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"items": [{"name": "my-team"}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer ts.Close()

	t.Setenv(auth.EnvVarCloudQueryAPIKey, "test-api-key")
	t.Setenv("CLOUDQUERY_ACCOUNTS_URL", ts.URL)
	t.Setenv("CLOUDQUERY_API_URL", ts.URL)
	t.Cleanup(func() {
		CloseLogFile()
		os.RemoveAll(configDir)
	})

	t.Setenv("XDG_CONFIG_HOME", configDir)
	xdg.Reload()

	// calling switch before a team is set should not result in an error
	cmd := NewCmdRoot()
	cmd.SetArgs([]string{"switch"})
	err := cmd.Execute()
	require.NoError(t, err)

	// now set the team
	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"switch", "my-team"})
	err = cmd.Execute()
	require.NoError(t, err)

	cmd = NewCmdRoot()
	cmd.SetArgs([]string{"switch"})
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	err = cmd.Execute()
	require.NoError(t, err)
	out, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)
	assert.Contains(t, string(out), "Your current team is set to my-team")
	assert.Contains(t, string(out), "Teams available to you: my-team")

	// check that the config file was created in the temporary directory,
	// not somewhere else
	_, err = os.Stat(path.Join(configDir, "cloudquery", "config.json"))
	require.NoError(t, err)
}
