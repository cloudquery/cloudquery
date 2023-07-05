package monitoring

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createAlertPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockMonitoringService(ctrl)

	var data []godo.AlertPolicy
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAlertPolicies(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Monitoring: m,
	}
}

func TestAlertPolicies(t *testing.T) {
	client.MockTestHelper(t, AlertPolicies(), createAlertPolicies, client.TestOptions{})
}
