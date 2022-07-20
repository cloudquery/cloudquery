package console

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/internal/firebase"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/core/state"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/vbauerster/mpb/v6/decor"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client console client is a wrapper around core.Client for console execution of CloudQuery
type Client struct {
	downloadProgress ui.Progress
	cfg              *config.Config
	Providers        registry.Providers
	Registry         registry.Registry
	PluginManager    *plugin.Manager
	Storage          database.Storage
	StateManager     *state.Client
	instanceId       uuid.UUID
}

func CreateClient(ctx context.Context, configPath string, allowDefaultConfig bool, configMutator func(*config.Config) error, instanceId uuid.UUID) (*Client, error) {
	cfg, ok := loadConfig(configPath)
	if !ok {
		if !allowDefaultConfig {
			return nil, diag.FromError(fmt.Errorf("config not read"), diag.USER)
		}
		cfg = &config.Config{
			CloudQuery: config.CloudQuery{
				PluginDirectory: "./.cq/providers",
				PolicyDirectory: "./.cq/policies",
				Connection:      &config.Connection{},
			},
		}
	}

	if configMutator != nil {
		if err := configMutator(cfg); err != nil {
			return nil, err
		}
	}

	setConfigAnalytics(cfg)
	return CreateClientFromConfig(ctx, cfg, instanceId)
}

func CreateClientFromConfig(ctx context.Context, cfg *config.Config, instanceId uuid.UUID) (*Client, error) {
	if cfg.CloudQuery.Connection == nil {
		return nil, errors.New("connection configuration is not set")
	}

	var progressUpdater ui.Progress

	if ui.DoProgress() {
		progressUpdater = NewProgress(ctx, func(o *ProgressOptions) {
			o.AppendDecorators = []decor.Decorator{decor.Percentage()}
		})
	}
	hub := registry.NewRegistryHub(firebase.CloudQueryRegistryURL, registry.WithPluginDirectory(cfg.CloudQuery.PluginDirectory), registry.WithProgress(progressUpdater))
	pm, err := plugin.NewManager(hub, plugin.WithAllowReattach())
	if err != nil {
		return nil, err
	}

	c := &Client{
		downloadProgress: progressUpdater,
		cfg:              cfg,
		Registry:         hub,
		PluginManager:    pm,
		instanceId:       instanceId,
	}

	if cfg.CloudQuery.Connection.DSN != "" {
		_, dialect, err := database.GetExecutor(cfg.CloudQuery.Connection.DSN)
		if err != nil {
			return nil, err
		}
		if ok, err := dialect.Validate(ctx); err != nil {
			if !ok {
				return nil, fmt.Errorf("validate: %w", err)
			}
			log.Warn().Err(err).Msg("database validation warning")
		} else if !ok {
			log.Warn().Msg("database validation warning")
		}

		if dbId, ok := dialect.Identifier(ctx); ok {
			setAnalyticsProperties(map[string]interface{}{"database_id": dbId})
			setUserId(dbId)
		}

		if info, err := dialect.Info(ctx); err == nil {
			setAnalyticsProperties(map[string]interface{}{"database_version": info.Version, "database_uptime": info.Uptime.Seconds(), "database_full_version": info.FullVersion})
		}

		c.Storage = database.NewStorage(cfg.CloudQuery.Connection.DSN, dialect)
		c.StateManager, err = state.NewClient(ctx, c.Storage.DSN())
		if err != nil {
			return nil, fmt.Errorf("could not init state: %w", err)
		}
	}
	c.Providers = make(registry.Providers, len(cfg.CloudQuery.Providers))
	for i, rp := range cfg.CloudQuery.Providers {
		src, name, err := core.ParseProviderSource(rp)
		if err != nil {
			return nil, err
		}
		c.Providers[i] = registry.Provider{Name: name, Version: rp.Version, Source: src}
	}

	c.checkForUpdate(ctx)
	return c, nil
}

// =====================================================================================================================
// 													Base Commands
// =====================================================================================================================

func (c Client) DownloadProviders(ctx context.Context) (diags diag.Diagnostics) {
	// make sure to print diagnostics, if any exist
	defer printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	ui.ColorizedOutput(ui.ColorProgress, "Initializing CloudQuery Providers...\n\n")
	_, diags = core.Download(ctx, c.PluginManager, &core.DownloadOptions{Providers: c.Providers, NoVerify: viper.GetBool("no-verify")})
	if c.downloadProgress != nil {
		c.downloadProgress.Wait()
	}

	if diags.HasErrors() {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ failed to initialize provider: %s.\n\n", diags.Error())
		return diags
	}
	ui.ColorizedOutput(ui.ColorProgress, " Finished provider initialization...\n\n")

	if viper.GetBool("no-provider-update") {
		log.Debug().Msg("Skipping provider update check")
		for _, provider := range c.Providers {
			// Colorized output is designed to match the output when user uses `--no-verify`
			ui.ColorizedOutput(ui.ColorInfo, "⚠️ %s@%s Skipped provider update check...\n", registry.ProviderRepoName(provider.Name), provider.Version)
		}
		return diags
	}

	ui.ColorizedOutput(ui.ColorProgress, "Checking available provider updates...\n\n")
	checkUpdateOpts := core.CheckUpdatesOptions{
		Providers: core.ManagedProviders(c.PluginManager, c.Providers),
	}

	updates, dd := core.CheckAvailableUpdates(ctx, c.Registry, &checkUpdateOpts)
	if dd.HasErrors() {
		return diags.Add(dd)
	}
	for _, u := range updates {
		ui.ColorizedOutput(ui.ColorInfo, fmt.Sprintf("Update available for provider %s: %s ➡️  %s\n\n", u.Name, u.CurrentVersion, u.AvailableVersion))
	}
	return diags
}

func (c Client) Fetch(ctx context.Context) (*core.FetchResponse, diag.Diagnostics) {
	if _, dd := c.SyncProviders(ctx, c.cfg.Providers.Names()...); dd.HasErrors() {
		return nil, dd
	}

	ui.ColorizedOutput(ui.ColorProgress, "Starting provider fetch...\n\n")
	var (
		fetchProgress ui.Progress
		fetchCallback core.FetchUpdateCallback
	)
	if ui.DoProgress() {
		fetchProgress, fetchCallback = buildFetchProgress(ctx, c.cfg.Providers)
	}

	providers := make([]core.ProviderInfo, len(c.cfg.Providers))
	for i, p := range c.cfg.Providers {
		rp, ok := c.Providers.Get(p.Name)
		if !ok {
			diags := diag.FromError(fmt.Errorf("failed to find provider %s in configuration", p.Name), diag.USER)
			printDiagnostics("Fetch", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
			return nil, diags
		}
		providers[i] = core.ProviderInfo{Provider: rp, Config: p}
	}
	result, diags := core.Fetch(ctx, c.StateManager, c.Storage, c.PluginManager, &core.FetchOptions{
		UpdateCallback: fetchCallback,
		ProvidersInfo:  providers,
		FetchId:        c.instanceId,
	})
	// first wait for progress to complete correctly
	if fetchProgress != nil {
		fetchProgress.MarkAllDone()
		fetchProgress.Wait()
	}
	// Check if any errors are found
	if diags.HasErrors() {
		// Ignore context cancelled error
		if st, ok := status.FromError(diags); ok && st.Code() == gcodes.Canceled {
			printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
			ui.ColorizedOutput(ui.ColorProgress, "Provider fetch canceled.\n\n")
			return result, diags
		}
	}
	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")
	printDiagnostics("Fetch", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	if result == nil {
		return nil, diags
	}
	for _, summary := range result.ProviderFetchSummary {
		PrintProviderSummary(summary)
	}
	return result, diags
}

// =====================================================================================================================
// 													Provider Commands
// =====================================================================================================================

func (c Client) SyncProviders(ctx context.Context, pp ...string) (results []*core.SyncResult, diags diag.Diagnostics) {
	defer printDiagnostics("Sync", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	providers := c.Providers
	if len(pp) > 0 {
		providers = c.Providers.GetMany(pp...)
	}
	ui.ColorizedOutput(ui.ColorProgress, "Syncing CloudQuery providers %s\n\n", providers)
	if len(providers) == 0 {
		return nil, diag.FromError(fmt.Errorf("one or more providers not found: %s", pp), diag.USER,
			diag.WithDetails("providers not found, are they defined in configuration?. Defined: %s", c.Providers))
	}
	diags = diags.Add(c.DownloadProviders(ctx))
	if diags.HasErrors() {
		return nil, diags
	}

	for _, p := range providers {
		sync, dd := core.Sync(ctx, c.StateManager, c.PluginManager, p)
		if dd.HasErrors() {
			ui.ColorizedOutput(ui.ColorError, "%s failed to sync provider %s.\n", emojiStatus[ui.StatusError], p.String())
			// TODO: should we just append diags and continue to sync others or stop syncing?
			return nil, dd
		}
		if sync.State != core.NoChange {
			ui.ColorizedOutput(ui.ColorSuccess, "%s sync provider %s to %s successfully. [%s]\n", emojiStatus[ui.StatusOK], p.Name, p.Version, sync.State)
		}
		diags = diags.Add(dd)
		if sync != nil {
			results = append(results, sync)
		}
	}
	ui.ColorizedOutput(ui.ColorProgress, "\nFinished syncing providers...\n\n")
	return results, diags
}

func (c Client) DropProvider(ctx context.Context, providerName string) (diags diag.Diagnostics) {
	defer printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	ui.ColorizedOutput(ui.ColorProgress, "Dropping CloudQuery provider %s schema...\n\n", providerName)
	if dd := c.DownloadProviders(ctx); dd.HasErrors() {
		return dd
	}
	p, ok := c.Providers.Get(providerName)
	if !ok {
		return diag.FromError(fmt.Errorf("failed to find provider %s in configuration", p.Name), diag.USER)
	}

	if diags = core.Drop(ctx, c.StateManager, c.PluginManager, p); diags.HasErrors() {
		ui.ColorizedOutput(ui.ColorError, "%s Failed to drop provider %s schema.\n\n", emojiStatus[ui.StatusError], providerName)
		return diags
	}
	ui.ColorizedOutput(ui.ColorSuccess, "%s provider %s schema dropped successfully.\n\n", emojiStatus[ui.StatusOK], providerName)
	return diags
}

func (c Client) RemoveStaleData(ctx context.Context, lastUpdate time.Duration, dryRun bool, providers []string) (diags diag.Diagnostics) {
	defer printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	if dd := c.DownloadProviders(ctx); dd.HasErrors() {
		return dd
	}
	pp := make([]registry.Provider, len(providers))
	for i, p := range providers {
		rp, ok := c.Providers.Get(p)
		if !ok {
			ui.ColorizedOutput(ui.ColorHeader, "unknown provider %s requested..\n\n", p)
		}
		pp[i] = rp
	}
	ui.ColorizedOutput(ui.ColorHeader, "Purging providers %s resources..\n\n", providers)
	defer printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	result, diags := core.PurgeProviderData(ctx, c.Storage, c.PluginManager, &core.PurgeProviderDataOptions{
		Providers:  pp,
		LastUpdate: lastUpdate,
		DryRun:     dryRun,
	})

	if dryRun && !diags.HasErrors() {
		ui.ColorizedOutput(ui.ColorWarning, "Expected resources to be purged: %d. Use --dry-run=false to purge these resources.\n", result.TotalAffected)
		for _, r := range result.Resources() {
			ui.ColorizedOutput(ui.ColorWarning, "\t%s: %d resources\n\n", r, result.AffectedResources[r])
		}
	}
	if diags.HasErrors() {
		ui.ColorizedOutput(ui.ColorProgress, "Purge for providers %s failed\n\n", providers)
		return diags
	}
	ui.ColorizedOutput(ui.ColorProgress, "Purge for providers %s was successful\n\n", providers)
	return diags
}

func (c Client) Close() {
	c.PluginManager.Shutdown()
	if c.StateManager != nil {
		c.StateManager.Close()
	}
}

func (Client) checkForUpdate(ctx context.Context) {
	v, err := core.CheckCoreUpdate(ctx, afero.Afero{Fs: afero.NewOsFs()}, time.Now().Unix(), core.UpdateCheckPeriod)
	if err != nil {
		log.Warn().Err(err).Msg("update check failed")
		return
	}
	if v != nil {
		ui.ColorizedOutput(ui.ColorInfo, "An update to CloudQuery core is available: %s!\n\n", v)
		log.Debug().Str("new_version", v.String()).Msg("update check succeeded")
	} else {
		log.Debug().Msg("update check succeeded, no new version")
	}
}

func buildFetchProgress(ctx context.Context, providers []*config.Provider) (*Progress, core.FetchUpdateCallback) {
	fetchProgress := NewProgress(ctx, func(o *ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.CountersNoUnit(" Finished Resources: %d/%d")}
	})

	for _, p := range providers {
		if len(p.Resources) == 0 {
			ui.ColorizedOutput(ui.ColorWarning, "%s Skipping provider %s[%s] configured with no resource to fetch\n", emojiStatus[ui.StatusWarn], p.Name, p.Alias)
			continue
		}

		if p.Alias != p.Name {
			fetchProgress.Add(fmt.Sprintf("%s_%s", p.Name, p.Alias), fmt.Sprintf("cq-provider-%s (%s)", p.Name, p.Alias), "fetching", int64(len(p.Resources)))
		} else {
			fetchProgress.Add(fmt.Sprintf("%s_%s", p.Name, p.Alias), fmt.Sprintf("cq-provider-%s", p.Name), "fetching", int64(len(p.Resources)))
		}
	}
	fetchCallback := func(update core.FetchUpdate) {
		name := fmt.Sprintf("%s_%s", update.Name, update.Name)
		if update.Alias != "" {
			name = fmt.Sprintf("%s_%s", update.Name, update.Alias)
		}
		if update.DiagnosticCount > 0 {
			fetchProgress.Update(name, ui.StatusWarn, fmt.Sprintf("diagnostics: %d", update.DiagnosticCount), 0)
		}
		bar := fetchProgress.GetBar(name)
		if bar == nil {
			fetchProgress.AbortAll()
			ui.ColorizedOutput(ui.ColorError, "❌ console UI failure, fetch will complete shortly\n")
			return
		}
		if bar.Total < int64(len(update.FinishedResources)) {
			bar.SetTotal(int64(len(update.FinishedResources)), false)
		}

		bar.b.IncrBy(update.DoneCount() - int(bar.b.Current()))

		if bar.Status == ui.StatusError {
			if update.AllDone() {
				bar.SetTotal(0, true)
			}
			return
		}
		if update.AllDone() && bar.Status != ui.StatusWarn {
			fetchProgress.Update(name, ui.StatusOK, "fetch complete", 0)
			return
		}
	}
	return fetchProgress, fetchCallback
}

func loadConfig(file string) (*config.Config, bool) {
	cfg, diags := config.NewParser(
		config.WithEnvironmentVariables(config.EnvVarPrefix, os.Environ()),
	).LoadConfigFile(file)
	if diags.HasDiags() {
		ui.ColorizedOutput(ui.ColorHeader, "Configuration Error Diagnostics:\n")
		for _, d := range diags {
			c := ui.ColorInfo
			switch d.Severity() {
			case diag.ERROR:
				c = ui.ColorError
			case diag.WARNING:
				c = ui.ColorWarning
			}
			ui.ColorizedOutput(c, "❌ %s; %s\n", d.Description().Summary, d.Description().Detail)
		}
		if diags.HasErrors() {
			return nil, false
		}
	}
	return cfg, true
}

func setAnalyticsProperties(props map[string]interface{}) {
	sprops := make(map[string]string, len(props))
	for k, v := range props {
		analytics.SetGlobalProperty(k, v)
		sprops[k] = fmt.Sprintf("%v", v)
	}
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(sprops)
	})
}

func setUserId(newId string) {
	analytics.SetUserId(newId)
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{
			ID: newId,
		})
	})
}

func setConfigAnalytics(cfg *config.Config) {
	cfgJSON, _ := json.Marshal(cfg)
	s := sha256.New()
	_, _ = s.Write(cfgJSON)
	cfgHash := fmt.Sprintf("%0x", s.Sum(nil))
	analytics.SetGlobalProperty("cfghash", cfgHash)

	const cfgf = "yaml"
	analytics.SetGlobalProperty("cfgformat", cfgf)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		if analytics.IsCI() {
			scope.SetUser(sentry.User{
				ID: cfgHash,
			})
		}
		scope.SetTags(map[string]string{
			"cfghash":   cfgHash,
			"cfgformat": cfgf,
		})
	})
}

func (c Client) ConvertRequiredToRegistry(providerName string) registry.Provider {
	rp := c.cfg.CloudQuery.Providers.Get(providerName)
	src, name, _ := core.ParseProviderSource(rp)
	return registry.Provider{Name: name, Version: rp.Version, Source: src}
}
