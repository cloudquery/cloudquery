package console

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
	"time"
)

type Client struct {
	c   *client.Client
	cfg *config.Config
}

func CreateClient(ctx context.Context, configPath string, opts ...client.Option) (*Client, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}
	return CreateClientFromConfig(ctx, cfg, opts...)
}

func CreateClientFromConfig(ctx context.Context, cfg *config.Config, opts ...client.Option) (*Client, error) {
	if IsTerminal() {
		opts = append(opts, func(c *client.Client) {
			c.Hub = registry.NewRegistryHub(registry.CloudQueryRegistryURl, func(h *registry.Hub) {
				h.ProgressUpdater = ProgressUpdater{
					mpb.NewWithContext(ctx, mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond)),
					make(map[string]string), make(map[string]*mpb.Bar), ""}

			})
		})
	}
	c, err := client.New(cfg, opts...)
	if err != nil {
		ColorizedOutput(ColorError, "❌ Failed to initialize client.\n\n")
		return nil, err
	}
	return &Client{c, cfg}, err
}

func (c Client) DownloadProviders(ctx context.Context) error {
	ColorizedOutput(ColorProgress, "\nInitializing CloudQuery Providers...\n\n")
	err := c.c.Initialize(ctx)
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		ColorizedOutput(ColorError, "❌ Failed to initialize provider: %s.\n\n", err.Error())
		return err
	}
	// sleep some extra 100 milliseconds for progress refresh
	time.Sleep(100 * time.Millisecond)
	ColorizedOutput(ColorProgress, "\nFinished provider initialization...\n")
	return nil
}

func (c Client) Fetch(ctx context.Context) error {
	if err := c.DownloadProviders(ctx); err != nil {
		return err
	}
	ColorizedOutput(ColorProgress, "\nStarting provider fetch...\n\n")
	fetchProgress, fetchCallback := buildFetchProgress(ctx, c.cfg.Providers)
	request := client.FetchRequest{Providers: c.cfg.Providers}
	if IsTerminal() {
		request.UpdateCallback = fetchCallback
	}
	if err := c.c.Fetch(ctx, request); err != nil {
		return err
	}
	if IsTerminal() {
		fetchProgress.Wait()
	}
	ColorizedOutput(ColorProgress, "\nProvider fetch complete.\n")
	return nil
}

func (c Client) Client() *client.Client {
	return c.c
}

func buildFetchProgress(ctx context.Context, providers []*config.Provider) (*UiProgress, client.FetchUpdateCallback) {
	fetchProgress := NewUiProgress(ctx, ProgressOptions{
		statusFunc:  DefaultStatusUpdater,
		messageHook: DefaultMessageUpdater,
		appendDecorators: []decor.Decorator{
			decor.CountersNoUnit(" Finished Resources: %d/%d"),
		},
	})

	for _, p := range providers {
		fetchProgress.Add(p.Name, fmt.Sprintf("cq-provider-%s@%s", p.Name, "latest"), "fetching", int64(len(p.Resources)))
	}
	fetchCallback := func(update client.FetchUpdate) {
		if update.Error != "" {
			fetchProgress.Step(update.Provider, StatusError, fmt.Sprintf("error: %s", update.Error))
			return
		}
		fetchProgress.Increment(update.Provider, 1)
		bar, _ := fetchProgress.bars[update.Provider]

		if bar.status == StatusError {
			if update.AllDone() {
				bar.b.SetTotal(0, true)
			}
			return
		}
		if update.AllDone() {
			fetchProgress.Step(update.Provider, StatusOK, "fetch complete")
			return
		}
	}
	return fetchProgress, fetchCallback
}

func loadConfig(path string) (*config.Config, error) {
	parser := config.NewParser(nil)
	cfg, diags := parser.LoadConfigFile(path)
	if diags != nil {
		ColorizedOutput(ColorHeader, "Configuration Error Diagnostics:\n")
		for _, d := range diags {
			ColorizedOutput(ColorError, "❌ %s\n", d.Error())
			ColorizedOutput(ColorInfo, "%s\n", d.Detail)
		}
		return nil, fmt.Errorf("bad configuration")
	}
	return cfg, nil
}
