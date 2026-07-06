// Package platform auto-injects a platform destination into syncs for teams
// with an active platform tenant.
package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	neturl "net/url"
	"os"
	"slices"
	"strconv"
	"strings"
	gosync "sync"
	"time"

	"github.com/Masterminds/semver"
	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	cqconfig "github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	cqauth "github.com/cloudquery/cloudquery/cli/v6/internal/auth"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/status"
)

const (
	envDisable  = "CQ_DISABLE_PLATFORM_DESTINATION"
	envTenantID = "CQ_PLATFORM_TENANT_ID"
	// EnvPlatformToken lets a user inject the platform destination from a
	// pre-minted cqpd_ token directly — no cloud login or session mint. The
	// token carries the tenant API URL, so it's all the destination needs.
	EnvPlatformToken = "CQ_PLATFORM_TOKEN"

	envPluginRegistry = "CQ_PLATFORM_PLUGIN_REGISTRY"
	envPluginPath     = "CQ_PLATFORM_PLUGIN_PATH"
	envPluginVersion  = "CQ_PLATFORM_PLUGIN_VERSION"

	destinationName = "platform"

	// cqpdPrefix marks a platform-destination token on the wire.
	cqpdPrefix = "cqpd_"

	requestTimeout = 10 * time.Second
)

// Tenant statuses that are eligible for platform destination injection.
var injectableStatuses = []cloudquery_api.PlatformTenantStatus{
	cloudquery_api.PlatformTenantStatusActive,
	cloudquery_api.PlatformTenantStatusCreated,
}

type pluginCoordinates struct {
	Registry string
	Path     string
	Version  string
}

var defaultPlugin = pluginCoordinates{
	Registry: "cloudquery",
	Path:     "cloudquery/platform",
	Version:  "v1.0.1",
}

func pluginCoords() pluginCoordinates {
	p := defaultPlugin
	if v := os.Getenv(envPluginRegistry); v != "" {
		p.Registry = v
	}
	if v := os.Getenv(envPluginPath); v != "" {
		p.Path = v
	}
	if v := os.Getenv(envPluginVersion); v != "" {
		p.Version = v
	}
	return p
}

// sourceVersion is one entry of the platform destination's `source_versions`
// spec field — the source plugin path+version the platform gates on. JSON tags
// match the platform's CreateExternalSync `sources` items.
type sourceVersion struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Version string `json:"version"`
}

// TenantInit carries what `init` needs about the CloudQuery Platform tenant a
// sync would auto-inject into, resolved in a single tenant lookup + mint: the URL
// to report, the pinned source versions to scaffold, and a session for further
// /external-syncs/* lookups (RecommendedTables). A nil *TenantInit from
// DetectTenantForInit means no platform-init scenario applies and init should use
// its normal source + destination flow.
type TenantInit struct {
	// APIURL is the tenant base URL to show the user (host only, no /api). May be
	// empty if a directly supplied CQ_PLATFORM_TOKEN predates url-carrying tokens.
	APIURL string
	// PinnedSourceVersions maps source plugin path -> pinned version. Best-effort:
	// nil when the lookup failed, so init falls back to the hub's latest.
	PinnedSourceVersions map[string]string

	// token + endpointBase reach /external-syncs/* for later per-plugin lookups,
	// reusing the same session so init mints at most once.
	token        string
	endpointBase string
}

// DetectTenantForInit resolves the platform tenant for `init` in one pass. Returns
// nil (no error) when no tenant applies — a team with no/ambiguous tenant, or the
// disable env — so init just uses its normal flow. A non-nil error means a tenant
// was detected but its session couldn't be minted: the platform sync can't run,
// so init should fail now rather than scaffold a spec that breaks at sync time.
func DetectTenantForInit(ctx context.Context, logger zerolog.Logger, cloudToken, teamName string) (*TenantInit, error) {
	if os.Getenv(envDisable) == "1" {
		return nil, nil
	}
	// A directly supplied cqpd_ token already identifies the tenant and carries its
	// URL; pins + recommended tables are fetched with that same token.
	if t := platformToken(); t != "" {
		u := apiURLFromToken(t)
		if u == "" {
			return &TenantInit{}, nil // tenant present, but no url to reach endpoints
		}
		// This same call validates the token: an explicit env token that the tenant
		// rejects (expired/revoked/wrong tenant) means the platform sync can't run,
		// so fail now rather than scaffold a spec that 401s at sync time. Other
		// failures stay best-effort (nil pins → hub latest).
		pins, err := fetchSupportedSourceVersions(ctx, logger, t, u)
		if errors.Is(err, errPlatformUnauthorized) {
			return nil, unauthorizedTokenError()
		}
		return &TenantInit{
			APIURL:               u,
			PinnedSourceVersions: pins,
			token:                t,
			endpointBase:         u,
		}, nil
	}
	cl, tenant, resolved := resolveCloudTenant(ctx, logger, cloudToken, teamName)
	if !resolved {
		return nil, nil
	}
	// Mint once to reach /external-syncs/*. A detected tenant that can't mint a
	// session can't run a platform sync either, so surface it rather than emit a
	// spec that fails later.
	session, _, err := mintSession(ctx, cl, tenant)
	if err != nil {
		return nil, fmt.Errorf("mint platform destination session for tenant %s: %w", tenant.TenantId.String(), err)
	}
	// The session was just minted, so a 401 here would be a server anomaly, not
	// user-fixable — keep pins best-effort rather than failing init.
	pins, _ := fetchSupportedSourceVersions(ctx, logger, session.Token, session.ApiUrl)
	return &TenantInit{
		APIURL:               "https://" + tenant.Host,
		PinnedSourceVersions: pins,
		token:                session.Token,
		endpointBase:         session.ApiUrl,
	}, nil
}

// RecommendedTables returns the tables the platform recommends syncing for the
// given source plugin path, from GET /external-syncs/recommended-tables, reusing
// the session resolved by DetectTenantForInit (no extra mint). Best-effort: nil
// when there's no session or the lookup fails / returns nothing, so init falls
// back to `tables: ['*']`.
func (ti *TenantInit) RecommendedTables(ctx context.Context, logger zerolog.Logger, sourcePath string) []string {
	if ti == nil || ti.token == "" || ti.endpointBase == "" || sourcePath == "" {
		return nil
	}
	base := externalSyncsURL(ti.endpointBase, "/external-syncs/recommended-tables")
	url := base + "?path=" + neturl.QueryEscape(sourcePath)

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform: failed to build recommended-tables request")
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+ti.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform: recommended-tables lookup failed")
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Debug().Int("status", resp.StatusCode).Str("url", url).Msg("platform: recommended-tables returned non-200")
		return nil
	}
	var body struct {
		Tables []string `json:"tables"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		logger.Debug().Err(err).Msg("platform: failed to decode recommended-tables")
		return nil
	}
	return body.Tables
}

// resolveCloudTenant resolves the single active tenant for the team (the
// logged-in path) plus an API client to act on it, in one enumeration. Uses the
// same selection as auto-injection (resolveTenant): the only active tenant, or
// the CQ_PLATFORM_TENANT_ID match. Best-effort: ok=false when there are no creds,
// no active tenant, or an ambiguous set with no override.
func resolveCloudTenant(ctx context.Context, logger zerolog.Logger, cloudToken, teamName string) (cl *cloudquery_api.ClientWithResponses, tenant cloudquery_api.PlatformTenantSummary, ok bool) {
	if cloudToken == "" || teamName == "" {
		return nil, cloudquery_api.PlatformTenantSummary{}, false
	}
	cl, err := api.NewClient(cloudToken)
	if err != nil {
		logger.Debug().Err(err).Msg("platform: api client init failed")
		return nil, cloudquery_api.PlatformTenantSummary{}, false
	}
	tenants, err := activeTenants(ctx, cl, teamName)
	if err != nil || len(tenants) == 0 {
		return nil, cloudquery_api.PlatformTenantSummary{}, false
	}
	tenant, err = resolveTenant(tenants)
	if err != nil {
		return nil, cloudquery_api.PlatformTenantSummary{}, false
	}
	return cl, tenant, true
}

// apiURLFromToken reads the api_url (`u`) claim from a cqpd_ token's payload
// without verifying the signature. Returns "" for a malformed token or one that
// carries no url. Mirrors the destination plugin's decoder (separate repos).
func apiURLFromToken(token string) string {
	apiURL, _ := decodeCQPDClaims(token)
	return apiURL
}

// platformToken returns the platform-destination cqpd_ token from its explicit
// CQ_PLATFORM_TOKEN env, or from CLOUDQUERY_API_KEY when that holds a cqpd_ (the
// standard credential env doubling as the platform token). "" when neither
// applies — i.e. no headless platform-destination token is configured. One
// helper so download, injection, and tenant detection treat both envs alike.
func platformToken() string {
	if t := os.Getenv(EnvPlatformToken); t != "" {
		warnTeamMismatchOnce(t)
		return t
	}
	if k := os.Getenv(cqapiauth.EnvVarCloudQueryAPIKey); strings.HasPrefix(k, cqpdPrefix) {
		warnTeamMismatchOnce(k)
		return k
	}
	return ""
}

// warnTeamMismatchOnce surfaces the team mismatch at the platformToken()
// chokepoint so every consumer (init, validate-config, sync, migrate) warns
// without each entry point remembering to — and only once per run, since
// several of them read the token during a single command.
var teamMismatchOnce gosync.Once

func warnTeamMismatchOnce(token string) {
	teamMismatchOnce.Do(func() {
		// stderr only — a parallel zlog.Warn() would double-print when console
		// logging is enabled (both land on the terminal). Matches login/logout,
		// which print the credential warnings once via cmd.Printf.
		if msg := teamMismatchWarning(TeamFromToken(token)); msg != "" {
			fmt.Fprintln(os.Stderr, msg)
		}
	})
}

// PropagatePluginCredential makes a headless cqpd_ token available to spawned
// source/destination plugins as CLOUDQUERY_API_KEY. Plugins validate premium
// tables and report usage against cloud using that env var (read from their own
// process env, which in local runs they inherit from us); a cqpd_ supplied only
// via CQ_PLATFORM_TOKEN would otherwise be invisible to them. No-op for a
// non-cqpd_ token, and it never overwrites an existing CLOUDQUERY_API_KEY — a
// user who set their own (e.g. a team key) keeps it for plugin auth.
func PropagatePluginCredential(token string) {
	if strings.HasPrefix(token, cqpdPrefix) && os.Getenv("CLOUDQUERY_API_KEY") == "" {
		_ = os.Setenv("CLOUDQUERY_API_KEY", token)
	}
}

// recommendedVersionFromWhoami asks the tenant's platform — whoami, reached via
// the token's api_url (`u`) and authed with the token — for the recommended
// destination plugin version. It lets the headless flow pin the right plugin
// without a session mint (which is where the non-headless flow gets it).
// Best-effort: "" on any failure, so the caller falls back to the CLI default.
func recommendedVersionFromWhoami(ctx context.Context, logger zerolog.Logger, cqpdToken string) string {
	apiURL := apiURLFromToken(cqpdToken)
	if apiURL == "" {
		logger.Debug().Msg("platform destination: token carries no api_url; skipping whoami version lookup, using default")
		return ""
	}
	url := externalSyncsURL(apiURL, "/external-syncs/whoami")
	logger.Debug().Str("url", url).Msg("platform destination: looking up recommended plugin version via whoami")

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform destination: failed to build whoami request; using default")
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+cqpdToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform destination: whoami lookup for recommended plugin version failed; using default")
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Debug().Int("status", resp.StatusCode).Str("url", url).Msg("platform destination: whoami returned non-200 for version lookup; using default")
		return ""
	}
	var body struct {
		PluginVersion *string `json:"plugin_version"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		logger.Debug().Err(err).Msg("platform destination: failed to decode whoami response; using default")
		return ""
	}
	if body.PluginVersion == nil {
		logger.Debug().Msg("platform destination: whoami returned no plugin_version; using default")
		return ""
	}
	logger.Debug().Str("plugin_version", *body.PluginVersion).Msg("platform destination: pinning recommended plugin version from whoami")
	return *body.PluginVersion
}

// externalSyncsURL builds the URL for an /external-syncs/* endpoint from a
// tenant's API base, which may or may not already carry the /api suffix (a
// minted session returns the bare host; the /external-syncs/* routes live under
// /api).
func externalSyncsURL(apiURL, path string) string {
	base := strings.TrimRight(apiURL, "/")
	if !strings.HasSuffix(base, "/api") {
		base += "/api"
	}
	return base + path
}

// resolvePlatformSession returns a cqpd_ token + tenant API base URL for reaching
// the /external-syncs/* endpoints, and whether it's a direct env token
// (CQ_PLATFORM_TOKEN / cqpd_ CLOUDQUERY_API_KEY, using its `u` claim) vs a
// freshly-minted session for the logged-in team's tenant. ok=false when there's no
// platform tenant or resolution fails, so callers fall back. `direct` matters for
// auth failures: a rejected env token is user-fixable and is the same token a sync
// reuses, whereas a rejected fresh-minted token is a transient server anomaly (a
// sync mints its own).
func resolvePlatformSession(ctx context.Context, logger zerolog.Logger, cloudToken, teamName string) (cqpdToken, apiURL string, direct, ok bool) {
	if os.Getenv(envDisable) == "1" {
		return "", "", false, false
	}
	if t := platformToken(); t != "" {
		u := apiURLFromToken(t)
		if u == "" {
			logger.Debug().Msg("platform: direct token carries no api_url; cannot reach external-syncs endpoints")
			return "", "", false, false
		}
		return t, u, true, true
	}
	cl, tenant, ok := resolveCloudTenant(ctx, logger, cloudToken, teamName)
	if !ok {
		return "", "", false, false
	}
	session, _, err := mintSession(ctx, cl, tenant)
	if err != nil {
		logger.Debug().Err(err).Msg("platform: session mint failed; cannot fetch pinned versions")
		return "", "", false, false
	}
	return session.Token, session.ApiUrl, false, true
}

// PinnedSourceVersions returns the platform-pinned source plugin versions
// (plugin path -> semver) the caller's tenant will accept, from
// GET /external-syncs/supported-source-versions. `init` scaffolds these so the
// generated spec matches what the tenant accepts, and `validate-config` gates
// against them — the same window CreateExternalSync enforces at sync time.
//
// Returns errPlatformUnauthorized only when a direct env token is rejected — the
// same user-fixable failure `init` fails on, and one a sync would hit too, so the
// gate shouldn't pass clean. Any other failure (no tenant, minted-session 401,
// network) yields (nil, nil): best-effort, so the gate stays fail-open.
func PinnedSourceVersions(ctx context.Context, logger zerolog.Logger, cloudToken, teamName string) (map[string]string, error) {
	token, apiURL, direct, ok := resolvePlatformSession(ctx, logger, cloudToken, teamName)
	if !ok {
		return nil, nil
	}
	versions, err := fetchSupportedSourceVersions(ctx, logger, token, apiURL)
	if errors.Is(err, errPlatformUnauthorized) && direct {
		return nil, errPlatformUnauthorized
	}
	return versions, nil
}

// errPlatformUnauthorized means the platform rejected the token (401/403) —
// expired, revoked, secret-rotated, or scoped to a dead tenant. Distinguished
// from other failures so an explicit CQ_PLATFORM_TOKEN can fail init early rather
// than silently degrade; every other failure stays best-effort (nil, nil).
var errPlatformUnauthorized = errors.New("platform rejected the token")

// unauthorizedTokenError is the user-facing error for a rejected direct env
// token, shared by init (DetectTenantForInit) and validate-config (GateSources)
// so both report the same actionable fix.
func unauthorizedTokenError() error {
	return fmt.Errorf("the platform token in your environment (%s, or a %s %s) was rejected — it is likely expired; mint a fresh token or unset it", EnvPlatformToken, cqpdPrefix, cqapiauth.EnvVarCloudQueryAPIKey)
}

// fetchSupportedSourceVersions GETs /external-syncs/supported-source-versions
// with an already-minted cqpd_ token and returns the pinned path->version map.
// Returns errPlatformUnauthorized on a 401/403; (nil, nil) on any other failure
// (best-effort). Shared by PinnedSourceVersions and the init tenant-detection
// path, which resolve the token/apiURL differently and treat the auth error
// differently.
func fetchSupportedSourceVersions(ctx context.Context, logger zerolog.Logger, cqpdToken, apiURL string) (map[string]string, error) {
	url := externalSyncsURL(apiURL, "/external-syncs/supported-source-versions")

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform: failed to build supported-source-versions request")
		return nil, nil
	}
	req.Header.Set("Authorization", "Bearer "+cqpdToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug().Err(err).Str("url", url).Msg("platform: supported-source-versions lookup failed")
		return nil, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		logger.Debug().Int("status", resp.StatusCode).Str("url", url).Msg("platform: supported-source-versions rejected the token")
		return nil, errPlatformUnauthorized
	}
	if resp.StatusCode != http.StatusOK {
		logger.Debug().Int("status", resp.StatusCode).Str("url", url).Msg("platform: supported-source-versions returned non-200")
		return nil, nil
	}
	var versions map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&versions); err != nil {
		logger.Debug().Err(err).Msg("platform: failed to decode supported-source-versions")
		return nil, nil
	}
	return versions, nil
}

// AnySourceTargetsPlatform reports whether any source opts into the platform
// destination (lists its reserved name in `destinations`). Exported so
// validate-config gates only the platform-bound sources — the same set a real
// sync would upload.
func AnySourceTargetsPlatform(sources []*specs.Source) bool {
	return anySourceTargetsPlatform(sources)
}

// sourceVersionSupported mirrors the server-side gate
// (api/externalsyncs.sourceSupported): a source version is accepted iff its
// plugin path is pinned and the version is the same major and not newer than the
// pin (a schema subset the corpus already knows).
func sourceVersionSupported(path, version string, pinned map[string]string) bool {
	pinnedVersion := pinned[path]
	if pinnedVersion == "" {
		return false
	}
	v, err := semver.NewVersion(version)
	if err != nil {
		return false
	}
	p, err := semver.NewVersion(pinnedVersion)
	if err != nil {
		return false
	}
	return v.Major() == p.Major() && !v.GreaterThan(p)
}

// GateSources returns an error naming any platform-targeted source whose version
// the tenant can't ingest — the same window CreateExternalSync enforces — so
// validate-config fails before a sync would. Fail-open (nil) when no source
// targets platform, or the pinned versions can't be resolved (mirrors the server,
// which opens the sync when versions are unavailable). The message matches the
// server gate's so users see the same text at validate and sync time.
func GateSources(ctx context.Context, logger zerolog.Logger, cloudToken, teamName string, sources []*specs.Source) error {
	targeted := make([]*specs.Source, 0, len(sources))
	for _, s := range sources {
		if slices.Contains(s.Destinations, destinationName) {
			targeted = append(targeted, s)
		}
	}
	if len(targeted) == 0 {
		return nil
	}
	pinned, err := PinnedSourceVersions(ctx, logger, cloudToken, teamName)
	if errors.Is(err, errPlatformUnauthorized) {
		// A rejected env token would 401 the sync too, so don't pass clean — this
		// is exactly what validate-config promises to catch.
		return unauthorizedTokenError()
	}
	if len(pinned) == 0 {
		logger.Debug().Msg("platform: pinned source versions unavailable; skipping version gate")
		return nil
	}
	unsupported := make([]string, 0, len(targeted))
	for _, s := range targeted {
		if sourceVersionSupported(s.Path, s.Version, pinned) {
			continue
		}
		if pv := pinned[s.Path]; pv != "" {
			unsupported = append(unsupported, fmt.Sprintf("%s (supported version: %s)", s.Name, pv))
		} else {
			unsupported = append(unsupported, fmt.Sprintf("%s (not a supported source)", s.Name))
		}
	}
	if len(unsupported) == 0 {
		return nil
	}
	return fmt.Errorf("unsupported source plugin version(s): %s", strings.Join(unsupported, ", "))
}

// DownloadAuth resolves the credential and team used to download (and meter)
// plugins. In the headless platform-destination flow — a cqpd_ token in
// CQ_PLATFORM_TOKEN or CLOUDQUERY_API_KEY (see platformToken) — it returns that
// token and the team from its `tm` claim, so a sync needs no `cloudquery login`;
// managedplugin then uses the team-scoped download endpoint and the team is
// recorded server-side. The cqpd_ is syncs-scoped and can't enumerate teams, so
// the team must come from the claim, not GetTeamForToken. Otherwise it falls
// back to the cloud login / team-API-key token and its team. Centralizing the
// env read keeps sync and migrate from drifting.
func DownloadAuth(ctx context.Context, logger zerolog.Logger, sources []*specs.Source, destinations []*specs.Destination, transformers []*specs.Transformer) (token, team string, err error) {
	if t := platformToken(); t != "" {
		return t, TeamFromToken(t), nil
	}
	authToken, err := cqauth.GetAuthTokenIfNeeded(logger, sources, destinations, transformers)
	if err != nil {
		return "", "", fmt.Errorf("failed to get auth token: %w", err)
	}
	teamName, err := cqauth.GetTeamForToken(ctx, authToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to get team name from token: %w", err)
	}
	return authToken.Value, teamName, nil
}

// teamMismatchWarning reports when a headless platform token routes plugin
// downloads and usage to a different team than the one the user is switched
// to (`cloudquery switch`) — the token's tm claim wins silently otherwise.
// Returns "" when there is nothing to warn about: no tm claim, no configured
// team, or the two match.
func teamMismatchWarning(tokenTeam string) string {
	if tokenTeam == "" {
		return ""
	}
	configTeam, err := cqconfig.GetValue("team")
	if err != nil || configTeam == "" || configTeam == tokenTeam {
		return ""
	}
	return fmt.Sprintf("Warning: the platform token in the environment (%s or a %s %s) belongs to team %q; plugin downloads and usage will be attributed to that team, not your currently selected team %q",
		EnvPlatformToken, cqpdPrefix, cqapiauth.EnvVarCloudQueryAPIKey, tokenTeam, configTeam)
}

// TeamFromToken returns the cloud team (`tm` claim) embedded in a cqpd_ token,
// or "" if absent/malformed. The CLI uses it to target the team-scoped
// plugin-download / usage endpoints (and premium entitlement) from the token
// alone — no `cloudquery login`. Read without verifying the signature; cloud
// still authenticates the token.
func TeamFromToken(token string) string {
	_, team := decodeCQPDClaims(token)
	return team
}

// decodeCQPDClaims reads the unverified claims payload of a cqpd_ token. The CLI
// only needs routing/identity hints (api_url, team) to decide where and as whom
// to call; the platform still authenticates the token. Wire format is
// "cqpd_" + base64url(claimsJSON) + "." + base64url(sig). Returns empty strings
// for a malformed or non-cqpd_ token. Mirrors the destination plugin's decoder
// (separate repos — keep the claim keys in sync).
func decodeCQPDClaims(token string) (apiURL, team string) {
	rest, ok := strings.CutPrefix(token, cqpdPrefix)
	if !ok {
		return "", ""
	}
	enc, _, ok := strings.Cut(rest, ".")
	if !ok {
		return "", ""
	}
	payload, err := base64.RawURLEncoding.DecodeString(enc)
	if err != nil {
		return "", ""
	}
	var claims struct {
		APIURL string `json:"u"`
		Team   string `json:"tm"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", ""
	}
	return claims.APIURL, claims.Team
}

// MaybeInjectDestination injects a `platform` destination carrying a freshly
// minted cqpd_ token — but only when the spec opts in by listing `platform` in
// a source's `destinations`. If the user already declares a `platform`
// destination themselves (e.g. for debugging), theirs is used as-is. With no
// opt-in, or on any credential/tenant failure, the spec is returned unchanged.
func MaybeInjectDestination(ctx context.Context, logger zerolog.Logger, token, teamName string, sources []*specs.Source, destinations []*specs.Destination) ([]*specs.Destination, error) {
	if os.Getenv(envDisable) == "1" {
		return destinations, nil
	}
	if env.IsCloud() {
		return destinations, nil
	}

	// Opt-in only: inject solely when a source targets the platform destination.
	// No source references it → nothing to do (no cloud calls, no surprise
	// dual-write).
	if !anySourceTargetsPlatform(sources) {
		return destinations, nil
	}
	// The user defined the `platform` destination themselves (debugging/override)
	// — respect it, don't inject over it.
	if hasPlatformDestination(destinations) {
		return destinations, nil
	}

	// Direct token: a pre-minted cqpd_ token supplied via env (CQ_PLATFORM_TOKEN
	// or a cqpd_ in CLOUDQUERY_API_KEY) injects the destination without cloud
	// login, tenant discovery or a session mint — the token already identifies
	// the tenant and carries its API URL.
	if t := platformToken(); t != "" {
		// Recommended plugin version: the env override wins (so skip the lookup),
		// otherwise ask the platform's whoami so the headless flow pins the right
		// version — the non-headless path gets this from the session mint instead.
		recommendedVersion := ""
		if os.Getenv(envPluginVersion) == "" {
			recommendedVersion = recommendedVersionFromWhoami(ctx, logger, t)
		}
		// No tenant id: the direct path doesn't parse the token's claims.
		return injectPlatformDestination(logger, destinations, sources, t, recommendedVersion, ""), nil
	}

	// The caller only fetches a token for cloudquery-registry specs; resolve
	// directly so source-only specs can still inject. Failure just skips.
	if token == "" {
		var err error
		if token, teamName, err = resolveCredentials(ctx); err != nil {
			logger.Debug().Err(err).Msg("platform destination: credentials unavailable, skipping auto-injection")
			return destinations, nil
		}
	}
	if token == "" || teamName == "" {
		return destinations, nil
	}

	cl, err := api.NewClient(token)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: api client init failed, skipping auto-injection")
		return destinations, nil
	}

	tenants, err := activeTenants(ctx, cl, teamName)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: tenant discovery failed, skipping auto-injection")
		return destinations, nil
	}
	tenant, err := resolveTenant(tenants)
	switch {
	case errors.Is(err, errNoActiveTenant):
		// No platform tenant for this team — nothing to inject; skip silently.
		return destinations, nil
	case err != nil:
		// Several active tenants and no usable CQ_PLATFORM_TENANT_ID. A source
		// opted into `platform`, so don't silently drop it — fail with the Hint.
		return destinations, err
	}

	session, platformPluginVersion, err := mintSession(ctx, cl, tenant)
	if err != nil {
		// A source opted into `platform` and we found its tenant, but the session
		// mint failed — the sync can't write to Platform. Fail with the reason
		// rather than silently dropping the opt-in and running a sync whose source
		// targets a destination that was never injected. (Same stance as the
		// ambiguous-tenant case above.) Reachable only past the opt-in guard, so a
		// non-platform sync never gets here.
		return destinations, fmt.Errorf("failed to set up CloudQuery Platform destination for tenant %s: %w", tenant.TenantId.String(), err)
	}

	return injectPlatformDestination(logger, destinations, sources, session.Token, platformPluginVersion, tenant.TenantId.String()), nil
}

// anySourceTargetsPlatform reports whether any source opts into the platform
// destination by listing its reserved name in `destinations`.
func anySourceTargetsPlatform(sources []*specs.Source) bool {
	for _, s := range sources {
		if slices.Contains(s.Destinations, destinationName) {
			return true
		}
	}
	return false
}

// hasPlatformDestination reports whether the spec already declares a `platform`
// destination (a user-provided one the CLI must not overwrite).
func hasPlatformDestination(destinations []*specs.Destination) bool {
	for _, d := range destinations {
		if d.Name == destinationName {
			return true
		}
	}
	return false
}

// injectPlatformDestination appends the reserved `platform` destination carrying
// the cqpd_ token. The caller guarantees a source already targets it (the opt-in)
// and that no user-defined `platform` destination exists. recommendedVersion,
// when set and not overridden by the env, pins the plugin version; an unknown
// registry skips injection (returns the spec unchanged).
func injectPlatformDestination(logger zerolog.Logger, destinations []*specs.Destination, sources []*specs.Source, token, recommendedVersion, tenantID string) []*specs.Destination {
	plugin := pluginCoords()
	// Version precedence: env override > platform-pinned > CLI default.
	// pluginCoords() already applied the env override (or the default), so only
	// let the platform's pin win when the env override isn't set.
	if os.Getenv(envPluginVersion) == "" && recommendedVersion != "" {
		plugin.Version = recommendedVersion
	}
	parsedRegistry, err := specs.RegistryFromString(plugin.Registry)
	if err != nil {
		logger.Warn().Err(err).Str("registry", plugin.Registry).Msg("platform destination: unknown plugin registry; skipping auto-injection")
		return destinations
	}

	// Report the path+version of the sources that target platform so it can
	// reject (before any upload) versions the asset view can't process.
	sourceVersions := make([]sourceVersion, 0, len(sources))
	for _, s := range sources {
		if slices.Contains(s.Destinations, destinationName) {
			sourceVersions = append(sourceVersions, sourceVersion{Name: s.Name, Path: s.Path, Version: s.Version})
		}
	}
	dest := &specs.Destination{
		Metadata: specs.Metadata{
			Name:     destinationName,
			Path:     plugin.Path,
			Registry: parsedRegistry,
			Version:  plugin.Version,
		},
		SyncSummary: true,
		// sync_group_id is rejected with the default overwrite-delete-stale mode.
		WriteMode: specs.WriteModeAppend,
		// Unique per invocation so concurrent runs don't wipe each other's rows.
		SyncGroupId: strconv.FormatUint(allocateSyncGroupID(time.Now()), 10),
		Spec: map[string]any{
			// api_url is omitted: the cqpd_ token carries the tenant's API URL,
			// and the platform destination derives it from the token.
			"token":           token,
			"source_versions": sourceVersions,
		},
	}
	dest.SetDefaults()
	destinations = append(destinations, dest)

	evt := logger.Info().
		Str("registry", plugin.Registry).
		Str("path", plugin.Path).
		Str("version", plugin.Version)
	if tenantID != "" {
		evt = evt.Str("tenant_id", tenantID)
	}
	evt.Msg("auto-injected platform destination")
	return destinations
}

// resolveCredentials fetches a token and team for best-effort injection when
// the sync command didn't authenticate. Overridable in tests.
var resolveCredentials = func(ctx context.Context) (token, team string, err error) {
	tok, err := cqapiauth.NewTokenClient().GetToken()
	if err != nil {
		return "", "", err
	}
	team, err = cqauth.GetTeamForToken(ctx, tok)
	if err != nil {
		return "", "", err
	}
	return tok.Value, team, nil
}

var (
	// errNoActiveTenant: the team has no active platform tenant to inject into.
	// A no-op for auto-injection (skip silently) — not a user error.
	errNoActiveTenant = errors.New("no active platform tenant")
	// errAmbiguousTenant: several active tenants and CQ_PLATFORM_TENANT_ID isn't
	// set to one of them. Surfaced to the user as a Hint (it's env-fixable) so an
	// explicit `platform` opt-in isn't silently dropped. Callers match it with
	// errors.Is to tell "ambiguous" apart from "none".
	errAmbiguousTenant = errors.New("multiple active CloudQuery Platform tenants for this team")
)

// resolveTenant picks the single tenant to act on: the only active one, or the
// CQ_PLATFORM_TENANT_ID match when several are active. Returns errNoActiveTenant
// when there are none, and an errAmbiguousTenant-wrapped error (carrying an
// actionable Hint) when several are active without a matching override. Pure (no
// logging) so every call site shares one decision.
func resolveTenant(tenants []cloudquery_api.PlatformTenantSummary) (cloudquery_api.PlatformTenantSummary, error) {
	switch len(tenants) {
	case 0:
		return cloudquery_api.PlatformTenantSummary{}, errNoActiveTenant
	case 1:
		return tenants[0], nil
	}
	want := os.Getenv(envTenantID)
	if want == "" {
		return cloudquery_api.PlatformTenantSummary{}, fmt.Errorf("%w. Hint: set %s to the tenant id you want to sync to", errAmbiguousTenant, envTenantID)
	}
	for _, t := range tenants {
		if t.TenantId.String() == want {
			return t, nil
		}
	}
	return cloudquery_api.PlatformTenantSummary{}, fmt.Errorf("%w: %s=%s matches none of them. Hint: set it to one of the team's active tenant ids", errAmbiguousTenant, envTenantID, want)
}

// YYYYMMDDhhmmssfff — same shape platform/syncs-transformer uses, so
// external-sync rows share the keyspace.
func allocateSyncGroupID(now time.Time) uint64 {
	t := now.UTC()
	base := t.Format("20060102150405") + fmt.Sprintf("%03d", t.Nanosecond()/1e6)
	u, _ := strconv.ParseUint(base, 10, 64)
	return u
}

func activeTenants(ctx context.Context, cl *cloudquery_api.ClientWithResponses, teamName string) ([]cloudquery_api.PlatformTenantSummary, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	resp, err := cl.ListUserPlatformTenantsWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, fmt.Errorf("unexpected status %d listing platform tenants: %s", resp.StatusCode(), strings.TrimSpace(string(resp.Body)))
	}
	active := make([]cloudquery_api.PlatformTenantSummary, 0, len(resp.JSON200.Items))
	for _, t := range resp.JSON200.Items {
		if t.TeamName == teamName && slices.Contains(injectableStatuses, t.Status) {
			active = append(active, t)
		}
	}
	return active, nil
}

func mintSession(ctx context.Context, cl *cloudquery_api.ClientWithResponses, tenant cloudquery_api.PlatformTenantSummary) (session *cloudquery_api.CreatePlatformDestinationSession201Response, pluginVersion string, err error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	resp, err := cl.CreatePlatformDestinationSessionWithResponse(ctx, cloudquery_api.CreatePlatformDestinationSessionRequest{TenantId: tenant.TenantId})
	if err != nil {
		return nil, "", err
	}
	if resp.JSON201 == nil {
		return nil, "", fmt.Errorf("unexpected status %d minting platform destination session: %s", resp.StatusCode(), strings.TrimSpace(string(resp.Body)))
	}
	if resp.JSON201.Token == "" || resp.JSON201.ApiUrl == "" {
		return nil, "", errors.New("platform destination session response missing token or api_url")
	}
	// plugin_version lets the platform pin the destination plugin version without
	// a CLI release. Optional: nil/empty → caller falls back to the CLI default.
	if resp.JSON201.PluginVersion != nil {
		pluginVersion = *resp.JSON201.PluginVersion
	}
	return resp.JSON201, pluginVersion, nil
}

// IsInjectedDestination reports whether a destination spec name is the
// auto-injected platform destination (a reserved name).
func IsInjectedDestination(name string) bool {
	return name == destinationName
}

// CleanInitError turns the gRPC-wrapped plugin-init error from the platform
// destination into a human-readable message: it unwraps the gRPC status (drops
// the "rpc error: code = ... desc =" prefix) and strips the plugin-sdk wrapper
// prefixes, leaving the destination's own message (e.g. the platform's 422 text).
// Scope this to the platform destination (see IsInjectedDestination) — the
// stripped prefixes are specific to that plugin's init path.
func CleanInitError(err error) string {
	msg := status.Convert(err).Message()
	for _, prefix := range []string{"failed to init plugin: ", "failed to initialize client: "} {
		msg = strings.TrimPrefix(msg, prefix)
	}
	return msg
}
