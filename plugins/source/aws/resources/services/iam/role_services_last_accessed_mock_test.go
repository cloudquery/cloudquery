package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamRoleServicesLastAccessed(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iam.ListRolesOutput{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	g.Marker = nil
	m.EXPECT().ListRoles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&g, nil)

	iamAccessDetailsMock(t, m)
	return client.Services{
		Iam: m,
	}
}

func TestIamRoleServicesLastAccessedMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, RoleServicesLastAccessed(), buildIamRoleServicesLastAccessed, client.TestOptions{})
}
