// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/ssoadmin.go -source=ssoadmin.go SsoadminClient
type SsoadminClient interface {
	DescribeAccountAssignmentCreationStatus(context.Context, *ssoadmin.DescribeAccountAssignmentCreationStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error)
	DescribeAccountAssignmentDeletionStatus(context.Context, *ssoadmin.DescribeAccountAssignmentDeletionStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error)
	DescribeInstanceAccessControlAttributeConfiguration(context.Context, *ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput, error)
	DescribePermissionSet(context.Context, *ssoadmin.DescribePermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetOutput, error)
	DescribePermissionSetProvisioningStatus(context.Context, *ssoadmin.DescribePermissionSetProvisioningStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error)
	GetInlinePolicyForPermissionSet(context.Context, *ssoadmin.GetInlinePolicyForPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error)
	GetPermissionsBoundaryForPermissionSet(context.Context, *ssoadmin.GetPermissionsBoundaryForPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput, error)
	ListAccountAssignmentCreationStatus(context.Context, *ssoadmin.ListAccountAssignmentCreationStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error)
	ListAccountAssignmentDeletionStatus(context.Context, *ssoadmin.ListAccountAssignmentDeletionStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error)
	ListAccountAssignments(context.Context, *ssoadmin.ListAccountAssignmentsInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountAssignmentsOutput, error)
	ListAccountsForProvisionedPermissionSet(context.Context, *ssoadmin.ListAccountsForProvisionedPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error)
	ListCustomerManagedPolicyReferencesInPermissionSet(context.Context, *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, error)
	ListInstances(context.Context, *ssoadmin.ListInstancesInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListInstancesOutput, error)
	ListManagedPoliciesInPermissionSet(context.Context, *ssoadmin.ListManagedPoliciesInPermissionSetInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error)
	ListPermissionSetProvisioningStatus(context.Context, *ssoadmin.ListPermissionSetProvisioningStatusInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error)
	ListPermissionSets(context.Context, *ssoadmin.ListPermissionSetsInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsOutput, error)
	ListPermissionSetsProvisionedToAccount(context.Context, *ssoadmin.ListPermissionSetsProvisionedToAccountInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error)
	ListTagsForResource(context.Context, *ssoadmin.ListTagsForResourceInput, ...func(*ssoadmin.Options)) (*ssoadmin.ListTagsForResourceOutput, error)
}
