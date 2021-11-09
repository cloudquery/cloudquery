package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

func k8sTestHelper(t *testing.T, table *schema.Table, builder func(t *testing.T, ctrl *gomock.Controller) client.Services) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resource := providertest.ResourceTestData{
		Table:  table,
		Config: client.Config{},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			c := &client.Client{
				Log: logging.New(&hclog.LoggerOptions{
					Level: hclog.Warn,
				}),
				Context: "testContext",
			}
			c.SetServices(map[string]client.Services{"testContext": builder(t, ctrl)})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}
