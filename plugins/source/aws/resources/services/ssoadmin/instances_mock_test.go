package ssoadmin

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	err := faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)
	}

	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListInstancesOutput{
			Instances: []types.InstanceMetadata{im},
		}, nil)

	return client.Services{
		Ssoadmin: mSSOAdmin,
	}
}

func TestSSOAdminInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildInstances, client.TestOptions{})
}
