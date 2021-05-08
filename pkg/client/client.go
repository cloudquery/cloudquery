package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	zerolog "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"io"
)

// UpgradeRequest is provided to the Client to execute an upgrade of one or more providers
type UpgradeRequest struct {
	Provider string
	Version  string
}

// FetchRequest is provided to the Client to execute a fetch on one or more providers
type FetchRequest struct {
	// UpdateCallback allows gets called when the client receives updates on fetch.
	UpdateCallback FetchUpdateCallback
	// Providers list of providers to call for fetching
	Providers []*config.Provider
}

type Option func(options *Client)

// ConnectionOptions is provided by Client consumers to control connection to database.
type ConnectionOptions struct {
	DriverName string
	DSN        string
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

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Optional: Logger framework can use to log.
	// default: default logger provided.
	Logger hclog.Logger

	// Optional: Hub client will use to download plugins, the Hub is used to download and pluginManager providers binaries
	// if not specified, default cloudquery registry is used.
	Hub registry.Registry

	// manager manages all plugins lifecycle
	Manager *plugin.Manager

	// pool is a list of connection that are used for policy/query execution
	pool *pgxpool.Pool

	// Configuration of CloudQuery Client
	config *config.Config

	// map of providers downloaded and are kept in the state for a fetch call
	providers map[string]registry.ProviderDetails
}

func New(config *config.Config, options ...Option) (*Client, error) {
	m, err := plugin.NewManager()
	if err != nil {
		return nil, err
	}
	c := &Client{
		config:    config,
		Logger:    logging.NewZHcLog(&zerolog.Logger, ""),
		Manager:   m,
		providers: make(map[string]registry.ProviderDetails),
	}
	for _, o := range options {
		o(c)
	}
	if c.Hub == nil {
		c.Hub = registry.NewRegistryHub(registry.CloudQueryRegistryURl)
	}
	for k, v := range c.Manager.ListUnmanaged() {
		c.providers[k] = v
	}
	return c, nil
}

// Initialize downloads all provider binaries
func (c *Client) Initialize(ctx context.Context) error {
	c.Logger.Info("Initializing required providers")
	for _, p := range c.config.CloudQuery.Providers {
		c.Logger.Info("Initializing provider", "name", p.Name, "version", p.Version)
		// TODO: when we will support multiple organization use the source attribute
		details, err := c.Hub.GetProvider(ctx, plugin.DefaultOrganization, p.Name, p.Version)
		if err != nil {
			return err
		}
		c.providers[p.Name] = details
	}
	return nil
}

func (c *Client) Fetch(ctx context.Context, request FetchRequest) error {

	errGroup, gctx := errgroup.WithContext(ctx)
	for _, provider := range request.Providers {
		details, ok := c.providers[provider.Name]
		if !ok {
			return fmt.Errorf("provider plugin %s missing from plugin directory", provider.Name)
		}
		c.Logger.Debug("creating provider plugin", "provider", provider.Name)
		providerCfg, err := c.config.GetProvider(provider.Name)
		if err != nil {
			return fmt.Errorf("failed to find provider %s inside config", provider.Name)
		}
		// TODO: pass filepath instead
		cqProvider, err := c.Manager.GetOrCreateProvider(provider.Name, details.Version)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", provider.Name, "error", err)
			return err
		}
		provider := provider
		errGroup.Go(func() error {
			cfg, err := convert.Body(providerCfg.Configuration, convert.Options{Simplify: false})
			if err != nil {
				return err
			}
			c.Logger.Info("requesting provider to configure", "provider", provider.Name, "version", details.Version)
			_, err = cqProvider.ConfigureProvider(gctx, &cqproto.ConfigureProviderRequest{
				CloudQueryVersion: "",
				Connection: cqproto.ConnectionDetails{
					DSN: c.config.CloudQuery.Connection.DSN,
				},
				Config: cfg,
			})
			if err != nil {
				c.Logger.Error("failed to configure provider", "error", err, "provider", provider.Name)
				return err
			}
			c.Logger.Info("provider configured successfully", "provider", provider.Name, "version", details.Version)
			c.Logger.Info("requesting provider fetch", "provider", provider.Name, "version", details.Version)
			stream, err := cqProvider.FetchResources(gctx, &cqproto.FetchResourcesRequest{Resources: provider.Resources})
			if err != nil {
				return err
			}
			c.Logger.Info("provider started fetching resources", "provider", providerCfg.Name, "version", details.Version)
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				update := FetchUpdate{
					Provider:          provider.Name,
					Version:           details.Version,
					FinishedResources: resp.FinishedResources,
					ResourceCount:     resp.ResourceCount,
					Error:             resp.Error,
				}
				c.Logger.Debug("fetch update", "provider", provider.Name, "resource_count", resp.ResourceCount, "finished", update.AllDone(), "finishCount", update.DoneCount())
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
	details, ok := c.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("provider plugin %s missing from plugin directory", providerName)
	}
	cqProvider, err := c.Manager.GetOrCreateProvider(providerName, details.Version)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to kill provider", "provider", providerName)
		}
	}()
	return cqProvider.GetProviderSchema(ctx, &cqproto.GetProviderSchemaRequest{})
}

func (c Client) GetProviderConfiguration(ctx context.Context, providerName string) (*cqproto.GetProviderConfigResponse, error) {
	details, ok := c.providers[providerName]
	if !ok {
		return nil, fmt.Errorf("provider plugin %s missing from plugin directory", providerName)
	}
	cqProvider, err := c.Manager.GetOrCreateProvider(providerName, details.Version)
	if err != nil {
		c.Logger.Error("failed to create provider plugin", "provider", providerName, "error", err)
		return nil, err
	}
	defer func() {
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to kill provider", "provider", providerName)
		}
	}()
	return cqProvider.GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
}

func (c Client) ExecutePolicy(ctx context.Context, request interface{}) (interface{}, error) {
	panic("implement me")
}

func (c Client) Query(ctx context.Context, query interface{}) (interface{}, error) {
	panic("not implemented")
}

func (c Client) Close() {
	c.Manager.Shutdown()
	c.pool.Close()
}
