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

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	require.NoError(t, faker.FakeObject(&im))
	group := iTypes.Group{}
	require.NoError(t, faker.FakeObject(&group))
	groupMembership := iTypes.GroupMembership{}
	require.NoError(t, faker.FakeObject(&groupMembership))
	groupMembership.MemberId = &iTypes.MemberIdMemberUserId{
		Value: "test",
	}

	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListInstancesOutput{
			Instances: []types.InstanceMetadata{im},
		}, nil)
	mIdentity.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&identitystore.ListGroupsOutput{
			Groups: []iTypes.Group{group},
		}, nil)
	mIdentity.EXPECT().ListGroupMemberships(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&identitystore.ListGroupMembershipsOutput{
			GroupMemberships: []iTypes.GroupMembership{groupMembership},
		}, nil)

	return client.Services{
		Ssoadmin:      mSSOAdmin,
		Identitystore: mIdentity,
	}
}

func TestIdentityStoreGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildGroups, client.TestOptions{})
}
