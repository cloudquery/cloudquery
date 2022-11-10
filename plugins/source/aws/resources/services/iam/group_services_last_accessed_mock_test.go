package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamGroupServicesLastAccessed(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)

	g := iamTypes.Group{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)

	iamAccessDetailsMock(t, m)
	return client.Services{
		Iam: m,
	}
}

func TestIamGroupServicesLastAccessedMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, GroupServicesLastAccessed(), buildIamGroupServicesLastAccessed, client.TestOptions{})
}
