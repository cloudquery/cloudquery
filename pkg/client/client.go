package client

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/client/state"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/client/database"
	"github.com/cloudquery/cloudquery/pkg/client/database/timescale"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/getsentry/sentry-go"
	"github.com/hashicorp/go-hclog"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

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

func (p ProviderFetchSummary) String() string {
	if p.ProviderAlias != "" {
		return fmt.Sprintf("%s(%s)", p.ProviderName, p.ProviderAlias)
	}
	return p.ProviderName
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
	// HistoryConfig defines configuration for CloudQuery history mode
	HistoryCfg *history.Config

	// metaStorage interacts with cloudquery core resources
	metaStorage     *state.Client
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

	return c, nil
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

func (c *Client) Close() {
	c.Manager.Shutdown()
	if c.db != nil {
		c.db.Close()
	}
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

	c.metaStorage = state.NewClient(c.db, c.Logger)
	// migrate cloudquery core tables to latest version
	if err := c.metaStorage.MigrateCore(ctx, c.dialectExecutor); err != nil {
		return fmt.Errorf("failed to migrate cloudquery_core tables: %w", err)
	}

	return nil
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
