package ssoadmin

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
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
	err := faker.FakeObject(&ps)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&as)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&im)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&pb)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&cmpr)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&amp)
	if err != nil {
		t.Fatal(err)
	}
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
