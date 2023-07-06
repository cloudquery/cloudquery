package billing_history

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createBillingHistory(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBillingHistoryService(ctrl)

	data := &godo.BillingHistory{}
	if err := faker.FakeObject(data); err != nil {
		t.Fatal(err)
	}
	data.Links = nil

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		BillingHistory: m,
	}
}

func TestBillingHistory(t *testing.T) {
	client.MockTestHelper(t, BillingHistory(), createBillingHistory, client.TestOptions{})
}
