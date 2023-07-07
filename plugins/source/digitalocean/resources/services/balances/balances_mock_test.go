package balances

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createBalance(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBalanceService(ctrl)

	var data godo.Balance
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().Get(gomock.Any()).Return(&data, nil, nil)

	return client.Services{
		Balance: m,
	}
}

func TestBalance(t *testing.T) {
	client.MockTestHelper(t, Balances(), createBalance, client.TestOptions{})
}
