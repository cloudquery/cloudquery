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
	sdkdb "github.com/cloudquery/cq-provider-sdk/database"
	"github.com/cloudquery/cq-provider-sdk/migration/migrator"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/fatih/color"
	"github.com/getsentry/sentry-go"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/vbauerster/mpb/v6/decor"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/module"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
)

// Client console client is a wrapper around client.Client for console execution of CloudQuery
type Client struct {
	c       *client.Client
	cfg     *config.Config
	updater *Progress
}

func CreateClient(ctx context.Context, configPath string, configMutator func(*config.Config) error, opts ...client.Option) (*Client, error) {
	cfg, ok := loadConfig(configPath)
	if !ok {
		// No explicit error string needed, user information is in diags
		return nil, &ExitCodeError{ExitCode: 1}
	}
	if configMutator != nil {
		if err := configMutator(cfg); err != nil {
			return nil, err
		}
	}

	return CreateClientFromConfig(ctx, cfg, opts...)
}

func CreateClientFromConfig(ctx context.Context, cfg *config.Config, opts ...client.Option) (*Client, error) {
	progressUpdater := NewProgress(ctx, func(o *ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.Percentage()}
	})
	opts = append(opts, func(c *client.Client) {
		if ui.IsTerminal() {
			c.HubProgressUpdater = progressUpdater
		}
		c.Providers = cfg.CloudQuery.Providers
		c.NoVerify = viper.GetBool("no-verify")
		c.PluginDirectory = cfg.CloudQuery.PluginDirectory
		c.PolicyDirectory = cfg.CloudQuery.PolicyDirectory
		c.DSN = cfg.CloudQuery.Connection.DSN
		c.SkipBuildTables = viper.GetBool("skip-build-tables")
		c.HistoryCfg = cfg.CloudQuery.History
	})
	c, err := client.New(ctx, opts...)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize client. Error: %s\n\n", err)
		return nil, err
	}
	cClient := &Client{c, cfg, progressUpdater}
	cClient.setTelemetryAttributes(trace.SpanFromContext(ctx))
	cClient.checkForUpdate(ctx)
	return cClient, err
}

func CreateNullClient(ctx context.Context, opts ...client.Option) (*Client, error) {
	progressUpdater := NewProgress(ctx, func(o *ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.Percentage()}
	})
	opts = append(opts, func(c *client.Client) {
		if ui.IsTerminal() {
			c.HubProgressUpdater = progressUpdater
			c.PluginDirectory = "./.cq/providers"
			c.PolicyDirectory = "./.cq/policies"
		}
	})
	c, err := client.New(ctx, opts...)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize client. Error: %s\n\n", err)
		return nil, err
	}
	cClient := &Client{c, nil, progressUpdater}
	return cClient, err
}

func ClientFactory(ctx context.Context, configPath *string, configMutator func(*config.Config) error, opts ...client.Option) (*Client, error) {
	if configPath == nil {
		return CreateNullClient(ctx)
	}
	if _, err := os.Stat(*configPath); errors.Is(err, os.ErrNotExist) {
		return CreateNullClient(ctx)
	}
	return CreateClient(ctx, *configPath, configMutator, opts...)
}

func (c Client) DownloadProviders(ctx context.Context) error {
	ui.ColorizedOutput(ui.ColorProgress, "Initializing CloudQuery Providers...\n\n")
	err := c.c.DownloadProviders(ctx)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize provider: %s.\n\n", err.Error())
		return err
	}
	// sleep some extra 500 milliseconds for progress refresh
	if ui.IsTerminal() {
		time.Sleep(500 * time.Millisecond)
		c.updater.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished provider initialization...\n\n")
	updates := c.c.CheckForProviderUpdates(ctx)
	for _, u := range updates {
		ui.ColorizedOutput(ui.ColorInfo, fmt.Sprintf("Update available for provider %s: %s ➡️ %s\n\n", u.Name, u.Version, u.LatestVersion))
	}
	return nil
}

func (c Client) Fetch(ctx context.Context, failOnError bool) error {
	if viper.GetBool("skip-schema-upgrade") {
		// only download providers and verify, no upgrade
		if err := c.DownloadProviders(ctx); err != nil {
			return err
		}
	} else if err := c.UpgradeProviders(ctx, c.cfg.Providers.Names()); err != nil {
		return err
	}

	if err := c.c.NormalizeResources(ctx, c.cfg.Providers); err != nil {
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Starting provider fetch...\n\n")
	var fetchProgress *Progress
	var fetchCallback client.FetchUpdateCallback

	if ui.IsTerminal() {
		fetchProgress, fetchCallback = buildFetchProgress(ctx, c.cfg.Providers)
	}
	request := client.FetchRequest{
		Providers:      c.cfg.Providers,
		UpdateCallback: fetchCallback,
	}
	response, err := c.c.Fetch(ctx, request)
	if err != nil {
		// Ignore context cancelled error

		if st, ok := status.FromError(err); !ok || st.Code() != gcodes.Canceled {
			return err
		}
	}

	if ui.IsTerminal() && fetchProgress != nil {
		fetchProgress.MarkAllDone()
		fetchProgress.Wait()
		printFetchResponse(response, viper.GetBool("redact-diags"), viper.GetBool("verbose"))
	}

	if response == nil {
		ui.ColorizedOutput(ui.ColorProgress, "Provider fetch canceled.\n\n")
		return nil
	}

	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")
	for _, summary := range response.ProviderFetchSummary {
		status := emojiStatus[ui.StatusOK]
		if summary.Status == "Canceled" {
			status = emojiStatus[ui.StatusError] + " (canceled)"
		}
		key := summary.ProviderName
		if summary.ProviderName != summary.ProviderAlias {
			key = fmt.Sprintf("%s(%s)", summary.ProviderName, summary.ProviderAlias)
		}
		diags := summary.Diagnostics().Squash()
		ui.ColorizedOutput(ui.ColorHeader, "Provider %s fetch summary: %s Total Resources fetched: %d\t ⚠️ Warnings: %s\t ❌ Errors: %s\n",
			key,
			status,
			summary.TotalResourcesFetched,
			countSeverity(diags, diag.WARNING),
			countSeverity(diags, diag.ERROR),
		)
		if failOnError && summary.HasErrors() {
			err = fmt.Errorf("provider fetch has one or more errors")
		}
	}
	return err
}

func (c Client) DownloadPolicy(ctx context.Context, args []string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Downloading CloudQuery Policy...\n")
	p, err := c.c.LoadPolicy(ctx, "policy", args[0])
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to Download policy: %s.\n\n", err.Error())
		return err
	}
	// sleep some extra 300 milliseconds for progress refresh
	if ui.IsTerminal() {
		time.Sleep(300 * time.Millisecond)
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
	c.c.Logger.Debug("run policy received params:", "policy", policySource, "outputDir", outputDir, "noResults", noResults)
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	policiesToRun, err := FilterPolicies(policySource, c.cfg.Policies)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return err
	}
	c.c.Logger.Debug("policies to run", "policies", policiesToRun)

	ui.ColorizedOutput(ui.ColorProgress, "Starting policies run...\n\n")

	var policyRunProgress *Progress
	var policyRunCallback policy.UpdateCallback

	// if we are running in a terminal, build the progress bar
	if ui.IsTerminal() {
		policyRunProgress, policyRunCallback = buildPolicyRunProgress(ctx, policiesToRun)
	}
	// Policies run request
	req := &client.PoliciesRunRequest{
		Policies:    policiesToRun,
		OutputDir:   outputDir,
		RunCallback: policyRunCallback,
	}
	results, err := c.c.RunPolicies(ctx, req)

	if ui.IsTerminal() && policyRunProgress != nil {
		policyRunProgress.MarkAllDone()
		// sleep some extra 500 milliseconds for progress refresh
		time.Sleep(500 * time.Millisecond)
		policyRunProgress.Wait()
		if !noResults {
			printPolicyResponse(results)
		}
	}

	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to run policies: %s.\n\n", err.Error())
		return err
	}

	ui.ColorizedOutput(ui.ColorProgress, "Finished policies run...\n\n")
	return nil
}

func (c Client) TestPolicies(ctx context.Context, policySource, snapshotDestination string) error {
	conn, err := sdkdb.New(ctx, hclog.NewNullLogger(), c.c.DSN)
	if err != nil {
		c.c.Logger.Error("failed to connect to new database", "err", err)
		return err
	}
	uniqueTempDir, err := os.MkdirTemp(os.TempDir(), "*-myOptionalSuffix")
	if err != nil {
		return err
	}

	policyManager := policy.NewManager(uniqueTempDir, conn, c.c.Logger)
	p, err := policyManager.Load(ctx, &policy.Policy{Name: "test-policy", Source: policySource})
	if err != nil {
		c.c.Logger.Error("failed to create policy manager", "err", err)
		return err
	}

	e := policy.NewExecutor(conn, c.c.Logger, nil)
	return p.Test(ctx, e, policySource, snapshotDestination, uniqueTempDir)

}

func (c Client) SnapshotPolicy(ctx context.Context, policySource, snapshotDestination string) error {
	policiesToSnapshot, err := FilterPolicies(policySource, c.cfg.Policies)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, err.Error())
		return err
	}
	c.c.Logger.Debug("policies to snapshot", "policies", policiesToSnapshot.All())
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
	c.c.Logger.Debug("policies to describe", "policies", policiesToDescribe.All())
	for _, p := range policiesToDescribe {
		if err := c.describePolicy(ctx, p, policySource); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) CallModule(ctx context.Context, req ModuleCallRequest) error {
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

	runReq := client.ModuleRunRequest{
		Name:      req.Name,
		Params:    req.Params,
		Providers: provs,
		Config:    cfg,
	}
	out, err := c.c.ExecuteModule(ctx, runReq)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
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
		if err := c.c.UpgradeProvider(ctx, p.Name); err != nil && err != migrate.ErrNoChange {
			if errors.Is(err, client.ErrMigrationsNotSupported) {
				ui.ColorizedOutput(ui.ColorWarning, "%s Failed to upgrade provider %s: %s.\n", emojiStatus[ui.StatusWarn], p.String(), err.Error())
				continue
			} else {
				ui.ColorizedOutput(ui.ColorError, "%s Failed to upgrade provider %s. Error: %s.\n", emojiStatus[ui.StatusError], p.String(), err.Error())
				return err
			}
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

	newCfg := *c.cfg
	newCfg.CloudQuery.Providers = providers
	conWithLatest, err := CreateClientFromConfig(ctx, &newCfg)
	if err != nil {
		return err
	}
	cqWithLatest := conWithLatest.Client()
	defer cqWithLatest.Close()

	ui.ColorizedOutput(ui.ColorProgress, "Will fetch latest providers first, before downgrade... %s\n\n", strings.Join(providers.Names(), ", "))

	if err := cqWithLatest.DownloadProviders(ctx); err != nil {
		return err
	}

	ui.ColorizedOutput(ui.ColorProgress, "Downgrading CloudQuery providers %s\n\n", strings.Join(providers.Names(), ", "))

	for _, p := range providers {
		ui.ColorizedOutput(ui.ColorSuccess, "%s Downgrading provider %s to %s...\n", emojiStatus[ui.StatusInProgress], p.Name, provVersions[p.Name])

		if err := cqWithLatest.DowngradeProvider(ctx, p.Name, provVersions[p.Name]); err != nil && err != migrate.ErrNoChange {
			if errors.Is(err, client.ErrMigrationsNotSupported) {
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
	if err := c.c.DropProvider(ctx, providerName); err != nil {
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

	for _, p := range c.c.Providers {
		if err := c.buildProviderTables(ctx, p.Name); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) buildProviderTables(ctx context.Context, providerName string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Building CloudQuery provider %s schema...\n\n", providerName)
	if err := c.c.BuildProviderTables(ctx, providerName); err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to build provider %s schema. Error: %s.\n\n", providerName, err.Error())
		return err
	} else {
		ui.ColorizedOutput(ui.ColorSuccess, "✓ provider %s schema built successfully.\n\n", providerName)
		color.GreenString("✓")
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished building provider schema...\n\n")
	return nil
}

func (c Client) Client() *client.Client {
	return c.c
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

		s, err := c.c.GetProviderSchema(ctx, p.Name)
		if err != nil {
			return nil, err
		}
		if s.Version == "" { // FIXME why?
			deets, err := c.c.Manager.GetPluginDetails(p.Name)
			if err != nil {
				c.c.Logger.Warn("GetPluginDetails failed", "error", err.Error())
			} else {
				s.Version = deets.Version
			}
		}
		list = append(list, s.GetProviderSchemaResponse)
	}

	return list, nil
}

func (c Client) checkForUpdate(ctx context.Context) {
	v, err := client.MaybeCheckForUpdate(ctx, afero.Afero{Fs: afero.NewOsFs()}, time.Now().Unix(), client.UpdateCheckPeriod)
	if err != nil {
		c.c.Logger.Warn("update check failed", "error", err)
		return
	}
	if v != nil {
		ui.ColorizedOutput(ui.ColorInfo, "An update to CloudQuery core is available: %s!\n\n", v)
		c.c.Logger.Debug("update check succeeded", "new_version", v.String())
	} else {
		c.c.Logger.Debug("update check succeeded, no new version")
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
	if c.c.HistoryCfg != nil {
		attrs = append(attrs, attribute.Bool("history_enabled", true))
	}
	span.SetAttributes(attrs...)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		if telemetry.IsCI() {
			scope.SetUser(sentry.User{
				ID: cfgHash,
			})
		}
		if c.c.HistoryCfg != nil {
			scope.SetTags(map[string]string{
				"history_enabled": strconv.FormatBool(true),
			})
		}
	})
}

func (c Client) snapshotControl(ctx context.Context, p *policy.Policy, fullSelector, destination string) error {
	p, err := c.c.LoadPolicy(ctx, p.Name, p.Source)
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
	return c.c.PolicyManager.Snapshot(ctx, &pol, destination, subPath)
}

func (c Client) describePolicy(ctx context.Context, p *policy.Policy, selector string) error {
	p, err := c.c.LoadPolicy(ctx, p.Name, p.Source)
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

func buildFetchProgress(ctx context.Context, providers []*config.Provider) (*Progress, client.FetchUpdateCallback) {
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
	fetchCallback := func(update client.FetchUpdate) {
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
			if d.Subject == nil {
				ui.ColorizedOutput(ui.ColorError, "❌ %s; %s\n", d.Summary, d.Detail)
				continue
			}

			ui.ColorizedOutput(ui.ColorError, "❌ %s; %s [%s]\n", d.Summary, d.Detail, d.Subject.String())
		}
		return nil, false
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
