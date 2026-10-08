package specs

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPluginPathToOrgName(t *testing.T) {
	tests := []struct {
		path    string
		org     string
		name    string
		wantErr bool
	}{
		{path: "cloudquery/aws", org: "cloudquery", name: "aws"},
		{path: "/.cq", wantErr: true},
		{path: "/.cq/plugins/source/cloudquery/aws/v1.0.0/plugin", wantErr: true},
		{path: "cloudquery/", wantErr: true},
		{path: "/aws", wantErr: true},
		{path: "aws", wantErr: true},
		{path: "", wantErr: true},
		{path: "cloudquery/aws/extra", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			org, name, err := pluginPathToOrgName(tt.path)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q, got org=%q name=%q", tt.path, org, name)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if org != tt.org || name != tt.name {
				t.Fatalf("got org=%q name=%q, want org=%q name=%q", org, name, tt.org, tt.name)
			}
		})
	}
}

func newHubServer(t *testing.T, latestVersions map[string]string) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		latestVersion, ok := latestVersions[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"latest_version":%q}`, latestVersion)
	}))
	t.Cleanup(server.Close)
	t.Setenv("CLOUDQUERY_API_URL", server.URL)
}

func TestWarnOnOutdatedVersionsRecommendsUpgradeCheck(t *testing.T) {
	tests := []struct {
		name          string
		latestVersion string
		configPaths   []string
		want          string
	}{
		{
			name:          "two major versions behind",
			latestVersion: "v32.0.0",
			configPaths:   []string{"config.yml"},
			want:          "Source aws-prod is 2 major versions behind (v30.1.0 → v32.0.0). Before you upgrade, run `cloudquery upgrade check config.yml --source aws-prod --to v32.0.0` to see the schema impact on your destinations.\n",
		},
		{
			name:          "one major version behind with several config paths",
			latestVersion: "v31.2.3",
			configPaths:   []string{"sources.yml", "my configs/destinations.yml"},
			want:          "Source aws-prod is 1 major version behind (v30.1.0 → v31.2.3). Before you upgrade, run `cloudquery upgrade check sources.yml \"my configs/destinations.yml\" --source aws-prod --to v31.2.3` to see the schema impact on your destinations.\n",
		},
		{
			name:          "minor version behind",
			latestVersion: "v30.2.0",
			configPaths:   []string{"config.yml"},
		},
		{
			name:          "patch version behind",
			latestVersion: "v30.1.1",
			configPaths:   []string{"config.yml"},
		},
		{
			name:          "same version",
			latestVersion: "v30.1.0",
			configPaths:   []string{"config.yml"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newHubServer(t, map[string]string{"/plugins/cloudquery/source/aws": tt.latestVersion})
			warner, err := managedplugin.NewPluginVersionWarner(zerolog.Nop(), "")
			require.NoError(t, err)
			sources := []*Source{{Metadata: Metadata{Name: "aws-prod", Path: "cloudquery/aws", Registry: RegistryCloudQuery, Version: "v30.1.0"}}}

			var output bytes.Buffer
			WarnOnOutdatedVersions(context.Background(), warner, sources, nil, nil, WithUpgradeCheckRecommendation(tt.configPaths, &output))

			assert.Equal(t, tt.want, output.String())
		})
	}
}

func TestWarnOnOutdatedVersionsDoesNotRecommendUpgradeCheckForDestinations(t *testing.T) {
	newHubServer(t, map[string]string{"/plugins/cloudquery/destination/postgresql": "v9.0.0"})
	warner, err := managedplugin.NewPluginVersionWarner(zerolog.Nop(), "")
	require.NoError(t, err)
	destinations := []*Destination{{Metadata: Metadata{Name: "postgresql", Path: "cloudquery/postgresql", Registry: RegistryCloudQuery, Version: "v7.0.0"}}}

	var output bytes.Buffer
	WarnOnOutdatedVersions(context.Background(), warner, nil, destinations, nil, WithUpgradeCheckRecommendation([]string{"config.yml"}, &output))

	assert.Empty(t, output.String())
}

func TestWarnOnOutdatedVersionsDoesNotRecommendUpgradeCheckForNonCloudQuerySources(t *testing.T) {
	newHubServer(t, map[string]string{"/plugins/community/source/aws": "v32.0.0"})
	warner, err := managedplugin.NewPluginVersionWarner(zerolog.Nop(), "")
	require.NoError(t, err)
	sources := []*Source{{Metadata: Metadata{Name: "aws", Path: "community/aws", Registry: RegistryGitHub, Version: "v30.1.0"}}}

	var output bytes.Buffer
	WarnOnOutdatedVersions(context.Background(), warner, sources, nil, nil, WithUpgradeCheckRecommendation([]string{"config.yml"}, &output))

	assert.Empty(t, output.String())
}

func TestWarnOnOutdatedVersionsRequestsSourceLatestVersionOnce(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"latest_version":"v32.0.0"}`))
	}))
	t.Cleanup(server.Close)
	t.Setenv("CLOUDQUERY_API_URL", server.URL)
	warner, err := managedplugin.NewPluginVersionWarner(zerolog.Nop(), "")
	require.NoError(t, err)
	sources := []*Source{{Metadata: Metadata{Name: "aws", Path: "cloudquery/aws", Registry: RegistryCloudQuery, Version: "v30.1.0"}}}

	var output bytes.Buffer
	WarnOnOutdatedVersions(context.Background(), warner, sources, nil, nil, WithUpgradeCheckRecommendation([]string{"config.yml"}, &output))

	assert.NotEmpty(t, output.String())
	assert.Equal(t, int32(1), requests.Load())
}
