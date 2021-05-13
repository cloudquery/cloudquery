package client

import (
	"context"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/internal/test/provider"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestClient_FetchTimeout(t *testing.T) {
	cancelServe := setupTestPlugin()
	defer cancelServe()

	cfg, diags := config.NewParser(nil).LoadConfigFromSource("config.hcl", []byte(testConfig))
	assert.Nil(t, diags)

	c, err := New(cfg, func(options *Client) {
		options.Hub = &MockRegistry{}
	})
	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}

	err = c.Initialize(context.Background())
	assert.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err = c.Fetch(ctx, FetchRequest{
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

func TestClient_Fetch(t *testing.T) {
	cancelServe := setupTestPlugin()
	defer cancelServe()

	cfg, diags := config.NewParser(nil).LoadConfigFromSource("config.hcl", []byte(testConfig))
	assert.Nil(t, diags)

	c, err := New(cfg, func(c *Client) {
		c.Hub = &MockRegistry{}
	})

	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}

	err = c.Initialize(context.Background())
	assert.Nil(t, err)

	ctx := context.Background()
	err = c.Fetch(ctx, FetchRequest{
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
	cancelServe := setupTestPlugin()
	defer cancelServe()

	cfg, diags := config.NewParser(nil).LoadConfigFromSource("config.hcl", []byte(testConfig))
	assert.Nil(t, diags)

	c, err := New(cfg, func(c *Client) {
		c.Hub = &MockRegistry{}
	})

	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}

	err = c.Initialize(context.Background())
	assert.Nil(t, err)

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
	cancelServe := setupTestPlugin()
	defer cancelServe()

	cfg, diags := config.NewParser(nil).LoadConfigFromSource("config.hcl", []byte(testConfig))
	assert.Nil(t, diags)

	c, err := New(cfg, func(c *Client) {
		c.Hub = &MockRegistry{}
	})

	assert.Nil(t, err)
	if c == nil {
		assert.FailNow(t, "failed to create client")
	}

	err = c.Initialize(context.Background())
	assert.Nil(t, err)

	ctx := context.Background()
	pConfig, err := c.GetProviderConfiguration(ctx, "test")
	if pConfig == nil {
		t.FailNow()
	}
	assert.NotNil(t, pConfig)
	assert.Equal(t, pConfig.Config, []byte(expectedProviderConfig))

	_, diags = hclparse.NewParser().ParseHCL(pConfig.Config, "testConfig.hcl")
	assert.Nil(t, diags)
}

type MockRegistry struct{}

func (m MockRegistry) VerifyProvider(_, _, _ string) bool {
	return true
}
func (m MockRegistry) GetProvider(_ context.Context, organization, providerName, providerVersion string) (registry.ProviderDetails, error) {
	return registry.ProviderDetails{
		Name:         providerName,
		Version:      providerVersion,
		Organization: organization,
	}, nil
}

const testConfig = `cloudquery {
  connection {
    dsn =  "host=localhost user=postgres password=pass DB.name=postgres port=5432"
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
    "slow_resource",
    "very_slow_resource",
    "error_resource"
  ]
}`

func setupTestPlugin() context.CancelFunc {
	debugCtx, cancelServe := context.WithCancel(context.Background())
	go provider.ServeTestPlugin(debugCtx)
	// sleep to allow test plugin to start
	time.Sleep(time.Second * 2)
	dir, _ := os.Getwd()
	_ = os.Setenv("CQ_REATTACH_PROVIDERS", path.Join(dir, ".cq_reattach"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cancelServe
}
