// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ram"
)

//go:generate mockgen -package=mocks -destination=../mocks/ram.go -source=ram.go RamClient
type RamClient interface {
	GetPermission(context.Context, *ram.GetPermissionInput, ...func(*ram.Options)) (*ram.GetPermissionOutput, error)
	GetResourcePolicies(context.Context, *ram.GetResourcePoliciesInput, ...func(*ram.Options)) (*ram.GetResourcePoliciesOutput, error)
	GetResourceShareAssociations(context.Context, *ram.GetResourceShareAssociationsInput, ...func(*ram.Options)) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareInvitations(context.Context, *ram.GetResourceShareInvitationsInput, ...func(*ram.Options)) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShares(context.Context, *ram.GetResourceSharesInput, ...func(*ram.Options)) (*ram.GetResourceSharesOutput, error)
	ListPendingInvitationResources(context.Context, *ram.ListPendingInvitationResourcesInput, ...func(*ram.Options)) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPermissionAssociations(context.Context, *ram.ListPermissionAssociationsInput, ...func(*ram.Options)) (*ram.ListPermissionAssociationsOutput, error)
	ListPermissionVersions(context.Context, *ram.ListPermissionVersionsInput, ...func(*ram.Options)) (*ram.ListPermissionVersionsOutput, error)
	ListPermissions(context.Context, *ram.ListPermissionsInput, ...func(*ram.Options)) (*ram.ListPermissionsOutput, error)
	ListPrincipals(context.Context, *ram.ListPrincipalsInput, ...func(*ram.Options)) (*ram.ListPrincipalsOutput, error)
	ListReplacePermissionAssociationsWork(context.Context, *ram.ListReplacePermissionAssociationsWorkInput, ...func(*ram.Options)) (*ram.ListReplacePermissionAssociationsWorkOutput, error)
	ListResourceSharePermissions(context.Context, *ram.ListResourceSharePermissionsInput, ...func(*ram.Options)) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceTypes(context.Context, *ram.ListResourceTypesInput, ...func(*ram.Options)) (*ram.ListResourceTypesOutput, error)
	ListResources(context.Context, *ram.ListResourcesInput, ...func(*ram.Options)) (*ram.ListResourcesOutput, error)
}
