package client

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudquery/cloudquery/internal/logging"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4/pgxpool"
	zerolog "github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

// FetchRequest is provided to the Client to execute a fetch on one or more providers
type FetchRequest struct {
	// UpdateCallback allows gets called when the client receives updates on fetch.
	UpdateCallback FetchUpdateCallback
	// Providers list of providers to call for fetching
	Providers []*config.Provider
}

type ExecutePolicyRequest struct {
	// Path to the policy, currently we still use the old .yml format, future versions will change to HCL
	PolicyPath string
	// UpdateCallback allows gets called when the client receives updates on policy execution.
	UpdateCallback PolicyExecutionCallback
	// if True policy execution will stop on first failure
	StopOnFailure bool
	// Path to save policy result
	OutputPath string
}

type PolicyExecutionResult struct {
	// True if all policies have passed
	Passed bool
	// Map of all query result sets
	Results map[string]*PolicyResult
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

type PolicyExecutionCallback func(name string, passed bool, resultCount int)

type Option func(options *Client)

// Client is the client for executing providers, fetching data and running queries and polices
type Client struct {
	// Optional: Logger framework can use to log.
	// default: global logger provided.
	Logger hclog.Logger

	// Optional: Hub client to use to download plugins, the Hub is used to download and pluginManager providers binaries
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

	poolCfg, err := pgxpool.ParseConfig(c.config.CloudQuery.Connection.DSN)
	if err != nil {
		return nil, err
	}
	poolCfg.LazyConnect = true
	c.pool, err = pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Initialize downloads all provider binaries
func (c *Client) Initialize(ctx context.Context) error {
	c.Logger.Info("Initializing required providers")
	for _, p := range c.config.CloudQuery.Providers {
		c.Logger.Info("Initializing provider", "name", p.Name, "version", p.Version)
		org, name, err := registry.ParseProviderName(p.Name)
		if err != nil {
			return err
		}
		details, err := c.Hub.GetProvider(ctx, org, name, p.Version)
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

		cqProvider, err := c.Manager.GetOrCreateProvider(&details)
		if err != nil {
			c.Logger.Error("failed to create provider plugin", "provider", provider.Name, "error", err)
			return err
		}
		provider := provider
		errGroup.Go(func() error {
			var cfg []byte
			if provider.Configuration != nil {
				cfg, err = convert.Body(providerCfg.Configuration, convert.Options{Simplify: true})
				if err != nil {
					return err
				}
			}
			c.Logger.Info("requesting provider to configure", "provider", provider.Name, "version", details.Version)
			_, err = cqProvider.ConfigureProvider(gctx, &cqproto.ConfigureProviderRequest{
				CloudQueryVersion: "", // TODO pass cloudquery version
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
			c.Logger.Debug("requesting provider fetch", "provider", provider.Name, "version", details.Version)
			stream, err := cqProvider.FetchResources(gctx, &cqproto.FetchResourcesRequest{Resources: provider.Resources})
			if err != nil {
				return err
			}
			c.Logger.Info("provider started fetching resources", "provider", providerCfg.Name, "version", details.Version)
			for {
				resp, err := stream.Recv()
				if err == io.EOF {
					c.Logger.Info("provider finished fetch", "provider", providerCfg.Name, "version", details.Version)
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
				if resp.Error != "" {
					c.Logger.Error("received error fetching", "provider", provider.Name, "error", resp.Error)
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
	cqProvider, err := c.Manager.GetOrCreateProvider(&details)
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
	cqProvider, err := c.Manager.GetOrCreateProvider(&details)
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

func (c Client) ExecutePolicy(ctx context.Context, request ExecutePolicyRequest) (*PolicyExecutionResult, error) {

	conn, err := c.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	data, err := afero.ReadFile(afero.NewOsFs(), request.PolicyPath)
	if err != nil {
		return nil, err
	}

	var policy config.Policy
	// TODO: convert to hcl
	if err := yaml.Unmarshal(data, &policy); err != nil {
		return nil, err
	}
	// Create Views
	c.Logger.Debug("creating policy views", "policy", request.PolicyPath)
	if err := createViews(ctx, conn, policy.Views); err != nil {
		return nil, fmt.Errorf("failed to create policy views %w", err)
	}
	exec := &PolicyExecutionResult{
		Passed:  true,
		Results: make(map[string]*PolicyResult, len(policy.Queries)),
	}

	for _, q := range policy.Queries {
		result, err := executePolicyQuery(ctx, conn, q)
		if err != nil {
			c.Logger.Error("failed to execute policy query", "policy", q.Name, "error", err)
			if request.StopOnFailure {
				return nil, fmt.Errorf("failed to execute policy query %s. Err: %w", q.Name, err)
			}
			if request.UpdateCallback != nil {
				request.UpdateCallback(q.Name, false, 0)
			}
			continue
		}
		if !result.Passed {
			exec.Passed = false
		}
		exec.Results[q.Name] = result
		if request.UpdateCallback != nil {
			request.UpdateCallback(q.Name, result.Passed, len(result.Data))
		}
	}
	if request.OutputPath != "" {
		if err := createPolicyOutput(request.OutputPath, exec); err != nil {
			return nil, fmt.Errorf("failed to create policy output %s. Err: %w", request.OutputPath, err)
		}
	}
	return exec, nil
}

func (c Client) Close() {
	c.Manager.Shutdown()
	c.pool.Close()
}
