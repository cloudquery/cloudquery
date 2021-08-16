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
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize client.\n\n")
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

func (c Client) Fetch(ctx context.Context) error {
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
	if err := c.c.Fetch(ctx, request); err != nil {
		return err
	}
	if ui.IsTerminal() && fetchProgress != nil {
		fetchProgress.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")
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

func (c Client) Client() *client.Client {
	return c.c
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
		bar := fetchProgress.GetBar(update.Provider)
		bar.b.IncrBy(update.DoneCount() - int(bar.b.Current()))

		if bar.Status == ui.StatusError {
			if update.AllDone() {
				bar.SetTotal(0, true)
			}
			return
		}
		if update.AllDone() {
			fetchProgress.Update(update.Provider, ui.StatusOK, "fetch complete", 0)
			return
		}
	}
	return fetchProgress, fetchCallback
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
