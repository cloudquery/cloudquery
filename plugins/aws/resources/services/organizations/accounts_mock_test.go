package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationsAccounts(t *testing.T) {
	client.AwsMockTestHelper(t, OrganizationsAccounts(), buildOrganizationsAccounts, client.TestOptions{})
}
