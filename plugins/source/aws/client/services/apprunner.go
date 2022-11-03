// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

//go:generate mockgen -package=mocks -destination=../mocks/apprunner.go -source=apprunner.go ApprunnerClient
type ApprunnerClient interface {
	DescribeAutoScalingConfiguration(context.Context, *apprunner.DescribeAutoScalingConfigurationInput, ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error)
	DescribeCustomDomains(context.Context, *apprunner.DescribeCustomDomainsInput, ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error)
	DescribeObservabilityConfiguration(context.Context, *apprunner.DescribeObservabilityConfigurationInput, ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error)
	DescribeService(context.Context, *apprunner.DescribeServiceInput, ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error)
	DescribeVpcConnector(context.Context, *apprunner.DescribeVpcConnectorInput, ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error)
	DescribeVpcIngressConnection(context.Context, *apprunner.DescribeVpcIngressConnectionInput, ...func(*apprunner.Options)) (*apprunner.DescribeVpcIngressConnectionOutput, error)
	ListAutoScalingConfigurations(context.Context, *apprunner.ListAutoScalingConfigurationsInput, ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error)
	ListConnections(context.Context, *apprunner.ListConnectionsInput, ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error)
	ListObservabilityConfigurations(context.Context, *apprunner.ListObservabilityConfigurationsInput, ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error)
	ListOperations(context.Context, *apprunner.ListOperationsInput, ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error)
	ListServices(context.Context, *apprunner.ListServicesInput, ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error)
	ListTagsForResource(context.Context, *apprunner.ListTagsForResourceInput, ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error)
	ListVpcConnectors(context.Context, *apprunner.ListVpcConnectorsInput, ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error)
	ListVpcIngressConnections(context.Context, *apprunner.ListVpcIngressConnectionsInput, ...func(*apprunner.Options)) (*apprunner.ListVpcIngressConnectionsOutput, error)
}
