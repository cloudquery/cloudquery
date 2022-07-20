package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func GcpMockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "gcp_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
				svc, err := createService()
				if err != nil {
					return nil, diag.FromError(err, diag.INTERNAL)
				}
				c := NewGcpClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), BackoffSettings{}, []string{"testProject"}, svc)
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
