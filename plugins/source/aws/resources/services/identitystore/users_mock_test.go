package identitystore

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	iTypes "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	require.NoError(t, faker.FakeObject(&im))
	users := iTypes.User{}
	require.NoError(t, faker.FakeObject(&users))

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
