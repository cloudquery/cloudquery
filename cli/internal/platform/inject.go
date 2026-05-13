// Package platform wires the CLI to the CloudQuery Platform's external-syncs
// flow. If the team is onboarded to the platform destination, every sync the
// user runs gets a platform destination appended automatically.
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	specs "github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
)

const (
	defaultAPIURL = "https://api.cloudquery.io"
	envAPIURL     = "CLOUDQUERY_API_URL"

	envDisable = "CQ_DISABLE_PLATFORM_DESTINATION"

	// Fallback defaults if cloud's /platform/status response is missing the
	// plugin coordinates (older cloud-api before the pinning fields shipped).
	fallbackRegistry = "cloudquery"
	fallbackPath     = "cloudquery/platform"
	fallbackVersion  = "v1.0.0"

	destinationName = "platform"
	statusActive    = "active"
)

type statusResponse struct {
	PlatformURL    string `json:"platform_url"`
	Status         string `json:"status"`
	TeamName       string `json:"team_name"`
	PluginRegistry string `json:"plugin_registry"`
	PluginPath     string `json:"plugin_path"`
	PluginVersion  string `json:"plugin_version"`
}

// MaybeInjectDestination ensures a `platform` destination exists in the spec
// when the team is onboarded. If an entry with that name already exists, its
// metadata + connection spec are overwritten with the values cloud's
// /platform/status returns — users get to keep a placeholder destination
// block in their config without needing to fill in real `api_url`/`token`
// values themselves. If no entry exists, a fresh one is appended and every
// source is wired to it. Failures fall through silently so a cloud API
// outage does not break local syncs.
func MaybeInjectDestination(ctx context.Context, logger zerolog.Logger, token, teamName string, sources []*specs.Source, destinations []*specs.Destination) []*specs.Destination {
	if os.Getenv(envDisable) == "1" {
		return destinations
	}
	if token == "" || teamName == "" {
		return destinations
	}

	info, err := fetchStatus(ctx, token, teamName)
	if err != nil {
		logger.Debug().Err(err).Msg("platform destination status check failed, skipping auto-injection")
		return destinations
	}
	if info == nil || info.Status != statusActive {
		return destinations
	}

	registry, path, version := info.PluginRegistry, info.PluginPath, info.PluginVersion
	if registry == "" {
		registry = fallbackRegistry
	}
	if path == "" {
		path = fallbackPath
	}
	if version == "" {
		version = fallbackVersion
	}

	parsedRegistry, err := specs.RegistryFromString(registry)
	if err != nil {
		logger.Warn().Err(err).Str("registry", registry).Msg("platform destination: unknown plugin registry; skipping auto-injection")
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
	existing.Path = path
	existing.Registry = parsedRegistry
	existing.Version = version
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
	existing.Spec["api_url"] = info.PlatformURL
	existing.Spec["token"] = token
	existing.SetDefaults()

	for _, s := range sources {
		if !slices.Contains(s.Destinations, destinationName) {
			s.Destinations = append(s.Destinations, destinationName)
		}
	}
	logger.Info().
		Str("platform_url", info.PlatformURL).
		Str("registry", registry).
		Str("path", path).
		Str("version", version).
		Msg("auto-injected platform destination")
	return destinations
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

func fetchStatus(ctx context.Context, token, teamName string) (*statusResponse, error) {
	base := env.GetEnvOrDefault(envAPIURL, defaultAPIURL)
	url := fmt.Sprintf("%s/teams/%s/platform/status", base, teamName)
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	cl := retryablehttp.NewClient()
	cl.Logger = nil
	cl.HTTPClient.Timeout = 5 * time.Second
	cl.RetryMax = 2
	cl.RetryWaitMin = 200 * time.Millisecond
	cl.RetryWaitMax = 2 * time.Second

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from /platform/status", resp.StatusCode)
	}

	var out statusResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode /platform/status response: %w", err)
	}
	return &out, nil
}
