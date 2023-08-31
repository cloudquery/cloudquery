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

func buildOrganizationalUnits(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := types.Root{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListRoots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListRootsOutput{
			Roots: []types.Root{g},
		}, nil)

	c := types.Child{}
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().ListChildren(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&organizations.ListChildrenOutput{
			Children: []types.Child{c},
		}, nil)

	ou := types.OrganizationalUnit{}
	require.NoError(t, faker.FakeObject(&ou))

	m.EXPECT().DescribeOrganizationalUnit(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.DescribeOrganizationalUnitOutput{
			OrganizationalUnit: &ou,
		}, nil)

	p := types.Parent{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().ListParents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListParentsOutput{
			Parents: []types.Parent{p},
		}, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationalUnits(t *testing.T) {
	client.AwsMockTestHelper(t, OrganizationalUnits(), buildOrganizationalUnits, client.TestOptions{})
}
