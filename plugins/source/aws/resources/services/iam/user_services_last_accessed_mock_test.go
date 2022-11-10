package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamUserServicesLastAccessed(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iam.ListUsersOutput{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	g.Marker = nil
	m.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&g, nil)

	iamAccessDetailsMock(t, m)
	return client.Services{
		Iam: m,
	}
}

func TestIamUserServicesLastAccessedMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, UserServicesLastAccessed(), buildIamUserServicesLastAccessed, client.TestOptions{})
}
