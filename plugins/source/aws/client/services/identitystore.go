// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
)

//go:generate mockgen -package=mocks -destination=../mocks/identitystore.go -source=identitystore.go IdentitystoreClient
type IdentitystoreClient interface {
	DescribeGroup(context.Context, *identitystore.DescribeGroupInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupOutput, error)
	DescribeGroupMembership(context.Context, *identitystore.DescribeGroupMembershipInput, ...func(*identitystore.Options)) (*identitystore.DescribeGroupMembershipOutput, error)
	DescribeUser(context.Context, *identitystore.DescribeUserInput, ...func(*identitystore.Options)) (*identitystore.DescribeUserOutput, error)
	GetGroupId(context.Context, *identitystore.GetGroupIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupIdOutput, error)
	GetGroupMembershipId(context.Context, *identitystore.GetGroupMembershipIdInput, ...func(*identitystore.Options)) (*identitystore.GetGroupMembershipIdOutput, error)
	GetUserId(context.Context, *identitystore.GetUserIdInput, ...func(*identitystore.Options)) (*identitystore.GetUserIdOutput, error)
	ListGroupMemberships(context.Context, *identitystore.ListGroupMembershipsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsOutput, error)
	ListGroupMembershipsForMember(context.Context, *identitystore.ListGroupMembershipsForMemberInput, ...func(*identitystore.Options)) (*identitystore.ListGroupMembershipsForMemberOutput, error)
	ListGroups(context.Context, *identitystore.ListGroupsInput, ...func(*identitystore.Options)) (*identitystore.ListGroupsOutput, error)
	ListUsers(context.Context, *identitystore.ListUsersInput, ...func(*identitystore.Options)) (*identitystore.ListUsersOutput, error)
}
