package ssoadmin

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mSSOAdmin := mocks.NewMockSsoadminClient(ctrl)
	im := types.InstanceMetadata{}
	ps := types.PermissionSet{}
	as := types.AccountAssignment{}
	pb := types.PermissionsBoundary{}
	cmpr := types.CustomerManagedPolicyReference{}
	amp := types.AttachedManagedPolicy{}
	ip := `{"key": "value"}`
	require.NoError(t, faker.FakeObject(&ps))

	require.NoError(t, faker.FakeObject(&as))

	require.NoError(t, faker.FakeObject(&im))

	require.NoError(t, faker.FakeObject(&pb))

	require.NoError(t, faker.FakeObject(&cmpr))

	require.NoError(t, faker.FakeObject(&amp))

	mSSOAdmin.EXPECT().ListInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListInstancesOutput{
			Instances: []types.InstanceMetadata{im},
		}, nil)

	mSSOAdmin.EXPECT().ListPermissionSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListPermissionSetsOutput{
			PermissionSets: []string{*ps.Name},
		}, nil)

	mSSOAdmin.EXPECT().DescribePermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.DescribePermissionSetOutput{
			PermissionSet: &ps,
		}, nil)
	mSSOAdmin.EXPECT().ListAccountsForProvisionedPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListAccountsForProvisionedPermissionSetOutput{
			AccountIds: []string{*as.AccountId},
		}, nil)
	mSSOAdmin.EXPECT().ListAccountAssignments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListAccountAssignmentsOutput{
			AccountAssignments: []types.AccountAssignment{as},
		}, nil)

	mSSOAdmin.EXPECT().GetInlinePolicyForPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.GetInlinePolicyForPermissionSetOutput{
			InlinePolicy: &ip,
		}, nil)

	mSSOAdmin.EXPECT().GetPermissionsBoundaryForPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.GetPermissionsBoundaryForPermissionSetOutput{
			PermissionsBoundary: &pb,
		}, nil)

	mSSOAdmin.EXPECT().ListCustomerManagedPolicyReferencesInPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput{
			CustomerManagedPolicyReferences: []types.CustomerManagedPolicyReference{cmpr},
		}, nil)
	mSSOAdmin.EXPECT().ListManagedPoliciesInPermissionSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ssoadmin.ListManagedPoliciesInPermissionSetOutput{
			AttachedManagedPolicies: []types.AttachedManagedPolicy{amp},
		}, nil)

	return client.Services{
		Ssoadmin: mSSOAdmin,
	}
}

func TestSSOAdminInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildInstances, client.TestOptions{})
}
