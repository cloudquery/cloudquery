package platform

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// Base URL env consumed by internal/api.NewClient.
const envAPIURL = "CLOUDQUERY_API_URL"

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

func tenantItem(id, status, team string) map[string]any {
	return map[string]any{"tenant_id": id, "status": status, "team_name": team}
}

func writeTenants(w http.ResponseWriter, items ...map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{"items": items})
}

func writeSession(w http.ResponseWriter, token, apiURL string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"token":              token,
		"api_url":            apiURL,
		"expires_in_seconds": 604800,
	})
}

func fakeCloud(t *testing.T, tenants func(w http.ResponseWriter, r *http.Request), session func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	t.Helper()
	if tenants == nil {
		tenants = func(w http.ResponseWriter, _ *http.Request) {
			writeTenants(w, tenantItem("11111111-1111-1111-1111-111111111111", "active", "team-x"))
		}
	}
	if session == nil {
		session = func(w http.ResponseWriter, _ *http.Request) {
			writeSession(w, "cqpd_payload.sig", "https://x.us.platform.cloudquery.io")
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/user/platform/tenants", tenants)
	mux.HandleFunc("/platform-destination/session", session)
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func TestInject_Active_AppendsDestinationAndWiresSources(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer tok", r.Header.Get("Authorization"))
		writeTenants(w, tenantItem("11111111-1111-1111-1111-111111111111", "active", "team-x"))
	}, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer tok", r.Header.Get("Authorization"))
		var body struct {
			TenantID string `json:"tenant_id"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		require.Equal(t, "11111111-1111-1111-1111-111111111111", body.TenantID)
		writeSession(w, "cqpd_payload.sig", "https://x.us.platform.cloudquery.io")
	})
	t.Setenv(envAPIURL, srv.URL)

	sources := testSources()
	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, destinations)

	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.Equal(t, "https://x.us.platform.cloudquery.io", got[1].Spec["api_url"])
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"], "destination must get the minted cqpd_ token, not the cloud credential")
	require.Equal(t, defaultPlugin.Version, got[1].Version)
	require.Equal(t, defaultPlugin.Path, got[1].Path)
	require.Equal(t, specs.RegistryCloudQuery, got[1].Registry)
	require.True(t, got[1].SyncSummary, "send_sync_summary must be set so the destination receives finalize signals")
	require.Contains(t, sources[0].Destinations, destinationName)
}

func TestInject_CreatedTenant_Injects(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w, tenantItem("11111111-1111-1111-1111-111111111111", "created", "team-x"))
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
}

func TestInject_ExistingPlatformBlockOverwrittenNotDuplicated(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)

	destinations := append(testDestinations(), &specs.Destination{
		Metadata: specs.Metadata{Name: destinationName, Path: "user/stale", Version: "v0.0.1", Registry: specs.RegistryLocal},
		Spec:     map[string]any{"api_url": "https://stale", "token": "stale", "keep": "me"},
	})
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), destinations)

	require.Len(t, got, 2)
	platformDest := got[1]
	require.Equal(t, destinationName, platformDest.Name)
	require.Equal(t, defaultPlugin.Path, platformDest.Path)
	require.Equal(t, specs.RegistryCloudQuery, platformDest.Registry)
	require.Equal(t, "https://x.us.platform.cloudquery.io", platformDest.Spec["api_url"])
	require.Equal(t, "cqpd_payload.sig", platformDest.Spec["token"])
	require.Equal(t, "me", platformDest.Spec["keep"], "unrelated spec keys must survive")
}

func TestInject_NoTenantForTeam_NoOp(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w,
			tenantItem("22222222-2222-2222-2222-222222222222", "active", "other-team"),
			tenantItem("33333333-3333-3333-3333-333333333333", "pending", "team-x"),
		)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	destinations := testDestinations()
	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), destinations)
	require.Len(t, got, 1)
}

func TestInject_TenantListError_NoOp(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
}

func TestInject_SessionMintError_NoOp(t *testing.T) {
	srv := fakeCloud(t, nil, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not a member", http.StatusNotFound)
	})
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
}

func TestInject_MultipleTenants_RequiresEnvSelection(t *testing.T) {
	tenants := func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w,
			tenantItem("11111111-1111-1111-1111-111111111111", "active", "team-x"),
			tenantItem("22222222-2222-2222-2222-222222222222", "active", "team-x"),
		)
	}

	t.Run("unset skips", func(t *testing.T) {
		srv := fakeCloud(t, tenants, nil)
		t.Setenv(envAPIURL, srv.URL)
		got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
		require.Len(t, got, 1)
	})

	t.Run("env picks", func(t *testing.T) {
		srv := fakeCloud(t, tenants, func(w http.ResponseWriter, r *http.Request) {
			var body struct {
				TenantID string `json:"tenant_id"`
			}
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			require.Equal(t, "22222222-2222-2222-2222-222222222222", body.TenantID)
			writeSession(w, "cqpd_x.y", "https://x")
		})
		t.Setenv(envAPIURL, srv.URL)
		t.Setenv(envTenantID, "22222222-2222-2222-2222-222222222222")
		got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
		require.Len(t, got, 2)
	})

	t.Run("env mismatch skips", func(t *testing.T) {
		srv := fakeCloud(t, tenants, nil)
		t.Setenv(envAPIURL, srv.URL)
		t.Setenv(envTenantID, "99999999-9999-9999-9999-999999999999")
		got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
		require.Len(t, got, 1)
	})
}

func TestInject_PluginCoordsEnvOverride(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)
	t.Setenv(envPluginRegistry, "local")
	t.Setenv(envPluginPath, "/abs/path/bin/platform")
	t.Setenv(envPluginVersion, "v0.0.0-dev")

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, specs.RegistryLocal, got[1].Registry)
	require.Equal(t, "/abs/path/bin/platform", got[1].Path)
	require.Equal(t, "v0.0.0-dev", got[1].Version)
}

func TestInject_DisableEnv_SkipsBeforeAnyCall(t *testing.T) {
	var calls atomic.Int32
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		writeTenants(w)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)
	t.Setenv(envDisable, "1")

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
	require.Zero(t, calls.Load())
}

func TestInject_CloudRun_SkipsBeforeAnyCall(t *testing.T) {
	var calls atomic.Int32
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		writeTenants(w)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)
	t.Setenv("CQ_CLOUD", "1")

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
	require.Zero(t, calls.Load())
}

func setResolveCredentials(t *testing.T, token, team string, err error) {
	t.Helper()
	prev := resolveCredentials
	resolveCredentials = func(context.Context) (string, string, error) {
		return token, team, err
	}
	t.Cleanup(func() { resolveCredentials = prev })
}

func TestInject_EmptyTokenOrTeam_NoOp(t *testing.T) {
	var calls atomic.Int32
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		writeTenants(w)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)
	// No token from the caller and none resolvable best-effort: stay a no-op.
	setResolveCredentials(t, "", "", errors.New("not logged in"))

	require.Len(t, MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "team-x", testSources(), testDestinations()), 1)
	require.Len(t, MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "", testSources(), testDestinations()), 1)
	require.Zero(t, calls.Load())
}

func TestInject_BestEffortCredentials_Injects(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)
	// Caller passed no token (spec pulls no cloudquery-registry plugin); the
	// best-effort resolver supplies one so injection still happens.
	setResolveCredentials(t, "tok", "team-x", nil)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"])
}

func TestAllocateSyncGroupID_TimeShaped(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	sgid := got[1].SyncGroupId
	require.Len(t, sgid, 17, "YYYYMMDDhhmmssfff")
	_, err := json.Number(sgid).Int64()
	require.NoError(t, err)
}
