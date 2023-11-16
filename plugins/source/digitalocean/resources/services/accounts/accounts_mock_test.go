package accounts

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createAccount(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAccountService(ctrl)

	var data godo.Account
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().Get(gomock.Any()).Return(&data, nil, nil)

	return client.Services{
		Account: m,
	}
}

func TestAccount(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccount, client.TestOptions{})
}
