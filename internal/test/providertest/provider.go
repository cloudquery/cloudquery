package providertest

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/cloudquery/cq-provider-sdk/serve"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

//go:embed migrations
var migrations embed.FS

type Configuration struct {
	Accounts []Account `hcl:"account,block"`
}

func (c Configuration) Example() string {
	return `
  configuration {
    account "1" {
      regions = ["asdas"]
      resources = ["ab", "c"]
    }

    regions = ["adsa"]
  }`
}

type Account struct {
	Name      string   `hcl:"name,label"`
	Id        string   `hcl:"id"`
	Regions   []string `hcl:"regions,optional"`
	Resources []string `hcl:"resources,optional"`
}

type TestClient struct {
	l hclog.Logger
}

func (t TestClient) Logger() hclog.Logger {
	return t.l
}

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:       "test",
		Version:    "v0.0.0",
		Migrations: migrations,
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
			return &TestClient{l: logger}, nil
		},
		ResourceMap: map[string]*schema.Table{
			"slow_resource": {
				Name: "slow_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					meta.Logger().Info("fetching")
					select {
					case <-ctx.Done():
						return nil
					case <-time.After(time.Second * 5):
						return nil
					}
				},
				Columns: []schema.Column{
					{
						Name: "some_bool",
						Type: schema.TypeBool,
					},
					{
						Name: "upgrade_column",
						Type: schema.TypeInt,
					},
					{
						Name: "upgrade_column_2",
						Type: schema.TypeInt,
					},
				},
			},
			"very_slow_resource": {
				Name: "very_slow_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					meta.Logger().Info("fetching very slow")
					select {
					case <-ctx.Done():
						return nil
					case <-time.After(time.Second * 8):
						return nil
					}
				},
			},
			"error_resource": {
				Name: "error_resource",
				Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
					return fmt.Errorf("error from provider")
				},
			},
		},
		Config: func() provider.Config {
			return &Configuration{}
		},
		Logger: hclog.NewNullLogger(),
	}
}

func ServeTestPlugin(_ context.Context) {
	opts := &serve.Options{
		Name:                "test",
		Provider:            Provider(),
		Logger:              hclog.NewNullLogger(),
		NoLogOutputOverride: true,
	}
	if err := serve.Debug(context.Background(), opts.Name, opts); err != nil {
		panic(fmt.Errorf("failed to run debug: %w", err))
	}
}
