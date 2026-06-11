// Package platform wires the CLI to the CloudQuery Platform's external-syncs
// flow. If the team has an active platform tenant, every sync the user runs
// gets a platform destination appended automatically, authenticated by a
// short-lived cqpd_ token minted from the caller's cloud credential.
package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"

	envDisable  = "CQ_DISABLE_PLATFORM_DESTINATION"
	envTenantID = "CQ_PLATFORM_TENANT_ID"

	destinationName = "platform"
	statusActive    = "active"
)

// pluginCoordinates pins the platform destination plugin the CLI injects.
// Currently hardcoded; kept as a struct so server-driven pinning can be wired
// back in without touching the injection logic.
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

type tenantSummary struct {
	TenantID string `json:"tenant_id"`
	Status   string `json:"status"`
	TeamName string `json:"team_name"`
}

type tenantListResponse struct {
	Items []tenantSummary `json:"items"`
}

type sessionResponse struct {
	Token            string `json:"token"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
	APIURL           string `json:"api_url"`
}

// MaybeInjectDestination ensures a `platform` destination exists in the spec
// when the team has an active platform tenant. It discovers the tenant via
// cloud's GET /user/platform/tenants, mints a tenant-scoped cqpd_ token via
// POST /platform-destination/session, and stamps the returned token + api_url
// into the destination spec — the cloud credential itself never reaches the
// plugin. If a destination named `platform` already exists, its metadata and
// connection spec are overwritten so users can keep a placeholder block in
// their config; otherwise a fresh entry is appended and every source is wired
// to it. Failures fall through silently so a cloud API outage does not break
// local syncs.
func MaybeInjectDestination(ctx context.Context, logger zerolog.Logger, token, teamName string, sources []*specs.Source, destinations []*specs.Destination) []*specs.Destination {
	if os.Getenv(envDisable) == "1" {
		return destinations
	}
	if token == "" || teamName == "" {
		return destinations
	}

	tenants, err := activeTenants(ctx, token, teamName)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination: tenant discovery failed, skipping auto-injection")
		return destinations
	}
	tenant, ok := selectTenant(logger, tenants)
	if !ok {
		return destinations
	}

	session, err := mintSession(ctx, token, tenant.TenantID)
	if err != nil {
		logger.Warn().Err(err).Str("tenant_id", tenant.TenantID).Msg("platform destination: session mint failed, skipping auto-injection")
		return destinations
	}

	parsedRegistry, err := specs.RegistryFromString(defaultPlugin.Registry)
	if err != nil {
		logger.Warn().Err(err).Str("registry", defaultPlugin.Registry).Msg("platform destination: unknown plugin registry; skipping auto-injection")
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
	existing.Path = defaultPlugin.Path
	existing.Registry = parsedRegistry
	existing.Version = defaultPlugin.Version
	existing.SyncSummary = true
	// Stamp a time-based numeric sync_group_id (YYYYMMDDhhmmssfff) so the CLI's
	// record transformer adds `_cq_sync_group_id` to every row. Assetview's
	// staging→main finalize keys on `(tenant, source, sync_group_id)`; without
	// a unique sgid per invocation, concurrent CLI runs from the same team
	// collide in finalize and wipe each other's rows.
	existing.SyncGroupId = strconv.FormatUint(allocateSyncGroupID(time.Now()), 10)
	if existing.Spec == nil {
		existing.Spec = map[string]any{}
	}
	existing.Spec["api_url"] = session.APIURL
	existing.Spec["token"] = session.Token
	existing.SetDefaults()

	for _, s := range sources {
		if !slices.Contains(s.Destinations, destinationName) {
			s.Destinations = append(s.Destinations, destinationName)
		}
	}
	logger.Info().
		Str("platform_url", session.APIURL).
		Str("tenant_id", tenant.TenantID).
		Str("registry", defaultPlugin.Registry).
		Str("path", defaultPlugin.Path).
		Str("version", defaultPlugin.Version).
		Msg("auto-injected platform destination")
	return destinations
}

// selectTenant picks the tenant to mint for. Zero tenants → not onboarded,
// skip quietly. Multiple → require CQ_PLATFORM_TENANT_ID to disambiguate;
// never guess, and never fail the sync.
func selectTenant(logger zerolog.Logger, tenants []tenantSummary) (tenantSummary, bool) {
	switch len(tenants) {
	case 0:
		return tenantSummary{}, false
	case 1:
		return tenants[0], true
	}
	want := os.Getenv(envTenantID)
	if want == "" {
		logger.Warn().Int("tenants", len(tenants)).Msgf("platform destination: team has multiple active tenants; set %s to choose one, skipping auto-injection", envTenantID)
		return tenantSummary{}, false
	}
	for _, t := range tenants {
		if t.TenantID == want {
			return t, true
		}
	}
	logger.Warn().Str("tenant_id", want).Msgf("platform destination: %s does not match any active tenant, skipping auto-injection", envTenantID)
	return tenantSummary{}, false
}

// allocateSyncGroupID returns a time-based uint64 (YYYYMMDDhhmmssfff). Same
// shape platform/syncs-transformer uses for managed syncs, so external-sync
// rows share the keyspace.
func allocateSyncGroupID(now time.Time) uint64 {
	t := now.UTC()
	base := t.Format("20060102150405") + fmt.Sprintf("%03d", t.Nanosecond()/1e6)
	u, _ := strconv.ParseUint(base, 10, 64)
	return u
}

func activeTenants(ctx context.Context, token, teamName string) ([]tenantSummary, error) {
	base := env.GetEnvOrDefault(envAPIURL, defaultAPIURL)
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, base+"/user/platform/tenants", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from /user/platform/tenants", resp.StatusCode)
	}
	var out tenantListResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode /user/platform/tenants response: %w", err)
	}
	active := make([]tenantSummary, 0, len(out.Items))
	for _, t := range out.Items {
		if t.TeamName == teamName && t.Status == statusActive {
			active = append(active, t)
		}
	}
	return active, nil
}

func mintSession(ctx context.Context, token, tenantID string) (*sessionResponse, error) {
	base := env.GetEnvOrDefault(envAPIURL, defaultAPIURL)
	body, err := json.Marshal(map[string]string{"tenant_id": tenantID})
	if err != nil {
		return nil, err
	}
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, base+"/platform-destination/session", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := newHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("unexpected status %d from /platform-destination/session: %s", resp.StatusCode, strings.TrimSpace(string(snippet)))
	}
	var out sessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode /platform-destination/session response: %w", err)
	}
	if out.Token == "" || out.APIURL == "" {
		return nil, fmt.Errorf("/platform-destination/session response missing token or api_url")
	}
	return &out, nil
}

func newHTTPClient() *retryablehttp.Client {
	cl := retryablehttp.NewClient()
	cl.Logger = nil
	cl.HTTPClient.Timeout = 5 * time.Second
	cl.RetryMax = 2
	cl.RetryWaitMin = 200 * time.Millisecond
	cl.RetryWaitMax = 2 * time.Second
	return cl
}
