package client

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

const (
	SnapshotsDirPath   = "./snapshots"
	TestSubscriptionID = "test_sub"
	FakeResourceGroup  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/cqprovidertest"
)

func AzureMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.Services, options TestOptions) {
	t.Helper()
	ctrl := gomock.NewController(t)

	cfg := `
		subscriptions = ["test_sub"]
	`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: &provider.Provider{
			Name:    "aws_mock_test_provider",
			Version: "development",
			Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
				c := NewAzureClient(logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}), []string{TestSubscriptionID})
				c.SetSubscriptionServices(TestSubscriptionID, builder(t, ctrl))
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

func AzureTestHelper(t *testing.T, table *schema.Table, snapshotDirPath string) {
	t.Helper()

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
		Table:        table,
		Config:       "",
		SnapshotsDir: snapshotDirPath,
	})

}
