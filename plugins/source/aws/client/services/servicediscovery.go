// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

//go:generate mockgen -package=mocks -destination=../mocks/servicediscovery.go -source=servicediscovery.go ServicediscoveryClient
type ServicediscoveryClient interface {
	GetInstance(context.Context, *servicediscovery.GetInstanceInput, ...func(*servicediscovery.Options)) (*servicediscovery.GetInstanceOutput, error)
	GetInstancesHealthStatus(context.Context, *servicediscovery.GetInstancesHealthStatusInput, ...func(*servicediscovery.Options)) (*servicediscovery.GetInstancesHealthStatusOutput, error)
	GetNamespace(context.Context, *servicediscovery.GetNamespaceInput, ...func(*servicediscovery.Options)) (*servicediscovery.GetNamespaceOutput, error)
	GetOperation(context.Context, *servicediscovery.GetOperationInput, ...func(*servicediscovery.Options)) (*servicediscovery.GetOperationOutput, error)
	GetService(context.Context, *servicediscovery.GetServiceInput, ...func(*servicediscovery.Options)) (*servicediscovery.GetServiceOutput, error)
	ListInstances(context.Context, *servicediscovery.ListInstancesInput, ...func(*servicediscovery.Options)) (*servicediscovery.ListInstancesOutput, error)
	ListNamespaces(context.Context, *servicediscovery.ListNamespacesInput, ...func(*servicediscovery.Options)) (*servicediscovery.ListNamespacesOutput, error)
	ListOperations(context.Context, *servicediscovery.ListOperationsInput, ...func(*servicediscovery.Options)) (*servicediscovery.ListOperationsOutput, error)
	ListServices(context.Context, *servicediscovery.ListServicesInput, ...func(*servicediscovery.Options)) (*servicediscovery.ListServicesOutput, error)
	ListTagsForResource(context.Context, *servicediscovery.ListTagsForResourceInput, ...func(*servicediscovery.Options)) (*servicediscovery.ListTagsForResourceOutput, error)
}
