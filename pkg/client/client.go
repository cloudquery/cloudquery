package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	zerolog "github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	otrace "go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/module/drift"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/schema/diag"
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
	// Optional: Disable deletion of data from tables.
	// Use this with caution, as it can create duplicates of data!
	DisableDataDelete bool
	// Optional: Adds extra fields to the provider, this is used for testing purposes.
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
	PartialFetchErrors    []*cqproto.FailedResourceFetch
	FetchErrors           []error
	TotalResourcesFetched uint64
	FetchResources        map[string]cqproto.ResourceFetchSummary
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

// PoliciesRunRequest is the request used to run a policy.
type PoliciesRunRequest struct {
	// Policies to run
	Policies []*config.Policy

	// PolicyName is optional attr to run specific policy
	PolicyName string

	// OutputDir is the output dir for policy execution output.
	OutputDir string

	// StopOnFailure signals policy execution to stop after first failure.
	StopOnFailure bool

	// RunCallBack is the callback method that is called after every policy execution.
	RunCallback policy.UpdateCallback

	// SkipVersioning if true policy will be executed without checking out the version of the policy repo using git tags
	SkipVersioning bool

	// FailOnViolation if true policy run will return error if there are violations
	FailOnViolation bool
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

// TableCreator creates tables based on schema received from providers
type TableCreator interface {
	CreateTable(ctx context.Context, conn *pgxpool.Conn, t *schema.Table, p *schema.Table) error
}

type FetchUpdateCallback func(update FetchUpdate)

type Option func(options *Client)

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Required: List of providers that are required, these providers will be downloaded if DownloadProviders is called.
	Providers []*config.RequiredProvider
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
	// Optional: Hub client to use to download plugins, the Hub is used to download and pluginManager providers binaries
	// if not specified, default cloudquery registry is used.
	Hub registry.Hub
	// manager manages all plugins lifecycle
	Manager *plugin.Manager
	// ModuleManager manages all modules lifecycle
	ModuleManager module.Manager
	// ModuleManager manages all modules lifecycle
	PolicyManager policy.Manager
	// TableCreator defines how table are created in the database
	TableCreator TableCreator
	// pool is a list of connection that are used for policy/query execution
	pool *pgxpool.Pool
}

func New(ctx context.Context, options ...Option) (*Client, error) {

	c := &Client{
		PluginDirectory:    filepath.Join(".", ".cq", "providers"),
		PolicyDirectory:    ".",
		NoVerify:           false,
		SkipBuildTables:    false,
		HubProgressUpdater: nil,
		RegistryURL:        registry.CloudQueryRegistryURl,
		Logger:             logging.NewZHcLog(&zerolog.Logger, ""),
	}
	for _, o := range options {
		o(c)
	}

	var err error
	c.Manager, err = plugin.NewManager(c.Logger, c.PluginDirectory, c.RegistryURL, c.HubProgressUpdater)
	if err != nil {
		return nil, err
	}
	c.Manager.LoadExisting(c.Providers)

	if c.TableCreator == nil {
		c.TableCreator = provider.NewTableCreator(c.Logger)
	}
	if c.DSN == "" {
		c.Logger.Warn("missing DSN, some commands won't work")
	} else {
		poolCfg, err := pgxpool.ParseConfig(c.DSN)
		if err != nil {
			return nil, err
		}
		poolCfg.LazyConnect = true
		c.pool, err = pgxpool.ConnectConfig(ctx, poolCfg)
		if err != nil {
			return nil, err
		}
		if err := validatePostgresVersion(ctx, c.pool, minPostgresVersion); err != nil {
			c.Logger.Warn(err.Error())
		}
	}

	c.initModules()

	c.PolicyManager = policy.NewManager(c.PolicyDirectory, c.pool, c.Logger)

	return c, nil
}

// DownloadProviders downloads all provider binaries
func (c *Client) DownloadProviders(ctx context.Context) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "DownloadProviders")
	defer spanEnder(retErr)

	c.Logger.Info("Downloading required providers")
	return c.Manager.DownloadProviders(ctx, c.Providers, c.NoVerify)
}

func (c *Client) TestProvider(ctx context.Context, providerCfg *config.Provider) error {
	providerPlugin, err := c.Manager.CreatePlugin(providerCfg.Name, providerCfg.Alias, providerCfg.Env)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerCfg.Name, "error", err)
		return err
	}
	defer providerPlugin.Close()
	c.Logger.Info("requesting provider to configure", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
	_, err = providerPlugin.Provider().ConfigureProvider(ctx, &cqproto.ConfigureProviderRequest{
		CloudQueryVersion: Version,
		Connection: cqproto.ConnectionDetails{
			DSN: c.DSN,
		},
		Config: providerCfg.Configuration,
	})
	if err != nil {
		return fmt.Errorf("provider test connection failed. Reason: %w", err)
	}
	return nil
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
	if !c.SkipBuildTables {
		for _, p := range request.Providers {
			if err := c.BuildProviderTables(ctx, p.Name); err != nil {
				return nil, err
			}
		}
	}

	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "Fetch")
	defer spanEnder(retErr)

	c.Logger.Info("received fetch request", "disable_delete", request.DisableDataDelete, "extra_fields", request.ExtraFields)

	fetchSummaries := make(chan ProviderFetchSummary, len(request.Providers))
	errGroup, gctx := errgroup.WithContext(ctx)
	for _, providerConfig := range request.Providers {
		providerConfig := providerConfig
		c.Logger.Debug("creating provider plugin", "provider", providerConfig.Name)
		providerPlugin, err := c.Manager.CreatePlugin(providerConfig.Name, providerConfig.Alias, providerConfig.Env)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", providerConfig.Name, "error", err)
			return nil, err
		}
		// TODO: move this into an outer function
		errGroup.Go(func() error {
			pLog := c.Logger.With("provider", providerConfig.Name, "alias", providerConfig.Alias, "version", providerPlugin.Version())
			pLog.Info("requesting provider to configure")
			_, err = providerPlugin.Provider().ConfigureProvider(gctx, &cqproto.ConfigureProviderRequest{
				CloudQueryVersion: Version,
				Connection: cqproto.ConnectionDetails{
					DSN: c.DSN,
				},
				Config:        providerConfig.Configuration,
				DisableDelete: request.DisableDataDelete,
				ExtraFields:   request.ExtraFields,
			})
			if err != nil {
				pLog.Error("failed to configure provider", "error", err)
				return err
			}
			pLog.Info("provider configured successfully")

			pLog.Info("requesting provider fetch", "partial_fetch_enabled", providerConfig.EnablePartialFetch)
			fetchStart := time.Now()
			stream, err := providerPlugin.Provider().FetchResources(gctx, &cqproto.FetchResourcesRequest{Resources: providerConfig.Resources, PartialFetchingEnabled: providerConfig.EnablePartialFetch})
			if err != nil {
				return err
			}
			pLog.Info("provider started fetching resources")
			var (
				fetchErrors         = make([]error, 0)
				partialFetchResults []*cqproto.FailedResourceFetch
				fetchedResources           = make(map[string]cqproto.ResourceFetchSummary, len(providerConfig.Resources))
				totalResources      uint64 = 0
			)
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					pLog.Info("provider finished fetch", "execution", time.Since(fetchStart).String())
					for _, fetchError := range partialFetchResults {
						pLog.Warn("received partial fetch error", parsePartialFetchKV(fetchError)...)
					}
					fetchSummaries <- ProviderFetchSummary{
						ProviderName:          providerConfig.Name,
						TotalResourcesFetched: totalResources,
						PartialFetchErrors:    partialFetchResults,
						FetchErrors:           fetchErrors,
						FetchResources:        fetchedResources,
					}
					return nil
				}
				if err != nil {
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

				totalResources = resp.ResourceCount
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
		response.ProviderFetchSummary[ps.ProviderName] = ps
	}

	collectFetchSummaryStats(otrace.SpanFromContext(ctx), response.ProviderFetchSummary)

	return response, nil
}

func (c *Client) GetProviderSchema(ctx context.Context, providerName string) (*cqproto.GetProviderSchemaResponse, error) {
	providerPlugin, err := c.Manager.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if providerPlugin.Version() == plugin.Unmanaged {
			c.Logger.Warn("Not closing unmanaged provider", "provider", providerName)
			return
		}
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to kill provider", "provider", providerName)
		}
	}()
	return providerPlugin.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
}

func (c *Client) GetProviderConfiguration(ctx context.Context, providerName string) (*cqproto.GetProviderConfigResponse, error) {
	providerPlugin, err := c.Manager.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if providerPlugin.Version() == plugin.Unmanaged {
			c.Logger.Warn("Not closing unmanaged provider", "provider", providerName)
			return
		}
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to close provider", "provider", providerName)
		}
	}()
	return providerPlugin.Provider().GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
}

func (c *Client) BuildProviderTables(ctx context.Context, providerName string) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "BuildProviderTables", otrace.WithAttributes(
		attribute.String("provider", providerName),
	))
	defer spanEnder(retErr)

	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	conn, err := c.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	for name, t := range s.ResourceTables {
		c.Logger.Debug("creating tables for resource for provider", "resource_name", name, "provider", s.Name, "version", s.Version)
		if err := c.TableCreator.CreateTable(ctx, conn, t, nil); err != nil {
			return err
		}
	}

	if s.Migrations == nil {
		c.Logger.Debug("provider doesn't support migrations", "provider", providerName)
		return nil
	}
	// create migration table and set it to version based on latest create table
	m, cfg, err := c.buildProviderMigrator(s.Migrations, providerName)
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			c.Logger.Error("failed to close migrator connection", "error", err)
		}
	}()
	if _, _, err := m.Version(); err == migrate.ErrNilVersion {
		mv, err := m.FindLatestMigration(cfg.Version)
		if err != nil {
			return err
		}
		c.Logger.Debug("setting provider schema migration version", "version", cfg.Version, "migration_version", mv)
		return m.SetVersion(cfg.Version)
	}
	return nil
}

func (c *Client) UpgradeProvider(ctx context.Context, providerName string) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "UpgradeProvider", otrace.WithAttributes(
		attribute.String("provider", providerName),
	))

	defer func() {
		if retErr != nil && (retErr == migrate.ErrNoChange || errors.Is(retErr, ErrMigrationsNotSupported)) {
			spanEnder(nil)
		} else {
			spanEnder(retErr)
		}
	}()

	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	if s.Migrations == nil {
		return ErrMigrationsNotSupported
	}
	m, cfg, err := c.buildProviderMigrator(s.Migrations, providerName)
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			c.Logger.Error("failed to close migrator connection", "error", err)
		}
	}()

	pVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("failed to get provider version: %w", err)
	}
	otrace.SpanFromContext(ctx).SetAttributes(attribute.String("old_version", pVersion))

	if dirty {
		return fmt.Errorf("provider schema is dirty, please drop provider and recreate")
	}
	if pVersion == "v0.0.0" {
		return c.BuildProviderTables(ctx, providerName)
	}
	c.Logger.Info("upgrading provider version", "version", cfg.Version, "provider", cfg.Name)
	otrace.SpanFromContext(ctx).SetAttributes(attribute.String("new_version", cfg.Version))

	return m.UpgradeProvider(cfg.Version)
}

func (c *Client) DowngradeProvider(ctx context.Context, providerName string) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "DowngradeProvider", otrace.WithAttributes(
		attribute.String("provider", providerName),
	))
	defer spanEnder(retErr)

	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	if s.Migrations == nil {
		return fmt.Errorf("provider doesn't support migrations")
	}
	m, cfg, err := c.buildProviderMigrator(s.Migrations, providerName)
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			c.Logger.Error("failed to close migrator connection", "error", err)
		}
	}()
	c.Logger.Info("downgrading provider version", "version", cfg.Version, "provider", cfg.Name)
	return m.DowngradeProvider(cfg.Version)
}

func (c *Client) DropProvider(ctx context.Context, providerName string) (retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "DropProvider", otrace.WithAttributes(
		attribute.String("provider", providerName),
	))
	defer spanEnder(retErr)

	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	m, cfg, err := c.buildProviderMigrator(s.Migrations, providerName)
	if err != nil {
		return err
	}
	defer func() {
		if err := m.Close(); err != nil {
			c.Logger.Error("failed to close migrator connection", "error", err)
		}
	}()
	c.Logger.Info("dropping provider tables", "version", cfg.Version, "provider", cfg.Name)
	return m.DropProvider(ctx, s.ResourceTables)
}

func (c *Client) DownloadPolicy(ctx context.Context, args []string) (pol *policy.RemotePolicy, retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "DownloadPolicy")
	defer spanEnder(retErr)

	c.Logger.Info("Downloading policy from GitHub", "args", args)

	remotePolicy, err := policy.ParsePolicyFromArgs(args)
	if err != nil {
		return nil, err
	}
	c.Logger.Debug("Parsed policy download input arguments", "policy", args)

	err = c.PolicyManager.DownloadPolicy(ctx, remotePolicy)
	if err != nil {
		return nil, err
	}
	return remotePolicy, nil
}

func (c *Client) RunPolicies(ctx context.Context, req *PoliciesRunRequest) ([]*policy.ExecutionResult, error) {
	results := make([]*policy.ExecutionResult, 0)

	for _, policyConfig := range req.Policies {
		result, err := c.runPolicy(ctx, policyConfig, req)

		c.Logger.Debug("Policy execution finished", "name", policyConfig.Name, "err", err)
		if err != nil {
			c.Logger.Error("Policy execution finished with error", "name", policyConfig.Name, "err", err)
			// update the ui with the error
			if req.RunCallback != nil {
				req.RunCallback(policy.Update{
					PolicyName:      policyConfig.Name,
					Version:         policyConfig.Version,
					FinishedQueries: 0,
					QueriesCount:    0,
					Error:           err.Error(),
				})
			}

			// add the execution error to the results
			results = append(results, &policy.ExecutionResult{
				PolicyName: policyConfig.Name,
				Passed:     false,
				Error:      err.Error(),
			})

			// if failOnViolation is set, we should stop the execution
			if req.FailOnViolation {
				return results, nil
			}
			continue
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *Client) runPolicy(ctx context.Context, policyConfig *config.Policy, req *PoliciesRunRequest) (*policy.ExecutionResult, error) {
	c.Logger.Info("Loading policy", "args", policyConfig)
	versions, err := collectProviderVersions(c.Providers, func(name string) (string, error) {
		d, err := c.Manager.GetPluginDetails(name)
		return d.Version, err
	})

	if err != nil {
		return nil, err
	}

	execReq := &policy.ExecuteRequest{
		Policy:           policyConfig,
		StopOnFailure:    req.StopOnFailure,
		SkipVersioning:   req.SkipVersioning,
		ProviderVersions: versions,
		UpdateCallback:   req.RunCallback,
	}

	// load the policy
	c.Logger.Info("Loading the policy", "args", policyConfig)
	policies, err := c.PolicyManager.Load(ctx, policyConfig, execReq)
	if err != nil {
		c.Logger.Error("failed loading the policy", "err", err)
		return nil, fmt.Errorf("failed to load policy: %w", err)
	}

	c.Logger.Info("Running policy", "args", policyConfig)
	result, err := c.PolicyManager.Run(ctx, execReq, policies)
	if err != nil {
		return nil, err
	}

	// execution was not finished
	if !result.Passed && req.StopOnFailure && req.RunCallback != nil {
		req.RunCallback(policy.Update{
			PolicyName:      policyConfig.Name,
			Version:         policyConfig.Version,
			FinishedQueries: 0,
			QueriesCount:    0,
			Error:           "Execution stops",
		})
	}

	// Store output in file if requested
	if req.OutputDir != "" {
		c.Logger.Info("Writing policy to output directory", "args", policyConfig)
		err = policy.GenerateExecutionResultFile(result, req.OutputDir)

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Client) ExecuteModule(ctx context.Context, req ModuleRunRequest) (res *module.ExecutionResult, retErr error) {
	ctx, spanEnder := telemetry.StartSpanFromContext(ctx, "ExecuteModule", otrace.WithAttributes(attribute.String("module", req.Name)))
	defer spanEnder(retErr)

	c.Logger.Info("Executing module", "module", req.Name, "params", req.Params)

	modReq := &module.ExecuteRequest{
		Providers: req.Providers,
		Params:    req.Params,
	}

	output, err := c.ModuleManager.ExecuteModule(ctx, req.Name, req.Config, modReq)
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
	if c.pool != nil {
		c.pool.Close()
	}
}

func (c *Client) SetProviderVersion(ctx context.Context, providerName, version string) error {
	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	if s.Migrations == nil {
		return fmt.Errorf("provider doesn't support migrations")
	}
	m, cfg, err := c.buildProviderMigrator(s.Migrations, providerName)
	if err != nil {
		return err
	}
	c.Logger.Info("set provider version", "version", version, "provider", cfg.Name)
	return m.SetVersion(version)
}

func (c *Client) initModules() {
	c.ModuleManager = module.NewManager(c.pool, c.Logger)
	c.ModuleManager.RegisterModule(drift.New(c.Logger))
}

func (c *Client) buildProviderMigrator(migrations map[string][]byte, providerName string) (*provider.Migrator, *config.RequiredProvider, error) {
	providerConfig, err := c.getProviderConfig(providerName)
	if err != nil {
		return nil, nil, err
	}
	org, name, err := registry.ParseProviderName(providerConfig.Name)
	if err != nil {
		return nil, nil, err
	}
	m, err := provider.NewMigrator(c.Logger, migrations, c.DSN, fmt.Sprintf("%s_%s", org, name))
	if err != nil {
		return nil, nil, err
	}
	return m, providerConfig, err
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

func FilterPolicies(args []string, configPolicies []*config.Policy, policyName, subPath string) ([]*config.Policy, error) {
	var policies []*config.Policy

	if len(args) > 0 {
		remotePolicy, err := policy.ParsePolicyFromArgs(args)
		if err != nil {
			return nil, err
		}
		policyConfig, err := remotePolicy.ToPolicyConfig()
		policyConfig.SubPath = subPath
		if err != nil {
			return nil, err
		}
		policies = append(policies, policyConfig)
	} else {
		policies = configPolicies
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf(`
Could not find policies to run.
Please add policy to block to your config file`)
	}
	policiesToRun := make([]*config.Policy, 0)

	// select policies to run
	for _, p := range policies {
		if policyName != "" {
			// request to run only specific policy
			if policyName == p.Name {
				// override subPath if specified
				if subPath != "" {
					p.SubPath = subPath
				}
				policiesToRun = append(policiesToRun, p)
				break
			}
		}
		policiesToRun = append(policiesToRun, p)
	}

	return policiesToRun, nil
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

// collectFetchSummaryStats reads provided fetch summaries and persists statistics into the span
func collectFetchSummaryStats(span otrace.Span, fetchSummaries map[string]ProviderFetchSummary) {
	var totalFetched, totalWarnings, totalErrors uint64

	for _, ps := range fetchSummaries {
		totalFetched += ps.TotalResourcesFetched
		totalWarnings += ps.Diagnostics().Warnings()
		totalErrors += ps.Diagnostics().Errors() + uint64(len(ps.PartialFetchErrors))

		span.SetAttributes(
			attribute.Int64("fetch.resources."+ps.ProviderName, int64(ps.TotalResourcesFetched)),
			attribute.Int64("fetch.warnings."+ps.ProviderName, int64(ps.Diagnostics().Warnings())),
			attribute.Int64("fetch.errors."+ps.ProviderName, int64(ps.Diagnostics().Errors())),
			attribute.Int("fetch.partial_errors."+ps.ProviderName, len(ps.PartialFetchErrors)),
		)
	}

	span.SetAttributes(
		attribute.Int64("fetch.resources.total", int64(totalFetched)),
		attribute.Int64("fetch.warnings.total", int64(totalWarnings)),
		attribute.Int64("fetch.errors.total", int64(totalErrors)),
	)
}
