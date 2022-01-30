package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func GcpMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "gcp_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				svc, err := createService()
				if err != nil {
					return nil, err
				}
				c := NewGcpClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), []string{"testProject"}, svc)
				return c, nil
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
			Config: func() provider.Config {
				return &Config{}
			},
		},
		Config: "",
	})
}

func GcpTestHelper(t *testing.T, table *schema.Table) {
	t.Helper()
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "gcp_mock_test_provider",
			Version:   "development",
			Configure: Configure,
			Config: func() provider.Config {
				return &Config{}
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
		},
	})
}
