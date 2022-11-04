package identitystore

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	iTypes "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	err := faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)
	}
	users := iTypes.User{}
	err = faker.FakeObject(&users)
	if err != nil {
		t.Fatal(err)
	}

	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListInstancesOutput{
			Instances: []types.InstanceMetadata{im},
		}, nil)
	mIdentity.EXPECT().ListUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&identitystore.ListUsersOutput{
			Users: []iTypes.User{users},
		}, nil)

	return client.Services{
		Ssoadmin:      mSSOAdmin,
		Identitystore: mIdentity,
	}
}

func TestIdentityStoreUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildUsers, client.TestOptions{})
}
