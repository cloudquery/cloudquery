// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
)

//go:generate mockgen -package=mocks -destination=../mocks/servicecatalogappregistry.go -source=servicecatalogappregistry.go ServicecatalogappregistryClient
type ServicecatalogappregistryClient interface {
	GetApplication(context.Context, *servicecatalogappregistry.GetApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetApplicationOutput, error)
	GetAssociatedResource(context.Context, *servicecatalogappregistry.GetAssociatedResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAssociatedResourceOutput, error)
	GetAttributeGroup(context.Context, *servicecatalogappregistry.GetAttributeGroupInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetAttributeGroupOutput, error)
	GetConfiguration(context.Context, *servicecatalogappregistry.GetConfigurationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.GetConfigurationOutput, error)
	ListApplications(context.Context, *servicecatalogappregistry.ListApplicationsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListApplicationsOutput, error)
	ListAssociatedAttributeGroups(context.Context, *servicecatalogappregistry.ListAssociatedAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedAttributeGroupsOutput, error)
	ListAssociatedResources(context.Context, *servicecatalogappregistry.ListAssociatedResourcesInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAssociatedResourcesOutput, error)
	ListAttributeGroups(context.Context, *servicecatalogappregistry.ListAttributeGroupsInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsOutput, error)
	ListAttributeGroupsForApplication(context.Context, *servicecatalogappregistry.ListAttributeGroupsForApplicationInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListAttributeGroupsForApplicationOutput, error)
	ListTagsForResource(context.Context, *servicecatalogappregistry.ListTagsForResourceInput, ...func(*servicecatalogappregistry.Options)) (*servicecatalogappregistry.ListTagsForResourceOutput, error)
}
