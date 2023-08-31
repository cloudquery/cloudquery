package accounts

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func buildAccounts(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var acc cloudflare.Account
	if err := faker.FakeObject(&acc); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Accounts(
		gomock.Any(),
		gomock.Any(),
	).Return(
		[]cloudflare.Account{acc},
		cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
		nil,
	)

	var accMem cloudflare.AccountMember
	if err := faker.FakeObject(&accMem); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().AccountMembers(
		gomock.Any(),
		acc.ID,
		gomock.Any(),
	).Return(
		[]cloudflare.AccountMember{accMem},
		cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), buildAccounts)
}
