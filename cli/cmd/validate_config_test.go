package cmd

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/stretchr/testify/require"
)

func TestValidateConfig(t *testing.T) {
	configs := []struct {
		name   string
		config string
		errors []string
	}{
		{
			name:   "multiple test sources should pass validation",
			config: "multiple-sources.yml",
		},
		{
			name:   "bad AWS and Postgres auth should fail validation",
			config: "validate-config-error.yml",
			errors: []string{"failed to validate source config cloudflare", "failed to validate destination config postgresql"},
		},
	}
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	for _, tc := range configs {
		t.Run(tc.name, func(t *testing.T) {
			cmd := NewCmdRoot()
			testConfig := path.Join(currentDir, "testdata", tc.config)
			baseArgs := testCommandArgs(t)

			args := append([]string{"validate-config", testConfig}, baseArgs...)
			cmd.SetArgs(args)
			err := cmd.Execute()
			// check that log was written and contains some lines from the plugin
			b, logFileError := os.ReadFile(baseArgs[3])
			logContent := string(b)
			require.NoError(t, logFileError, "failed to read cloudquery.log")
			require.NotEmpty(t, logContent, "cloudquery.log empty; expected some logs")
			require.NotContains(t, logContent, "skipping validation")

			if tc.errors != nil {
				for _, e := range tc.errors {
					require.Contains(t, err.Error(), e)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSplitHubPath(t *testing.T) {
	team, name, err := splitHubPath("cloudquery/aws")
	require.NoError(t, err)
	require.Equal(t, "cloudquery", team)
	require.Equal(t, "aws", name)

	team, name, err = splitHubPath("myteam/my-plugin")
	require.NoError(t, err)
	require.Equal(t, "myteam", team)
	require.Equal(t, "my-plugin", name)

	_, _, err = splitHubPath("no-slash")
	require.Error(t, err)

	_, _, err = splitHubPath("/missing-team")
	require.Error(t, err)

	_, _, err = splitHubPath("missing-name/")
	require.Error(t, err)

	_, _, err = splitHubPath("")
	require.Error(t, err)
}

func TestValidateConfig_HubAPI(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	// fakeHub serves canned responses for GetPluginVersion.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/plugins/cloudquery/source/aws/versions/v1.0.0":
			body, _ := json.Marshal(cloudquery_api.PluginVersionDetails{
				Name: "v1.0.0",
				SpecJsonSchema: strPtr(`{
					"$schema": "https://json-schema.org/draft/2020-12/schema",
					"type": ["object","null"],
					"properties": {"use_paid_apis": {"type": "boolean"}},
					"additionalProperties": false
				}`),
			})
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(body)
		case "/plugins/cloudquery/destination/pg/versions/v1.0.0":
			body, _ := json.Marshal(cloudquery_api.PluginVersionDetails{
				Name:           "v1.0.0",
				SpecJsonSchema: strPtr(`{"$schema":"https://json-schema.org/draft/2020-12/schema","type":["object","null"]}`),
			})
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(body)
		case "/plugins/cloudquery/source/missing/versions/v9.9.9":
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"not found"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	t.Setenv("CLOUDQUERY_API_URL", srv.URL)

	t.Run("good spec validates against Hub-returned schema (no plugin spawn)", func(t *testing.T) {
		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-hub-good.yml")
		baseArgs := testCommandArgs(t)
		args := append([]string{"validate-config", testConfig}, baseArgs...)
		cmd.SetArgs(args)
		err := cmd.Execute()
		require.NoError(t, err)

		logContent, readErr := os.ReadFile(baseArgs[3])
		require.NoError(t, readErr)
		require.Contains(t, string(logContent), "Fetching spec schema from Hub API")
		require.NotContains(t, string(logContent), "Initializing source")
	})

	t.Run("schema-violating spec fails validation", func(t *testing.T) {
		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-hub-bad.yml")
		baseArgs := testCommandArgs(t)
		args := append([]string{"validate-config", testConfig}, baseArgs...)
		cmd.SetArgs(args)
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to validate source config aws")
	})

	t.Run("--license bypasses Hub schema-fetch path and forces plugin spawn", func(t *testing.T) {
		licensePath := path.Join(t.TempDir(), "fake.license")
		require.NoError(t, os.WriteFile(licensePath, []byte("not-a-real-license"), 0o600))

		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-hub-good.yml")
		baseArgs := testCommandArgs(t)
		args := append([]string{"validate-config", testConfig, "--license", licensePath}, baseArgs...)
		cmd.SetArgs(args)
		// The spawn path will fail in this hermetic environment — we only care
		// that the Hub schema-fetch helper is not entered. Its identifying log
		// line "Fetching spec schema from Hub API" must therefore be absent.
		_ = cmd.Execute()
		logContent, readErr := os.ReadFile(baseArgs[3])
		require.NoError(t, readErr)
		require.NotContains(t, string(logContent), "Fetching spec schema from Hub API",
			"--license must skip the Hub schema-fetch code path")
	})

	t.Run("missing plugin version surfaces Hub 404", func(t *testing.T) {
		cmd := NewCmdRoot()
		testConfig := path.Join(currentDir, "testdata", "validate-config-hub-404.yml")
		baseArgs := testCommandArgs(t)
		args := append([]string{"validate-config", testConfig}, baseArgs...)
		cmd.SetArgs(args)
		err := cmd.Execute()
		require.Error(t, err)
		require.Contains(t, err.Error(), "404")
	})
}

func strPtr(s string) *string { return &s }

// A source-only spec that targets the platform destination (as `init` scaffolds)
// has no destination block. validate-config must mirror sync — auto-inject the
// platform destination and validate — rather than reject it for having none.
func TestValidateConfig_PlatformSourceOnly(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filename)

	// Hub API: return the source's schema so the aws source validates.
	hub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/plugins/cloudquery/source/aws/versions/v1.0.0" {
			body, _ := json.Marshal(cloudquery_api.PluginVersionDetails{
				Name:           "v1.0.0",
				SpecJsonSchema: strPtr(`{"$schema":"https://json-schema.org/draft/2020-12/schema","type":["object","null"]}`),
			})
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(body)
			return
		}
		http.NotFound(w, r)
	}))
	t.Cleanup(hub.Close)

	// Platform tenant API: the source-version gate and the injected destination's
	// whoami version lookup. A request to neither path fails the test.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/external-syncs/supported-source-versions":
			_ = json.NewEncoder(w).Encode(map[string]string{"cloudquery/aws": "v1.0.0"})
		case "/api/external-syncs/whoami":
			_ = json.NewEncoder(w).Encode(map[string]any{"tenant_id": "11111111-1111-1111-1111-111111111111", "plugin_version": "v1.0.1"})
		default:
			http.Error(w, "unexpected "+r.URL.Path, http.StatusNotFound)
		}
	}))
	t.Cleanup(es.Close)

	// A cqpd_ token in the standard key env drives the headless platform path: it
	// carries the tenant API URL (u claim) and doubles as the Hub credential.
	payload, _ := json.Marshal(map[string]any{"u": es.URL, "tm": "team-x"})
	token := "cqpd_" + base64.RawURLEncoding.EncodeToString(payload) + ".sig"
	t.Setenv("CLOUDQUERY_API_KEY", token)
	t.Setenv("CLOUDQUERY_API_URL", hub.URL)

	cmd := NewCmdRoot()
	testConfig := path.Join(currentDir, "testdata", "validate-config-platform-source-only.yml")
	baseArgs := testCommandArgs(t)
	args := append([]string{"validate-config", testConfig}, baseArgs...)
	cmd.SetArgs(args)
	require.NoError(t, cmd.Execute(), "a source-only platform spec validates without a destination block")

	logContent, readErr := os.ReadFile(baseArgs[3])
	require.NoError(t, readErr)
	require.NotContains(t, string(logContent), "expecting at least one destination")
}
