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

type TestOptions struct{}

func AwsMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Services, _ TestOptions) {
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := `
regions: ["us-east-1"]
accounts:
  - id: testAccount
    role_arn: ""
aws_debug: false
max_retries: 3
max_backoff: 60
`
	// accounts := []Account{
	// 	{ID: "testAccount"},
	// }

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "aws_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, diag.Diagnostics) {
				c := NewAwsClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}))
				c.ServicesManager.InitServicesForPartitionAccountAndRegion("aws", "testAccount", "us-east-1", builder(t, ctrl))
				c.Partition = "aws"
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
