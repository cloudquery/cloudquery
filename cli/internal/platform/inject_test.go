package platform

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func testSources() []*specs.Source {
	return []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
		Destinations: []string{"pg"},
	}}
}

func testDestinations() []*specs.Destination {
	return []*specs.Destination{{
		Metadata: specs.Metadata{Name: "pg", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
	}}
}

func TestInject_Active_AppendsDestinationAndWiresSources(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/teams/team-x/platform/status", r.URL.Path)
		require.Equal(t, "Bearer tok", r.Header.Get("Authorization"))
		_ = json.NewEncoder(w).Encode(statusResponse{
			PlatformURL:    "https://x.mycloudquery.com",
			Status:         statusActive,
			TeamName:       "team-x",
			PluginRegistry: "cloudquery",
			PluginPath:     "cloudquery/platform",
			PluginVersion:  "v1.2.3",
		})
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.Equal(t, "https://x.mycloudquery.com", got[1].Spec["api_url"])
	require.Equal(t, "tok", got[1].Spec["token"])
	require.Equal(t, "v1.2.3", got[1].Version)
	require.Equal(t, "cloudquery/platform", got[1].Path)
	require.Equal(t, specs.RegistryCloudQuery, got[1].Registry)
	require.True(t, got[1].SyncSummary, "send_sync_summary must be set so the destination receives finalize signals")
	require.Contains(t, sources[0].Destinations, destinationName)
}

func TestInject_LocalRegistryFromStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(statusResponse{
			PlatformURL:    "https://x.mycloudquery.com",
			Status:         statusActive,
			PluginRegistry: "local",
			PluginPath:     "/abs/path/bin/platform",
		})
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, specs.RegistryLocal, got[1].Registry)
	require.Equal(t, "/abs/path/bin/platform", got[1].Path)
}

func TestInject_FallbackWhenStatusMissingPluginFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(statusResponse{
			PlatformURL: "https://x.mycloudquery.com",
			Status:      statusActive,
		})
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, fallbackVersion, got[1].Version)
	require.Equal(t, fallbackPath, got[1].Path)
	require.Equal(t, specs.RegistryCloudQuery, got[1].Registry)
}

func TestInject_NotOnboarded_NoOp(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 1)
	require.NotContains(t, sources[0].Destinations, destinationName)
}

func TestInject_Pending_NoOp(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_ = json.NewEncoder(w).Encode(statusResponse{
			PlatformURL: "https://x.mycloudquery.com",
			Status:      "pending",
		})
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 1)
	require.NotContains(t, sources[0].Destinations, destinationName)
}

func TestInject_Disabled_NoOp(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)
	t.Setenv(envDisable, "1")

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 1)
	require.EqualValues(t, 0, hits.Load(), "should not call cloud API when disabled")
}

func TestInject_AlreadyPresent_OverwritesInPlace(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
		_ = json.NewEncoder(w).Encode(statusResponse{
			PlatformURL:    "https://x.mycloudquery.com",
			Status:         statusActive,
			PluginRegistry: "cloudquery",
			PluginPath:     "cloudquery/platform",
			PluginVersion:  "v9.9.9",
		})
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := append(testDestinations(), &specs.Destination{
		Metadata: specs.Metadata{Name: destinationName, Path: "placeholder", Version: "placeholder", Registry: specs.RegistryCloudQuery},
		Spec:     map[string]any{"api_url": "placeholder", "token": "placeholder"},
	})
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 2, "should not duplicate the destination block")
	require.EqualValues(t, 1, hits.Load())
	require.Equal(t, "v9.9.9", got[1].Version)
	require.Equal(t, "cloudquery/platform", got[1].Path)
	require.Equal(t, "https://x.mycloudquery.com", got[1].Spec["api_url"])
	require.Equal(t, "tok", got[1].Spec["token"])
	require.True(t, got[1].SyncSummary)
	require.Contains(t, sources[0].Destinations, destinationName)
}

func TestInject_ServerError_FailOpen(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 1, "5xx should not break sync; fall-through silently")
}
