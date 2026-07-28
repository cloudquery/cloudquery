package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"

	cqconfig "github.com/cloudquery/cloudquery-api-go/config"
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

func cqpdTokenWithClaims(t *testing.T, claims map[string]any) string {
	t.Helper()
	payload, err := json.Marshal(claims)
	require.NoError(t, err)
	return "cqpd_" + base64.RawURLEncoding.EncodeToString(payload) + ".sig"
}

func TestTeamFromToken(t *testing.T) {
	t.Parallel()
	require.Equal(t, "acme", TeamFromToken(cqpdTokenWithClaims(t, map[string]any{"tm": "acme", "u": "https://x"})),
		"reads the tm claim so the CLI can target team-scoped endpoints from the token alone")
	require.Empty(t, TeamFromToken(cqpdTokenWithClaims(t, map[string]any{"u": "https://x"})), "no tm claim -> empty")
	require.Empty(t, TeamFromToken("not-a-cqpd-token"), "non-cqpd_ token -> empty")
	require.Empty(t, TeamFromToken("cqpd_@@@.sig"), "malformed payload -> empty")
}

func TestDownloadAuth_HeadlessCQPDToken(t *testing.T) {
	tok := cqpdTokenWithClaims(t, map[string]any{"tm": "acme", "u": "https://x"})
	t.Setenv(EnvPlatformToken, tok)
	// The cqpd_ branch resolves from the token alone — no cloud is wired, so any
	// cloud call would fail the test.
	gotTok, team, err := DownloadAuth(context.Background(), zerolog.Nop(), nil, nil, nil)
	require.NoError(t, err)
	require.Equal(t, tok, gotTok, "headless flow uses the cqpd_ token as the download credential")
	require.Equal(t, "acme", team, "team comes from the token's tm claim")
}

func TestPropagatePluginCredential(t *testing.T) {
	t.Run("cqpd_ with no api key exports it for plugins", func(t *testing.T) {
		t.Setenv("CLOUDQUERY_API_KEY", "")
		PropagatePluginCredential("cqpd_x.y")
		require.Equal(t, "cqpd_x.y", os.Getenv("CLOUDQUERY_API_KEY"))
	})
	t.Run("never overwrites an existing api key", func(t *testing.T) {
		t.Setenv("CLOUDQUERY_API_KEY", "cq_existing")
		PropagatePluginCredential("cqpd_x.y")
		require.Equal(t, "cq_existing", os.Getenv("CLOUDQUERY_API_KEY"), "a user-set key (e.g. a team key) wins")
	})
	t.Run("non-cqpd token is a no-op", func(t *testing.T) {
		t.Setenv("CLOUDQUERY_API_KEY", "")
		PropagatePluginCredential("not-a-cqpd-token")
		require.Empty(t, os.Getenv("CLOUDQUERY_API_KEY"))
	})
}

func TestTeamMismatchWarning(t *testing.T) {
	setConfigTeam := func(t *testing.T, team string) {
		t.Helper()
		require.NoError(t, cqconfig.SetConfigHome(t.TempDir()))
		t.Cleanup(func() { _ = cqconfig.UnsetConfigHome() })
		if team != "" {
			require.NoError(t, cqconfig.SetValue("team", team))
		}
	}
	t.Run("configured team differs from tm claim -> warns with both teams", func(t *testing.T) {
		setConfigTeam(t, "team-a")
		msg := teamMismatchWarning("acme")
		require.Contains(t, msg, `"acme"`)
		require.Contains(t, msg, `"team-a"`)
	})
	t.Run("configured team matches -> no warning", func(t *testing.T) {
		setConfigTeam(t, "acme")
		require.Empty(t, teamMismatchWarning("acme"))
	})
	t.Run("no configured team -> no warning", func(t *testing.T) {
		setConfigTeam(t, "")
		require.Empty(t, teamMismatchWarning("acme"))
	})
	t.Run("no tm claim -> no warning", func(t *testing.T) {
		setConfigTeam(t, "team-a")
		require.Empty(t, teamMismatchWarning(""))
	})
}

func TestDownloadAuth_CQPDViaAPIKeyEnv(t *testing.T) {
	tok := cqpdTokenWithClaims(t, map[string]any{"tm": "acme", "u": "https://x"})
	t.Setenv("CLOUDQUERY_API_KEY", tok)
	// A cloudquery-registry source makes GetAuthTokenIfNeeded resolve the token
	// from CLOUDQUERY_API_KEY. Its cqpd_ prefix must route to the tm claim, not
	// to GetTeamForToken (which would call cloud and fail for a syncs-scoped key).
	src := []*specs.Source{{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery}}}
	gotTok, team, err := DownloadAuth(context.Background(), zerolog.Nop(), src, nil, nil)
	require.NoError(t, err)
	require.Equal(t, tok, gotTok, "a cqpd_ in CLOUDQUERY_API_KEY is used as the download credential")
	require.Equal(t, "acme", team, "team resolved locally from tm, no cloud call")
}

func TestDetectTenantForInit_DirectToken(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0"}
	es := supportedSourceVersionsServer(t, "", pinned)
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err)
	require.NotNil(t, ti, "a CQ_PLATFORM_TOKEN means a tenant is present")
	require.Equal(t, es.URL, ti.APIURL, "url comes from the token's u claim")
	require.Equal(t, pinned, ti.PinnedSourceVersions, "pins fetched with the direct token")
}

func TestDetectTenantForInit_DirectToken_NoURL(t *testing.T) {
	// A legacy token with no url claim still identifies a tenant, but there's
	// nowhere to fetch pins/configs from.
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"tm": "acme"}))
	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err)
	require.NotNil(t, ti)
	require.Empty(t, ti.APIURL)
	require.Nil(t, ti.PinnedSourceVersions)
	_, cfgErr := ti.RecommendedSourceConfig(context.Background(), "cloudquery/aws")
	require.Error(t, cfgErr, "no session → no config")
}

func TestDetectTenantForInit_Disabled(t *testing.T) {
	t.Setenv(envDisable, "1")
	t.Setenv(EnvPlatformToken, cqpdTokenWithURL(t, "https://x.example.com"))
	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err)
	require.Nil(t, ti, "disable env suppresses detection")
}

func TestDetectTenantForInit_NoCredentials(t *testing.T) {
	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err)
	require.Nil(t, ti, "no token and no cloud creds → not detected")
}

func TestDetectTenantForInit_CloudPath(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0"}
	es := supportedSourceVersionsServer(t, "cqpd_minted.sig", pinned)
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
		}})
	}, func(w http.ResponseWriter, _ *http.Request) {
		writeSession(w, "cqpd_minted.sig", es.URL)
	})
	t.Setenv(envAPIURL, srv.URL)

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
	require.NoError(t, err)
	require.NotNil(t, ti)
	require.Equal(t, "https://acme.us.platform.cloudquery.io", ti.APIURL, "url is built from the active tenant's host")
	require.Equal(t, pinned, ti.PinnedSourceVersions, "pins fetched via the minted session")
}

func TestDetectTenantForInit_MintFails_Errors(t *testing.T) {
	// A detected tenant whose session can't be minted can't run a platform sync,
	// so init must fail rather than scaffold a source-only spec that would break.
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
		}})
	}, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	})
	t.Setenv(envAPIURL, srv.URL)

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
	require.Error(t, err, "mint failure is surfaced")
	require.Nil(t, ti)
}

func TestDetectTenantForInit_MintOK_PinsUnavailable_NoError(t *testing.T) {
	// Mint succeeds but the versions lookup fails → best-effort: no error, nil
	// pins, tenant detected. init scaffolds with the hub's latest.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer es.Close()
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
		}})
	}, func(w http.ResponseWriter, _ *http.Request) {
		writeSession(w, "cqpd_minted.sig", es.URL)
	})
	t.Setenv(envAPIURL, srv.URL)

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
	require.NoError(t, err)
	require.NotNil(t, ti)
	require.Equal(t, "https://acme.us.platform.cloudquery.io", ti.APIURL)
	require.Nil(t, ti.PinnedSourceVersions, "pins unavailable → nil, init falls back to hub latest")
}

func TestDetectTenantForInit_DirectToken_Unauthorized_Errors(t *testing.T) {
	// An expired/rejected env CQ_PLATFORM_TOKEN must fail init early rather than
	// silently scaffold a spec that 401s at sync time.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer es.Close()
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.Error(t, err, "a rejected env token fails init early")
	require.ErrorContains(t, err, "expired")
	require.Nil(t, ti)
}

func TestDetectTenantForInit_DirectToken_ServerError_BestEffort(t *testing.T) {
	// A non-auth failure (500) stays best-effort — init proceeds without pins.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer es.Close()
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err, "a non-auth failure doesn't fail init")
	require.NotNil(t, ti)
	require.Nil(t, ti.PinnedSourceVersions)
}

func TestDetectTenantForInit_MintOK_PinsUnauthorized_BestEffort(t *testing.T) {
	// The session was just minted, so a 401 from the pins lookup is a server
	// anomaly, not user-fixable — init still proceeds (unlike the direct-token
	// path, which the user can fix by refreshing/unsetting the env token).
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer es.Close()
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []map[string]any{
			{"tenant_id": "11111111-1111-1111-1111-111111111111", "status": "active", "team_name": "team-x", "host": "acme.us.platform.cloudquery.io", "subdomain": "acme"},
		}})
	}, func(w http.ResponseWriter, _ *http.Request) {
		writeSession(w, "cqpd_minted.sig", es.URL)
	})
	t.Setenv(envAPIURL, srv.URL)

	ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
	require.NoError(t, err, "a fresh-minted session's 401 doesn't fail init")
	require.NotNil(t, ti)
	require.Nil(t, ti.PinnedSourceVersions)
}

// DetectTenantForInit must make the SAME multi-tenant decision auto-injection
// does: skip (report nothing) when a team has several active tenants and no
// CQ_PLATFORM_TENANT_ID override — otherwise `init` would point the user at a
// tenant a real sync would refuse to inject into.
func TestDetectTenantForInit_MultipleActiveTenants(t *testing.T) {
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
		ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
		require.NoError(t, err, "ambiguous tenant is not an error; init falls back to its normal flow")
		require.Nil(t, ti, "several active tenants + no override is ambiguous; a sync would skip, so detection must too")
	})

	t.Run("override picks the matching tenant", func(t *testing.T) {
		pinned := map[string]string{"cloudquery/gcp": "v18.0.0"}
		es := supportedSourceVersionsServer(t, "", pinned)
		srv := fakeCloud(t, twoActive, func(w http.ResponseWriter, _ *http.Request) {
			writeSession(w, "cqpd_minted.sig", es.URL)
		})
		t.Setenv(envAPIURL, srv.URL)
		t.Setenv(envTenantID, "22222222-2222-2222-2222-222222222222")
		ti, err := DetectTenantForInit(context.Background(), zerolog.Nop(), "tok", "team-x")
		require.NoError(t, err)
		require.NotNil(t, ti)
		require.Equal(t, "https://beta.us.platform.cloudquery.io", ti.APIURL)
		require.Equal(t, pinned, ti.PinnedSourceVersions)
	})
}

func TestTenantInit_RecommendedSourceConfig(t *testing.T) {
	// The recommended-source-config lookup reuses the session (token + endpoint)
	// from DetectTenantForInit — no extra mint.
	const config = "kind: source\nspec:\n  name: aws\n  destinations: [\"platform\"]\n"
	var gotPath string
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/external-syncs/recommended-source-config", r.URL.Path)
		require.Equal(t, "Bearer cqpd_direct.sig", r.Header.Get("Authorization"))
		gotPath = r.URL.Query().Get("path")
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"config": config, "version": "v33.0.0", "tables": []string{"aws_s3_buckets"}})
	}))
	defer es.Close()

	ti := &TenantInit{token: "cqpd_direct.sig", endpointBase: es.URL}
	got, err := ti.RecommendedSourceConfig(context.Background(), "cloudquery/aws")
	require.NoError(t, err)
	require.Equal(t, config, got, "the platform's config is returned verbatim")
	require.Equal(t, "cloudquery/aws", gotPath, "the source path is passed through")

	// No session and empty path are errors — a platform scaffold is required.
	_, err = (&TenantInit{}).RecommendedSourceConfig(context.Background(), "cloudquery/aws")
	require.Error(t, err)
	_, err = ti.RecommendedSourceConfig(context.Background(), "")
	require.Error(t, err)
}

func TestTenantInit_RecommendedSourceConfig_Errors(t *testing.T) {
	newTI := func(handler http.HandlerFunc) *TenantInit {
		es := httptest.NewServer(handler)
		t.Cleanup(es.Close)
		return &TenantInit{token: "cqpd_direct.sig", endpointBase: es.URL}
	}

	t.Run("non-200", func(t *testing.T) {
		ti := newTI(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "boom", http.StatusInternalServerError)
		})
		_, err := ti.RecommendedSourceConfig(context.Background(), "cloudquery/aws")
		require.ErrorContains(t, err, "status 500")
	})

	t.Run("empty config is an error, not a fallback", func(t *testing.T) {
		ti := newTI(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"config": ""}`))
		})
		_, err := ti.RecommendedSourceConfig(context.Background(), "cloudquery/aws")
		require.ErrorContains(t, err, "no recommended source config for cloudquery/aws")
	})
}

func TestInject_DirectToken_InjectsWithoutCloud(t *testing.T) {
	// A pre-minted cqpd_ token via env injects the destination with no cloud
	// login, tenant discovery or session mint. No fake cloud is wired, so any
	// such call would fail the test.
	t.Setenv(EnvPlatformToken, "cqpd_payload.sig")

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

func TestInject_DirectToken_PinsVersionFromWhoami(t *testing.T) {
	var tok string // captured by the handler; assigned after the server URL is known
	whoami := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/external-syncs/whoami", r.URL.Path)
		require.Equal(t, "Bearer "+tok, r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"tenant_id": "11111111-1111-1111-1111-111111111111", "plugin_version": "v9.9.9"})
	}))
	defer whoami.Close()
	tok = cqpdTokenWithClaims(t, map[string]any{"u": whoami.URL, "tm": "acme"}) // u → whoami server
	t.Setenv(EnvPlatformToken, tok)

	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", testSources(), testDestinations())
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, "v9.9.9", got[1].Version, "headless flow pins the version whoami recommends")
}

func TestInject_DirectToken_ViaAPIKeyEnv(t *testing.T) {
	// A cqpd_ in the standard CLOUDQUERY_API_KEY env injects just like
	// CQ_PLATFORM_TOKEN — same headless path, no cloud calls.
	t.Setenv("CLOUDQUERY_API_KEY", "cqpd_payload.sig")

	sources := testSources()
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", sources, testDestinations())
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, "cqpd_payload.sig", got[1].Spec["token"], "the cqpd_ from CLOUDQUERY_API_KEY is used directly")
	require.NotContains(t, got[1].Spec, "api_url")
}

func TestInject_NoPlatformTarget_NoOp(t *testing.T) {
	// Even with a token available, no injection happens unless a source opts in
	// by listing `platform` in its destinations.
	t.Setenv(EnvPlatformToken, "cqpd_payload.sig")

	sources := []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
		Destinations: []string{"pg"},
	}}
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "", "", sources, testDestinations())
	require.NoError(t, err)
	require.Len(t, got, 1, "no source targets platform → nothing injected")
}

func TestInject_DirectToken_ExistingPlatformDestination_UsesTheirs(t *testing.T) {
	t.Setenv(EnvPlatformToken, "cqpd_payload.sig")

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

func TestInject_SessionMintError_Fails(t *testing.T) {
	// A source opted into `platform` and its tenant was found, but the mint fails
	// → hard error rather than silently dropping the opt-in (mirrors the
	// ambiguous-tenant case). Destinations are left unchanged.
	srv := fakeCloud(t, nil, func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "not a member", http.StatusNotFound)
	})
	t.Setenv(envAPIURL, srv.URL)

	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", testSources(), testDestinations())
	require.Error(t, err, "mint failure on an opted-in sync fails rather than skipping")
	require.Len(t, got, 1, "destinations unchanged when injection errors")
}

func TestInject_NonPlatformSync_MintErrorIrrelevant(t *testing.T) {
	// The safety guarantee: a sync that does NOT opt into `platform` returns
	// before any tenant/mint call, so a broken session server can't affect it —
	// no injection, no error. Any cloud call would surface via the error.
	var calls atomic.Int32
	srv := fakeCloud(t, func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		writeTenants(w, tenantItem("11111111-1111-1111-1111-111111111111", "active", "team-x"))
	}, func(w http.ResponseWriter, _ *http.Request) {
		calls.Add(1)
		http.Error(w, "boom", http.StatusInternalServerError)
	})
	t.Setenv(envAPIURL, srv.URL)

	sources := []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0", Registry: specs.RegistryCloudQuery},
		Destinations: []string{"pg"}, // no platform opt-in
	}}
	got, err := MaybeInjectDestination(context.Background(), zerolog.Nop(), "tok", "team-x", sources, testDestinations())
	require.NoError(t, err, "a non-platform sync is never affected by platform-destination setup")
	require.Len(t, got, 1, "nothing injected")
	require.Zero(t, calls.Load(), "no cloud/mint call is even attempted without an opt-in")
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

func supportedSourceVersionsServer(t *testing.T, wantToken string, versions map[string]string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/api/external-syncs/supported-source-versions", r.URL.Path)
		if wantToken != "" {
			require.Equal(t, "Bearer "+wantToken, r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(versions)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestPinnedSourceVersions_DirectToken(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0"}
	var tok string
	es := supportedSourceVersionsServer(t, "", pinned) // token asserted below via closure capture
	tok = cqpdTokenWithClaims(t, map[string]any{"u": es.URL})
	t.Setenv(EnvPlatformToken, tok)

	got, err := PinnedSourceVersions(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err)
	require.Equal(t, pinned, got, "a direct cqpd_ token resolves pinned versions without cloud")
}

func TestPinnedSourceVersions_LoginMintsAndFetches(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0", "cloudquery/gcp": "v18.2.1"}
	es := supportedSourceVersionsServer(t, "cqpd_minted.sig", pinned)
	srv := fakeCloud(t, nil, func(w http.ResponseWriter, _ *http.Request) {
		// The minted session's api_url points at the external-syncs server.
		writeSession(w, "cqpd_minted.sig", es.URL)
	})
	t.Setenv(envAPIURL, srv.URL)

	got, err := PinnedSourceVersions(context.Background(), zerolog.Nop(), "tok", "team-x")
	require.NoError(t, err)
	require.Equal(t, pinned, got, "logged-in flow mints a session then fetches pinned versions")
}

func TestPinnedSourceVersions_NoTenant_Nil(t *testing.T) {
	got, err := PinnedSourceVersions(context.Background(), zerolog.Nop(), "", "")
	require.NoError(t, err, "no token and no cloud creds → nothing to resolve, no error")
	require.Nil(t, got)
}

func TestPinnedSourceVersions_DirectToken_Unauthorized(t *testing.T) {
	// A rejected direct env token propagates errPlatformUnauthorized so the gate
	// (validate-config) fails instead of silently passing.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer es.Close()
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	got, err := PinnedSourceVersions(context.Background(), zerolog.Nop(), "", "")
	require.ErrorIs(t, err, errPlatformUnauthorized)
	require.Nil(t, got)
}

func TestSourceVersionSupported(t *testing.T) {
	t.Parallel()
	pinned := map[string]string{"cloudquery/aws": "v33.28.0"}
	// Mirrors api/externalsyncs.sourceSupported: same major, not newer than pinned.
	require.True(t, sourceVersionSupported("cloudquery/aws", "v33.28.0", pinned), "exact pin is supported")
	require.True(t, sourceVersionSupported("cloudquery/aws", "v33.0.0", pinned), "older same-major is supported")
	require.False(t, sourceVersionSupported("cloudquery/aws", "v33.29.0", pinned), "newer than pin rejected")
	require.False(t, sourceVersionSupported("cloudquery/aws", "v32.0.0", pinned), "older major rejected")
	require.False(t, sourceVersionSupported("cloudquery/aws", "v34.0.0", pinned), "newer major rejected")
	require.False(t, sourceVersionSupported("cloudquery/gcp", "v1.0.0", pinned), "unpinned path rejected")
	require.False(t, sourceVersionSupported("cloudquery/aws", "not-semver", pinned), "unparseable version rejected")
}

func TestAnySourceTargetsPlatform(t *testing.T) {
	t.Parallel()
	require.True(t, AnySourceTargetsPlatform(testSources()))
	require.False(t, AnySourceTargetsPlatform([]*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws"},
		Destinations: []string{"pg"},
	}}))
}

func TestGateSources_NoPlatformTarget_NoNetwork(t *testing.T) {
	// A source that doesn't target platform must not trigger any cloud/tenant
	// call — no server is wired, so one would fail the test.
	t.Setenv(envAPIURL, "http://127.0.0.1:0")
	sources := []*specs.Source{{
		Metadata:     specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v99.0.0"},
		Destinations: []string{"pg"},
	}}
	require.NoError(t, GateSources(context.Background(), zerolog.Nop(), "tok", "team-x", sources))
}

func TestGateSources_UnsupportedVersion_Errors(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0"}
	es := supportedSourceVersionsServer(t, "", pinned)
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	sources := []*specs.Source{
		{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v34.0.0"}, Destinations: []string{"platform"}},
		{Metadata: specs.Metadata{Name: "custom", Path: "cloudquery/unknown", Version: "v1.0.0"}, Destinations: []string{"platform"}},
	}
	err := GateSources(context.Background(), zerolog.Nop(), "", "", sources)
	require.Error(t, err)
	require.ErrorContains(t, err, "unsupported source plugin version(s)")
	require.ErrorContains(t, err, "aws (supported version: v33.0.0)", "names the accepted version")
	require.ErrorContains(t, err, "custom (not a supported source)", "flags an unrecognized source path")
}

func TestGateSources_SupportedVersion_OK(t *testing.T) {
	pinned := map[string]string{"cloudquery/aws": "v33.0.0"}
	es := supportedSourceVersionsServer(t, "", pinned)
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	sources := []*specs.Source{
		{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v33.0.0"}, Destinations: []string{"platform"}},
	}
	require.NoError(t, GateSources(context.Background(), zerolog.Nop(), "", "", sources))
}

func TestGateSources_VersionsUnavailable_FailOpen(t *testing.T) {
	// A platform-targeted source, but the pinned-versions lookup fails (500) →
	// gate opens (nil), mirroring the server's fail-open when versions are down.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer es.Close()
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	sources := []*specs.Source{
		{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v99.0.0"}, Destinations: []string{"platform"}},
	}
	require.NoError(t, GateSources(context.Background(), zerolog.Nop(), "", "", sources))
}

func TestGateSources_RejectedToken_Errors(t *testing.T) {
	// A rejected direct env token would 401 the sync too, so the gate must fail
	// (not pass clean) — validate-config's "catch what sync would hit" contract.
	es := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}))
	defer es.Close()
	t.Setenv(EnvPlatformToken, cqpdTokenWithClaims(t, map[string]any{"u": es.URL}))

	sources := []*specs.Source{
		{Metadata: specs.Metadata{Name: "aws", Path: "cloudquery/aws", Version: "v1.0.0"}, Destinations: []string{"platform"}},
	}
	err := GateSources(context.Background(), zerolog.Nop(), "", "", sources)
	require.Error(t, err)
	require.ErrorContains(t, err, "expired")
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

func TestOnlyPlatformDestinations(t *testing.T) {
	platformDest := specs.Destination{Metadata: specs.Metadata{Name: destinationName, Path: "cloudquery/platform", Registry: specs.RegistryCloudQuery}}
	pgDest := specs.Destination{Metadata: specs.Metadata{Name: "pg", Path: "cloudquery/postgresql", Registry: specs.RegistryCloudQuery}}

	require.False(t, OnlyPlatformDestinations(nil), "no destinations is not a platform-only sync")
	require.False(t, OnlyPlatformDestinations([]specs.Destination{pgDest}))
	require.False(t, OnlyPlatformDestinations([]specs.Destination{platformDest, pgDest}), "mixed destinations keep the CLI event")
	require.True(t, OnlyPlatformDestinations([]specs.Destination{platformDest}))
}
