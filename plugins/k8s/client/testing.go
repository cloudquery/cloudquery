package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

func K8sTestHelper(t *testing.T, table *schema.Table, snapshotDirPath string) {
	cfg := ``
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "k8s_mock_test_provider",
			Version:   "development",
			Configure: Configure,
			Config: func() provider.Config {
				return &Config{}
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
		},
		Table:           table,
		Config:          cfg,
		SkipEmptyRows:   false,
		SkipEmptyColumn: true,
	})
}

type TestOptions struct {
	SkipEmptyJsonB bool
}

func K8sMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, options TestOptions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := ``

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "k8s_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
				c := &Client{
					Log:     logger,
					Context: "testContext",
				}
				c.SetServices(map[string]Services{"testContext": builder(t, ctrl)})
				return c, nil
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
			Config: func() provider.Config {
				return &Config{}
			},
		},
		Table:          table,
		Config:         cfg,
		SkipEmptyJsonB: options.SkipEmptyJsonB,
	})
}
