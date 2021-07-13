package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudquery/cq-provider-sdk/provider"
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
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
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
	// Optional: Disable deletion of data from tables.
	// Use this with caution, as it can create duplicates of data!
	DisableDataDelete bool
	// Optional: Adds extra fields to the provider, this is used for testing purposes.
	ExtraFields map[string]interface{}
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

// PolicyRunRequest is the request used to run a policy.
type PolicyRunRequest struct {
	// Args are the given arguments from the policy run command.
	Args []string

	// SubPath is the optional sub path for sub policy/query execution only.
	SubPath string

	// OutputPath is the output path for policy execution output.
	OutputPath string

	// StopOnFailure signals policy execution to stop after first failure.
	StopOnFailure bool

	// RunCallBack is the callback method that is called after every policy execution.
	RunCallBack policy.ExecutionCallback

	// SkipVersioning if true policy will be executed without checking out the version of the policy repo using git tags
	SkipVersioning bool
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

// TableCreator creates tables based on schema received from providers
type TableCreator interface {
	CreateTable(ctx context.Context, conn *pgxpool.Conn, t *schema.Table, p *schema.Table) error
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
	// Optional: Where to save downloaded policies, by default current working directory, defaults to ./cq/policy
	PolicyDirectory string
	// Optional: If true cloudquery just runs policy files without using git tag to select a version
	SkipVersioning bool
	// Optional: if this flag is true, plugins downloaded from URL won't be verified when downloaded
	NoVerify bool
	// Optional: DSN connection information for database client will connect to
	DSN string
	// Optional: Skips Building tables on fetch execution
	SkipBuildTables bool
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
	// TableCreator defines how table are created in the database
	TableCreator TableCreator
	// pool is a list of connection that are used for policy/query execution
	pool *pgxpool.Pool
}

func New(ctx context.Context, options ...Option) (*Client, error) {

	c := &Client{
		PluginDirectory:    filepath.Join(".", ".cq", "providers"),
		PolicyDirectory:    ".",
		NoVerify:           false,
		SkipBuildTables:    false,
		HubProgressUpdater: nil,
		RegistryURL:        registry.CloudQueryRegistryURl,
		Logger:             logging.NewZHcLog(&zerolog.Logger, ""),
	}
	for _, o := range options {
		o(c)
	}

	var err error

	if c.DSN == "" {
		return nil, errors.New("missing DSN, make sure to pass it either cq config connection block or CLI arg --dsn")
	}

	c.Manager, err = plugin.NewManager(c.Logger, c.PluginDirectory, c.RegistryURL, c.HubProgressUpdater)
	if err != nil {
		return nil, err
	}
	if c.TableCreator == nil {
		c.TableCreator = provider.NewMigrator(c.Logger)
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
	if !c.SkipBuildTables {
		for _, p := range request.Providers {
			if err := c.BuildProviderTables(ctx, p.Name); err != nil {
				return err
			}
		}
	}
	c.Logger.Info("received fetch request", "disable_delete", request.DisableDataDelete, "extra_fields", request.ExtraFields)
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
				Config:        cfg,
				DisableDelete: request.DisableDataDelete,
				ExtraFields:   request.ExtraFields,
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
		if providerPlugin.Version() == plugin.Unmanaged {
			c.Logger.Warn("Not closing unmanaged provider", "provider", providerName)
			return
		}
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
		if providerPlugin.Version() == plugin.Unmanaged {
			c.Logger.Warn("Not closing unmanaged provider", "provider", providerName)
			return
		}
		if err := c.Manager.KillProvider(providerName); err != nil {
			c.Logger.Warn("failed to close provider", "provider", providerName)
		}
	}()
	return providerPlugin.Provider().GetProviderConfig(ctx, &cqproto.GetProviderConfigRequest{})
}

func (c *Client) BuildProviderTables(ctx context.Context, providerName string) error {
	s, err := c.GetProviderSchema(ctx, providerName)
	if err != nil {
		return err
	}
	conn, err := c.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	for name, t := range s.ResourceTables {
		c.Logger.Debug("creating tables for resource for provider", "resource_name", name, "provider", s.Name, "version", s.Version)
		if err := c.TableCreator.CreateTable(ctx, conn, t, nil); err != nil {
			return err
		}
	}
	return nil
}

func (c Client) DownloadPolicy(ctx context.Context, args []string) error {
	c.Logger.Info("Downloading policy from GitHub", "args", args)
	m := policy.NewManager(c.PolicyDirectory, c.pool, c.Logger)

	// Parse input args
	p, err := m.ParsePolicyHubPath(args, "")
	if err != nil {
		return err
	}
	c.Logger.Debug("Parsed policy download input arguments", "policy", p)
	return m.DownloadPolicy(ctx, p)
}

func (c Client) RunPolicy(ctx context.Context, req PolicyRunRequest) error {
	c.Logger.Info("Running policy", "args", req.Args)
	m := policy.NewManager(c.PolicyDirectory, c.pool, c.Logger)

	// Parse input args
	p, err := m.ParsePolicyHubPath(req.Args, req.SubPath)
	if err != nil {
		return err
	}
	c.Logger.Debug("Parsed policy run input arguments", "policy", p)
	output, err := m.RunPolicy(ctx, &policy.ExecuteRequest{Policy: p, StopOnFailure: req.StopOnFailure, SkipVersioning: req.SkipVersioning, UpdateCallback: req.RunCallBack})
	if err != nil {
		return err
	}

	// Store output in file if requested
	if req.OutputPath != "" {
		fs := afero.NewOsFs()
		f, err := fs.OpenFile(req.OutputPath, os.O_RDWR|os.O_CREATE, 0644)
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
