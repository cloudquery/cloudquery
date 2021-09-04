package client

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4"

	"github.com/cloudquery/cloudquery/internal/test/provider"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	requiredTestProviders = []*config.RequiredProvider{
		{
			Name:    "test",
			Source:  "cloudquery",
			Version: "latest",
		},
	}
)

func TestClient_PartialFetch(t *testing.T) {
	ctx := context.Background()
	c, err := New(ctx, func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	// download test provider if it doesn't already exist
	err = c.DownloadProviders(ctx)
	assert.Nil(t, err)

	result, err := c.Fetch(ctx, FetchRequest{
		UpdateCallback: nil,
		Providers: []*config.Provider{{
			Name:               "test",
			Alias:              "test_alias",
			EnablePartialFetch: true,
			Resources:          []string{"slow_resource", "panic_resource", "error_resource", "very_slow_resource"},
			Env:                nil,
			Configuration:      nil,
		},
		},
	})
	assert.Nil(t, err)
	testSummary, ok := result.ProviderFetchSummary["test"]
	assert.True(t, ok)
	assert.Len(t, testSummary.PartialFetchErrors, 2)
}

func TestClient_TestNoDownload(t *testing.T) {
	_ = os.RemoveAll(".cq/downloadTest")
	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
		options.PluginDirectory = ".cq/downloadTest"
	})
	assert.Nil(t, err)
	_, err = c.Manager.GetPluginDetails("test")
	assert.Error(t, err)

	_, err = c.GetProviderSchema(context.Background(), "test")
	assert.Error(t, err)
	err = c.DownloadProviders(context.Background())
	assert.Nil(t, err)
	// Should work after provider was downloaded
	_, err = c.GetProviderSchema(context.Background(), "test")
	assert.Nil(t, err)
	pd, err := c.Manager.GetPluginDetails("test")
	assert.Nil(t, err)
	assert.Equal(t, "test", pd.Name)
	c, err = New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
		options.PluginDirectory = ".cq/downloadTest"
	})
	assert.Nil(t, err)
	pd2, err := c.Manager.GetPluginDetails("test")
	assert.Nil(t, err)
	assert.Equal(t, pd2.FilePath, pd.FilePath)
	// Should work without download
	_, err = c.GetProviderSchema(context.Background(), "test")
	assert.Nil(t, err)
}

func TestClient_FetchTimeout(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	assert.Nil(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = c.Fetch(ctx, FetchRequest{
		Providers: []*config.Provider{
			{
				Name:      "test",
				Resources: []string{"slow_resource"},
			},
		},
	})
	_, ok := ctx.Deadline()
	assert.EqualError(t, err, "rpc error: code = DeadlineExceeded desc = context deadline exceeded")
	assert.True(t, ok)
}

// Test Client fetch but with a nil configuration, provider won't crash but use it's default configuration method
func TestClient_FetchNilConfig(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	cfg, diags := config.NewParser(nil).LoadConfigFromSource("config.hcl", []byte(testConfig))
	assert.Nil(t, diags)
	// Set configuration to nil
	cfg.Providers[0].Configuration = nil
	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	_, err = c.Fetch(ctx, FetchRequest{
		Providers: []*config.Provider{
			{
				Name:      "test",
				Resources: []string{"slow_resource"},
			},
		},
	})
	assert.Nil(t, err)

}

func TestClient_Fetch(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	assert.Nil(t, err)

	ctx := context.Background()
	_, err = c.Fetch(ctx, FetchRequest{
		Providers: []*config.Provider{
			{
				Name:      "test",
				Resources: []string{"slow_resource", "very_slow_resource"},
			},
		},
	})
	assert.Nil(t, err)
}

func TestClient_GetProviderSchema(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()
	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	schema, err := c.GetProviderSchema(ctx, "test")
	if schema == nil {
		t.FailNow()
	}
	assert.Equal(t, "test", schema.Name)
	assert.Equal(t, "v0.0.0", schema.Version)
	assert.Equal(t, 3, len(schema.ResourceTables))
	assert.Nil(t, err)
}

func TestClient_GetProviderConfig(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}

	ctx := context.Background()
	pConfig, err := c.GetProviderConfiguration(ctx, "test")
	if pConfig == nil {
		t.FailNow()
	}
	assert.NotNil(t, pConfig)
	assert.Equal(t, string(pConfig.Config), expectedProviderConfig)
	_, diags := hclparse.NewParser().ParseHCL(pConfig.Config, "testConfig.hcl")
	assert.Nil(t, diags)
}

func TestClient_ProviderMigrations(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	err = c.DropProvider(ctx, "test")
	assert.Nil(t, err)
	err = c.BuildProviderTables(ctx, "test")
	assert.Nil(t, err)
	err = c.UpgradeProvider(ctx, "test")
	assert.ErrorIs(t, err, migrate.ErrNoChange)

	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Nil(t, err)

	c.Providers[0].Version = "v0.0.1"
	err = c.DowngradeProvider(ctx, "test")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column from slow_resource")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Error(t, err)

	c.Providers[0].Version = "v0.0.2"
	err = c.UpgradeProvider(ctx, "test")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Nil(t, err)
}

func TestClient_ProviderSkipVersionMigrations(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	err = c.DropProvider(ctx, "test")
	assert.Nil(t, err)
	err = c.BuildProviderTables(ctx, "test")
	assert.Nil(t, err)
	err = c.UpgradeProvider(ctx, "test")
	assert.ErrorIs(t, err, migrate.ErrNoChange)

	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Nil(t, err)

	c.Providers[0].Version = "v0.0.1"
	err = c.DowngradeProvider(ctx, "test")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column from slow_resource")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Error(t, err)

	c.Providers[0].Version = "v0.0.5"
	// latest migration should be to v0.0.2
	err = c.UpgradeProvider(ctx, "test")
	assert.Nil(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Nil(t, err)

	// insert dummy migration files like test provider just for version number return
	m, _, err := c.buildProviderMigrator(map[string][]byte{
		"1_v0.0.1.up.sql":   []byte(""),
		"1_v0.0.1.down.sql": []byte(""),
		"2_v0.0.2.up.sql":   []byte(""),
		"2_v0.0.2.down.sql": []byte(""),
	}, "test")
	if err != nil {
		t.Fatal(err)
	}
	// migrations should be in 2 i.e v0.0.2
	v, dirty, err := m.Version()
	assert.Equal(t, []interface{}{"v0.0.2", false, nil}, []interface{}{v, dirty, err})

}

const testConfig = `cloudquery {
  connection {
    dsn = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
  }
  provider "test" {
    source = "cloudquery"
    version = "v0.0.0"
  }
}

provider "test" {
  configuration {
	account "dev" {
	    id = 123
		regions = ["us-east1"]
		resources = ["ec2"]
	}
  }
  resources = ["slow_resource"]
}`

const expectedProviderConfig = `
provider "test" {

  configuration {
    account "1" {
      regions   = ["asdas"]
      resources = ["ab", "c"]
    }

    regions = ["adsa"]
  }
  resources = [
    "error_resource",
    "slow_resource",
    "very_slow_resource"
  ]
}`

func setupTestPlugin(t *testing.T) context.CancelFunc {
	debugCtx, cancelServe := context.WithCancel(context.Background())
	dir, _ := os.Getwd()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatal(err)
	}
	if err := watcher.Add(dir); err != nil {
		t.Fatal(err)
	}
	defer watcher.Close()

	go provider.ServeTestPlugin(debugCtx)
	_ = os.Setenv("CQ_REATTACH_PROVIDERS", filepath.Join(dir, ".cq_reattach"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	_ = <-watcher.Events

	unmanaged, err := serve.ParseReattachProviders(os.Getenv("CQ_REATTACH_PROVIDERS"))
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range unmanaged {
		_, err := net.DialTimeout(u.Addr.Network(), u.Addr.String(), time.Second*5)
		if err != nil {
			t.Fatal(err)
		}
	}

	return cancelServe
}
