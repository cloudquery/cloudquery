package console

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/vbauerster/mpb/v6/decor"
)

// Client console client is a wrapper around client.Client for console execution of CloudQuery
type Client struct {
	c       *client.Client
	cfg     *config.Config
	updater *Progress
}

func CreateClient(ctx context.Context, configPath string, opts ...client.Option) (*Client, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
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
	})
	c, err := client.New(ctx, opts...)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize client. Error: %s\n\n", err)
		return nil, err
	}
	return &Client{c, cfg, progressUpdater}, err
}

func (c Client) DownloadProviders(ctx context.Context) error {
	ui.ColorizedOutput(ui.ColorProgress, "Initializing CloudQuery Providers...\n\n")
	err := c.c.DownloadProviders(ctx)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize provider: %s.\n\n", err.Error())
		return err
	}
	// sleep some extra 300 milliseconds for progress refresh
	if ui.IsTerminal() {
		time.Sleep(300 * time.Millisecond)
		c.updater.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished provider initialization...\n\n")
	return nil
}

func (c Client) Fetch(ctx context.Context, failOnError bool) error {
	if err := c.DownloadProviders(ctx); err != nil {
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
		return err
	}

	if ui.IsTerminal() && fetchProgress != nil {
		fetchProgress.MarkAllDone()
		fetchProgress.Wait()
		printFetchResponse(response)
	}

	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")

	if failOnError {
		for _, summary := range response.ProviderFetchSummary {
			if len(summary.FetchErrors) > 0 {
				return fmt.Errorf("provider fetch has one or more errors")
			}
		}
	}
	return nil
}

func (c Client) DownloadPolicy(ctx context.Context, args []string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Downloading CloudQuery Policy...\n\n")
	err := c.c.DownloadPolicy(ctx, args)
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
	ui.ColorizedOutput(ui.ColorProgress, "Finished downloading policy...\n\n")
	return nil
}

func (c Client) RunPolicy(ctx context.Context, args []string, subPath, outputPath string, stopOnFailure bool, skipVersioning bool) error {
	ui.ColorizedOutput(ui.ColorProgress, "Starting policy run...\n")
	req := client.PolicyRunRequest{
		Args:           args,
		SubPath:        subPath,
		OutputPath:     outputPath,
		StopOnFailure:  stopOnFailure,
		SkipVersioning: skipVersioning,
		RunCallBack: func(name string, qtype config.QueryType, passed bool) {
			switch {
			case passed:
				ui.ColorizedOutput(ui.ColorInfo, "\t%s  %-140s %5s\n", emojiStatus[ui.StatusOK], name, color.GreenString("passed"))
			case qtype == config.ManualQuery:
				ui.ColorizedOutput(ui.ColorInfo, "\t%s  %-140s %5s\n", emojiStatus[ui.StatusWarn], name, color.YellowString("manual"))
			default:
				ui.ColorizedOutput(ui.ColorInfo, "\t%s %-140s %5s\n", emojiStatus[ui.StatusError], name, color.RedString("failed"))
			}
		},
	}
	err := c.c.RunPolicy(ctx, req)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to run policy: %s.\n\n", err.Error())
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished policy run...\n\n")
	return nil
}

func (c Client) UpgradeProviders(ctx context.Context, args []string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Upgrading CloudQuery providers %s\n\n", args)
	providers, err := c.getRequiredProviders(args)
	if err != nil {
		return err
	}
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	for _, p := range providers {

		if err := c.c.UpgradeProvider(ctx, p.Name); err != nil {
			ui.ColorizedOutput(ui.ColorError, "❌ Failed to upgrade provider %s. Error: %s.\n\n", p.String(), err.Error())
			return err
		} else {
			ui.ColorizedOutput(ui.ColorSuccess, "✓ Upgraded provider %s to %s successfully.\n\n", p.Name, p.Version)
			color.GreenString("✓")
		}
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished upgrading providers...\n\n")
	return nil
}

func (c Client) DowngradeProviders(ctx context.Context, args []string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Downgrading CloudQuery providers %s\n\n", args)
	providers, err := c.getRequiredProviders(args)
	if err != nil {
		return err
	}
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	for _, p := range providers {
		if err := c.c.DowngradeProvider(ctx, p.Name); err != nil {
			ui.ColorizedOutput(ui.ColorError, "❌ Failed to downgrade provider %s. Error: %s.\n\n", p.String(), err.Error())
			return err
		} else {
			ui.ColorizedOutput(ui.ColorSuccess, "✓ Downgraded provider %s to %s successfully.\n\n", p.Name, p.Version)
			color.GreenString("✓")
		}
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished downgrading providers...\n\n")
	return nil
}

func (c Client) DropProvider(ctx context.Context, providerName string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Dropping CloudQuery provider %s schema...\n\n", providerName)
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	if err := c.c.DropProvider(ctx, providerName); err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to drop provider %s schema. Error: %s.\n\n", providerName, err.Error())
		return err
	} else {
		ui.ColorizedOutput(ui.ColorSuccess, "✓ provider %s schema dropped successfully.\n\n", providerName)
		color.GreenString("✓")
	}
	ui.ColorizedOutput(ui.ColorProgress, "Finished downgrading providers...\n\n")
	return nil
}

func (c Client) Client() *client.Client {
	return c.c
}

func (c Client) getRequiredProviders(providerNames []string) ([]*config.RequiredProvider, error) {
	if len(providerNames) == 0 {
		// if no providers are given we will return all providers
		return c.cfg.CloudQuery.Providers, nil
	}
	providers := make([]*config.RequiredProvider, len(providerNames))
	for i, p := range providerNames {
		pCfg, err := c.cfg.CloudQuery.GetRequiredProvider(p)
		if err != nil {
			return nil, err
		}
		providers[i] = pCfg
	}
	return providers, nil
}

func buildFetchProgress(ctx context.Context, providers []*config.Provider) (*Progress, client.FetchUpdateCallback) {
	fetchProgress := NewProgress(ctx, func(o *ProgressOptions) {
		o.AppendDecorators = []decor.Decorator{decor.CountersNoUnit(" Finished Resources: %d/%d")}
	})

	for _, p := range providers {
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
			fetchProgress.Update(update.Provider, ui.StatusWarn, fmt.Sprintf("errors: %d", len(update.PartialFetchResults)), 0)
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

func printFetchResponse(summary *client.FetchResponse) {
	if summary == nil {
		return
	}
	for _, pfs := range summary.ProviderFetchSummary {
		ui.ColorizedOutput(ui.ColorHeader, "Partial Fetch Errors for Provider %s:\n\n", pfs.ProviderName)
		for _, r := range pfs.PartialFetchErrors {
			if r.RootTableName != "" {
				ui.ColorizedOutput(ui.ColorErrorBold,
					"Parent-Resource: %-64s Parent-Primary-Keys: %v, Table: %s, Error: %s\n",
					r.RootTableName,
					r.RootPrimaryKeyValues,
					r.TableName,
					r.Error)
			} else {
				ui.ColorizedOutput(ui.ColorErrorBold,
					"Table: %-64s Error: %s\n",
					r.TableName,
					r.Error)
			}
		}
		ui.ColorizedOutput(ui.ColorWarning, "\n")
	}
}

func loadConfig(path string) (*config.Config, error) {
	parser := config.NewParser(nil)
	cfg, diags := parser.LoadConfigFile(path)
	if diags != nil {
		ui.ColorizedOutput(ui.ColorHeader, "Configuration Error Diagnostics:\n")
		for _, d := range diags {
			ui.ColorizedOutput(ui.ColorError, "❌ %s; %s\n", d.Summary, d.Detail)
		}
		return nil, fmt.Errorf("bad configuration")
	}
	return cfg, nil
}
