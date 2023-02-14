package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildOrganizationalUnits(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := types.Root{}
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListRoots(gomock.Any(), gomock.Any()).Return(
		&organizations.ListRootsOutput{
			Roots: []types.Root{g},
		}, nil)

	c := types.Child{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListChildren(gomock.Any(), gomock.Any()).MinTimes(1).Return(
		&organizations.ListChildrenOutput{
			Children: []types.Child{c},
		}, nil)

	ou := types.OrganizationalUnit{}
	if err := faker.FakeObject(&ou); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeOrganizationalUnit(gomock.Any(), gomock.Any()).Return(
		&organizations.DescribeOrganizationalUnitOutput{
			OrganizationalUnit: &ou,
		}, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationalUnits(t *testing.T) {
	client.AwsMockTestHelper(t, OrganizationalUnits(), buildOrganizationalUnits, client.TestOptions{})
}
