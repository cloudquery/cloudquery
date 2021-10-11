package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

func clearMultiplexers(f func() *provider.Provider) func() *provider.Provider {
	return func() *provider.Provider {
		p := f()
		for _, table := range p.ResourceMap {
			table.Multiplex = nil
		}
		return p
	}
}

func k8sTestHelper(t *testing.T, table *schema.Table, builder func(t *testing.T, ctrl *gomock.Controller) client.Services) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resource := providertest.ResourceTestData{
		Table:  table,
		Config: client.Config{},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			return &client.Client{
				Log: logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}),
				Services: builder(t, ctrl),
			}, nil
		},
	}
	providertest.TestResource(t, clearMultiplexers(Provider), resource)
}
