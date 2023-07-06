package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDelegatedAdministrators(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	da := types.DelegatedAdministrator{}
	require.NoError(t, faker.FakeObject(&da))

	m.EXPECT().ListDelegatedAdministrators(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListDelegatedAdministratorsOutput{
			DelegatedAdministrators: []types.DelegatedAdministrator{da},
		}, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestDelegatedAdministrators(t *testing.T) {
	client.AwsMockTestHelper(t, DelegatedAdministrators(), buildDelegatedAdministrators, client.TestOptions{})
}
