package core

import (
	"context"
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/cloudquery/cloudquery/internal/analytics"
	cqsort "github.com/cloudquery/cloudquery/internal/sort"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/core/state"
	cqerrors "github.com/cloudquery/cloudquery/pkg/errors"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/database/dsn"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thoas/go-funk"
)

type FetchStatus int

// ProviderFetchSummary represents a request for the FetchFinishCallback
type ProviderFetchSummary struct {
	Name                  string                          `json:"name,omitempty"`
	Alias                 string                          `json:"alias,omitempty"`
	Version               string                          `json:"version,omitempty"`
	TotalResourcesFetched uint64                          `json:"total_resources_fetched,omitempty"`
	FetchedResources      map[string]ResourceFetchSummary `json:"fetch_resources,omitempty"`
	Status                FetchStatus                     `json:"status,omitempty"`
	Duration              time.Duration                   `json:"duration,omitempty"`
}

type ResourceFetchSummary struct {
	// Execution status of resource
	Status string `json:"status,omitempty"`
	// Total Amount of resources collected by this resource
	ResourceCount uint64 `json:"resource_count,omitempty"`
	// TelemetryEvents is a list of telemetry events that occurred during the fetch
	TelemetryEvents []analytics.TelemetryEvent `json:"-"`
	// Duration in seconds
	Duration time.Duration `json:"duration,omitempty"`
}

type FetchUpdateCallback func(update FetchUpdate)

type FetchUpdate struct {
	Name    string
	Alias   string
	Version string
	// Map of resources that have finished fetching
	FinishedResources map[string]bool
	// Amount of resources collected so far
	ResourceCount uint64
	// Error if any returned by the provider
	Error string
	// Diagnostic count
	DiagnosticCount int
}

// FetchResponse is returned after a successful fetch execution, it holds a fetch summary for each provider that was executed.
type FetchResponse struct {
	FetchId              uuid.UUID                        `json:"fetch_id,omitempty"`
	ProviderFetchSummary map[string]*ProviderFetchSummary `json:"provider_fetch_summary,omitempty"`
	TotalFetched         uint64                           `json:"total_fetched,omitempty"`
	Duration             time.Duration                    `json:"total_fetch_time,omitempty"`
	TelemetryEvents      []analytics.TelemetryEvent       `json:"-"`
}

type ProviderInfo struct {
	Provider registry.Provider
	Config   *config.Provider
}

// FetchOptions is provided to the Client to execute a fetch on one or more providers
type FetchOptions struct {
	// UpdateCallback allows gets called when the client receives updates on fetch.
	UpdateCallback FetchUpdateCallback
	// Providers list of providers to call for fetching
	ProvidersInfo []ProviderInfo
	// Optional: Adds extra fields to the provider
	ExtraFields map[string]interface{}
	// Optional: unique identifier for the fetch, if this isn't given, a random one is generated.
	FetchId uuid.UUID
}

type fetchResult struct {
	summary *ProviderFetchSummary
}

const (
	FetchFailed FetchStatus = iota + 1
	FetchConfigureFailed
	FetchCanceled
	FetchFinished
	FetchPartial
)

func (fs FetchStatus) String() string {
	switch fs {
	case FetchFailed:
		return "failed"
	case FetchCanceled:
		return "canceled"
	case FetchFinished:
		return "successful"
	case FetchPartial:
		return "partial"
	case FetchConfigureFailed:
		return "configure_failed"
	default:
		return "unknown"
	}
}

func (p ProviderFetchSummary) Resources() []string {
	rr := make([]string, 0, len(p.FetchedResources))
	for r := range p.FetchedResources {
		rr = append(rr, r)
	}
	return rr
}

func (p ProviderFetchSummary) String() string {
	if p.Alias != "" {
		return fmt.Sprintf("%s(%s)", p.Name, p.Alias)
	}
	return p.Name
}

func (p ProviderFetchSummary) Properties() map[string]interface{} {
	rd := make(map[string]float64, len(p.FetchedResources))
	for rn, r := range p.FetchedResources {
		rd[rn] = math.Round(r.Duration.Seconds()*100) / 100
	}
	return map[string]interface{}{
		"fetch_provider":              p.Name,
		"fetch_provider_version":      p.Version,
		"fetch_resources":             p.Resources(),
		"fetch_total_resources_count": p.TotalResourcesFetched,
		"fetch_resources_durations":   rd,
		"fetch_duration":              math.Round(p.Duration.Seconds()*100) / 100,
		"fetch_status":                p.Status.String(),
	}
}

func (f FetchUpdate) AllDone() bool {
	for _, v := range f.FinishedResources {
		if !v {
			return false
		}
	}
	return true
}

func (f FetchUpdate) DoneCount() int {
	count := 0
	for _, v := range f.FinishedResources {
		if v {
			count++
		}
	}
	return count
}

func Fetch(ctx context.Context, sta *state.Client, storage database.Storage, pm *plugin.Manager, opts *FetchOptions) (*FetchResponse, error) {
	var err error
	fetchId := opts.FetchId
	if fetchId == uuid.Nil {
		fetchId, err = uuid.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate fetch id: %w", err)
		}
	}
	// set metadata we want to pass to
	metadata := map[string]interface{}{schema.FetchIdMetaKey: fetchId}
	log.Info().Interface("extra_fields", opts.ExtraFields).Msg("received fetch request")

	var (
		fetchSummaries = make(chan fetchResult, len(opts.ProvidersInfo))
		wg             sync.WaitGroup
		start          = time.Now()
	)

	dsnURI, err := parseDSN(storage)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %w", err)
	}
	for _, providerInfo := range opts.ProvidersInfo {
		if len(providerInfo.Config.Resources) == 0 {
			log.Warn().Str("provider", providerInfo.Config.Name).Str("alias", providerInfo.Config.Alias).Msg("skipping provider which configured with 0 resources to fetch")
			continue
		}
		wg.Add(1)
		go func(info ProviderInfo) {
			defer wg.Done()
			start := time.Now()
			s, _ := runProviderFetch(ctx, pm, info, dsnURI, metadata, opts)
			if _, ok := ctx.Deadline(); ok {
				fetchSummaries <- fetchResult{s}
				return
			}
			// TODO: if context deadline exceeds in fetch, do we still want to run the save?
			if err := sta.SaveFetchSummary(ctx, createFetchSummary(fetchId, start, s)); err != nil {
				log.Error().Err(err).Msg("failed to save fetch summary")
			}
			fetchSummaries <- fetchResult{s}
		}(providerInfo)
	}
	wg.Wait()
	response := &FetchResponse{
		FetchId:              fetchId,
		ProviderFetchSummary: make(map[string]*ProviderFetchSummary, len(opts.ProvidersInfo)),
		Duration:             time.Since(start),
	}
	close(fetchSummaries)
	for ps := range fetchSummaries {
		response.ProviderFetchSummary[ps.summary.String()] = ps.summary
		response.TotalFetched += ps.summary.TotalResourcesFetched
	}

	return response, nil
}

func runProviderFetch(ctx context.Context, pm *plugin.Manager, info ProviderInfo, dsnURI string, metadata map[string]interface{}, opts *FetchOptions) (*ProviderFetchSummary, error) {
	cfg := info.Config
	pLog := log.With().Str("provider", cfg.Name).Str("alias", cfg.Alias).Logger()

	pLog.Debug().Str("name", info.Provider.String()).Str("alias", cfg.Alias).Msg("creating provider plugin")
	providerPlugin, err := pm.CreatePlugin(&plugin.CreationOptions{
		Provider: info.Provider,
		Alias:    cfg.Alias,
		Env:      cfg.Env,
	})
	if err != nil {
		pLog.Error().Err(err).Msg("failed to create provider plugin")
		return nil, fmt.Errorf("failed to create provider plugin: %w", err)
	}
	defer pm.ClosePlugin(providerPlugin)

	pLog.Info().Msg("requesting provider to configure")
	resp, err := providerPlugin.Provider().ConfigureProvider(ctx, &cqproto.ConfigureProviderRequest{
		CloudQueryVersion: Version,
		Connection: cqproto.ConnectionDetails{
			DSN: dsnURI,
		},
		Config: cfg.ConfigBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to configure provider: %w", err)
	}
	if resp.Error != "" {
		return nil, fmt.Errorf("failed to configure provider with message: %w", resp.Error)
	}

	pLog.Info().Msg("provider configured successfully")
	summary, err := executeFetch(ctx, pLog, providerPlugin, info, metadata, opts.UpdateCallback)

	return summary, err
}

func executeFetch(ctx context.Context, pLog zerolog.Logger, providerPlugin plugin.Plugin, info ProviderInfo, metadata map[string]interface{}, callback FetchUpdateCallback) (*ProviderFetchSummary, error) {
	var (
		start   = time.Now()
		summary = &ProviderFetchSummary{
			Name:                  info.Provider.Name,
			Alias:                 info.Config.Alias,
			Version:               providerPlugin.Version(),
			FetchedResources:      make(map[string]ResourceFetchSummary),
			Status:                FetchFinished,
			TotalResourcesFetched: 0,
		}
	)
	// Set fetch duration one function end
	defer func() {
		summary.Duration = time.Since(start)
	}()

	var resources []string
	resources, err := normalizeResources(ctx, providerPlugin, info.Config.Resources, info.Config.SkipResources)
	if err != nil {
		return nil, err
	}

	pLog.Info().Msg("provider started fetching resources")
	stream, err := providerPlugin.Provider().FetchResources(ctx,
		&cqproto.FetchResourcesRequest{
			Resources:             resources,
			ParallelFetchingLimit: info.Config.MaxParallelResourceFetchLimit,
			MaxGoroutines:         info.Config.MaxGoroutines,
			Timeout:               time.Duration(info.Config.ResourceTimeout) * time.Second,
			Metadata:              metadata,
		})
	if err != nil {
		summary.Status = FetchFailed
		return summary, fmt.Errorf("failed to fetch resources: %w", err)
	}

	for {
		resp, err := stream.Recv()
		switch err {
		case nil:
			// We didn't receive an error we received a response
			pLog.Debug().Str("resource", resp.ResourceName).Uint64("fetched", resp.ResourceCount).Msg("resource fetched successfully")
			if callback != nil {
				callback(FetchUpdate{
					Name:              info.Provider.Name,
					Alias:             info.Config.Alias,
					Version:           providerPlugin.Version(),
					FinishedResources: resp.FinishedResources,
					ResourceCount:     resp.ResourceCount,
				})
				// pLog.Debug().Str("resource", resp.ResourceName).Uint64("finishedCount", resp.ResourceCount).
				//	Bool("finished", update.AllDone()).Int("finishCount", update.DoneCount()).Msg("received fetch update")
			}
			summary.TotalResourcesFetched += resp.ResourceCount
			summary.FetchedResources[resp.ResourceName] = ResourceFetchSummary{}
			if resp.Error != "" {
				pLog.Warn().Err(err).Str("resource", resp.ResourceName).Msg("received resource fetch error")
			}
			// TODO: print diags, specific to resource into log?
		case io.EOF:
			// This case means the stream closed peacefully, i.e the provider finished without any error
			pLog.Info().TimeDiff("execution", time.Now(), start).Msg("provider finished fetch")
			return summary, nil
		default:
			if callback != nil {
				callback(FetchUpdate{
					Name:    info.Provider.Name,
					Alias:   info.Config.Alias,
					Version: providerPlugin.Version(),
					Error:   err.Error(),
				})
			}
			// We received an error, first lets check if we got canceled, if not we log the error and add to diags
			if cqerrors.IsCancelation(err) {
				pLog.Warn().TimeDiff("execution", time.Now(), start).Msg("provider fetch was canceled")
				summary.Status = FetchCanceled
				return summary, err
			}
			pLog.Error().Err(err).Msg("received unexpected provider fetch error")
			summary.Status = FetchFailed
			return summary, err
		}
	}
}

// NormalizeResources walks over all given providers and in place normalizes their resources list:
//
// * wildcard expansion
// * verify no unknown resources
// * verify no duplicate resources
func normalizeResources(ctx context.Context, provider plugin.Plugin, resources, skip []string) ([]string, error) {
	s, err := provider.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get provider schema: %w", err)
	}

	return doNormalizeResources(resources, skip, s.ResourceTables)
}

// doNormalizeResources matches the given two resource lists to all provider resources and returns the requested resources (excluding skip resources) as another list.
func doNormalizeResources(resources, skip []string, all map[string]*schema.Table) ([]string, error) {
	useRes, err := doGlobResources(resources, false, all)
	if err != nil {
		return nil, err
	}
	skipRes, err := doGlobResources(skip, true, all)
	if err != nil {
		return nil, err
	}
	return funk.Subtract(useRes, skipRes).([]string), nil
}

// doGlobResources returns a canonical list of resources given a list of requested and all known resources.
// It replaces wildcard resource with all resources in non-wild mode. Error is returned if:
//
// * wildcard is present and other explicit resource is requested;
// * one of explicitly requested resources is not present in all known;
// * some resource is specified more than once (duplicate).
func doGlobResources(requested []string, allowWild bool, all map[string]*schema.Table) ([]string, error) {
	if allowWild {
		for _, s := range requested {
			if s == "*" {
				return nil, fmt.Errorf("wildcard resource is not allowed")
			}
		}
	} else if len(requested) == 1 && requested[0] == "*" {
		requested = make([]string, 0, len(all))
		for k := range all {
			requested = append(requested, k)
		}
	}

	result := make([]string, 0, len(requested))
	seen := make(map[string]struct{})
	for _, r := range requested {
		if r == "" {
			return nil, fmt.Errorf("empty not allowed")
		}

		if _, ok := seen[r]; ok {
			return nil, fmt.Errorf("resource %q is duplicate", r)
		}
		seen[r] = struct{}{}

		if _, ok := all[r]; ok {
			result = append(result, r)
			continue
		}

		if r == "*" {
			return nil, fmt.Errorf("wildcard resource must be the only one in the list")
		}

		switch globMatches, err := matchResourceGlob(r, all); {
		case err != nil:
			return nil, err
		case len(globMatches) == 0:
			return nil, fmt.Errorf("resource %q does not exist", r)
		default:
			result = append(result, globMatches...)
		}
	}

	return cqsort.Unique(result), nil
}

// matchResourceGlob matches pattern to the given resources, returns matched resources or diags
// pattern should end with .*, exact matches are not handled.
func matchResourceGlob(pattern string, all map[string]*schema.Table) ([]string, error) {
	var result []string
	wildPos := strings.Index(pattern, ".*")

	if wildPos > 0 {
		if wildPos != len(pattern)-2 { // make sure it ends with .*
			return nil, fmt.Errorf("resource match should end with `.*`")
		}
		for k := range all {
			if strings.HasPrefix(k, pattern[:wildPos+1]) { // include the "." in the match
				result = append(result, k)
			}
		}
	} else if wildPos == 0 || strings.Contains(pattern, "*") {
		return nil, fmt.Errorf("invalid wildcard syntax")
	}

	return result, nil
}

func parseDSN(storage database.Storage) (string, error) {
	parsed, err := dsn.ParseConnectionString(storage.DSN())
	if err != nil {
		return "", err
	}
	return parsed.String(), nil
}

func createFetchSummary(fetchId uuid.UUID, start time.Time, ps *ProviderFetchSummary) *state.FetchSummary {
	return &state.FetchSummary{
		FetchId:            fetchId,
		CreatedAt:          time.Now().UTC(),
		Start:              start,
		Finish:             time.Now().UTC(),
		TotalResourceCount: ps.TotalResourcesFetched,
		ProviderName:       ps.Name,
		ProviderAlias:      ps.Alias,
		ProviderVersion:    ps.Version,
		CoreVersion:        Version,
		Resources:          parseFetchedResources(ps.FetchedResources),
	}
}

func parseFetchedResources(resources map[string]ResourceFetchSummary) []state.ResourceFetchSummary {
	rfs := make([]state.ResourceFetchSummary, 0, len(resources))
	for k, v := range resources {
		rfs = append(rfs, state.ResourceFetchSummary{
			ResourceName:  k,
			Status:        v.Status,
			ResourceCount: v.ResourceCount,
		})
	}
	return rfs
}
