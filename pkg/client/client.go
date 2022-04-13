package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cloudquery/pkg/client/database/timescale"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	"github.com/cloudquery/cloudquery/pkg/client/meta_storage"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/database/dsn"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMigrationsNotSupported = errors.New("provider doesn't support migrations")
)

// FetchRequest is provided to the Client to execute a fetch on one or more providers
type FetchRequest struct {
	// UpdateCallback allows gets called when the client receives updates on fetch.
	UpdateCallback FetchUpdateCallback
	// Providers list of providers to call for fetching
	Providers []*config.Provider
	// Optional: Adds extra fields to the provider, this is used for history mode and testing purposes.
	ExtraFields map[string]interface{}
}

// FetchResponse is returned after a successful fetch execution, it holds a fetch summary for each provider that was executed.
type FetchResponse struct {
	ProviderFetchSummary map[string]ProviderFetchSummary
}

type FetchUpdate struct {
	Provider string
	Version  string
	// Map of resources that have finished fetching
	FinishedResources map[string]bool
	// Amount of resources collected so far
	ResourceCount uint64
	// Error if any returned by the provider
	Error string
	// PartialFetchResults contains the partial fetch results for this update
	PartialFetchResults []*cqproto.FailedResourceFetch
}

// ProviderFetchSummary represents a request for the FetchFinishCallback
type ProviderFetchSummary struct {
	ProviderName          string
	ProviderAlias         string
	Version               string
	PartialFetchErrors    []*cqproto.FailedResourceFetch
	FetchErrors           []error
	TotalResourcesFetched uint64
	FetchResources        map[string]cqproto.ResourceFetchSummary
	Status                string
}

func (p ProviderFetchSummary) Diagnostics() diag.Diagnostics {
	var allDiags diag.Diagnostics
	for _, s := range p.FetchResources {
		allDiags = append(allDiags, s.Diagnostics...)
	}
	return allDiags
}

func (p ProviderFetchSummary) HasErrors() bool {
	if len(p.FetchErrors) > 0 || len(p.PartialFetchErrors) > 0 {
		return true
	}
	return false
}

func (p ProviderFetchSummary) Metrics() map[string]int64 {
	type diagCount map[diag.DiagnosticType]int64
	sevCounts := make(map[diag.Severity]diagCount)

	for _, d := range p.Diagnostics() {
		if _, ok := sevCounts[d.Severity()]; !ok {
			tc := make(diagCount)
			tc[d.Type()]++
			sevCounts[d.Severity()] = tc
		} else {
			sevCounts[d.Severity()][d.Type()]++
		}
	}

	ret := make(map[string]int64, len(sevCounts)+1)
	for severity, typeCount := range sevCounts {
		var sevName string
		switch severity {
		case diag.IGNORE:
			sevName = "ignore"
		case diag.WARNING:
			sevName = "warning"
		case diag.ERROR:
			sevName = "error"
		default:
			sevName = "unknown"
		}

		prefix := "fetch.diag." + sevName + "."
		var total int64
		for typ, count := range typeCount {
			ret[prefix+strings.ToLower(typ.String())+"."+p.ProviderName] = count
			total += count
		}
		ret[prefix+"total."+p.ProviderName] = total
	}

	return ret
}

// PoliciesRunRequest is the request used to run a policy.
type PoliciesRunRequest struct {
	// Policies to run
	Policies policy.Policies

	// OutputDir is the output dir for policy execution output.
	OutputDir string

	// RunCallBack is the callback method that is called after every policy execution.
	RunCallback policy.UpdateCallback
}

// ModuleRunRequest is the request used to run a module.
type ModuleRunRequest struct {
	// Name of the module
	Name string

	// Params are the invocation parameters specific to the module
	Params interface{}

	// Providers is the list of providers to process
	Providers []*cqproto.GetProviderSchemaResponse

	// Config is the config profile provided by the user
	Config hcl.Body
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
			count += 1
		}
	}
	return count
}

type FetchDoneResult struct {
	// Map of providers and resources that have finished fetching
	DoneResources map[string]map[string]bool
	// Amount of resources collected so far
	ResourceCount string
}

type FetchUpdateCallback func(update FetchUpdate)

type Option func(options *Client)

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Required: List of providers that are required, these providers will be downloaded if DownloadProviders is called.
	Providers config.RequiredProviders
	// Optional: Registry url to verify plugins from, defaults to CloudQuery hub
	RegistryURL string
	// Optional: Where to save downloaded providers, by default current working directory, defaults to ./cq/providers
	PluginDirectory string
	// Optional: Where to save downloaded policies, by default current working directory, defaults to ./cq/policy
	PolicyDirectory string
	// Optional: If true cloudquery just runs policy files without using git tag to select a version
	SkipVersioning bool
	// Optional: if this flag is true, plugins downloaded from URL won't be verified when downloaded
	NoVerify bool
	// Optional: DSN connection information for database client will connect to
	DSN string
	// Optional: Skips Building tables on fetch execution
	SkipBuildTables bool
	// Optional: HubProgressUpdater allows the client creator to get called back on download progress and completion.
	HubProgressUpdater ui.Progress
	// Optional: Logger framework can use to log.
	// default: global logger provided.
	Logger hclog.Logger
	// Hub client to use to download plugins, the Hub is used to download and pluginManager providers binaries
	// if not specified, default cloudquery registry is used.
	Hub registry.Hub
	// manager manages all plugins lifecycle
	Manager *plugin.Manager
	// ModuleManager manages all modules lifecycle
	ModuleManager module.Manager
	// ModuleManager manages all modules lifecycle
	PolicyManager policy.Manager
	// HistoryConfig defines configuration for CloudQuery history mode
	HistoryCfg *history.Config

	// metaStorage interacts with cloudquery core resources
	metaStorage     *meta_storage.Client
	db              *sdkdb.DB
	dialectExecutor database.DialectExecutor
}

func New(ctx context.Context, options ...Option) (*Client, error) {
	c := &Client{
		PluginDirectory:    filepath.Join(".", ".cq", "providers"),
		PolicyDirectory:    ".",
		NoVerify:           false,
		SkipBuildTables:    false,
		HubProgressUpdater: nil,
		HistoryCfg:         nil,
		RegistryURL:        registry.CloudQueryRegistryURL,
		Logger:             logging.NewZHcLog(&zerolog.Logger, ""),
		Hub:                *registry.NewRegistryHub(registry.CloudQueryRegistryURL),
	}
	for _, o := range options {
		o(c)
	}

	var err error
	c.Manager, err = plugin.NewManager(registry.NewRegistryHub(c.RegistryURL,
		registry.WithPluginDirectory(c.PluginDirectory), registry.WithProgress(c.HubProgressUpdater)), plugin.WithAllowReattach())
	if err != nil {
		return nil, err
	}

	if c.DSN == "" {
		c.Logger.Warn("missing DSN, some commands won't work")
	} else if err := c.initDatabase(ctx); err != nil {
		return nil, err
	}

	c.initModules()

	c.PolicyManager = policy.NewManager(c.PolicyDirectory, c.db, c.Logger)

	return c, nil
}

// DownloadProviders downloads all provider binaries
func (c *Client) DownloadProviders(ctx context.Context) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "DownloadProviders")
	defer spanEnder(retErr)

	c.Logger.Info("Downloading required providers")
	pp := make([]registry.Provider, len(c.Providers))
	for i, rp := range c.Providers {
		src, name, err := ParseProviderSource(rp)
		if err != nil {
			return err
		}
		pp[i] = registry.Provider{
			Name:    name,
			Version: rp.Version,
			Source:  src,
		}
	}
	_, diags := Download(ctx, c.Manager, &DownloadOptions{
		Providers: pp,
		NoVerify:  c.NoVerify,
	})
	if diags.HasErrors() {
		return diags
	}
	return nil
}

type ProviderUpdateSummary struct {
	Name          string
	Version       string
	LatestVersion string
}

// NormalizeResources walks over all given providers and in place normalizes their resources list:
//
// * wildcard expansion
// * no unknown resources
// * no duplicate resources
func (c *Client) NormalizeResources(ctx context.Context, providers []*config.Provider) error {
	for _, p := range providers {
		if err := c.normalizeProvider(ctx, p); err != nil {
			return fmt.Errorf("provider %s: %w", p.Name, err)
		}
	}
	return nil
}

func (c *Client) normalizeProvider(ctx context.Context, p *config.Provider) error {
	s, err := c.GetProviderSchema(ctx, p.Name)
	if err != nil {
		return err
	}
	p.Resources, err = normalizeResources(p.Resources, s.ResourceTables)
	return err
}

func (c *Client) Fetch(ctx context.Context, request FetchRequest) (res *FetchResponse, retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "Fetch")
	defer spanEnder(retErr)

	reportNumProviders(ctx, request.Providers)

	c.Logger.Info("received fetch request", "extra_fields", request.ExtraFields, "history_enabled", c.HistoryCfg != nil)

	var dsnURI string
	if c.HistoryCfg != nil {
		var err error
		dsnURI, err = history.TransformDSN(c.DSN)
		if err != nil {
			return nil, err
		}
	} else {
		parsed, err := dsn.ParseConnectionString(c.DSN)
		if err != nil {
			return nil, err
		}
		dsnURI = parsed.String()
	}

	fetchSummaries := make(chan ProviderFetchSummary, len(request.Providers))
	// Ignoring gctx since we don't want to stop other running providers if one provider fails with an error
	// future refactor should probably use a something else rather than error group.
	errGroup, _ := errgroup.WithContext(ctx)
	fetchId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	for _, providerConfig := range request.Providers {
		if len(providerConfig.Resources) == 0 {
			c.Logger.Warn("skipping provider which configured with 0 resources to fetch", "provider", providerConfig.Name, "alias", providerConfig.Alias)
			continue
		}

		c.Logger.Debug("creating provider plugin", "provider", providerConfig.Name)
		providerPlugin, err := c.CreatePlugin(providerConfig.Name, providerConfig.Alias, providerConfig.Env)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", providerConfig.Name, "error", err)
			return nil, err
		}

		providerConfig := providerConfig
		createdAt := time.Now().UTC()
		fetchSummary := meta_storage.FetchSummary{
			FetchId:         fetchId,
			ProviderName:    providerConfig.Name,
			ProviderAlias:   providerConfig.Alias,
			ProviderVersion: providerPlugin.Version(),
			CreatedAt:       &createdAt,
			CoreVersion:     Version,
		}

		saveFetchSummary := func() {
			if err := c.metaStorage.SaveFetchSummary(ctx, &fetchSummary); err != nil {
				c.Logger.Error("failed to save fetch summary", "err", err)
			}
		}

		// TODO: move this into an outer function
		errGroup.Go(func() error {
			defer saveFetchSummary()
			pLog := c.Logger.With("provider", providerConfig.Name, "alias", providerConfig.Alias, "version", providerPlugin.Version())
			pLog.Info("requesting provider to configure")

			metadata := map[string]interface{}{
				"cq_fetch_id": fetchId.String(),
			}

			if c.HistoryCfg != nil {
				fd := c.HistoryCfg.FetchDate()
				pLog.Info("history enabled adding fetch date", "fetch_date", fd.Format(time.RFC3339))
				metadata["cq_fetch_date"] = fd

				// TODO Remove(Compatibility): Code below is for providers using the old SDK version, where metadata isn't available in FetchRequest
				// Removing this without updating provider will set cq_fetch_date to the time of execution start, which HistoryCfg.TimeTruncation doesn't apply
				if request.ExtraFields == nil {
					request.ExtraFields = make(map[string]interface{})
				}
				request.ExtraFields["cq_fetch_date"] = fd
			}
			_, err = providerPlugin.Provider().ConfigureProvider(ctx, &cqproto.ConfigureProviderRequest{
				CloudQueryVersion: Version,
				Connection: cqproto.ConnectionDetails{
					DSN: dsnURI,
				},
				Config:      providerConfig.Configuration,
				ExtraFields: request.ExtraFields,
			})
			if err != nil {
				pLog.Error("failed to configure provider", "error", err)
				return err
			}
			pLog.Info("provider configured successfully")

			pLog.Info("requesting provider fetch")
			fetchStart := time.Now()
			fetchSummary.Start = &fetchStart
			stream, err := providerPlugin.Provider().FetchResources(ctx,
				&cqproto.FetchResourcesRequest{
					Resources:              providerConfig.Resources,
					PartialFetchingEnabled: true,
					ParallelFetchingLimit:  providerConfig.MaxParallelResourceFetchLimit,
					MaxGoroutines:          providerConfig.MaxGoroutines,
					Timeout:                time.Duration(providerConfig.ResourceTimeout) * time.Second,
					Metadata:               metadata,
				})
			if err != nil {
				return err
			}
			pLog.Info("provider started fetching resources")
			var (
				fetchErrors         = make([]error, 0)
				partialFetchResults []*cqproto.FailedResourceFetch
				fetchedResources           = make(map[string]cqproto.ResourceFetchSummary, len(providerConfig.Resources))
				totalResources      uint64 = 0
				totalErrors         uint64 = 0
			)
			for {
				resp, err := stream.Recv()

				if err != nil {
					st, ok := status.FromError(err)

					if (ok && st.Code() == gcodes.Canceled) || err == io.EOF {
						message := "provider finished fetch"
						fetchStatus := "Finished"
						if ok && st.Code() == gcodes.Canceled {
							message = "provider fetch canceled"
							fetchStatus = "Canceled"
						}

						pLog.Info(message, "execution", time.Since(fetchStart).String())
						for _, fetchError := range partialFetchResults {
							pLog.Warn("received partial fetch error", parsePartialFetchKV(fetchError)...)
						}
						fetchSummaries <- ProviderFetchSummary{
							ProviderName:          providerConfig.Name,
							ProviderAlias:         providerConfig.Alias,
							Version:               providerPlugin.Version(),
							TotalResourcesFetched: totalResources,
							PartialFetchErrors:    partialFetchResults,
							FetchErrors:           fetchErrors,
							FetchResources:        fetchedResources,
							Status:                fetchStatus,
						}
						t := time.Now().UTC()
						fetchSummary.Finish = &t
						fetchSummary.IsSuccess = true
						fetchSummary.TotalErrorsCount = totalErrors
						fetchSummary.TotalResourceCount = totalResources
						return nil
					}
					pLog.Error("received provider fetch error", "error", err)
					return err
				}
				update := FetchUpdate{
					Provider:            providerPlugin.Name(),
					Version:             providerPlugin.Version(),
					FinishedResources:   resp.FinishedResources,
					ResourceCount:       resp.ResourceCount,
					Error:               resp.Error,
					PartialFetchResults: partialFetchResults,
				}

				if len(resp.PartialFetchFailedResources) != 0 {
					partialFetchResults = append(partialFetchResults, resp.PartialFetchFailedResources...)
				}

				totalResources += resp.ResourceCount
				totalErrors += uint64(len(resp.PartialFetchFailedResources))
				fetchedResources[resp.ResourceName] = resp.Summary

				if resp.Error != "" {
					fetchErrors = append(fetchErrors, fmt.Errorf("fetch error: %s", resp.Error))
					pLog.Error("received provider fetch update error", "error", resp.Error)
				}
				pLog.Debug("received fetch update",
					"resource", resp.ResourceName, "finishedCount", resp.ResourceCount, "finished", update.AllDone(), "finishCount", update.DoneCount())
				if request.UpdateCallback != nil {
					request.UpdateCallback(update)
				}

				fetchSummary.Resources = append(fetchSummary.Resources, meta_storage.ResourceFetchSummary{
					ResourceName:      resp.ResourceName,
					FinishedResources: resp.FinishedResources,
					Status:            resp.Summary.Status.String(),
					Error:             resp.Error,
					ResourceCount:     resp.ResourceCount,
				})
			}
		})
	}

	response := &FetchResponse{ProviderFetchSummary: make(map[string]ProviderFetchSummary, len(request.Providers))}
	// TODO: kill all providers on end, add defer on top loop
	if err := errGroup.Wait(); err != nil {
		close(fetchSummaries)
		return nil, err
	}
	close(fetchSummaries)

	for ps := range fetchSummaries {
		response.ProviderFetchSummary[fmt.Sprintf("%s(%s)", ps.ProviderName, ps.ProviderAlias)] = ps
	}

	reportFetchSummaryErrors(trace.SpanFromContext(ctx), response.ProviderFetchSummary)

	return response, nil
}

func (c *Client) GetProviderSchema(ctx context.Context, providerName string) (*ProviderSchema, error) {
	providerPlugin, err := c.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer c.Manager.ClosePlugin(providerPlugin)

	providerSchema, err := providerPlugin.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
	if err != nil {
		return nil, err
	}
	return &ProviderSchema{
		GetProviderSchemaResponse: providerSchema,
		ProtocolVersion:           providerPlugin.ProtocolVersion(),
	}, nil
}

func (c *Client) GetProviderConfiguration(ctx context.Context, providerName string) (*cqproto.GetProviderConfigResponse, error) {
	providerPlugin, err := c.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer c.Manager.ClosePlugin(providerPlugin)
	return providerPlugin.Provider().GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
}

func (c *Client) GetProviderModule(ctx context.Context, providerName string, req cqproto.GetModuleRequest) (*cqproto.GetModuleResponse, error) {
	providerPlugin, err := c.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer c.Manager.ClosePlugin(providerPlugin)
	inf, err := providerPlugin.Provider().GetModuleInfo(ctx, &req)
	if err != nil && strings.Contains(err.Error(), `unknown method GetModuleInfo`) {
		return &cqproto.GetModuleResponse{}, nil
	}

	return inf, err
}

func (c *Client) LoadPolicy(ctx context.Context, name, source string) (pol *policy.Policy, retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "LoadPolicy")
	defer spanEnder(retErr)
	c.Logger.Info("Downloading policy from remote source", "name", name, "source", source)
	return c.PolicyManager.Load(ctx, &policy.Policy{Name: name, Source: source})
}

func (c *Client) RunPolicies(ctx context.Context, req *PoliciesRunRequest) ([]*policy.ExecutionResult, error) {
	results := make([]*policy.ExecutionResult, 0)

	for _, p := range req.Policies {
		c.Logger = c.Logger.With("policy", p.Name, "version", p.Version(), "subPath", p.SubPolicy())
		result, err := c.runPolicy(ctx, p, req)

		c.Logger.Info("policy execution finished", "err", err)
		if err != nil {
			// this error means error in execution and not policy violation
			// we should exit immediately as this is a non-recoverable error
			// might mean schema is incorrect, provider version
			c.Logger.Error("policy execution finished with error", "err", err)
			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *Client) runPolicy(ctx context.Context, p *policy.Policy, req *PoliciesRunRequest) (res *policy.ExecutionResult, retErr error) {
	executionTime := time.Now()

	spanAttrs := []attribute.KeyValue{
		attribute.String("policy_name", p.Name),
	}

	if strings.HasPrefix(p.Name, policy.CloudQueryOrg) {
		spanAttrs = append(spanAttrs,
			attribute.String("policy_version", p.Version()),
			attribute.String("policy_subpath", p.SubPolicy()),
		)
	}

	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "runPolicy", trace.WithAttributes(spanAttrs...))
	defer spanEnder(retErr)

	c.Logger.Info("preparing to run policy")

	if err := c.ensureConnection(); err != nil {
		return nil, err
	}

	versions, err := collectProviderVersions(c.Providers, func(name string) (string, error) {
		d, err := c.Manager.GetPluginDetails(name)
		return d.Version, err
	})
	c.Logger.Debug("collected policy versions", "versions", versions)

	if err != nil {
		return nil, err
	}

	execReq := &policy.ExecuteRequest{
		Policy:           p,
		ProviderVersions: versions,
		UpdateCallback:   req.RunCallback,
	}

	execReq.Policy, err = c.PolicyManager.Load(ctx, p)
	if err != nil {
		return nil, err
	}

	c.Logger.Info("running the policy")
	result, err := c.PolicyManager.Run(ctx, execReq)
	if err != nil {
		return nil, err
	}
	result.ExecutionTime = executionTime

	// Store output in file if requested
	if req.OutputDir != "" {
		c.Logger.Info("writing policy to output directory")
		err = policy.GenerateExecutionResultFile(result, req.OutputDir)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Client) ExecuteModule(ctx context.Context, req ModuleRunRequest) (res *module.ExecutionResult, retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "ExecuteModule", trace.WithAttributes(attribute.String("module", req.Name)))
	defer spanEnder(retErr)

	c.Logger.Info("Executing module", "module", req.Name, "params", req.Params)

	if err := c.ensureConnection(); err != nil {
		return nil, err
	}

	modReq := &module.ExecuteRequest{
		Module:        req.Name,
		ProfileConfig: req.Config,
		Providers:     req.Providers,
		Params:        req.Params,
	}

	output, err := c.ModuleManager.ExecuteModule(ctx, modReq)
	if err != nil {
		return nil, err
	}

	if output.Error != nil {
		c.Logger.Error("Module execution failed with error", "error", output.Error)
	} else {
		c.Logger.Info("Module execution finished")
		c.Logger.Debug("Module execution results", "data", output)
	}

	return output, nil
}

func (c *Client) Close() {
	c.Manager.Shutdown()
	if c.db != nil {
		c.db.Close()
	}
}

func (c *Client) initModules() {
	c.ModuleManager = module.NewManager(c.db, c.Logger, c)
	c.ModuleManager.RegisterModule(drift.New(c.Logger))
}

func (c *Client) ensureConnection() error {
	if c.dialectExecutor != nil {
		return nil
	}
	return fmt.Errorf("missing connection info in config.hcl")
}

func (c *Client) getProviderConfig(providerName string) (*config.RequiredProvider, error) {
	var providerConfig *config.RequiredProvider
	for _, p := range c.Providers {
		if p.Name == providerName {
			providerConfig = p
			break
		}
	}
	if providerConfig == nil {
		return nil, fmt.Errorf("provider %s doesn't exist in configuration", providerName)
	}
	return providerConfig, nil
}

func (c *Client) CreatePlugin(providerName, alias string, env []string) (plugin.Plugin, error) {
	rp := c.Providers.Get(providerName)
	if rp == nil {
		return nil, fmt.Errorf("failed to find provider in configuration %s", providerName)
	}
	return c.Manager.CreatePlugin(&plugin.CreationOptions{
		Provider: registry.Provider{
			Name:    providerName,
			Version: rp.Version,
		},
		Alias: alias,
		Env:   env,
	})
}

func parsePartialFetchKV(r *cqproto.FailedResourceFetch) []interface{} {
	kv := []interface{}{"table", r.TableName, "err", r.Error}
	if r.RootTableName != "" {
		kv = append(kv, "root_table", r.RootTableName, "root_table_pks", r.RootPrimaryKeyValues)
	}
	return kv
}

// normalizeResources returns a canonical list of resources given a list of requested and all known resources.
// It replaces wildcard resource with all resources. Error is returned if:
//
// * wildcard is present and other explicit resource is requested;
// * one of explicitly requested resources is not present in all known;
// * some resource is specified more than once (duplicate).
func normalizeResources(requested []string, all map[string]*schema.Table) ([]string, error) {
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

// collectProviderVersions walks over the list of required providers, determines currently loaded version of each provider
// through getVersion function and returns a map from provider name to its version.
func collectProviderVersions(providers []*config.RequiredProvider, getVersion func(providerName string) (string, error)) (map[string]*version.Version, error) {
	ver := make(map[string]*version.Version, len(providers))
	for _, p := range providers {
		s, err := getVersion(p.Name)
		if err != nil {
			return nil, err
		}
		v, err := version.NewVersion(s)
		if err != nil {
			return nil, err
		}
		ver[p.Name] = v
	}
	return ver, nil
}

// reportFetchSummaryErrors reads provided fetch summaries, persists statistics into the span and sends the errors to sentry
func reportFetchSummaryErrors(span trace.Span, fetchSummaries map[string]ProviderFetchSummary) {
	var totalFetched, totalWarnings, totalErrors uint64

	allowUnmanaged := Version == DevelopmentVersion && viper.GetBool("debug-sentry")

	for _, ps := range fetchSummaries {
		totalFetched += ps.TotalResourcesFetched
		totalWarnings += ps.Diagnostics().Warnings()
		totalErrors += ps.Diagnostics().Errors()

		span.SetAttributes(
			attribute.Int64("fetch.resources."+ps.ProviderName, int64(ps.TotalResourcesFetched)),
			attribute.Int64("fetch.warnings."+ps.ProviderName, int64(ps.Diagnostics().Warnings())),
			attribute.Int64("fetch.errors."+ps.ProviderName, int64(ps.Diagnostics().Errors())),
		)
		span.SetAttributes(telemetry.MapToAttributes(ps.Metrics())...)

		if ps.Version == plugin.Unmanaged && !allowUnmanaged {
			continue
		}

		for _, e := range ps.Diagnostics().Squash() {
			if telemetry.ShouldIgnoreDiag(e) {
				continue
			}

			if rd, ok := e.(diag.Redactable); ok {
				if r := rd.Redacted(); r != nil {
					e = r
				}
			}

			if e.Severity() == diag.IGNORE {
				continue
			}

			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetTags(map[string]string{
					"diag_type":        e.Type().String(),
					"provider":         ps.ProviderName,
					"provider_version": ps.Version,
					"resource":         e.Description().Resource,
				})
				scope.SetExtra("detail", e.Description().Detail)
				switch e.Severity() {
				case diag.IGNORE:
					scope.SetLevel(sentry.LevelDebug)
				case diag.WARNING:
					scope.SetLevel(sentry.LevelWarning)
				case diag.PANIC:
					scope.SetLevel(sentry.LevelFatal)
				}
				sentry.CaptureException(e)
			})
		}
	}

	span.SetAttributes(
		attribute.Int64("fetch.resources.total", int64(totalFetched)),
		attribute.Int64("fetch.warnings.total", int64(totalWarnings)),
		attribute.Int64("fetch.errors.total", int64(totalErrors)),
	)
}

func (c *Client) initDatabase(ctx context.Context) error {
	var err error
	c.db, err = sdkdb.New(ctx, c.Logger, c.DSN)
	if err != nil {
		return err
	}

	var dt schema.DialectType
	dt, c.dialectExecutor, err = database.GetExecutor(c.DSN, c.HistoryCfg)
	if err != nil {
		return fmt.Errorf("getExecutor: %w", err)
	}

	if c.HistoryCfg != nil && dt != schema.TSDB {
		// check if we're already on TSDB but the dsn is wrong
		ts, err := timescale.New(c.DSN, c.HistoryCfg)
		if err != nil {
			return err
		}
		if ok, err := ts.Validate(ctx); ok && err == nil {
			return fmt.Errorf("you must update the dsn to use tsdb:// prefix")
		}

		return fmt.Errorf("history is only supported on timescaledb")
	}

	if ok, err := c.dialectExecutor.Validate(ctx); err != nil {
		if !ok {
			return fmt.Errorf("validate: %w", err)
		}
		c.Logger.Warn("database validation warning", "message", err.Error())
	} else if !ok {
		c.Logger.Warn("database validation warning")
	}

	c.metaStorage = meta_storage.NewClient(c.db, c.Logger)
	// migrate cloudquery core tables to latest version
	if err := c.metaStorage.MigrateCore(ctx, c.dialectExecutor); err != nil {
		return fmt.Errorf("failed to migrate cloudquery_core tables: %w", err)
	}

	return nil
}

// reportNumProviders counts multiple (aliased) providers and sets tracing and sentry specific attributes
func reportNumProviders(ctx context.Context, provs []*config.Provider) {
	numProviders := make(map[string]int, len(provs))
	for _, p := range provs {
		numProviders[p.Name]++
	}
	var multiProviders []string
	for k, v := range numProviders {
		if v > 1 {
			multiProviders = append(multiProviders, k+":"+strconv.Itoa(v))
		}
	}
	if len(multiProviders) == 0 {
		return
	}

	sort.Strings(multiProviders)
	trace.SpanFromContext(ctx).SetAttributes(
		attribute.StringSlice("multi_providers", multiProviders),
	)
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(map[string]string{
			"multi_providers": strings.Join(multiProviders, ","),
		})
	})
}

func ParseProviderSource(requestedProvider *config.RequiredProvider) (string, string, error) {
	var requestedSource string
	if requestedProvider.Source == nil || *requestedProvider.Source == "" {
		requestedSource = requestedProvider.Name
	} else {
		requestedSource = *requestedProvider.Source
	}
	return registry.ParseProviderName(requestedSource)
}
