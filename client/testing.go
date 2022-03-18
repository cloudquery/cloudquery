package client

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"

	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

type TestOptions struct {
}

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := `
		regions = ["us-east-1"]
		accounts "testAccount" {
			role_arn = ""
		}
		aws_debug = false
		max_retries = 3
		max_backoff = 60
	`
	accounts := []Account{
		{ID: "testAccount"},
	}

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "aws_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				c := NewAwsClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), accounts)
				c.ServicesManager.InitServicesForAccountAndRegion("testAccount", "us-east-1", builder(t, ctrl))
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

func AWSTestHelper(t *testing.T, table *schema.Table) {
	t.Helper()
	cfg := `
	aws_debug = false
	`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:      "aws_mock_test_provider",
			Version:   "development",
			Configure: Configure,
			Config: func() provider.Config {
				return &Config{}
			},
			ResourceMap: map[string]*schema.Table{
				"test_resource": table,
			},
		},
		Config:           cfg,
		SkipIgnoreInTest: true,
	})

}
