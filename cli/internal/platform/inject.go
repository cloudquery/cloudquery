// Package platform auto-injects a platform destination into syncs for teams
// with an active platform tenant.
package platform

import (
	"context"
	"errors"
	"fmt"
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

	envPluginRegistry = "CQ_PLATFORM_PLUGIN_REGISTRY"
	envPluginPath     = "CQ_PLATFORM_PLUGIN_PATH"
	envPluginVersion  = "CQ_PLATFORM_PLUGIN_VERSION"

	destinationName = "platform"

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
	Version:  "v1.0.0",
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

// MaybeInjectDestination appends a `platform` destination carrying a freshly
// minted cqpd_ token when the team has an active platform tenant. Tenant/network
// failures skip injection silently; a pre-existing `platform` destination is a
// hard error rather than a silent overwrite.
func MaybeInjectDestination(ctx context.Context, logger zerolog.Logger, token, teamName string, sources []*specs.Source, destinations []*specs.Destination) ([]*specs.Destination, error) {
	if os.Getenv(envDisable) == "1" {
		return destinations, nil
	}
	if env.IsCloud() {
		return destinations, nil
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
	tenant, ok := selectTenant(logger, tenants)
	if !ok {
		return destinations, nil
	}

	// Injecting: a pre-existing `platform` destination collides with the
	// reserved name — fail rather than overwrite.
	for _, d := range destinations {
		if d.Name == destinationName {
			return destinations, fmt.Errorf("a destination named %q already exists, but this name is reserved for the auto-injected CloudQuery Platform destination; remove it from your spec", destinationName)
		}
	}

	session, platformPluginVersion, err := mintSession(ctx, cl, tenant)
	if err != nil {
		logger.Warn().Err(err).Str("tenant_id", tenant.TenantId.String()).Msg("platform destination: session mint failed, skipping auto-injection")
		return destinations, nil
	}

	plugin := pluginCoords()
	// Version precedence: env override > platform-pinned > CLI default.
	// pluginCoords() already applied the env override (or the default), so only
	// let the platform's pin win when the env override isn't set.
	if os.Getenv(envPluginVersion) == "" && platformPluginVersion != "" {
		plugin.Version = platformPluginVersion
	}
	parsedRegistry, err := specs.RegistryFromString(plugin.Registry)
	if err != nil {
		logger.Warn().Err(err).Str("registry", plugin.Registry).Msg("platform destination: unknown plugin registry; skipping auto-injection")
		return destinations, nil
	}

	// Report each source's plugin path+version so the platform can reject (before
	// any upload) sources whose version the asset view can't process.
	sourceVersions := make([]sourceVersion, 0, len(sources))
	for _, s := range sources {
		sourceVersions = append(sourceVersions, sourceVersion{Name: s.Name, Path: s.Path, Version: s.Version})
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
			"token":           session.Token,
			"source_versions": sourceVersions,
		},
	}
	dest.SetDefaults()
	destinations = append(destinations, dest)

	for _, s := range sources {
		if !slices.Contains(s.Destinations, destinationName) {
			s.Destinations = append(s.Destinations, destinationName)
		}
	}
	logger.Info().
		Str("platform_url", session.ApiUrl).
		Str("tenant_id", tenant.TenantId.String()).
		Str("registry", plugin.Registry).
		Str("path", plugin.Path).
		Str("version", plugin.Version).
		Msg("auto-injected platform destination")
	return destinations, nil
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

func selectTenant(logger zerolog.Logger, tenants []cloudquery_api.PlatformTenantSummary) (cloudquery_api.PlatformTenantSummary, bool) {
	switch len(tenants) {
	case 0:
		return cloudquery_api.PlatformTenantSummary{}, false
	case 1:
		return tenants[0], true
	}
	want := os.Getenv(envTenantID)
	if want == "" {
		logger.Warn().Int("tenants", len(tenants)).Msgf("platform destination: team has multiple active tenants; set %s to choose one, skipping auto-injection", envTenantID)
		return cloudquery_api.PlatformTenantSummary{}, false
	}
	for _, t := range tenants {
		if t.TenantId.String() == want {
			return t, true
		}
	}
	logger.Warn().Str("tenant_id", want).Msgf("platform destination: %s does not match any active tenant, skipping auto-injection", envTenantID)
	return cloudquery_api.PlatformTenantSummary{}, false
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
