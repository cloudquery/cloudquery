// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	
)

//go:generate mockgen -package=mocks -destination=../mocks/resourcegroups.go -source=resourcegroups.go ResourcegroupsClient
type ResourcegroupsClient interface {
	GetAccountSettings(context.Context, *resourcegroups.GetAccountSettingsInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetAccountSettingsOutput, error)
	GetGroup(context.Context, *resourcegroups.GetGroupInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error)
	GetGroupConfiguration(context.Context, *resourcegroups.GetGroupConfigurationInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupConfigurationOutput, error)
	GetGroupQuery(context.Context, *resourcegroups.GetGroupQueryInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error)
	GetTags(context.Context, *resourcegroups.GetTagsInput, ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error)
	ListGroupResources(context.Context, *resourcegroups.ListGroupResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupResourcesOutput, error)
	ListGroups(context.Context, *resourcegroups.ListGroupsInput, ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error)
	SearchResources(context.Context, *resourcegroups.SearchResourcesInput, ...func(*resourcegroups.Options)) (*resourcegroups.SearchResourcesOutput, error)
}
