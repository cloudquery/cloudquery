package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)

	tt := make([]organizationsTypes.Tag, 3)
	if err := faker.FakeObject(&tt); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListTagsForResourceOutput{
			Tags: tt,
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationsAccounts(t *testing.T) {
	client.AwsMockTestHelper(t, Accounts(), buildOrganizationsAccounts, client.TestOptions{})
}
