package client

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/config"
	"github.com/cloudquery/cloudquery/hub"
	"github.com/cloudquery/cloudquery/logging"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	zerolog "github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

// UpgradeRequest is provided to the Client to execute an upgrade of one or more providers
type UpgradeRequest struct {
	Provider string
	Version  string
}

// FetchRequest is provided to the Client to execute a fetch on one or more providers
type FetchRequest struct {
	Providers []config.Provider
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
	ResourceCount string
	// True if provider has finished fetching
	Done bool
}

type FetchDoneResult struct {
	// Map of providers and resources that have finished fetching
	DoneResources map[string]map[string]bool
	// Amount of resources collected so far
	ResourceCount string
}

type NoOpFetchHooks struct{}

func (n NoOpFetchHooks) OnFetchStart(_ context.Context, _ FetchRequest) error { return nil }

func (n NoOpFetchHooks) OnFetchUpdate(_ context.Context, _ FetchUpdate) error { return nil }

func (n NoOpFetchHooks) OnFetchDone(_ context.Context, _ FetchDoneResult) error { return nil }

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Optional: database connection options. If not specified client will connect to a local postgres
	ConnectionOptions ConnectionOptions

	// Optional: Logger framework can use to log.
	// default: default logger provided.
	Logger hclog.Logger

	// Optional: client will download missing providers in fetch execution
	// default: true.
	DownloadMissingProviders bool

	// Optional: Hub client will use to download plugins, the Hub is used to download and pluginManager providers binaries
	Hub hub.Registry

	// manager manages all plugins lifecycle
	manager *pluginManager

	// pool is a list of connection that are used for policy/query execution
	pool *pgxpool.Pool

	// map of providers downloaded and are kept in the state for a fetch call
	providers map[string]hub.ProviderDetails
}

func New(options ...Option) (*Client, error) {
	m, err := newManager()
	if err != nil {
		return nil, err
	}
	c := &Client{
		ConnectionOptions: ConnectionOptions{
			DriverName: "postgres",
			DSN:        "",
		},
		Logger:                   logging.NewZHcLog(&zerolog.Logger, ""),
		DownloadMissingProviders: true,
		manager:                  m,
		providers:                make(map[string]hub.ProviderDetails),
	}
	for _, o := range options {
		o(c)
	}
	if c.Hub == nil {
		c.Hub = hub.NewRegistryHub(hub.CloudQueryRegistryURl)
	}

	return c, nil

}

func (c *Client) Initialize(ctx context.Context, providers []config.RequiredProvider) error {
	c.Logger.Info("Initializing required providers")
	for _, p := range providers {
		c.Logger.Info("Initializing provider", "name", p.Name, "version", p.Version)
		// TODO: when we support multiple organization use the source attribute
		details, err := c.Hub.GetProvider(ctx, defaultOrganization, p.Name, p.Version)
		if err != nil {
			return err
		}
		c.providers[p.Name] = details
	}
	return nil
}

func (c *Client) Fetch(ctx context.Context, request FetchRequest) error {
	errGroup, _ := errgroup.WithContext(ctx)
	for _, provider := range request.Providers {
		details, ok := c.providers[provider.Name]
		if !ok {
			return fmt.Errorf("provider plugin %s missing from plugin directory", provider.Name)
		}
		c.Logger.Debug("creating provider plugin", "provider", provider.Name)
		// TODO: pass filepath instead
		cqProvider, err := c.manager.GetOrCreateProvider(provider.Name, details.Version)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", provider.Name)
			continue
		}
		// create intermediate variable
		provider := provider
		errGroup.Go(func() error {
			c.Logger.Info("requesting provider initialize", "provider", provider.Name, "version", details.Version)
			err = cqProvider.Init(c.ConnectionOptions.DriverName, c.ConnectionOptions.DSN, true)
			if err != nil {
				return err
			}
			d, err := yaml.Marshal(&provider)
			if err != nil {
				return err
			}
			c.Logger.Info("requesting provider fetch", "provider", provider.Name, "version", details.Version)
			err = cqProvider.Fetch(d)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}

func (c Client) ExecutePolicy(ctx context.Context, request interface{}) (interface{}, error) {
	panic("implement me")
}

func (c Client) Query(ctx context.Context, query interface{}) (interface{}, error) {
	panic("not implemented")
}

func (c Client) Close() {
	c.manager.Shutdown()
	c.pool.Close()
}