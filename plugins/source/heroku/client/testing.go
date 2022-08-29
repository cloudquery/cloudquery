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

type TestOptions struct{}

func HerokuMockTestHelper(t *testing.T, table *schema.Table, builder func() (HerokuService, error), _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	cfg := ``

	heroku, err := builder()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "heroku_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
				c := Client{
					logger: logging.New(&hclog.LoggerOptions{
						Level: hclog.Warn,
					}),
					Heroku: heroku,
				}
				return &c, nil
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
			Config: func() provider.Config {
				return &Config{}
			},
		},
		Config:           cfg,
		SkipIgnoreInTest: true,
	})
}
