package resources_test

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

const testSubscriptionID = "test_sub"

func azureTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.Services) {
	ctrl := gomock.NewController(t)
	providertest.TestResource(t, resources.Provider, providertest.ResourceTestData{
		Table:  table,
		Config: client.Config{Subscriptions: []string{testSubscriptionID}},
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
			c := client.NewAzureClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{testSubscriptionID})
			c.SetSubscriptionServices(testSubscriptionID, builder(t, ctrl))
			return c, nil
		},
	})
}
