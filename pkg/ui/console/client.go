package console

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/internal/getter"
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"

	"github.com/cloudquery/cq-provider-sdk/cqproto"
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/fatih/color"
	"github.com/getsentry/sentry-go"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/vbauerster/mpb/v6/decor"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client console client is a wrapper around core.Client for console execution of CloudQuery
type Client struct {
	updater       ui.Progress
	cfg           *config.Config
	Registry      registry.Registry
	PluginManager *plugin.Manager
	Storage       database.Storage
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

	cClient := &Client{progressUpdater, cfg, hub, pm, storage}
	cClient.setTelemetryAttributes(trace.SpanFromContext(ctx))
	cClient.checkForUpdate(ctx)
	return cClient, err
}

func (c Client) Close() {
	c.PluginManager.Shutdown()
}

func (c Client) DownloadProviders(ctx context.Context) error {
	ui.ColorizedOutput(ui.ColorProgress, "Initializing CloudQuery Providers...\n\n")
	pp := make([]registry.Provider, len(c.cfg.CloudQuery.Providers))
	for i, rp := range c.cfg.CloudQuery.Providers {
		src, name, err := core.ParseProviderSource(rp)
		if err != nil {
			return err
		}
		pp[i] = registry.Provider{
			Name:    name,
			Version: rp.Version,
			Source:  src,
		}
	}
	_, diags := core.Download(ctx, c.PluginManager, &core.DownloadOptions{Providers: pp, NoVerify: viper.GetBool("no-verify")})
	if diags.HasErrors() {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ failed to initialize provider: %s.\n\n", diags.Error())
		return diags
	}
	if c.updater != nil {
		c.updater.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished provider initialization...\n\n")
	updates, diags := core.CheckAvailableUpdates(ctx, c.Registry, &core.CheckUpdatesOptions{Providers: pp})
	if diags.HasErrors() {
		printDiagnostics("Diagnostics", "", diags, true, false)
	}
	for _, u := range updates {
		ui.ColorizedOutput(ui.ColorInfo, fmt.Sprintf("Update available for provider %s: %s ➡️ %s\n\n", u.Name, u.CurrentVersion, u.AvailableVersion))
	}
	return nil
}

func (c Client) Fetch(ctx context.Context, failOnError bool) error {

	if err := c.UpgradeProviders(ctx, c.cfg.Providers.Names()); err != nil {
		return err
	}

	ui.ColorizedOutput(ui.ColorProgress, "Starting provider fetch...\n\n")
	var fetchProgress ui.Progress
	var fetchCallback core.FetchUpdateCallback

	if ui.DoProgress() {
		fetchProgress, fetchCallback = buildFetchProgress(ctx, c.cfg.Providers)
	}

	providers := make([]core.ProviderInfo, len(c.cfg.Providers))
	for i, p := range c.cfg.Providers {
		providers[i] = core.ProviderInfo{Provider: c.ConvertRequiredToRegistry(p.Name), Config: p}
	}

	response, diags := core.Fetch(ctx, c.Storage, c.PluginManager, &core.FetchOptions{
		UpdateCallback: fetchCallback,
		ProvidersInfo:  providers,
		History:        c.cfg.CloudQuery.History,
	})
	if diags.HasErrors() {
		// Ignore context cancelled error
		if st, ok := status.FromError(diags); !ok || st.Code() != gcodes.Canceled {
			return diags
		}
	}

	if fetchProgress != nil {
		fetchProgress.MarkAllDone()
		fetchProgress.Wait()
	}
	printFetchResponse(response, viper.GetBool("redact-diags"), viper.GetBool("verbose"))

	if response == nil {
		ui.ColorizedOutput(ui.ColorProgress, "Provider fetch canceled.\n\n")
		return nil
	}

	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")
	for _, summary := range response.ProviderFetchSummary {
		s := emojiStatus[ui.StatusOK]
		if summary.Status == "Canceled" {
			s = emojiStatus[ui.StatusError] + " (canceled)"
		}
		key := summary.ProviderName
		if summary.ProviderName != summary.ProviderAlias {
			key = fmt.Sprintf("%s(%s)", summary.ProviderName, summary.ProviderAlias)
		}
		diags := summary.Diagnostics().Squash()
		ui.ColorizedOutput(ui.ColorHeader, "Provider %s fetch summary: %s Total Resources fetched: %d\t ⚠️ Warnings: %s\t ❌ Errors: %s\n",
			key,
			s,
			summary.TotalResourcesFetched,
			countSeverity(diags, diag.WARNING),
			countSeverity(diags, diag.ERROR),
		)
	}
	if failOnError && response.HasErrors() {
		return fmt.Errorf("provider fetch has one or more errors")
	}

	return nil
}

func (c Client) DownloadPolicy(ctx context.Context, args []string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Downloading CloudQuery Policy...\n")
	p, err := policy.Load(ctx, c.cfg.CloudQuery.PolicyDirectory, &policy.Policy{Name: "policy", Source: args[0]})
	if err != nil {
		ui.SleepBeforeError(ctx)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to Download policy: %s.\n\n", err.Error())
		return err
	}
	if c.updater != nil {
		c.updater.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished downloading policy...\n")
	// Show policy instructions
	ui.ColorizedOutput(ui.ColorHeader, fmt.Sprintf(`
Add this block into your CloudQuery config file:

policy "%s" {
    source = "%s"
}

`, p.Name, p.Source))
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

	var policyRunProgress ui.Progress
	var policyRunCallback policy.UpdateCallback

	// if we are running in a terminal, build the progress bar
	if ui.DoProgress() {
		policyRunProgress, policyRunCallback = buildPolicyRunProgress(ctx, policiesToRun)
	}
	// Policies run request
	req := &policy.RunRequest{
		Policies:    policiesToRun,
		Directory:   c.cfg.CloudQuery.PolicyDirectory,
		OutputDir:   outputDir,
		RunCallback: policyRunCallback,
	}
	results, err := policy.Run(ctx, c.Storage, req)

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

func (c Client) CallModule(ctx context.Context, req ModuleCallRequest) error {

	// TODO: move this in module package
	provs, err := c.getModuleProviders(ctx)
	if err != nil {
		return err
	}

	profiles, err := config.ReadModuleConfigProfiles(req.Name, c.cfg.Modules)
	if err != nil {
		return err
	}

	cfg, err := c.selectProfile(req.Profile, profiles)
	if err != nil {
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Starting module...\n")

	m := module.NewManager(c.Storage, c.PluginManager)
	out, err := m.Execute(ctx, &module.ExecuteRequest{
		Module:        req.Name,
		ProfileConfig: cfg,
		Params:        req.Params,
		Providers:     provs,
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

	type stringer interface {
		String() string
	}

	if outString, ok := out.Result.(stringer); ok {
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

func (c Client) UpgradeProviders(ctx context.Context, args []string) error {
	providers, err := c.getRequiredProviders(args)
	if err != nil {
		return err
	}
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Upgrading CloudQuery providers %s\n\n", strings.Join(providers.Names(), ", "))
	for _, p := range providers {
		if sync, diags := core.Sync(ctx, c.Storage, c.PluginManager, &core.SyncOptions{Provider: c.ConvertRequiredToRegistry(p.Name), DownloadLatest: false}); diags.HasErrors() && sync.State != core.NoChange {
			if errors.Is(diags, core.ErrMigrationsNotSupported) {
				ui.ColorizedOutput(ui.ColorWarning, "%s Failed to upgrade provider %s: %s.\n", emojiStatus[ui.StatusWarn], p.String(), err.Error())
				continue
			}
			ui.ColorizedOutput(ui.ColorError, "%s Failed to upgrade provider %s. Error: %s.\n", emojiStatus[ui.StatusError], p.String(), err.Error())
			return err
		}

		ui.ColorizedOutput(ui.ColorSuccess, "%s Upgraded provider %s to %s successfully.\n", emojiStatus[ui.StatusOK], p.Name, p.Version)
	}
	ui.ColorizedOutput(ui.ColorProgress, "\nFinished upgrading providers...\n\n")
	return nil
}

func (c Client) DowngradeProviders(ctx context.Context, args []string) error {
	providers, err := c.getRequiredProviders(args)
	if err != nil {
		return err
	}

	// download requested versions
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}

	provVersions := make(map[string]string, len(providers))
	for _, p := range providers {
		provVersions[p.Name] = p.Version
		p.Version = migrator.Latest
	}

	ui.ColorizedOutput(ui.ColorProgress, "Downgrading CloudQuery providers %s\n\n", strings.Join(providers.Names(), ", "))

	for _, p := range providers {
		ui.ColorizedOutput(ui.ColorSuccess, "%s Downgrading provider %s to %s...\n", emojiStatus[ui.StatusInProgress], p.Name, provVersions[p.Name])

		if _, diags := core.Sync(ctx, c.Storage, c.PluginManager, &core.SyncOptions{Provider: c.ConvertRequiredToRegistry(p.Name), DownloadLatest: true}); diags.HasErrors() {
			if errors.Is(err, core.ErrMigrationsNotSupported) {
				ui.ColorizedOutput(ui.ColorWarning, "%s Failed to downgrade provider %s: %s.\n", emojiStatus[ui.StatusWarn], p.Name, err.Error())
				continue
			} else {
				ui.ColorizedOutput(ui.ColorError, "%s Failed to downgrade provider %s. Error: %s.\n", emojiStatus[ui.StatusError], p.Name, err.Error())
				return err
			}
		}

		ui.ColorizedOutput(ui.ColorSuccess, "%s Downgraded provider %s to %s successfully.\n", emojiStatus[ui.StatusOK], p.Name, provVersions[p.Name])
	}
	ui.ColorizedOutput(ui.ColorProgress, "\nFinished downgrading providers...\n\n")
	return nil
}

func (c Client) DropProvider(ctx context.Context, providerName string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Dropping CloudQuery provider %s schema...\n\n", providerName)
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}

	if err := core.Drop(ctx, c.Storage, c.PluginManager, c.ConvertRequiredToRegistry(providerName)); err != nil {
		ui.ColorizedOutput(ui.ColorError, "%s Failed to drop provider %s schema. Error: %s.\n\n", emojiStatus[ui.StatusError], providerName, err.Error())
		return err
	} else {
		ui.ColorizedOutput(ui.ColorSuccess, "%s provider %s schema dropped successfully.\n\n", emojiStatus[ui.StatusOK], providerName)
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished downgrading providers...\n\n")
	return nil
}

func (c Client) BuildProviderTables(ctx context.Context, providerName string) error {
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}

	if err := c.buildProviderTables(ctx, providerName); err != nil {
		return err
	}

	return nil
}

func (c Client) BuildAllProviderTables(ctx context.Context) error {
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}

	for _, p := range c.cfg.Providers {
		if err := c.buildProviderTables(ctx, p.Name); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) RemoveStaleData(ctx context.Context, lastUpdate time.Duration, dryRun bool, providers []string) error {
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	pp := make([]registry.Provider, len(providers))
	for i, p := range providers {
		pp[i] = c.ConvertRequiredToRegistry(p)
	}
	ui.ColorizedOutput(ui.ColorHeader, "Purging providers %s resources..\n\n", providers)
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

	if len(diags) > 0 {
		printDiagnostics("Purge", "", diags, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
		return diags
	} else {
		ui.ColorizedOutput(ui.ColorProgress, "Purge for providers %s was successful\n\n", providers)
	}
	return nil
}

func (c Client) buildProviderTables(ctx context.Context, providerName string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Building CloudQuery provider %s schema...\n\n", providerName)
	if err := c.UpgradeProviders(ctx, []string{providerName}); err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to build provider %s schema. Error: %s.\n\n", providerName, err.Error())
		return err
	} else {
		ui.ColorizedOutput(ui.ColorSuccess, "✓ provider %s schema built successfully.\n\n", providerName)
		color.GreenString("✓")
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished building provider schema...\n\n")
	return nil
}

func (c Client) selectProfile(profileName string, profiles map[string]hcl.Body) (hcl.Body, error) {
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

func (c Client) getRequiredProviders(providerNames []string) (config.RequiredProviders, error) {
	if len(providerNames) == 0 {
		// if no providers are given we will return all providers
		return c.cfg.CloudQuery.Providers.Distinct(), nil
	}

	providers := make(config.RequiredProviders, len(providerNames))
	for i, p := range providerNames {
		pCfg, err := c.cfg.CloudQuery.GetRequiredProvider(p)
		if err != nil {
			return nil, err
		}
		providers[i] = pCfg
	}
	return providers, nil
}

func (c Client) getModuleProviders(ctx context.Context) ([]*cqproto.GetProviderSchemaResponse, error) {
	if err := c.DownloadProviders(ctx); err != nil {
		return nil, err
	}
	list := make([]*cqproto.GetProviderSchemaResponse, 0, len(c.cfg.Providers))
	dupes := make(map[string]struct{})
	for _, p := range c.cfg.Providers {
		if _, ok := dupes[p.Name]; ok {
			continue
		}
		dupes[p.Name] = struct{}{}

		s, err := core.GetProviderSchema(ctx, c.PluginManager, &core.GetProviderSchemaOptions{Provider: c.ConvertRequiredToRegistry(p.Name)})
		if err != nil {
			return nil, err
		}
		if s.Version == "" { // FIXME why?
			deets, err := c.PluginManager.GetPluginDetails(p.Name)
			if err != nil {
				log.Warn().Err(err).Msg("GetPluginDetails failed")
			} else {
				s.Version = deets.Version
			}
		}
		list = append(list, s.GetProviderSchemaResponse)
	}

	return list, nil
}

func (c Client) checkForUpdate(ctx context.Context) {
	v, err := core.MaybeCheckForUpdate(ctx, afero.Afero{Fs: afero.NewOsFs()}, time.Now().Unix(), core.UpdateCheckPeriod)
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

func (c Client) setTelemetryAttributes(span trace.Span) {
	cfgJSON, _ := json.Marshal(c.cfg)
	s := sha1.New()
	_, _ = s.Write(cfgJSON)
	cfgHash := fmt.Sprintf("%0x", s.Sum(nil))
	attrs := []attribute.KeyValue{
		attribute.String("cfghash", cfgHash),
	}
	if c.cfg.CloudQuery.History != nil {
		attrs = append(attrs, attribute.Bool("history_enabled", true))
	}
	span.SetAttributes(attrs...)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		if telemetry.IsCI() {
			scope.SetUser(sentry.User{
				ID: cfgHash,
			})
		}
		if c.cfg.CloudQuery.History != nil {
			scope.SetTags(map[string]string{
				"history_enabled": strconv.FormatBool(true),
			})
		}
	})
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

func countSeverity(d diag.Diagnostics, sev diag.Severity) string {
	basicCount := d.CountBySeverity(sev, false)

	if !viper.GetBool("verbose") {
		return fmt.Sprintf("%d", basicCount)
	}

	deepCount := d.CountBySeverity(sev, true)
	if basicCount == deepCount {
		return fmt.Sprintf("%d", basicCount)
	}

	return fmt.Sprintf("%d(%d)", basicCount, deepCount)
}

func (c Client) ConvertRequiredToRegistry(providerName string) registry.Provider {
	rp := c.cfg.CloudQuery.Providers.Get(providerName)
	src, name, _ := core.ParseProviderSource(rp)
	return registry.Provider{Name: name, Version: rp.Version, Source: src}
}
