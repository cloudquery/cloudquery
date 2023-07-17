// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
)

//go:generate mockgen -package=mocks -destination=../mocks/appconfig.go -source=appconfig.go AppconfigClient
type AppconfigClient interface {
	GetApplication(context.Context, *appconfig.GetApplicationInput, ...func(*appconfig.Options)) (*appconfig.GetApplicationOutput, error)
	GetConfiguration(context.Context, *appconfig.GetConfigurationInput, ...func(*appconfig.Options)) (*appconfig.GetConfigurationOutput, error)
	GetConfigurationProfile(context.Context, *appconfig.GetConfigurationProfileInput, ...func(*appconfig.Options)) (*appconfig.GetConfigurationProfileOutput, error)
	GetDeployment(context.Context, *appconfig.GetDeploymentInput, ...func(*appconfig.Options)) (*appconfig.GetDeploymentOutput, error)
	GetDeploymentStrategy(context.Context, *appconfig.GetDeploymentStrategyInput, ...func(*appconfig.Options)) (*appconfig.GetDeploymentStrategyOutput, error)
	GetEnvironment(context.Context, *appconfig.GetEnvironmentInput, ...func(*appconfig.Options)) (*appconfig.GetEnvironmentOutput, error)
	GetExtension(context.Context, *appconfig.GetExtensionInput, ...func(*appconfig.Options)) (*appconfig.GetExtensionOutput, error)
	GetExtensionAssociation(context.Context, *appconfig.GetExtensionAssociationInput, ...func(*appconfig.Options)) (*appconfig.GetExtensionAssociationOutput, error)
	GetHostedConfigurationVersion(context.Context, *appconfig.GetHostedConfigurationVersionInput, ...func(*appconfig.Options)) (*appconfig.GetHostedConfigurationVersionOutput, error)
	ListApplications(context.Context, *appconfig.ListApplicationsInput, ...func(*appconfig.Options)) (*appconfig.ListApplicationsOutput, error)
	ListConfigurationProfiles(context.Context, *appconfig.ListConfigurationProfilesInput, ...func(*appconfig.Options)) (*appconfig.ListConfigurationProfilesOutput, error)
	ListDeploymentStrategies(context.Context, *appconfig.ListDeploymentStrategiesInput, ...func(*appconfig.Options)) (*appconfig.ListDeploymentStrategiesOutput, error)
	ListDeployments(context.Context, *appconfig.ListDeploymentsInput, ...func(*appconfig.Options)) (*appconfig.ListDeploymentsOutput, error)
	ListEnvironments(context.Context, *appconfig.ListEnvironmentsInput, ...func(*appconfig.Options)) (*appconfig.ListEnvironmentsOutput, error)
	ListExtensionAssociations(context.Context, *appconfig.ListExtensionAssociationsInput, ...func(*appconfig.Options)) (*appconfig.ListExtensionAssociationsOutput, error)
	ListExtensions(context.Context, *appconfig.ListExtensionsInput, ...func(*appconfig.Options)) (*appconfig.ListExtensionsOutput, error)
	ListHostedConfigurationVersions(context.Context, *appconfig.ListHostedConfigurationVersionsInput, ...func(*appconfig.Options)) (*appconfig.ListHostedConfigurationVersionsOutput, error)
	ListTagsForResource(context.Context, *appconfig.ListTagsForResourceInput, ...func(*appconfig.Options)) (*appconfig.ListTagsForResourceOutput, error)
}
