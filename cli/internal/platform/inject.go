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

// platformAPIURL is the base URL the destination plugin uses for
// /external-syncs/*. Those endpoints are served under /api, which is appended
// to the minted session's api_url unless already present.
func platformAPIURL(sessionURL string) string {
	url := strings.TrimRight(sessionURL, "/")
	if !strings.HasSuffix(url, "/api") {
		url += "/api"
	}
	return url
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

// MaybeInjectDestination ensures a `platform` destination exists in the spec
// when the team has an active platform tenant, carrying a freshly minted
// cqpd_ token — the cloud credential itself never reaches the plugin, and
// failures never break the sync.
func MaybeInjectDestination(ctx context.Context, logger zerolog.Logger, token, teamName string, sources []*specs.Source, destinations []*specs.Destination) []*specs.Destination {
	if os.Getenv(envDisable) == "1" {
		return destinations
	}
	// Cloud-run syncs compose their spec server-side.
	if env.IsCloud() {
		return destinations
	}
	// The caller only fetches a token when the spec pulls a cloudquery-registry
	// plugin. A source-only (or fully non-cloudquery) spec relies on injection,
	// so resolve credentials directly here; any failure just skips injection.
	if token == "" {
		var err error
		if token, teamName, err = resolveCredentials(ctx); err != nil {
			logger.Debug().Err(err).Msg("platform destination: credentials unavailable, skipping auto-injection")
			return destinations
		}
	}
	if token == "" || teamName == "" {
		return destinations
	}

	cl, err := api.NewClient(token)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: api client init failed, skipping auto-injection")
		return destinations
	}

	tenants, err := activeTenants(ctx, cl, teamName)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: tenant discovery failed, skipping auto-injection")
		return destinations
	}
	tenant, ok := selectTenant(logger, tenants)
	if !ok {
		return destinations
	}

	session, err := mintSession(ctx, cl, tenant)
	if err != nil {
		logger.Warn().Err(err).Str("tenant_id", tenant.TenantId.String()).Msg("platform destination: session mint failed, skipping auto-injection")
		return destinations
	}

	plugin := pluginCoords()
	parsedRegistry, err := specs.RegistryFromString(plugin.Registry)
	if err != nil {
		logger.Warn().Err(err).Str("registry", plugin.Registry).Msg("platform destination: unknown plugin registry; skipping auto-injection")
		return destinations
	}

	var existing *specs.Destination
	for _, d := range destinations {
		if d.Name == destinationName {
			existing = d
			break
		}
	}
	if existing == nil {
		existing = &specs.Destination{Metadata: specs.Metadata{Name: destinationName}}
		destinations = append(destinations, existing)
	}
	existing.Path = plugin.Path
	existing.Registry = parsedRegistry
	existing.Version = plugin.Version
	existing.SyncSummary = true
	// sync_group_id is incompatible with the default overwrite-delete-stale
	// write mode; append accumulates each sync group's external-sync rows.
	existing.WriteMode = specs.WriteModeAppend
	// Unique per invocation: assetview finalize keys on (tenant, source,
	// sync_group_id); concurrent runs would otherwise wipe each other's rows.
	existing.SyncGroupId = strconv.FormatUint(allocateSyncGroupID(time.Now()), 10)
	if existing.Spec == nil {
		existing.Spec = map[string]any{}
	}
	apiURL := platformAPIURL(session.ApiUrl)
	existing.Spec["api_url"] = apiURL
	existing.Spec["token"] = session.Token
	existing.SetDefaults()

	for _, s := range sources {
		if !slices.Contains(s.Destinations, destinationName) {
			s.Destinations = append(s.Destinations, destinationName)
		}
	}
	logger.Info().
		Str("platform_url", apiURL).
		Str("tenant_id", tenant.TenantId.String()).
		Str("registry", plugin.Registry).
		Str("path", plugin.Path).
		Str("version", plugin.Version).
		Msg("auto-injected platform destination")
	return destinations
}

// resolveCredentials fetches a cloud token and its team for best-effort
// injection when the sync command did not need to authenticate. Overridable in
// tests; a returned error is non-fatal and simply skips injection.
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
		if string(t.TeamName) == teamName && slices.Contains(injectableStatuses, t.Status) {
			active = append(active, t)
		}
	}
	return active, nil
}

func mintSession(ctx context.Context, cl *cloudquery_api.ClientWithResponses, tenant cloudquery_api.PlatformTenantSummary) (*cloudquery_api.CreatePlatformDestinationSession201Response, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	resp, err := cl.CreatePlatformDestinationSessionWithResponse(ctx, cloudquery_api.CreatePlatformDestinationSessionRequest{TenantId: tenant.TenantId})
	if err != nil {
		return nil, err
	}
	if resp.JSON201 == nil {
		return nil, fmt.Errorf("unexpected status %d minting platform destination session: %s", resp.StatusCode(), strings.TrimSpace(string(resp.Body)))
	}
	if resp.JSON201.Token == "" || resp.JSON201.ApiUrl == "" {
		return nil, errors.New("platform destination session response missing token or api_url")
	}
	return resp.JSON201, nil
}
