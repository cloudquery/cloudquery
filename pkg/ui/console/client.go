package console

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/fatih/color"
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
		c.Hub = registry.NewRegistryHub(registry.CloudQueryRegistryURl, func(h *registry.Hub) {
			if ui.IsTerminal() {
				h.ProgressUpdater = progressUpdater
			}
			h.NoVerify = viper.GetBool("no-verify")
			h.PluginDirectory = cfg.CloudQuery.PluginDirectory
		})
	})
	c, err := client.New(cfg, opts...)
	if err != nil {
		ui.ColorizedOutput(ui.ColorError, "❌ Failed to initialize client.\n\n")
		return nil, err
	}
	return &Client{c, cfg, progressUpdater}, err
}

func (c Client) DownloadProviders(ctx context.Context) error {
	ui.ColorizedOutput(ui.ColorProgress, "Initializing CloudQuery Providers...\n\n")
	err := c.c.Initialize(ctx)
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
	request := client.FetchRequest{Providers: c.cfg.Providers, UpdateCallback: fetchCallback}
	if err := c.c.Fetch(ctx, request); err != nil {
		return err
	}
	if ui.IsTerminal() && fetchProgress != nil {
		fetchProgress.Wait()
	}
	ui.ColorizedOutput(ui.ColorProgress, "Provider fetch complete.\n\n")
	return nil
}

func (c Client) ExecutePolicy(ctx context.Context, policyPath string, output string) error {
	ui.ColorizedOutput(ui.ColorProgress, "Executing Policy %s...\n", policyPath)
	_, err := c.c.ExecutePolicy(ctx, client.ExecutePolicyRequest{OutputPath: output, PolicyPath: policyPath, UpdateCallback: func(name string, passed bool, resultCount int) {
		if passed {
			ui.ColorizedOutput(ui.ColorInfo, "\t%s  %-120s %5s\n", emojiStatus[ui.StatusOK], name, color.GreenString("passed"))
		} else {
			ui.ColorizedOutput(ui.ColorInfo, "\t%s %-120s %5s\n", emojiStatus[ui.StatusError], name, color.RedString("failed"))
		}
	}})
	if err != nil {
		return err
	}
	ui.ColorizedOutput(ui.ColorProgress, "Policy Executed successfully\n")
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
		fetchProgress.Add(p.Name, fmt.Sprintf("cq-provider-%s@%s", p.Name, "latest"), "fetching", int64(len(p.Resources)))
	}
	fetchCallback := func(update client.FetchUpdate) {
		if update.Error != "" {
			fetchProgress.Update(update.Provider, ui.StatusError, fmt.Sprintf("error: %s", update.Error), 0)
			return
		}
		fetchProgress.Increment(update.Provider, 1)
		bar := fetchProgress.GetBar(update.Provider)

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
			ui.ColorizedOutput(ui.ColorError, "❌ %s\n", d.Error())
		}
		return nil, fmt.Errorf("bad configuration")
	}
	return cfg, nil
}
