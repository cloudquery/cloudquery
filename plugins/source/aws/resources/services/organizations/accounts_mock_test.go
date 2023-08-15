package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)

	tt := make([]organizationsTypes.Tag, 3)
	require.NoError(t, faker.FakeObject(&tt))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListTagsForResourceOutput{
			Tags: tt,
		}, nil)

	ds := organizationsTypes.DelegatedService{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().ListDelegatedServicesForAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListDelegatedServicesForAccountOutput{
			DelegatedServices: []organizationsTypes.DelegatedService{ds},
		}, nil)
	p := organizationsTypes.Parent{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().ListParents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListParentsOutput{
			Parents: []organizationsTypes.Parent{p},
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationsAccounts(t *testing.T) {
	client.AwsMockTestHelper(t, Accounts(), buildOrganizationsAccounts, client.TestOptions{})
}
