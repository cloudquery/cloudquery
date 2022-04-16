package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sort"
	"sync"
	"time"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/client/state"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"

	"github.com/rs/zerolog"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/rs/zerolog/log"

	"github.com/cloudquery/cloudquery/pkg/config"

	"github.com/cloudquery/cloudquery/pkg/plugin"

	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/database/dsn"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

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
	// Optional: Adds extra fields to the provider, this is used for history mode and testing purposes.
	ExtraFields map[string]interface{}

	History *history.Config
}

// FetchResponse is returned after a successful fetch execution, it holds a fetch summary for each provider that was executed.
type FetchResponse struct {
	ProviderFetchSummary map[string]ProviderFetchSummary
}

func (fr FetchResponse) HasErrors() bool {
	for _, p := range fr.ProviderFetchSummary {
		if p.HasErrors() {
			return true
		}
	}
	return false
}

type fetchResult struct {
	summary ProviderFetchSummary
	diags   diag.Diagnostics
}

func parseDSN(storage database.Storage, cfg *history.Config) (string, error) {
	if cfg == nil {
		parsed, err := dsn.ParseConnectionString(storage.DSN())
		if err != nil {
			return "", err
		}
		return parsed.String(), nil
	}
	return history.TransformDSN(storage.DSN())
}

func Fetch(ctx context.Context, storage database.Storage, pm *plugin.Manager, opts FetchOptions) (res *FetchResponse, diagnostics diag.Diagnostics) {
	fetchId, err := uuid.NewUUID()
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	// set metadata we want to pass to
	metadata := map[string]interface{}{"cq_fetch_id": fetchId.String()}
	if opts.History != nil {
		fd := opts.History.FetchDate()
		log.Info().Str("fetch_date", fd.Format(time.RFC3339)).Stringer("fetch_id", fetchId).Msg("history enabled adding fetch date")
		metadata["cq_fetch_date"] = fd
		// TODO Remove(Compatibility): Code below is for providers using the old SDK version, where metadata isn't available in FetchRequest
		// Removing this without updating provider will set cq_fetch_date to the time of execution start, which HistoryCfg.TimeTruncation doesn't apply
		if opts.ExtraFields == nil {
			opts.ExtraFields = make(map[string]interface{})
		}
		opts.ExtraFields["cq_fetch_date"] = fd
	}
	log.Info().Interface("extra_fields", opts.ExtraFields).Bool("history_enabled", opts.History != nil).Msg("received fetch request")

	// TODO: in future more components will want state, so make state more generic and passable via database.Storage
	db, err := sdkdb.New(ctx, logging.NewZHcLog(&log.Logger, "fetch"), storage.DSN())
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	defer db.Close()
	stateClient := state.NewClient(db, logging.NewZHcLog(&log.Logger, "fetch"))
	// migrate cloudquery core tables to latest version
	if err := stateClient.MigrateCore(ctx, storage.DialectExecutor()); err != nil {
		return nil, diag.FromError(err, diag.DATABASE, diag.WithSummary("failed to migrate cloudquery_core tables"))
	}

	var (
		diags          diag.Diagnostics
		fetchSummaries = make(chan fetchResult, len(opts.ProvidersInfo))
		wg             sync.WaitGroup
	)

	dsnURI, err := parseDSN(storage, opts.History)
	if err != nil {
		return nil, diag.FromError(err, diag.INTERNAL)
	}
	for _, providerInfo := range opts.ProvidersInfo {
		if len(providerInfo.Config.Resources) == 0 {
			log.Warn().Str("provider", providerInfo.Config.Name).Str("alias", providerInfo.Config.Alias).Msg("skipping provider which configured with 0 resources to fetch")
			diags = diags.Add(diag.FromError(nil, diag.INTERNAL, diag.WithSeverity(diag.WARNING), diag.WithSummary("skipping provider %s which configured with 0 resources to fetch", providerInfo.Config.Name)))
			continue
		}
		wg.Add(1)
		go func(info ProviderInfo) {
			defer wg.Done()
			start := time.Now()
			s, d := runProviderFetch(ctx, pm, info, dsnURI, metadata, opts)
			if _, ok := ctx.Deadline(); ok {
				fetchSummaries <- fetchResult{s, d}
				return
			}
			// TODO: if context deadline exceeds in fetch, do we still want to run the save?
			if err := stateClient.SaveFetchSummary(ctx, &state.FetchSummary{
				FetchId:            fetchId,
				CreatedAt:          time.Now().UTC(),
				Start:              start,
				Finish:             time.Now().UTC(),
				IsSuccess:          diags.HasErrors(),
				TotalResourceCount: s.TotalResourcesFetched,
				TotalErrorsCount:   diags.Errors(),
				ProviderName:       info.Config.Name,
				ProviderAlias:      info.Config.Alias,
				ProviderVersion:    info.Provider.Version,
				CoreVersion:        Version,
				Resources:          parseFetchedResources(s.FetchResources),
			}); err != nil {
				d = d.Add(diag.FromError(err, diag.INTERNAL))
			}
			fetchSummaries <- fetchResult{s, d}
		}(providerInfo)
	}
	wg.Wait()
	response := &FetchResponse{ProviderFetchSummary: make(map[string]ProviderFetchSummary, len(opts.ProvidersInfo))}
	close(fetchSummaries)
	for ps := range fetchSummaries {
		response.ProviderFetchSummary[ps.summary.String()] = ps.summary
		if ps.diags.HasDiags() {
			diags = diags.Add(ps.diags)
		}
	}
	reportFetchSummaryErrors(trace.SpanFromContext(ctx), response.ProviderFetchSummary)

	return response, diags
}

func runProviderFetch(ctx context.Context, pm *plugin.Manager, info ProviderInfo, dsn string, metadata map[string]interface{}, opts FetchOptions) (ProviderFetchSummary, diag.Diagnostics) {
	cfg := info.Config
	pLog := log.With().Str("provider", cfg.Name).Str("alias", cfg.Alias).Logger()

	pLog.Debug().Msg("creating provider plugin")
	providerPlugin, err := pm.CreatePlugin(&plugin.CreationOptions{
		Provider: info.Provider,
		Alias:    cfg.Alias,
		Env:      cfg.Env,
	})
	if err != nil {
		pLog.Error().Err(err).Msg("failed to create provider plugin")
		return ProviderFetchSummary{}, diag.FromError(err, diag.INTERNAL)
	}
	defer pm.ClosePlugin(providerPlugin)

	pLog.Info().Msg("requesting provider to configure")
	_, err = providerPlugin.Provider().ConfigureProvider(ctx, &cqproto.ConfigureProviderRequest{
		CloudQueryVersion: Version,
		Connection: cqproto.ConnectionDetails{
			DSN: dsn,
		},
		Config:      cfg.Configuration,
		ExtraFields: opts.ExtraFields,
	})
	if err != nil {
		pLog.Error().Err(err).Msg("failed to configure provider")
		return ProviderFetchSummary{}, diag.FromError(err, diag.INTERNAL)
	}
	pLog.Info().Msg("provider configured successfully")
	summary, diags := executeFetch(ctx, pLog, providerPlugin, info, metadata, opts.UpdateCallback)
	return summary, diags
}

func executeFetch(ctx context.Context, pLog zerolog.Logger, plugin plugin.Plugin, info ProviderInfo, metadata map[string]interface{}, callback FetchUpdateCallback) (ProviderFetchSummary, diag.Diagnostics) {
	var (
		start   = time.Now()
		summary = ProviderFetchSummary{
			ProviderName:          info.Provider.Name,
			ProviderAlias:         info.Config.Alias,
			Version:               info.Provider.Version,
			FetchResources:        make(map[string]cqproto.ResourceFetchSummary),
			Status:                "Finished",
			TotalResourcesFetched: 0,
		}
		diags diag.Diagnostics
	)
	resources, err := normalizeResources(ctx, plugin, info.Config.Resources)
	if err != nil {
		summary.Status = "Failed"
		return summary, diag.FromError(err, diag.INTERNAL)
	}

	pLog.Info().Msg("provider started fetching resources")
	stream, err := plugin.Provider().FetchResources(ctx,
		&cqproto.FetchResourcesRequest{
			Resources:             resources,
			ParallelFetchingLimit: info.Config.MaxParallelResourceFetchLimit,
			MaxGoroutines:         info.Config.MaxGoroutines,
			Timeout:               time.Duration(info.Config.ResourceTimeout) * time.Second,
			Metadata:              metadata,
		})
	if err != nil {
		summary.Status = "Failed"
		return summary, diag.FromError(err, diag.INTERNAL)
	}

	for {
		resp, err := stream.Recv()
		switch err {
		case nil:
			// We didn't receive an error we received a response
			pLog.Debug().Str("resource", resp.ResourceName).Uint64("fetched", resp.ResourceCount).Msg("resource fetched successfully")
			if callback != nil {
				callback(FetchUpdate{
					Provider:          plugin.Name(),
					Version:           plugin.Version(),
					FinishedResources: resp.FinishedResources,
					ResourceCount:     resp.ResourceCount,
				})
				// pLog.Debug().Str("resource", resp.ResourceName).Uint64("finishedCount", resp.ResourceCount).
				//	Bool("finished", update.AllDone()).Int("finishCount", update.DoneCount()).Msg("received fetch update")
			}
			summary.TotalResourcesFetched += resp.ResourceCount
			summary.FetchResources[resp.ResourceName] = resp.Summary
			if resp.Error != "" {
				pLog.Warn().Err(err).Str("resource", resp.ResourceName).Msg("received resource fetch error")
				diags = diags.Add(diag.FromError(errors.New(resp.Error), diag.RESOLVING, diag.WithResourceName(resp.ResourceName)))
			}
			// TODO: print diags, specific to resource into log?
			if resp.Summary.Diagnostics.HasDiags() {
				pLog.Warn().Str("resource", resp.ResourceName).Msg("received resource fetch diagnostics")
				diags = diags.Add(resp.Summary.Diagnostics)
			}
		case io.EOF:
			// This case means the stream closed peacefully, i.e the provider finished without any error
			pLog.Info().TimeDiff("execution", time.Now(), start).Msg("provider finished fetch")
			return summary, diags
		default:
			// We received an error, first lets check if we got canceled, if not we log the error and add to diags
			if st, ok := status.FromError(err); ok && st.Code() == gcodes.Canceled {
				pLog.Warn().TimeDiff("execution", time.Now(), start).Msg("provider fetch was canceled")
				summary.Status = "Canceled"
				return summary, diags
			}
			pLog.Error().Err(err).Msg("received unexpected provider fetch error")
			summary.Status = "Failed"
			return summary, diags.Add(diag.FromError(err, diag.INTERNAL))
		}
	}
}

func parseFetchedResources(resources map[string]cqproto.ResourceFetchSummary) []state.ResourceFetchSummary {
	rfs := make([]state.ResourceFetchSummary, 0, len(resources))
	for k, v := range resources {
		rfs = append(rfs, state.ResourceFetchSummary{
			ResourceName:  k,
			Status:        v.Status.String(),
			Error:         v.Diagnostics.Error(),
			ResourceCount: v.ResourceCount,
		})
	}
	return rfs
}

// NormalizeResources walks over all given providers and in place normalizes their resources list:
//
// * wildcard expansion
// * verify no unknown resources
// * verify no duplicate resources
func normalizeResources(ctx context.Context, provider plugin.Plugin, resources []string) ([]string, error) {
	s, err := provider.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
	if err != nil {
		return nil, err
	}
	return doNormalizeResources(resources, s.ResourceTables)
}

// doNormalizeResources returns a canonical list of resources given a list of requested and all known resources.
// It replaces wildcard resource with all resources. Error is returned if:
//
// * wildcard is present and other explicit resource is requested;
// * one of explicitly requested resources is not present in all known;
// * some resource is specified more than once (duplicate).
func doNormalizeResources(requested []string, all map[string]*schema.Table) ([]string, error) {
	if len(requested) == 1 && requested[0] == "*" {
		requested = make([]string, 0, len(all))
		for k := range all {
			requested = append(requested, k)
		}
	}
	result := make([]string, 0, len(requested))
	seen := make(map[string]struct{})
	for _, r := range requested {
		if _, ok := seen[r]; ok {
			return nil, fmt.Errorf("resource %s is duplicate", r)
		}
		seen[r] = struct{}{}
		if _, ok := all[r]; !ok {
			if r == "*" {
				return nil, fmt.Errorf("wildcard resource must be the only one in the list")
			}
			return nil, fmt.Errorf("resource %s does not exist", r)
		}
		result = append(result, r)
	}
	sort.Strings(result)
	return result, nil
}
