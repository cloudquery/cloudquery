package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

const (
	TestAccountID = "test_account"
	TestZoneID    = "test_zone"
)

func CFMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Api) {
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := ""

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "cloudflare_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
				c := NewClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), builder(t, ctrl), AccountZones{
					TestAccountID: {
						AccountId: TestAccountID,
						Zones:     []string{TestZoneID},
					},
				})
				return &c, nil
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
			Config: func() provider.Config {
				return &Config{}
			},
		},
		Config: cfg,
	})
}
