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
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
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

// DetectTenant reports the CloudQuery Platform tenant a sync would auto-inject
// into — for commands (e.g. init) that want to skip the destination and tell
// the user where data lands. ok is true when a tenant is found; apiURL is its
// base URL (host only, no /api), which may be empty if a directly supplied
// CQ_PLATFORM_TOKEN predates url-carrying tokens. Best-effort: any lookup
// failure returns ("", false) so callers fall back to normal behavior.
func DetectTenant(ctx context.Context, token, teamName string) (apiURL string, ok bool) {
	if os.Getenv(envDisable) == "1" {
		return "", false
	}
	// A directly supplied cqpd_ token already identifies the tenant; read its URL.
	if t := platformToken(); t != "" {
		return apiURLFromToken(t), true
	}
	if token == "" || teamName == "" {
		return "", false
	}
	cl, err := api.NewClient(token)
	if err != nil {
		return "", false
	}
	tenants, err := activeTenants(ctx, cl, teamName)
	if err != nil || len(tenants) == 0 {
		return "", false
	}
	// Use the same selection as auto-injection (resolveTenant): the only active
	// tenant, or the CQ_PLATFORM_TENANT_ID match. None or ambiguous (several
	// active, no override) → report nothing; init is informational, so it never
	// errors here, and it won't point at a tenant a real sync would skip.
	tenant, err := resolveTenant(tenants)
	if err != nil {
		return "", false
	}
	return "https://" + tenant.Host, true
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
		return t
	}
	if k := os.Getenv("CLOUDQUERY_API_KEY"); strings.HasPrefix(k, cqpdPrefix) {
		return k
	}
	return ""
}

// recommendedVersionFromWhoami asks the tenant's platform — whoami, reached via
// the token's api_url (`u`) and authed with the token — for the recommended
// destination plugin version. It lets the headless flow pin the right plugin
// without a session mint (which is where the non-headless flow gets it).
// Best-effort: "" on any failure, so the caller falls back to the CLI default.
func recommendedVersionFromWhoami(ctx context.Context, logger zerolog.Logger, cqpdToken string) string {
	apiURL := apiURLFromToken(cqpdToken)
	if apiURL == "" {
		return ""
	}
	base := strings.TrimRight(apiURL, "/")
	if !strings.HasSuffix(base, "/api") { // /external-syncs/* is served under /api
		base += "/api"
	}
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+"/external-syncs/whoami", nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+cqpdToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: whoami lookup for recommended plugin version failed; using default")
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	var body struct {
		PluginVersion *string `json:"plugin_version"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil || body.PluginVersion == nil {
		return ""
	}
	return *body.PluginVersion
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
		logger.Warn().Err(err).Str("tenant_id", tenant.TenantId.String()).Msg("platform destination: session mint failed, skipping auto-injection")
		return destinations, nil
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
