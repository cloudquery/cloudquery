package services

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAccounts(t *testing.T, ctrl *gomock.Controller) client.Api {
	mock := mocks.NewMockApi(ctrl)

	var acc cloudflare.Account
	if err := faker.FakeData(&acc); err != nil {
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
	if err := faker.FakeData(&accMem); err != nil {
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

	return mock
}

func TestAccounts(t *testing.T) {
	client.CFMockTestHelper(t, Accounts(), buildAccounts)
}
