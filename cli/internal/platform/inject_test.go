package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Base URL env consumed by internal/api.NewClient.
const envAPIURL = "CLOUDQUERY_API_URL"

// testSources opts into the platform destination (lists it in `destinations`),
// the trigger for injection.
func testSources() []*specs.Source {
	return []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
		Destinations: []string{"pg", "platform"},
	}}
}

func testDestinations() []*specs.Destination {
	return []*specs.Destination{{
		Metadata: specs.Metadata{Name: "pg", Path: "cloudquery/postgresql", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
	}}
}

// mustInject runs injection and fails the test on a (hard) error.
func mustInject(t *testing.T, token, team string, sources []*specs.Source, destinations []*specs.Destination) []*specs.Destination {
	t.Helper()
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), token, team, sources, destinations)
	require.NoError(t, err)
	return got
}

func tenantItem(id, tenantStatus, team string) map[string]any {
	return map[string]any{"tenant_id": id, "status": tenantStatus, "team_name": team}
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

func TestInject_OptIn_AppendsDestination(t *testing.T) {
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
	got := mustInject(t, "tok", "team-x", sources, testDestinations())

	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.NotContains(t, got[1].Spec, "api_url", "api_url is not injected — the cqpd_ token carries it")
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"], "destination must get the minted cqpd_ token, not the cloud credential")
	srcVersionsJSON, err := json.Marshal(got[1].Spec["source_versions"])
	require.NoError(t, err)
	require.JSONEq(t, `[{"name":"aws","path":"cloudquery/aws","version":"v1.0.0"}]`, string(srcVersionsJSON),
		"each source's path+version must be reported for the platform gate")
	require.Equal(t, defaultPlugin.Version, got[1].Version)
	require.Equal(t, defaultPlugin.Path, got[1].Path)
	require.Equal(t, specs.RegistryCloudQuery, got[1].Registry)
	require.True(t, got[1].SyncSummary, "send_sync_summary must be set so the destination receives finalize signals")
	require.Equal(t, specs.WriteModeAppend, got[1].WriteMode, "sync_group_id requires a write mode other than overwrite-delete-stale")
	require.NotEmpty(t, got[1].SyncGroupId)
	require.Contains(t, sources[0].Destinations, destinationName, "the opted-in source still targets platform")

	// Multiple platform-targeting sources are reported in order, none dropped.
	twoGot := mustInject(t, "tok", "team-x", []*specs.Source{
		{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery}, Destinations: []string{"platform"}},
		{Metadata: specs.Metadata{Name: "gcp", Path: "cloudquery/gcp", Version: "v2.3.4", Registry: specs.RegistryCloudQuery}, Destinations: []string{"platform"}},
	}, testDestinations())
	twoJSON, err := json.Marshal(twoGot[1].Spec["source_versions"])
	require.NoError(t, err)
	require.JSONEq(t, `[{"name":"aws","path":"cloudquery/aws","version":"v1.0.0"},{"name":"gcp","path":"cloudquery/gcp","version":"v2.3.4"}]`,
		string(twoJSON), "sources reported in order, none dropped")
}

func cqpdTokenWithURL(t *testing.T, apiURL string) string {
	t.Helper()
	payload, err := json.Marshal(map[string]any{"u": apiURL})
	require.NoError(t, err)
	return "cqpd_" + base64.RawURLEncoding.EncodeToString(payload) + ".sig"
}

func TestDetectTenant_DirectToken(t *testing.T) {
	t.Setenv(envPlatformToken, cqpdTokenWithURL(t, "https://acme.us.platform.cloudquery.io"))
	url, ok := DetectTenant(context.Background(), "", "")
	require.True(t, ok, "a CQ_PLATFORM_TOKEN means a tenant is present")
	require.Equal(t, "https://acme.us.platform.cloudquery.io", url, "url comes from the token's u claim")
}

func TestDetectTenant_Disabled(t *testing.T) {
	t.Setenv(envDisable, "1")
	t.Setenv(envPlatformToken, cqpdTokenWithURL(t, "https://x.example.com"))
	_, ok := DetectTenant(context.Background(), "", "")
	require.False(t, ok, "disable env suppresses detection")
}

func TestDetectTenant_NoCredentials(t *testing.T) {
	_, ok := DetectTenant(context.Background(), "", "")
	require.False(t, ok, "no token and no cloud creds → not detected")
}

func TestDetectTenant_CloudPath(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
		}})
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	url, ok := DetectTenant(context.Background(), "tok", "team-x")
	require.True(t, ok)
	require.Equal(t, "https://acme.us.platform.cloudquery.io", url, "url is built from the active tenant's host")
}

// DetectTenant must make the SAME multi-tenant decision auto-injection does:
// skip (report nothing) when a team has several active tenants and no
// CQ_PLATFORM_TENANT_ID override — otherwise `init` would point the user at a
// tenant a real sync would refuse to inject into.
func TestDetectTenant_MultipleActiveTenants(t *testing.T) {
	twoActive := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
			{"tenant_id": "22222222-2222-2222-2222-222222222222", "status": "active", "team_name": "team-x", "host": "beta.us.platform.cloudquery.io", "subdomain": "beta"},
		}})
	}

	t.Run("ambiguous without override reports nothing", func(t *testing.T) {
		srv := fakeCloud(t, twoActive, nil)
		t.Setenv(envAPIURL, srv.URL)
		_, ok := DetectTenant(context.Background(), "tok", "team-x")
		require.False(t, ok, "several active tenants + no override is ambiguous; a sync would skip, so DetectTenant must too")
	})

	t.Run("override picks the matching tenant", func(t *testing.T) {
		srv := fakeCloud(t, twoActive, nil)
		t.Setenv(envAPIURL, srv.URL)
		t.Setenv(envTenantID, "22222222-2222-2222-2222-222222222222")
		url, ok := DetectTenant(context.Background(), "tok", "team-x")
		require.True(t, ok)
		require.Equal(t, "https://beta.us.platform.cloudquery.io", url)
	})
}

func TestInject_DirectToken_InjectsWithoutCloud(t *testing.T) {
	// A pre-minted cqpd_ token via env injects the destination with no cloud
	// login, tenant discovery or session mint. No fake cloud is wired, so any
	// such call would fail the test.
	t.Setenv(envPlatformToken, "cqpd_payload.sig")

	sources := testSources()
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", sources, testDestinations())
	require.NoError(t, err)

	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"], "the supplied cqpd_ token is used directly")
	require.NotContains(t, got[1].Spec, "api_url", "api_url is derived from the token, not injected")
	require.NotEmpty(t, got[1].Spec["source_versions"], "sources are still reported for the gate")
	require.Equal(t, defaultPlugin.Version, got[1].Version)
	require.Contains(t, sources[0].Destinations, destinationName)
}

func TestInject_NoPlatformTarget_NoOp(t *testing.T) {
	// Even with a token available, no injection happens unless a source opts in
	// by listing `platform` in its destinations.
	t.Setenv(envPlatformToken, "cqpd_payload.sig")

	sources := []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
		Destinations: []string{"pg"},
	}}
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", sources, testDestinations())
	require.NoError(t, err)
	require.Len(t, got, 1, "no source targets platform → nothing injected")
}

func TestInject_DirectToken_ExistingPlatformDestination_UsesTheirs(t *testing.T) {
	t.Setenv(envPlatformToken, "cqpd_payload.sig")

	userDest := &specs.Destination{Metadata: specs.Metadata{Name: destinationName, Path: "user/custom"}}
	destinations := append(testDestinations(), userDest)
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", testSources(), destinations)
	require.NoError(t, err, "a user-defined platform destination is used, not an error")
	require.Len(t, got, 2, "nothing injected on top of the user's platform destination")
	require.Equal(t, "user/custom", got[1].Path, "the user's platform destination is left untouched")
}

func TestInject_CreatedTenant_Injects(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w, tenantItem("11111111-1111-1111-1111-111111111111", "created", "team-x"))
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
}

func TestInject_ExistingPlatformDestination_UsesTheirs(t *testing.T) {
	// A user-defined platform destination (e.g. for debugging) is respected; the
	// CLI doesn't mint or inject over it — so no cloud call is even attempted.
	t.Setenv(envAPIURL, "http://127.0.0.1:0") // any cloud call would fail

	userDest := &specs.Destination{Metadata: specs.Metadata{Name: destinationName, Path: "user/custom", Version: "v9.9.9"}}
	destinations := append(testDestinations(), userDest)
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), destinations)

	require.NoError(t, err)
	require.Len(t, got, 2, "spec returned unchanged; nothing injected")
	require.Equal(t, "user/custom", got[1].Path, "the user's platform destination is left untouched")
}

func TestInject_NoTenantForTeam_NoOp(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w,
			tenantItem("22222222-2222-2222-2222-222222222222", "active", "other-team"),
			tenantItem("33333333-3333-3333-3333-333333333333", "pending", "team-x"),
		)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
}

func TestInject_TenantListError_NoOp(t *testing.T) {
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
}

func TestInject_SessionMintError_NoOp(t *testing.T) {
	srv := fakeCloud(t, nil, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not a member", http.StatusNotFound)
	})
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 1)
}

func TestInject_MultipleTenants_RequiresEnvSelection(t *testing.T) {
	tenants := func(w http.ResponseWriter, _ *http.Request) {
		writeTenants(w,
			tenantItem("11111111-1111-1111-1111-111111111111", "active", "team-x"),
			tenantItem("22222222-2222-2222-2222-222222222222", "active", "team-x"),
		)
	}

	t.Run("unset errors with a hint", func(t *testing.T) {
		srv := fakeCloud(t, tenants, nil)
		t.Setenv(envAPIURL, srv.URL)
		// A source opted into `platform` but the tenant is ambiguous → fail with
		// an actionable hint rather than silently dropping the opt-in.
		got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
		require.ErrorIs(t, err, errAmbiguousTenant)
		require.ErrorContains(t, err, "Hint:")
		require.ErrorContains(t, err, envTenantID)
		require.Len(t, got, 1, "destinations unchanged when injection errors")
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
		got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
		require.Len(t, got, 2)
	})

	t.Run("env mismatch errors with a hint", func(t *testing.T) {
		srv := fakeCloud(t, tenants, nil)
		t.Setenv(envAPIURL, srv.URL)
		t.Setenv(envTenantID, "99999999-9999-9999-9999-999999999999")
		got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
		require.ErrorIs(t, err, errAmbiguousTenant)
		require.ErrorContains(t, err, "Hint:")
		require.Len(t, got, 1, "destinations unchanged when injection errors")
	})
}

func TestInject_PluginCoordsEnvOverride(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)
	t.Setenv(envPluginRegistry, "local")
	t.Setenv(envPluginPath, "/abs/path/bin/platform")
	t.Setenv(envPluginVersion, "v0.0.0-dev")

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
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

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
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

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
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
	// No token and none resolvable: no-op.
	setResolveCredentials(t, "", "", errors.New("not logged in"))

	require.Len(t, mustInject(t, "", "team-x", testSources(), testDestinations()), 1)
	require.Len(t, mustInject(t, "tok", "", testSources(), testDestinations()), 1)
	require.Zero(t, calls.Load())
}

func TestInject_BestEffortCredentials_Injects(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)
	// No caller token; best-effort resolver supplies one.
	setResolveCredentials(t, "tok", "team-x", nil)

	got := mustInject(t, "", "", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, destinationName, got[1].Name)
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"])
}

func TestAllocateSyncGroupID_TimeShaped(t *testing.T) {
	srv := fakeCloud(t, nil, nil)
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	sgid := got[1].SyncGroupId
	require.Len(t, sgid, 17, "YYYYMMDDhhmmssfff")
	_, err := json.Number(sgid).Int64()
	require.NoError(t, err)
}

func TestIsInjectedDestination(t *testing.T) {
	require.True(t, IsInjectedDestination(destinationName))
	require.False(t, IsInjectedDestination("postgresql"))
}

func TestCleanInitError(t *testing.T) {
	// gRPC-wrapped plugin-init chain → strip the rpc + plugin-sdk prefixes,
	// leaving the destination's own message.
	wrapped := status.Error(codes.Internal,
		"failed to init plugin: failed to initialize client: failed to start sync with CloudQuery Platform: unsupported source plugin version(s): aws (supported version: v33.28.0) (HTTP 422)")
	require.Equal(t,
		"failed to start sync with CloudQuery Platform: unsupported source plugin version(s): aws (supported version: v33.28.0) (HTTP 422)",
		CleanInitError(wrapped))

	// A plain (non-gRPC) error passes through unchanged.
	require.Equal(t, "boom", CleanInitError(errors.New("boom")))
}

func sessionWithPluginVersion(version string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"token":              "cqpd_payload.sig",
			"api_url":            "https://x.us.platform.cloudquery.io",
			"expires_in_seconds": 604800,
			"plugin_version":     version,
		})
	}
}

func TestInject_PlatformPinnedVersion(t *testing.T) {
	srv := fakeCloud(t, nil, sessionWithPluginVersion("v2.5.0"))
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Len(t, got, 2)
	require.Equal(t, "v2.5.0", got[1].Version, "platform-pinned destination version overrides the CLI default")
}

func TestInject_EnvVersionBeatsPlatformPin(t *testing.T) {
	t.Setenv(envPluginVersion, "v9.9.9")
	srv := fakeCloud(t, nil, sessionWithPluginVersion("v2.5.0"))
	t.Setenv(envAPIURL, srv.URL)

	got := mustInject(t, "tok", "team-x", testSources(), testDestinations())
	require.Equal(t, "v9.9.9", got[1].Version, "env override wins over the platform pin")
}
