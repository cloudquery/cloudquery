package client

import (
	"context"
	"errors"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/internal/test/provider"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/fsnotify/fsnotify"
	"github.com/golang-migrate/migrate/v4"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestClient_FailOnFetchWithPartialFetch(t *testing.T) {
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
	assert.True(t, testSummary.HasErrors())
	assert.Len(t, testSummary.PartialFetchErrors, 2)
	assert.Len(t, testSummary.FetchErrors, 0)
}

func TestClient_FailOnFetch(t *testing.T) {
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
			EnablePartialFetch: false,
			Resources:          []string{"slow_resource", "panic_resource", "error_resource", "very_slow_resource"},
			Env:                nil,
			Configuration:      nil,
		},
		},
	})
	assert.Nil(t, err)
	testSummary, ok := result.ProviderFetchSummary["test"]
	assert.True(t, ok)
	assert.True(t, testSummary.HasErrors())
	assert.Len(t, testSummary.FetchErrors, 2)
}

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
	cfg, diags := config.NewParser().LoadConfigFromSource("config.hcl", []byte(testConfig))
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
		return
	}
	ctx := context.Background()
	s, err := c.GetProviderSchema(ctx, "test")
	if s == nil {
		t.FailNow()
	}
	assert.Equal(t, "test", s.Name)
	assert.Equal(t, "v0.0.0", s.Version)
	assert.Equal(t, 3, len(s.ResourceTables))
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
		return
	}

	ctx := context.Background()
	pConfig, err := c.GetProviderConfiguration(ctx, "test")
	if err != nil || pConfig == nil {
		t.FailNow()
	}
	assert.NotNil(t, pConfig)
	assert.Equal(t, string(pConfig.Config), expectedProviderConfig)
	_, diags := hclparse.NewParser().ParseHCL(pConfig.Config, "testConfig.hcl")
	assert.Nil(t, diags)
}

// TestClient_ProviderUpgradeNoBuild tests doing an upgrade but without
func TestClient_ProviderUpgradeNoBuild(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.NoError(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	err = c.DropProvider(ctx, "test")
	assert.NoError(t, err)
	err = c.UpgradeProvider(ctx, "test")
	assert.NoError(t, nil)
}

func TestClient_ProviderMigrations(t *testing.T) {
	cancelServe := setupTestPlugin(t)
	defer cancelServe()

	c, err := New(context.Background(), func(options *Client) {
		options.DSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
		options.Providers = requiredTestProviders
	})
	assert.NoError(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}
	ctx := context.Background()
	err = c.DropProvider(ctx, "test")
	assert.NoError(t, err)
	err = c.BuildProviderTables(ctx, "test")
	assert.NoError(t, err)
	err = c.UpgradeProvider(ctx, "test")
	assert.ErrorIs(t, err, migrate.ErrNoChange)

	conn, err := pgx.Connect(ctx, "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.NoError(t, err)

	c.Providers[0].Version = "v0.0.1"
	err = c.DowngradeProvider(ctx, "test")
	assert.NoError(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column from slow_resource")
	assert.NoError(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.Error(t, err)

	c.Providers[0].Version = "v0.0.2"
	err = c.UpgradeProvider(ctx, "test")
	assert.NoError(t, err)
	_, err = conn.Exec(ctx, "select some_bool, upgrade_column, upgrade_column_2 from slow_resource")
	assert.NoError(t, err)

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
  // list of resources to fetch
  resources = [
    "error_resource",
    "slow_resource",
    "very_slow_resource"
  ]
  // enables partial fetching, allowing for any failures to not stop full resource pull
  enable_partial_fetch = true
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
	<-watcher.Events

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

func Test_normalizeResources(t *testing.T) {

	tests := []struct {
		name      string
		requested []string
		all       map[string]*schema.Table
		want      []string
		wantErr   bool
	}{
		{
			"wilcard",
			[]string{"*"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2", "3"},
			false,
		},
		{
			"wilcard with explicit",
			[]string{"*", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"unknown resource",
			[]string{"1", "2", "x"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"duplicate resource",
			[]string{"1", "2", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"ok, all explicit",
			[]string{"2", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeResources(tt.requested, tt.all)
			if (err != nil) != tt.wantErr {
				t.Errorf("doInterpolate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doInterpolate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collectProviderVersions(t *testing.T) {
	tests := []struct {
		name       string
		providers  []*config.RequiredProvider
		getVersion func(providerName string) (string, error)
		want       map[string]*version.Version
		wantErr    bool
	}{
		{
			"no required providers",
			nil,
			func(providerName string) (string, error) { panic("test") },
			map[string]*version.Version{},
			false,
		},
		{
			"failed to get a version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "1.0", errors.New("test") },
			nil,
			true,
		},
		{
			"failed to parse version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "xyz", nil },
			nil,
			true,
		},
		{
			"failed to parse version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "xyz", nil },
			nil,
			true,
		},
		{
			"ok",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) {
				if providerName == "aws" {
					return "1.2.3", nil
				} else {
					return "v4.5.6", nil
				}
			},
			map[string]*version.Version{
				"aws": version.Must(version.NewVersion("1.2.3")),
				"gcp": version.Must(version.NewVersion("v4.5.6")),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collectProviderVersions(tt.providers, tt.getVersion)
			require.Equal(t, tt.wantErr, err != nil, "collectProviderVersions() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "collectProviderVersions() = %v, want %v", got, tt.want)
		})
	}
}
