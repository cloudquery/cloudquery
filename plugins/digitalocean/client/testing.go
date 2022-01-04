package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func DOTestHelper(t *testing.T, table *schema.Table) {
	cfg := ``
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "digitalocean_test_provider",
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
