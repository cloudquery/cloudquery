package client

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cloudquery/pkg/ui/console"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/fatih/color"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
)

// FetchRequest is provided to the Client to execute a fetch on one or more providers
type FetchRequest struct {
	// UpdateCallback allows gets called when the client receives updates on fetch.
	UpdateCallback FetchUpdateCallback
	// Providers list of providers to call for fetching
	Providers []*config.Provider
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

type FetchUpdateCallback func(update FetchUpdate)

type Option func(options *Client)

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Required: List of providers that are required, these providers will be download if DownloadProviders is called.
	Providers []*config.RequiredProvider
	// Optional: Registry url to verify plugins from, defaults to CloudQuery hub
	RegistryURL string
	// Optional: Where to save downloaded providers, by default current working directory, defaults to ./cq/providers
	PluginDirectory string
	// Optional: if this flag is true, plugins downloaded from URL won't be verified when downloaded
	NoVerify bool
	// Optional: DSN connection information for database client will connect to
	DSN string
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
	// pool is a list of connection that are used for policy/query execution
	pool *pgxpool.Pool
}

func New(ctx context.Context, options ...Option) (*Client, error) {

	c := &Client{
		PluginDirectory:    filepath.Join(".", ".cq", "providers"),
		NoVerify:           false,
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

	poolCfg, err := pgxpool.ParseConfig(c.DSN)
	if err != nil {
		return nil, err
	}
	poolCfg.LazyConnect = true
	c.pool, err = pgxpool.ConnectConfig(ctx, poolCfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// DownloadProviders downloads all provider binaries
func (c *Client) DownloadProviders(ctx context.Context) error {
	c.Logger.Info("Downloading required providers")
	return c.Manager.DownloadProviders(ctx, c.Providers, c.NoVerify)
}

func (c *Client) Fetch(ctx context.Context, request FetchRequest) error {
	errGroup, gctx := errgroup.WithContext(ctx)
	for _, providerConfig := range request.Providers {
		providerConfig := providerConfig
		c.Logger.Debug("creating provider plugin", "provider", providerConfig.Name)
		providerPlugin, err := c.Manager.CreatePlugin(providerConfig.Name, providerConfig.Alias, providerConfig.Env)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", providerConfig.Name, "error", err)
			return err
		}
		errGroup.Go(func() error {
			var cfg []byte
			if providerConfig.Configuration != nil {
				cfg, err = convert.Body(providerConfig.Configuration, convert.Options{Simplify: true})
				if err != nil {
					return err
				}
			}
			c.Logger.Info("requesting provider to configure", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
			_, err = providerPlugin.Provider().ConfigureProvider(gctx, &cqproto.ConfigureProviderRequest{
				CloudQueryVersion: "", // TODO pass cloudquery version
				Connection: cqproto.ConnectionDetails{
					DSN: c.DSN,
				},
				Config: cfg,
			})
			if err != nil {
				c.Logger.Error("failed to configure provider", "error", err, "provider", providerPlugin.Name())
				return err
			}
			c.Logger.Info("provider configured successfully", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
			c.Logger.Debug("requesting provider fetch", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
			stream, err := providerPlugin.Provider().FetchResources(gctx, &cqproto.FetchResourcesRequest{Resources: providerConfig.Resources})
			if err != nil {
				return err
			}
			c.Logger.Info("provider started fetching resources", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					c.Logger.Info("provider finished fetch", "provider", providerPlugin.Name(), "version", providerPlugin.Version())
					return nil
				}
				if err != nil {
					return err
				}
				update := FetchUpdate{
					Provider:          providerPlugin.Name(),
					Version:           providerPlugin.Version(),
					FinishedResources: resp.FinishedResources,
					ResourceCount:     resp.ResourceCount,
					Error:             resp.Error,
				}
				if resp.Error != "" {
					c.Logger.Error("received error fetching", "provider", providerPlugin.Name(), "error", resp.Error)
				}
				c.Logger.Debug("fetch update", "provider", providerPlugin.Name(), "resource_count", resp.ResourceCount, "finished", update.AllDone(), "finishCount", update.DoneCount())
				if request.UpdateCallback != nil {
					request.UpdateCallback(update)
				}
			}
		})
	}
	// TODO: kill all providers on end, add defer on top loop
	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (c Client) GetProviderSchema(ctx context.Context, providerName string) (*cqproto.GetProviderSchemaResponse, error) {
	providerPlugin, err := c.Manager.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to kill provider", "provider", providerName)
		}
	}()
	return providerPlugin.Provider().GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
}

func (c Client) GetProviderConfiguration(ctx context.Context, providerName string) (*cqproto.GetProviderConfigResponse, error) {
	providerPlugin, err := c.Manager.CreatePlugin(providerName, "", nil)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to kill provider", "provider", providerName)
		}
	}()
	return providerPlugin.Provider().GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
}

func (c Client) DownloadPolicy(ctx context.Context, args []string) error {
	c.Logger.Info("Downloading policy from GitHub", "args", args)
	m := policy.NewManager(c.config, c.pool)

	// Parse input args
	p, err := m.ParsePolicyHubPath(args, "")
	if err != nil {
		return err
	}
	c.Logger.Debug("Parsed policy download input arguments", "policy", p)
	return m.DownloadPolicy(ctx, p)
}

func (c Client) RunPolicy(ctx context.Context, args []string, subPath, outputPath string, stopOnFailure bool) error {
	c.Logger.Info("Running policy", "args", args)
	m := policy.NewManager(c.config, c.pool)

	// Parse input args
	p, err := m.ParsePolicyHubPath(args, subPath)
	if err != nil {
		return err
	}
	c.Logger.Debug("Parsed policy run input arguments", "policy", p)
	output, err := m.RunPolicy(ctx, &policy.ExecuteRequest{Policy: p, StopOnFailure: stopOnFailure, UpdateCallback: func(name string, passed bool) {
		if passed {
			ui.ColorizedOutput(ui.ColorInfo, "\t%s  %-140s %5s\n", console.EmojiStatus[ui.StatusOK], name, color.GreenString("passed"))
		} else {
			ui.ColorizedOutput(ui.ColorInfo, "\t%s %-140s %5s\n", console.EmojiStatus[ui.StatusError], name, color.RedString("failed"))
		}
	}})

	// Store output in file if requested
	if outputPath != "" {
		fs := afero.NewOsFs()
		f, err := fs.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()

		data, err := json.Marshal(&output)
		if err != nil {
			return err
		}
		if _, err := f.Write(data); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) Close() {
	c.Manager.Shutdown()
	c.pool.Close()
}
