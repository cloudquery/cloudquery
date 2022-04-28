package console

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cloudquery/cloudquery/pkg/module/drift"

	"github.com/cloudquery/cloudquery/internal/analytics"
	"github.com/cloudquery/cloudquery/internal/getter"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/getsentry/sentry-go"

	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/vbauerster/mpb/v6/decor"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	fetchSummary       = "Provider %s fetch summary: %s Total Resources fetched: %d\t ⚠️ Warnings: %s\t ❌ Errors: %s\n\n"
	policyConfigFormat = `
Add this block into your CloudQuery config file:

policy "%s" {
    source = "%s"
}

`
)

// Client console client is a wrapper around core.Client for console execution of CloudQuery
type Client struct {
	downloadProgress ui.Progress
	cfg              *config.Config
	Providers        registry.Providers
	Registry         registry.Registry
	PluginManager    *plugin.Manager
	Storage          database.Storage
}

func CreateClient(ctx context.Context, configPath string, allowDefaultConfig bool, configMutator func(*config.Config) error) (*Client, error) {
	cfg, ok := loadConfig(configPath)
	if !ok {
		if !allowDefaultConfig {
			return nil, &ExitCodeError{ExitCode: 1}
		}
		cfg = &config.Config{
			CloudQuery: config.CloudQuery{
				PluginDirectory: "./.cq/providers",
				PolicyDirectory: "./.cq/policies",
				Connection:      &config.Connection{DSN: ""},
			},
		}
	}

	if configMutator != nil {
		if err := configMutator(cfg); err != nil {
			return nil, err
		}
	}
	setConfigAnalytics(cfg, cfg.CloudQuery.History != nil)
	return CreateClientFromConfig(ctx, cfg)
}

func CreateClientFromConfig(ctx context.Context, cfg *config.Config) (*Client, error) {
	if cfg.CloudQuery.Connection == nil {
		return nil, errors.New("connection configuration is not set")
	}
	var (
		progressUpdater ui.Progress
		dialect         database.DialectExecutor
		err             error
	)
	if ui.DoProgress() {
		progressUpdater = NewProgress(ctx, func(o *ProgressOptions) {
			o.AppendDecorators = []decor.Decorator{decor.Percentage()}
		})
	}
	hub := registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(cfg.CloudQuery.PluginDirectory), registry.WithProgress(progressUpdater))
	pm, err := plugin.NewManager(hub, plugin.WithAllowReattach())
	if err != nil {
		return nil, err
	}

	var storage database.Storage
	if cfg.CloudQuery.Connection.DSN != "" {
		_, dialect, err = database.GetExecutor(cfg.CloudQuery.Connection.DSN, cfg.CloudQuery.History)
		if err != nil {
			return nil, err
		}
		storage = database.NewStorage(cfg.CloudQuery.Connection.DSN, dialect)
	}
	pp := make(registry.Providers, len(cfg.CloudQuery.Providers))
	for i, rp := range cfg.CloudQuery.Providers {
		src, name, err := core.ParseProviderSource(rp)
		if err != nil {
			return nil, err
		}
		pp[i] = registry.Provider{Name: name, Version: rp.Version, Source: src}
	}

	c := &Client{progressUpdater, cfg, pp, hub, pm, storage}
	c.checkForUpdate(ctx)
	return c, err
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
	ui.ColorizedOutput(ui.ColorProgress, "Finished provider initialization...\n\n")

	ui.ColorizedOutput(ui.ColorProgress, "Checking available provider updates...\n\n")
	updates, dd := core.CheckAvailableUpdates(ctx, c.Registry, &core.CheckUpdatesOptions{Providers: c.Providers})
	if dd.HasErrors() {
		return diags.Add(dd)
	}
	for _, u := range updates {
		ui.ColorizedOutput(ui.ColorInfo, fmt.Sprintf("Update available for provider %s: %s ➡️ %s\n\n", u.Name, u.CurrentVersion, u.AvailableVersion))
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
			return nil, diag.FromError(fmt.Errorf("failed to find provider %s in configuration", p.Name), diag.INTERNAL)
		}
		providers[i] = core.ProviderInfo{Provider: rp, Config: p}
	}
	result, diags := core.Fetch(ctx, c.Storage, c.PluginManager, &core.FetchOptions{
		UpdateCallback: fetchCallback,
		ProvidersInfo:  providers,
		History:        c.cfg.CloudQuery.History,
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
	for _, summary := range result.ProviderFetchSummary {
		s := emojiStatus[ui.StatusOK]
		if summary.Status == core.FetchCanceled {
			s = emojiStatus[ui.StatusError] + " (canceled)"
		}
		key := summary.Name
		if summary.Name != summary.Alias {
			key = fmt.Sprintf("%s(%s)", summary.Name, summary.Alias)
		}
		diags := summary.Diagnostics().Squash()
		ui.ColorizedOutput(ui.ColorHeader, fetchSummary, key, s, summary.TotalResourcesFetched, countSeverity(diags, diag.WARNING), countSeverity(diags, diag.ERROR, diag.PANIC))
	}
	return result, diags
}

// =====================================================================================================================
// 													Provider Commands
// =====================================================================================================================

func (c Client) SyncProviders(ctx context.Context, pp ...string) (results []*core.SyncResult, diags diag.Diagnostics) {
	defer printDiagnostics("Sync", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	ui.ColorizedOutput(ui.ColorProgress, "Syncing CloudQuery providers %s\n\n", pp)
	providers := c.Providers
	if pp == nil {
		providers = c.Providers.GetMany(pp...)
	}
	if len(providers) == 0 {
		return nil, diag.FromError(fmt.Errorf("one or more providers not found: %s", pp), diag.USER,
			diag.WithDetails("providers not found, are they defined in configuration?. Defined: %s", c.Providers))
	}
	diags = diags.Add(c.DownloadProviders(ctx))
	if diags.HasErrors() {
		return nil, diags
	}

	for _, p := range providers {
		sync, dd := core.Sync(ctx, c.Storage, c.PluginManager, &core.SyncOptions{Provider: p, DownloadLatest: false})
		if dd.HasErrors() {
			if errors.Is(dd, core.ErrMigrationsNotSupported) {
				ui.ColorizedOutput(ui.ColorWarning, "%s failed to sync provider, migrations not supported %s.\n", emojiStatus[ui.StatusWarn], p.String())
				continue
			}
			ui.ColorizedOutput(ui.ColorError, "%s failed to sync provider %s.\n", emojiStatus[ui.StatusError], p.String())
			// TODO: should we just append diags and continue to sync others or stop syncing?
			return nil, dd
		}
		ui.ColorizedOutput(ui.ColorSuccess, "%s sync provider %s to %s successfully.\n", emojiStatus[ui.StatusOK], p.Name, p.Version)
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
		return diag.FromError(fmt.Errorf("failed to find provider %s in configuration", p.Name), diag.INTERNAL)
	}

	if diags = core.Drop(ctx, c.Storage, c.PluginManager, p); diags.HasErrors() {
		ui.ColorizedOutput(ui.ColorError, "%s Failed to drop provider %s schema.\n\n", emojiStatus[ui.StatusError], providerName)
		return diags
	}
	ui.ColorizedOutput(ui.ColorSuccess, "%s provider %s schema dropped successfully.\n\n", emojiStatus[ui.StatusOK], providerName)
	return diags
}

func (c Client) RemoveStaleData(ctx context.Context, lastUpdate time.Duration, dryRun bool, providers []string) (diags diag.Diagnostics) {
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
	if !diags.HasErrors() {
		ui.ColorizedOutput(ui.ColorProgress, "Purge for providers %s failed\n\n", providers)
		return diags
	}
	ui.ColorizedOutput(ui.ColorProgress, "Purge for providers %s was successful\n\n", providers)
	return diags
}

// =====================================================================================================================
// 													Policy Commands
// =====================================================================================================================

func (c Client) DownloadPolicy(ctx context.Context, args []string) (diags diag.Diagnostics) {
	ui.ColorizedOutput(ui.ColorProgress, "Downloading CloudQuery Policy...\n")
	defer printDiagnostics("", &diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	p, err := policy.Load(ctx, c.cfg.CloudQuery.PolicyDirectory, &policy.Policy{Name: "policy", Source: args[0]})
	if err != nil {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to Download policy: %s.\n\n", err.Error())
		return diags.Add(diag.FromError(err, diag.RESOLVING))
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished downloading policy...\n")
	// Show policy instructions
	ui.ColorizedOutput(ui.ColorHeader, fmt.Sprintf(policyConfigFormat, p.Name, p.Source))
	return nil
}

func (c Client) RunPolicies(ctx context.Context, policySource, outputDir string, noResults bool) error {
	log.Debug().Str("policy", policySource).Str("output_dir", outputDir).Bool("noResults", noResults).Msg("run policy received params")
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	policiesToRun, err := FilterPolicies(policySource, c.cfg.Policies)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return err
	}
	log.Debug().Interface("policies", policiesToRun).Msg("policies to run")

	ui.ColorizedOutput(ui.ColorProgress, "Starting policies run...\n\n")

	var (
		policyRunProgress ui.Progress
		policyRunCallback policy.UpdateCallback
	)
	// if we are running in a terminal, build the progress bar
	if ui.DoProgress() {
		policyRunProgress, policyRunCallback = buildPolicyRunProgress(ctx, policiesToRun)
	}
	// Policies run request
	results, err := policy.Run(ctx, c.Storage, &policy.RunRequest{
		Policies:    policiesToRun,
		Directory:   c.cfg.CloudQuery.PolicyDirectory,
		OutputDir:   outputDir,
		RunCallback: policyRunCallback,
	})

	if err := analytics.Capture("policy run", c.Providers, policiesToRun, diag.FromError(err, diag.INTERNAL)); err != nil {
		log.Warn().Err(err).Msg("analytic send failed")
	}

	if policyRunProgress != nil {
		policyRunProgress.MarkAllDone()
		policyRunProgress.Wait()
	}
	if !noResults {
		printPolicyResponse(results)
	}

	if err != nil {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to run policies: %s.\n\n", err.Error())
		return err
	}

	ui.ColorizedOutput(ui.ColorProgress, "Finished policies run...\n\n")
	return nil
}

func (c Client) TestPolicies(ctx context.Context, policySource, snapshotDestination string) error {
	conn, err := sdkdb.New(ctx, hclog.NewNullLogger(), c.Storage.DSN())
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to new database")
		return err
	}
	defer conn.Close()
	uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
	if err != nil {
		return err
	}

	p, err := policy.Load(ctx, c.cfg.CloudQuery.PolicyDirectory, &policy.Policy{Name: "test-policy", Source: policySource})
	if err != nil {
		log.Error().Err(err).Msg("failed to create policy manager")
		return err
	}

	e := policy.NewExecutor(conn, nil)
	return p.Test(ctx, e, policySource, snapshotDestination, uniqueTempDir)

}

func (c Client) SnapshotPolicy(ctx context.Context, policySource, snapshotDestination string) error {
	policiesToSnapshot, err := FilterPolicies(policySource, c.cfg.Policies)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return err
	}
	log.Debug().Strs("policies", policiesToSnapshot.All()).Msg("policies to snapshot")
	for _, p := range policiesToSnapshot {
		if err := c.snapshotControl(ctx, p, policySource, snapshotDestination); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) DescribePolicies(ctx context.Context, policySource string) error {
	policiesToDescribe, err := FilterPolicies(policySource, []*policy.Policy{})
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return err
	}
	log.Debug().Strs("policies", policiesToDescribe.All()).Msg("policies to describe")
	for _, p := range policiesToDescribe {
		if err := c.describePolicy(ctx, p, policySource); err != nil {
			return err
		}
	}
	return nil
}

// =====================================================================================================================
// 													Module Commands
// =====================================================================================================================

func (c Client) CallModule(ctx context.Context, req ModuleCallRequest) error {
	if _, diags := c.SyncProviders(ctx); diags.HasErrors() {
		return diags
	}

	profiles, err := config.ReadModuleConfigProfiles(req.Name, c.cfg.Modules)
	if err != nil {
		return err
	}
	cfg, err := selectProfile(req.Profile, profiles)
	if err != nil {
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Starting module...\n")

	m := module.NewManager(c.Storage, c.PluginManager)
	m.Register(drift.New())
	out, err := m.Execute(ctx, &module.ExecutionOptions{
		Module:        req.Name,
		ProfileConfig: cfg,
		Params:        req.Params,
		Providers:     c.Providers,
	})
	if err != nil {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to execute module: %s.\n\n", err.Error())
		return err
	} else if out == nil {
		ui.ColorizedOutput(ui.ColorSuccess, "Finished module, no results\n\n")
		return nil
	}

	if out.ErrorMsg != "" {
		ui.ColorizedOutput(ui.ColorError, "Finished module with error: %s\n\n", out.ErrorMsg)
		return out.Error
	}

	if req.OutputPath != "" {
		// Store output in file if requested
		fs := afero.NewOsFs()
		f, err := fs.OpenFile(req.OutputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()

		data, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return err
		}
		if _, err := f.Write(data); err != nil {
			return err
		}

		ui.ColorizedOutput(ui.ColorProgress, "Wrote JSON output to %q\n", req.OutputPath)
	}
	if outString, ok := out.Result.(fmt.Stringer); ok {
		ui.ColorizedOutput(ui.ColorInfo, "Module output: \n%s\n", outString.String())
	} else {
		b, _ := json.MarshalIndent(out.Result, "", "  ")
		ui.ColorizedOutput(ui.ColorInfo, "Module output: \n%s\n", string(b))
	}

	ui.ColorizedOutput(ui.ColorSuccess, "Finished module\n\n")

	if exitCoder, ok := out.Result.(module.ExitCoder); ok {
		return &ExitCodeError{ExitCode: exitCoder.ExitCode()}
	}

	return nil
}

func (c Client) Close() {
	c.PluginManager.Shutdown()
}

func (c Client) checkForUpdate(ctx context.Context) {
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

func (c Client) snapshotControl(ctx context.Context, p *policy.Policy, fullSelector, destination string) error {
	p, err := policy.Load(ctx, c.cfg.CloudQuery.PolicyDirectory, &policy.Policy{Name: p.Name, Source: p.Source})
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return fmt.Errorf("failed to load policies: %w", err)
	}
	if !p.HasChecks() {
		return errors.New("no checks loaded")
	}

	_, subPath := getter.ParseSourceSubPolicy(fullSelector)
	pol := p.Filter(subPath)
	if pol.TotalQueries() != 1 {
		return errors.New("selector must specify only a single control")
	}
	return policy.Snapshot(ctx, c.Storage, &pol, destination, subPath)
}

func (c Client) describePolicy(ctx context.Context, p *policy.Policy, selector string) error {
	p, err := policy.Load(ctx, c.cfg.CloudQuery.PolicyDirectory, &policy.Policy{Name: p.Name, Source: p.Source})
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return fmt.Errorf("failed to load policies: %w", err)
	}
	ui.ColorizedOutput(ui.ColorHeader, "Describe Policy %s output:\n\n", p.String())
	t := &Table{writer: tablewriter.NewWriter(os.Stdout)}
	t.SetHeaders("Path", "Description")

	policyName, subPath := getter.ParseSourceSubPolicy(selector)

	// The `buildDescribePolicyTable` builds the output based on Policy Name and Path
	// In the case of no path, the PolicyName is just the root policy
	if subPath == "" {
		policyName = ""
	}
	pol := p.Filter(subPath)
	buildDescribePolicyTable(t, policy.Policies{&pol}, policyName)
	t.Render()
	ui.ColorizedOutput(ui.ColorInfo, "To execute any policy use the path defined in the table above.\nFor example `cloudquery policy run %s`\n", buildPolicyPath(p.Name, getNestedPolicyExample(p.Policies[0], "")))
	return nil
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
			fetchProgress.Add(fmt.Sprintf("%s_%s", p.Name, p.Alias), fmt.Sprintf("cq-provider-%s@%s-%s", p.Name, "latest", p.Alias), "fetching", int64(len(p.Resources)))
		} else {
			fetchProgress.Add(fmt.Sprintf("%s_%s", p.Name, p.Alias), fmt.Sprintf("cq-provider-%s@%s", p.Name, "latest"), "fetching", int64(len(p.Resources)))
		}
	}
	fetchCallback := func(update core.FetchUpdate) {
		if update.Error != "" {
			fetchProgress.Update(update.Provider, ui.StatusError, fmt.Sprintf("error: %s", update.Error), 0)
			return
		}
		if len(update.PartialFetchResults) > 0 {
			fetchProgress.Update(update.Provider, ui.StatusWarn, fmt.Sprintf("diagnostics: %d", len(update.PartialFetchResults)), 0)
		}
		bar := fetchProgress.GetBar(update.Provider)
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
			fetchProgress.Update(update.Provider, ui.StatusOK, "fetch complete", 0)
			return
		}
	}
	return fetchProgress, fetchCallback
}

func buildPolicyRunProgress(ctx context.Context, policies policy.Policies) (*Progress, policy.UpdateCallback) {
	policyRunProgress := NewProgress(ctx, func(o *ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.CountersNoUnit(" Finished Checks: %d/%d")}
	})

	for _, p := range policies {
		policyRunProgress.Add(p.Name, fmt.Sprintf("policy \"%s\" - ", p.Name), "evaluating - ", 1)
	}

	policyRunCallback := func(update policy.Update) {
		bar := policyRunProgress.GetBar(update.PolicyName)
		// try to get with policy source
		if bar == nil {
			bar = policyRunProgress.GetBar(update.Source)
		}
		if bar == nil {
			policyRunProgress.AbortAll()
			ui.ColorizedOutput(ui.ColorError, "❌ console UI failure, policy run will complete shortly\n")
			return
		}
		if update.Error != "" {
			policyRunProgress.Update(update.PolicyName, ui.StatusError, fmt.Sprintf("error: %s", update.Error), 0)
			return
		}

		// set the total queries to track
		if update.QueriesCount > 0 {
			bar.SetTotal(int64(update.QueriesCount), false)
		}

		if bar == nil {
			policyRunProgress.AbortAll()
			ui.ColorizedOutput(ui.ColorError, "❌ console UI failure, fetch will complete shortly\n")
			return
		}

		bar.b.IncrBy(update.DoneCount() - int(bar.b.Current()))

		if update.AllDone() && bar.Status != ui.StatusWarn {
			policyRunProgress.Update(update.PolicyName, ui.StatusOK, "policy run complete - ", 0)
			bar.Done()
			return
		}
	}

	return policyRunProgress, policyRunCallback
}

func loadConfig(file string) (*config.Config, bool) {
	parser := config.NewParser(
		config.WithEnvironmentVariables(config.EnvVarPrefix, os.Environ()),
		config.WithFileFunc(filepath.Dir(file)),
	)
	cfg, diags := parser.LoadConfigFile(file)
	if diags != nil {
		ui.ColorizedOutput(ui.ColorHeader, "Configuration Error Diagnostics:\n")
		for _, d := range diags {
			c := ui.ColorInfo
			switch d.Severity {
			case hcl.DiagError:
				c = ui.ColorError
			case hcl.DiagWarning:
				c = ui.ColorWarning
			}
			if d.Subject == nil {
				ui.ColorizedOutput(c, "❌ %s; %s\n", d.Summary, d.Detail)
				continue
			}
			ui.ColorizedOutput(c, "❌ %s; %s [%s]\n", d.Summary, d.Detail, d.Subject.String())
		}
		if diags.HasErrors() {
			return nil, false
		}
	}
	return cfg, true
}

func countSeverity(d diag.Diagnostics, sevs ...diag.Severity) string {
	var basicCount uint64
	for _, sev := range sevs {
		basicCount += d.CountBySeverity(sev, false)
	}

	if !viper.GetBool("verbose") {
		return fmt.Sprintf("%d", basicCount)
	}

	var deepCount uint64
	for _, sev := range sevs {
		deepCount += d.CountBySeverity(sev, true)
	}
	if basicCount == deepCount {
		return fmt.Sprintf("%d", basicCount)
	}

	return fmt.Sprintf("%d(%d)", basicCount, deepCount)
}

func selectProfile(profileName string, profiles map[string]hcl.Body) (hcl.Body, error) {
	if profileName == "" && len(profiles) > 1 {
		return nil, fmt.Errorf("multiple profiles detected, choose one with --profile")
	}

	if profileName != "" {
		chosenProfile, ok := profiles[profileName]
		if !ok {
			return nil, fmt.Errorf("specified profile doesn't exist in config")
		}
		return chosenProfile, nil
	}

	for k, v := range profiles {
		ui.ColorizedOutput(ui.ColorDebug, "Using profile %s\n", k)
		return v, nil
	}

	return nil, nil
}

func setConfigAnalytics(cfg *config.Config, history bool) {
	cfgJSON, _ := json.Marshal(cfg)
	s := sha256.New()
	_, _ = s.Write(cfgJSON)
	cfgHash := fmt.Sprintf("%0x", s.Sum(nil))
	analytics.SetGlobalProperty("cfghash", cfgHash)
	analytics.SetGlobalProperty("history", history)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		if analytics.IsCI() {
			scope.SetUser(sentry.User{
				ID: cfgHash,
			})
		}
		scope.SetTags(map[string]string{
			"history": strconv.FormatBool(history),
		})
	})
}

func (c Client) ConvertRequiredToRegistry(providerName string) registry.Provider {
	rp := c.cfg.CloudQuery.Providers.Get(providerName)
	src, name, _ := core.ParseProviderSource(rp)
	return registry.Provider{Name: name, Version: rp.Version, Source: src}
}
