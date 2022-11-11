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

func buildGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mIdentity := mocks.NewMockIdentitystoreClient(ctrl)
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	err := faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)
	}
	group := iTypes.Group{}
	err = faker.FakeObject(&group)
	if err != nil {
		t.Fatal(err)
	}
	groupMembership := iTypes.GroupMembership{}
	err = faker.FakeObject(&groupMembership)
	if err != nil {
		t.Fatal(err)
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
